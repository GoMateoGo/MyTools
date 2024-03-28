package main

import (
	"fmt"
	"proc_guard/cmd"
	"proc_guard/global"
	"proc_guard/utils"
	"time"
)

func main() {
	//初始化
	global.Logger = cmd.InitLogger()
	// 读取守护的进程数据
	cmd.LoadRunTxt()
	// 启动守护进程

	go func() {
		for {
			for name, path := range global.RunPath {
				isFine, appName := cmd.IsProcessExist(name)
				if !isFine {
					fmt.Println("进程不存在,启动进程...", appName)
					utils.RunGame(appName, path)
					global.Logger.Info("进程不存在,启动进程名称:" + appName)
				}
			}
			<-time.After(time.Second * 5)
			fmt.Println(time.Now().Format(time.DateTime) + ":每5秒检查进程...")
		}
	}()

	select {}
}
