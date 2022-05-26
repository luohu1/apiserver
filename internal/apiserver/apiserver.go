package apiserver

import (
	"github.com/gin-gonic/gin"

	"github.com/luohu1/gin-apiserver/internal/apiserver/controller/v1/user"
)

func RunServer() error {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin")
	})

	v1 := r.Group("/api/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := new(user.UserController)
			userv1.POST("/", userController.Create)
			userv1.GET("/:uid", userController.Retrieve)
			userv1.PUT("/:uid", userController.Update)
			userv1.DELETE("/:uid", userController.Delete)
			userv1.GET("/", userController.List)
		}
	}

	if err := r.Run(":8080"); err != nil {
		return err
	} else {
		return nil
	}
}
