package main

import (
	"fmt"
	"github.com/go-gin-example/models"
	"github.com/go-gin-example/pkg/setting"
	"github.com/go-gin-example/routers"
	"net/http"
)

func init()  {
	setting.SetUpSetting()
	models.ModelInit()
}

func main() {
	//fmt.Println(setting.DatabaseS.Type)
	router:=routers.InitRouter()

	s:=&http.Server{
		Addr: fmt.Sprintf(":%d",setting.ServerS.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerS.ReadTimeout,
		WriteTimeout: setting.ServerS.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}

