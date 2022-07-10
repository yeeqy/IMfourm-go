package v1

import (
	"IMfourm-go/app/models/category"
	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context){
	req := requests.PaginationRequest{}
	if ok := requests.Validate(c,&req,requests.Pagination); !ok {
		return
	}
	data, pager := category.Paginate(c,10)
	response.JSON(c,gin.H{
		"data":data,
		"pager":pager,
	})
}
//
//func(ctrl *CategoriesController) Show(c *gin.Context){
//	categoryModel := category.Get(c.Param("id"))
//	if categoryModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//	response.Data(c,categoryModel)
//}

func (ctrl *CategoriesController) Store(c *gin.Context) {
	var req = requests.CategoryRequest{}
	if ok := requests.Validate(c,&req,requests.CategorySave);!ok{
		return
	}
	categoryModel := category.Category{
		Name: req.Name,
		Description: req.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0{
		response.Created(c,categoryModel)
	} else {
		response.Abort500(c,"创建失败，请稍后再试")
	}
}

func(ctrl *CategoriesController) Update(c *gin.Context){
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	req := requests.CategoryRequest{}
	if ok := requests.Validate(c,&req,requests.CategorySave); !ok{
		return
	}
	categoryModel.Name = req.Name
	categoryModel.Description = req.Description
	rowsAffected := categoryModel.Save()
	if rowsAffected > 0 {
		response.Data(c,categoryModel)
	} else {
		response.Abort500(c)
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context){
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}
	response.Abort500(c,"删除失败，请稍后再试")
}