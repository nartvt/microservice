package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	conf "health-service/app/config"
	"health-service/app/domain/usercases/user/repo"
)

const (
	expireTime  = 86400
	prefixToken = "Bearer"
	roleKey     = "role"
	userKey     = "user"
)

type jwtToken struct {
}

type userAuth struct {
	userId   int
	userName string
}

var Jwt jwtToken
var (
	jwtSecretKey      = conf.Config.SecretKey // replace with your own secret key
	jwtTokenDuration  = time.Hour * 24        // customize token duration based on your needs
	jwtIssuer         = "health-service"      // replace with your own issuer name
	jwtRefreshExpires = time.Hour * 24 * 30   // customize refresh token duration based on your needs
)

func (j *jwtToken) RequireLogin() fiber.Handler {
	return j.GetCurrentUserLogin()
}
func (*jwtToken) GetCurrentUserLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Get("Authorization")
		if authToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization header",
			})
		}

		tokenString := authToken[len(prefixToken):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
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
			c.Locals(userKey, claims[userKey])
			c.Locals(roleKey, claims[roleKey])
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
