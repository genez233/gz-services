package main

import (
	"fmt"
	"github.com/genez233/go-utils/gen"
	"github.com/genez233/go-utils/glog"
	"gz-services/servers/user_server/global"
	"gz-services/servers/user_server/model"
	"gz-services/servers/user_server/pkg/sett"
	"log"
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

	global.JWT.Expire *= time.Hour
	global.Server.ReadTimeout *= time.Second
	global.Server.WriteTimeout *= time.Second
	// endregion

	err = setupDBEngine()
	if err != nil {
		return
	}

	global.Log = glog.New(&glog.Config{
		ServerName:       global.Server.ServerName,
		Version:          global.Server.Version,
		RunMode:          global.Server.RunMode,
		ConsoleLog:       true,
		IsUpload:         true,
		LogUrl:           "",
		OpenobserveToken: "",
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

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	gen.GenEntity("D:\\Documents\\GitHub\\gz-services\\servers\\user_server\\model")

	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
