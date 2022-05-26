package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) List(c *gin.Context) {
	c.String(http.StatusOK, "UserController List")
}
