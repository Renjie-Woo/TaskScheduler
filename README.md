# README

本项目用于批量定时调度python脚本

使用方式：
```shell
go run ./cmd/main.go --work_dir=your/work/dir --config_path=your/config/path --frequency=10
```
或编译代码为执行程序
```shell
# WINDOWS task_scheduler.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o task_scheduler.exe cmd/main.go
# LINUX task_scheduler: TODO
# MAC task_scheduler: TODO
```
后通过以下方式运行
```shell
task_scheduler --work_dir=your/work/dir --config_path=your/config/path --frequency=10
```

其中参数：
```shell
work_dir: 需要批量运行的python脚本的目录绝对路径
config_path: 脚本调度配置文件路径(JSON)
frequency: 调度脚本刷新频率, 默认10(SECOND)
```

配置文件格式：
```json lines
[
  {
    "task_name": "你的任务名称",
    "script_path": "任务对应脚本",
    "params": [], // 默认参数, 当不同规则使用相同参数时，在此设置
    "rules": [
      {
        "rule": "* * * * * *", // cron 规则，支持秒级
        "params": [] // 规则对应脚本独立运行参数，如果想要使用默认参数，删除该字段或者参数值值为null
      }
    ]
  }
]
```