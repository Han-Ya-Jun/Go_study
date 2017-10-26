package service

import (
	"github.com/emicklei/go-restful"
	log "github.com/Sirupsen/logrus"
	"time"
)

/**
 *CreateBy 01305155
 *Date:10:56
 *Description::
*/
type UserController struct {
	WebService  *restful.WebService
	UserMgr *UserManager
}


func NewUserController(userMgr *UserManager) *UserController {
	return &UserController{
		WebService:  new(restful.WebService),
		UserMgr:userMgr,
	}
}
func (controller *UserController) Register(container *restful.Container) {
	controller.WebService.
		Path("/user/").
		Doc("user service").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	controller.WebService.Route(controller.WebService.
		GET("/getusers").
		To(controller.GetUsers).
		Doc("get all users").
		Operation("get users").
		Writes(RespWrapper{}))

	controller.WebService.Route(controller.WebService.
		GET("/get/{id}").
		To(controller.GetUsersById).
		Doc("get user by id").
		Operation("get user").
		Writes(RespWrapper{}))
	container.Add(controller.WebService)
}
func (controller *UserController) GetUsers(request *restful.Request, response *restful.Response) {
	user,err:=controller.UserMgr.GetUsers()
    if err!=nil{
		log.Error(err)
	}
	controller.responseSuccessWithObj(response, "", user)
}

func (controller *UserController) GetUsersById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	user,err:=controller.UserMgr.GetUsersById(id)
	if err!=nil{
		log.Error(err)
	}
	controller.responseSuccessWithObj(response, "", user)
}
func (controller *UserController) responseSuccessWithObj(response *restful.Response, requestId string, obj interface{}) {
	if response == nil {
		return
	}
	resp := &RespWrapper{
		RequestId: requestId,
		Success:   true,
		Date:      time.Now().Format("2006-01-02 15:04:05"),
		Obj:       obj,
	}
	respErr := response.WriteAsJson(resp)
	if respErr != nil {
		log.Errorf("write http response error: %s", respErr.Error())
	}
}