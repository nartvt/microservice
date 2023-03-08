package router

import (
	"github.com/gofiber/fiber/v2"

	aboutRepo "health-service/app/domain/usercases/about/repo"
	dishRepo "health-service/app/domain/usercases/dish/repo"
	bodyRecordRepo "health-service/app/domain/usercases/record/body/repo"
	diaryRecordRepo "health-service/app/domain/usercases/record/diary/repo"
	exerciseRecordRepo "health-service/app/domain/usercases/record/exercise/repo"
	sectionRepo "health-service/app/domain/usercases/section/repo"

	"health-service/app/middleware"
	dishHandler "health-service/app/transport/dish/handler"
	sectionHandler "health-service/app/transport/section/handler"

	bodyRecordHandler "health-service/app/transport/record/body/handler"
	diaryRecordHandler "health-service/app/transport/record/diary/handler"
	exerciseRecordHandler "health-service/app/transport/record/exercise/handler"

	aboutHandler "health-service/app/transport/about/handler"
)

func SetupRoutes(app *fiber.App) {

	groupSectionHandler := sectionHandler.SectionHandler{
		SectionDomain: sectionRepo.NewNewsfeedSectionRepo(),
	}
	groupDishHandler := dishHandler.DishHandler{
		DishDomain: dishRepo.NewDishRepo(),
	}

	groupBodyRecordHandler := bodyRecordHandler.BodyRecordHandler{
		BodyRecordDomain: bodyRecordRepo.NewUserBodyRecordRepo(),
	}

	groupDiaryRecordHandler := diaryRecordHandler.DiaryHandler{
		DiaryDomain: diaryRecordRepo.NewUserBodyRecordRepoRepo(),
	}

	groupExerciseRecordHandler := exerciseRecordHandler.ExerciseHandler{
		ExerciseDomain: exerciseRecordRepo.NewUserExerciseRecordRepoRepo(),
	}

	groupAboutHandler := aboutHandler.AboutHandler{
		AboutDomain: aboutRepo.NewAboutRepo(),
	}

	v1 := app.Group("/api/v1")
	v1.Use(middleware.CorsFilter())
	v1.Use(middleware.RateLimit())

	groupNewsfeedSection := v1.Group("/newsfeed/sections")
	groupNewsfeedSection.Use(middleware.RequireLogin())
	{
		GET(groupNewsfeedSection, "", groupSectionHandler.GetSections)
	}

	groupDishes := groupNewsfeedSection.Group("dishes")
	{
		GET(groupDishes, "", groupDishHandler.CreateDish)
		GET(groupDishes, "/:sectionId", groupDishHandler.GetDishBySectionId)
	}

	groupBodyRecords := v1.Group("/body-records")
	{
		GET(groupBodyRecords, "/:userId", groupBodyRecordHandler.GetBodyRecordsByUserId)
	}

	groupDiary := v1.Group("/diary-records")
	{
		GET(groupDiary, "/:userId", groupDiaryRecordHandler.GetDiariesByUserId)
	}

	groupExercise := v1.Group("/exercise-records")
	{
		GET(groupExercise, "/:userId", groupExerciseRecordHandler.GetExerciseByUserId)
	}

	groupAbout := v1.Group("/abouts")
	{
		GET(groupAbout, "/:sectionId", groupAboutHandler.GetAboutBySectionId)
	}

}

func GET(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "GET", relativePath, f)
}

func POST(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "POST", relativePath, f)
}
func PUT(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "PUT", relativePath, f)
}

func DELETE(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "DELETE", relativePath, f)
}
func route(app fiber.Router, method string, relativePath string, f fiber.Handler) {
	switch method {
	case "POST":
		app.Post(relativePath, f)
	case "GET":
		app.Get(relativePath, f)
	case "PUT":
		app.Put(relativePath, f)
	case "DELETE":
		app.Delete(relativePath, f)
	}
}
