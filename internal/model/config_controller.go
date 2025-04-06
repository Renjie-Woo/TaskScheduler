package model

import (
	"doraemon/TaskScheduler/pkg/tool"
	"time"
)

type configController interface {
	GetTicker() *time.Ticker
	Refresh() (map[string]singleTask, error)
}

type FileConfig struct {
	CfgPath   string
	FreshFreq int64
	ticker    *time.Ticker
}

func NewFileConfig(cfgPath string, freshFreq int64) *FileConfig {

	return &FileConfig{
		CfgPath:   cfgPath,
		FreshFreq: freshFreq,
		ticker:    time.NewTicker(time.Duration(freshFreq) * time.Second),
	}
}

func (f *FileConfig) Refresh() (map[string]singleTask, error) {
	var configs = make(configList, 0)
	var err = tool.ReadStruct(f.CfgPath, &configs)
	if err != nil {
		return nil, err
	}

	return configs.ToTask()
}

func (f *FileConfig) GetTicker() *time.Ticker {
	return f.ticker
}
