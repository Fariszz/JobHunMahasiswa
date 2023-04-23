package router

import (
	"database/sql"
	"mahasiswa/handler"
	"mahasiswa/modules/mahasiswa"

	"github.com/gin-gonic/gin"
)

func Init(db *sql.DB) *gin.Engine {

	mahasiswaRepository := mahasiswa.NewRepository(db)
	mahasiswaService := mahasiswa.NewService(mahasiswaRepository)
	handler := handler.NewMahasiswaHandler(mahasiswaService)
	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/mahasiswa", handler.GetAllMahasiswa)
	api.GET("/mahasiswa/:id", handler.GetDetailMahasiswa)
	api.POST("/mahasiswa", handler.StoreMahasiswa)
	api.PUT("/mahasiswa/:id", handler.UpdateMahasiswa)
	api.DELETE("/mahasiswa/:id", handler.DeleteMahasiswa)
	
	return router
}
