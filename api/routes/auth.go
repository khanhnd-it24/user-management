package routes

import (
	"github.com/gin-gonic/gin"
	"user-management/adapter/databases"
	"user-management/adapter/repositories"
	"user-management/api/controllers"
	"user-management/api/services"
)

func AuthRoutesSetup(r *gin.Engine) {
	authRepository := repositories.New(databases.DB)
	authService := services.New(authRepository)
	authController := controllers.New(authService)

	r.POST("/api/users/admin", authController.Register)
}
