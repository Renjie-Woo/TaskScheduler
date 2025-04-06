package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
)

type singleTask struct {
	base
	id   cron.EntryID
	Rule rule
	uuid string
}

func (t *singleTask) GetUUID() (string, error) {
	if t.uuid == "" {
		var task, err = json.Marshal(t)
		if err != nil {
			return "", err
		}
		t.uuid = fmt.Sprintf("%x", md5.Sum(task))
	}

	return t.uuid, nil
}

func (t *singleTask) SetTaskID(id cron.EntryID) {
	t.id = id
}

func (t *singleTask) GetTaskID() cron.EntryID {
	return t.id
}
