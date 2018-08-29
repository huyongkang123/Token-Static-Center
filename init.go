

package main

import (
	"github.com/LuRenJiasWorld/Token-Static-Center/util"
	"flag"
	"strconv"
	"github.com/LuRenJiasWorld/Token-Static-Center/core"
	"net/http"
	"fmt"
)

func main() {
	// Step 1. 从命令行参数获取配置文件路径
	configFilePath := flag.String("config", "", "配置文件路径，必须指定，无默认值！")
	flag.Parse()

	// Step 2. 加载配置文件
	err := util.ReadConfig(*configFilePath)

	if err != nil {
		util.ErrorLog("main", err.Error(), "main->ReadConfig")
		return
	}

	util.OperationLog("main", "配置文件加载成功，路径为" + *configFilePath, "main->ReadConfig")

	// Step 3. 从配置文件读取监听端口
	// 注意，此处defaultPort返回回来的是一个Interface
	defaultPortInterface, err := util.GetConfig("Global", "Port")

	// 将Interface转换成int
	listenPort := defaultPortInterface.(uint32)

	if err != nil {
		util.ErrorLog("main", "无法读取默认配置！" + err.Error(), "main->GetConfig")
		return
	}

	util.OperationLog("main", "使用配置文件中的端口" + strconv.Itoa(int(listenPort)), "main->listenPort")

	// 输出欢迎信息
	util.AccessLog("main", "欢迎使用 Token-Static-Center v1.02, 初始化完成，开始接受外部请求，启动服务", "main")


	// Step 4. 开始监听
	server := core.NewServer()
	util.ErrorLog("main", http.ListenAndServe(fmt.Sprintf(":%d", listenPort), server).Error(), "main->ListenAndServe")

}