package v1

import (
	"IMfourm-go/app/models/user"
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

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context)  {
	data := user.All()
	response.Data(c,data)

}