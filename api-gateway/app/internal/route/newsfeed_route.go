package router

import (
	productRepo "api-gateway/app/domain/usercases/product/repo"
	bodyRecordRepo "api-gateway/app/domain/usercases/record/body/repo"
	diaryRecordRepo "api-gateway/app/domain/usercases/record/diary/repo"
	exerciseRecordRepo "api-gateway/app/domain/usercases/record/exercise/repo"
	sectionRepo "api-gateway/app/domain/usercases/section/repo"
	productHandler "api-gateway/app/transport/product/handler"
	bodyRecordHandler "api-gateway/app/transport/record/body/handler"
	diaryRecordHandler "api-gateway/app/transport/record/diary/handler"
	exerciseRecordHandler "api-gateway/app/transport/record/exercise/handler"
	sectionHandler "api-gateway/app/transport/section/handler"
)

func setupNewsfeedRoute(route fiber.Router) {

	groupSectionHandler := sectionHandler.SectionHandler{
		SectionDomain: sectionRepo.NewNewsfeedSectionRepo(),
	}
	groupProductHandler := productHandler.ProductHandler{
		ProductDomain: productRepo.NewProductRepo(),
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

	groupNewsfeedSection := route.Group("/newsfeed/sections")
	{
		GET(groupNewsfeedSection, "", groupSectionHandler.GetSections)
	}

	groupProduct := groupNewsfeedSection.Group("product")
	{
		GET(groupProduct, "", groupProductHandler.CreateProduct)
		GET(groupProduct, "/:sectionId", groupProductHandler.GetProductBySectionId)
	}

	groupBodyRecords := route.Group("/body-records")
	{
		GET(groupBodyRecords, "", groupBodyRecordHandler.GetBodyRecordsByUserId)
	}

	groupDiary := route.Group("/diary-records")
	{
		GET(groupDiary, "", groupDiaryRecordHandler.GetDiariesByUserId)
	}

	groupExercise := route.Group("/exercise-records")
	{
		GET(groupExercise, "", groupExerciseRecordHandler.GetExerciseByUserId)
	}
}
