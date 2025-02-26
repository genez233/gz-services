package main

import (
	"github.com/genez233/go-utils/glog"
	"gz-services/servers/user_server/global"
	"gz-services/servers/user_server/model"
	"gz-services/servers/user_server/pkg/sett"
	"gz-services/servers/user_server/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	// region 初始化配置文件
	st, err := sett.NewSetting()
	if err != nil {
		log.Fatal(err)
	}
	err = st.ReadSection("Server", &global.Server)
	if err != nil {
		return
	}
	err = st.ReadSection("App", &global.App)
	if err != nil {
		return
	}
	err = st.ReadSection("Database", &global.Database)
	if err != nil {
		return
	}
	err = st.ReadSection("Jwt", &global.JWT)
	if err != nil {
		return
	}
	err = st.ReadSection("Log", &global.LogSetting)

	global.JWT.Expire *= time.Hour
	global.Server.ReadTimeout *= time.Second
	global.Server.WriteTimeout *= time.Second
	// endregion

	err = setupDBEngine()
	if err != nil {
		return
	}

	global.Logger = glog.New(&glog.Config{
		ServerName:       global.Server.ServerName,
		Version:          global.Server.Version,
		RunMode:          global.Server.RunMode,
		ConsoleLog:       true,
		IsUpload:         true,
		LogUrl:           global.LogSetting.LogUrl,
		OpenobserveToken: global.LogSetting.OpenobserveToken,
	})
}

// setup connect db
func setupDBEngine() error {
	var err error
	global.DB, err = model.NewEngine()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//gen.GenEntity("D:\\Documents\\GitHub\\gz-services\\servers\\user_server\\model", global.DB)

	global.Logger.Info(global.Server.ServerName + " start with port " + global.Server.HttpPort)

	router := routers.GetRouter()

	s := &http.Server{
		Addr:           ":" + global.Server.HttpPort,
		Handler:        router,
		ReadTimeout:    global.Server.ReadTimeout,
		WriteTimeout:   global.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
