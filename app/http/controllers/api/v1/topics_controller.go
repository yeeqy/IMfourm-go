package v1

import (
	"IMfourm-go/app/models/topic"
	"IMfourm-go/pkg/auth"

	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Index(c *gin.Context){
	topics := topic.All()
	response.Data(c,topics)
}

func(ctrl *TopicsController) Show(c *gin.Context){
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c,topicModel)
}

func (ctrl *TopicsController) Store(c *gin.Context) {
	var req = requests.TopicRequest{}
	if ok := requests.Validate(c,&req,requests.TopicSave);!ok{
		return
	}
	topicModel := topic.Topic{
		Title: req.Title,
		Body: req.Body,
		CategoryID: req.CategoryID,
		UserId: auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0{
		response.Created(c,topicModel)
	} else {
		response.Abort500(c,"创建失败，请稍后再试")
	}
}

//func(ctrl *TopicsController) Update(c *gin.Context){
//	topicModel := topic.Get(c.Param("id"))
//	if topicModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//
//	req := requests.TopicRequest{}
//	bindOK,errs := requests.Validate(c,&req,requests.TopicSave)
//	if !bindOK{
//		return
//	}
//	if len(errs) > 0 {
//		response.ValidationError(c,20101,errs)
//		return
//	}
//	topicModel.FieldName = req.FieldName
//	rowsAffected := topicModel.Save()
//	if rowsAffected > 0 {
//		response.Data(c,topicModel)
//	} else {
//		response.Abort500(c,"更新失败，请稍后再试")
//	}
//}
//
//func (ctrl *TopicsController) Delete(c *gin.Context){
//	topicModel := topic.Get(c.Param("id"))
//	if topicModel.ID == 0 {
//		response.Abort404(c)
//		return
//	}
//	if ok := policies.CanModifyTopic(c,VariableNameModel); !ok{
//		response.Abort403(c)
//		return
//	}
//	rowsAffected := topicModel.Delete()
//	if rowsAffected > 0 {
//		response.Success(c)
//		return
//	}
//	response.Abort500(c,"删除失败，请稍后再试")
//}