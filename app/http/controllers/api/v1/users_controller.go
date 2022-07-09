package v1

import (
	"IMfourm-go/app/models/user"
	"IMfourm-go/app/requests"
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

	req := requests.PaginationRequest{}
	if ok := requests.Validate(c,&req,requests.Pagination); !ok{
		return
	}

	data,pager := user.Paginate(c,10)
	response.JSON(c,gin.H{
		"data":data,
		"pager":pager,
	})

}