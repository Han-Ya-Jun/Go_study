package service

import (
	"github.com/emicklei/go-restful"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"net/http"
)

/**
 *CreateBy 01305155
 *Date:10:43
 *Description:操作数据库
*/

type UserService struct {
	Controller *UserController
	Container  *restful.Container
	Port       string
}

func NewScheduleServer(port string, controller *UserController) *UserService {
	container := restful.NewContainer()
	server := &UserService{
		Controller: controller,
		Container:  container,
		Port:       port,
	}
	server.Controller.Register(container)
	return server
}

func (server *UserService) Start() {
	addr := fmt.Sprintf(":%s", server.Port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: server.Container,
	}
	log.Infof("schedule server[port=%s] started!", server.Port)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Errorf("start server error: ", err)
	}
}
