package routes

import (
	"github.com/gin-gonic/gin"
)

func registerDiaryRoutes(router *gin.RouterGroup) {
	router.GET("/detail", diaryController.GetDiaryWithImages)
	router.GET("/all", diaryController.GetAllDiaries)
	router.POST("/new", diaryController.CreateDiary)
	router.GET("/latest", diaryController.GetLatestDiary)
	router.PUT("/update", diaryController.UpdateDiary)
	router.DELETE("/delete", diaryController.DeleteDiary)

	router.GET("/search/t", diaryController.SearchDiaryByTitle) // t for title
	router.GET("/search/c", diaryController.SearchDiaryByContent) // c for content
	router.GET("/search/d", diaryController.SearchDiaryByDiaryDate) // d for diary date (diary_date)
}
