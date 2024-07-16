package routers

import (
	"UjianKetrampilan/controllers"
	"UjianKetrampilan/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.CreateUser)
	r.POST("/users/login", controllers.Login)
	r.PUT("/users", middleware.Authentication, controllers.UpdateUser)
	r.DELETE("/users", middleware.Authentication, controllers.DeleteUser)

	r.GET(
		"/photos",
		middleware.Authentication,
		controllers.GetPhotos,
	)

	r.POST(
		"/photos",
		middleware.Authentication,
		controllers.CreatePhoto,
	)

	r.PUT(
		"/photos/:photoID",
		middleware.Authentication,
		middleware.Authorization,
		controllers.UpdatePhoto,
	)

	r.DELETE(
		"/photos/:photoID",
		middleware.Authentication,
		middleware.Authorization,
		controllers.DeletePhoto,
	)

	return r
}
