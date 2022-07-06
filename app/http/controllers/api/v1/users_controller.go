package v1

import (
	"IMfourm-go/pkg/auth"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context){
	userModel := auth.CurrentUser(c)
	response.Data(c,userModel)
}