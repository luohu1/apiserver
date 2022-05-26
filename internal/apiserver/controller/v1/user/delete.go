package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "UserController Delete")
}
