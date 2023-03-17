package middleware

import (
	"fmt"
	"strings"
	"time"

	conf "api-gateway/app/config"
	"api-gateway/app/domain/usercases/user/repo"
)

const (
	expireTime  = 86400
	prefixToken = "Bearer"
	roleKey     = "role"
	userKey     = "user"
	userNameKey = "userName"
)

type jwtToken struct {
}

type userAuth struct {
	userId   int
	userName string
}

var Auth *jwtToken
var (
	jwtSecretKey      = conf.Config.SecretKey // replace with your own secret key
	jwtTokenDuration  = time.Hour * 24        // customize token duration based on your needs
	jwtIssuer         = "api-gateway"         // replace with your own issuer name
	jwtRefreshExpires = time.Hour * 24 * 30   // customize refresh token duration based on your needs
)

func init() {
	Auth = &jwtToken{}
}

func (t *jwtToken) RequireLogin() fiber.Handler {
	return t.deserializeUser()
}
func (*jwtToken) deserializeUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := ""
		authToken := c.Get("Authorization")
		if strings.HasPrefix(authToken, prefixToken) {
			tokenString = authToken[len(prefixToken):]
		}
		if c.Cookies("token") != "" {
			tokenString = c.Cookies("token")
		}

		if authToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization header",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid authorization token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expiration := int64(claims["exp"].(float64))
			if time.Unix(expiration, 0).Sub(time.Now()) < 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token has expired",
				})
			}

			user, err := repo.User.GetUserByUserName(userNameKey)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
			}
			if user == nil {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
			}
			c.Locals(userKey, user)
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
}
func (*jwtToken) generateAccessToken(user *userAuth) (string, error) {
	claims := jwt.MapClaims{
		"id":     user.userId,
		"iss":    fmt.Sprintf("%s-%s", jwtIssuer, jwtSecretKey),
		"exp":    time.Now().Add(jwtTokenDuration).Unix(),
		"sub":    user.userName,
		"scopes": []string{"user"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
func (*jwtToken) generateRefreshToken() (string, error) {
	claims := jwt.MapClaims{
		"iss":    fmt.Sprintf("%s-%s", jwtIssuer, jwtSecretKey),
		"exp":    time.Now().Add(jwtRefreshExpires).Unix(),
		"scopes": []string{"refresh_token"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

// for login
func (j *jwtToken) authenticate() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&request); err != nil {
			return fiber.ErrBadRequest
		}

		// validate current user system
		currentUser, err := repo.User.GetUserByUserName(request.Username)
		if err != nil {
			return err
		}
		if currentUser == nil || currentUser.Password != request.Password {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid authorization token",
			})
		}
		// Authenticate user (e.g., check if username and password are valid)
		// ...

		user := &userAuth{
			userId:   currentUser.Id,
			userName: request.Username,
		}
		// Generate access token
		accessToken, err := j.generateAccessToken(user)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		// Generate refresh token
		refreshToken, err := j.generateRefreshToken()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: accessToken,
			Path:  "/",
		})
		c.Cookie(&fiber.Cookie{
			Name:  "refresh_token",
			Value: refreshToken,
			Path:  "/",
		})

		return c.Next()
	}
}
