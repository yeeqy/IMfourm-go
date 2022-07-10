package v1

import (
	"IMfourm-go/app/models/link"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}


func (ctrl *LinksController) Index(c *gin.Context){
	//调用缓存后的数据
	response.Data(c,link.AllCache())
}

//func(ctrl *LinksController) Show(c *gin.Context){
//	linkModel := link.Get(c.Param("id"))
//	if linkModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//	response.Data(c,linkModel)
//}
//
//func (ctrl *LinksController) Store(c *gin.Context) {
//	var req = requests.LinkRequest{}
//	if ok := requests.Validate(c,&req,requests.LinkSave);!ok{
//		return
//	}
//	linkModel := link.Link{
//		FieldName: req.FieldName,
//	}
//	linkModel.Create()
//	if linkModel.ID > 0{
//		response.Created(c,linkModel)
//	} else {
//		response.Abort500(c,"创建失败，请稍后再试")
//	}
//}
//
//func(ctrl *LinksController) Update(c *gin.Context){
//	linkModel := link.Get(c.Param("id"))
//	if linkModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//	if ok:= policies.CanModifyLink(c,VariableNameModel); !ok{
//		response.Abort403(c)
//		return
//	}
//	req := requests.LinkRequest{}
//	bindOK,errs := requests.Validate(c,&req,requests.LinkSave)
//	if !bindOK{
//		return
//	}
//	if len(errs) > 0 {
//		response.ValidationError(c,20101,errs)
//		return
//	}
//	linkModel.FieldName = req.FieldName
//	rowsAffected := linkModel.Save()
//	if rowsAffected > 0 {
//		response.Data(c,linkModel)
//	} else {
//		response.Abort500(c,"更新失败，请稍后再试")
//	}
//}
//
//func (ctrl *LinksController) Delete(c *gin.Context){
//	linkModel := link.Get(c.Param("id"))
//	if linkModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//	if ok := policies.CanModifyLink(c,VariableNameModel); !ok{
//		response.Abort403(c)
//		return
//	}
//	rowsAffected := linkModel.Delete()
//	if rowsAffected > 0 {
//		response.Success(c)
//		return
//	}
//	response.Abort500(c,"删除失败，请稍后再试")
//}