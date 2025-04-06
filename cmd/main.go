package main

import (
	"doraemon/TaskScheduler/internal/model"
	log "doraemon/TaskScheduler/pkg/logger"
	"flag"
	"fmt"
)

var (
	workDir   = flag.String("work_dir", "", "待运行脚本的工作目录")
	cfgPath   = flag.String("config_path", "", "配置文件绝对路径")
	freshFreq = flag.Int("frequency", 10, "配置轮询频率，秒")
)

func main() {
	flag.Parse()

	var err error

	if *workDir == "" {
		fmt.Println("没有配置有效工作目录，请使用--work_dir配置")
		return
		//*workDir, err = os.Getwd()
		//if err != nil {
		//	fmt.Println("没有配置有效工作目录，请使用--work_dir配置")
		//	return
		//}
	}

	if *cfgPath == "" {
		fmt.Println("没有配置运行参数文件路径，请使用--config_path配置")
		return
	}

	fmt.Printf(">>> 工作目录: %s\n>>> 配置文件路径: %s\n>>> 配置文件刷新频率: %d\n", *workDir, *cfgPath, *freshFreq)

	var logger = log.NewConsoleLogger()
	defer func() {
		err = logger.Sync()
		if err != nil {
			_ = fmt.Errorf("关闭日志失败 %v", err)
		}
	}()

	var cfgController = model.NewFileConfig(*cfgPath, int64(*freshFreq))
	var taskScheduler = model.NewTaskScheduler(*workDir, cfgController, logger)

	taskScheduler.Run()
}
