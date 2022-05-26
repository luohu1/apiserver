package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Update(c *gin.Context) {
	c.String(http.StatusOK, "UserController Update")
}
