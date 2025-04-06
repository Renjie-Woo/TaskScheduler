package model

type config struct {
	base
	Rules []rule `json:"rules"`
}

func (t *config) toTask() (map[string]singleTask, error) {
	var tasks = make(map[string]singleTask)
	for _, r := range t.Rules {
		if r.Params == nil {
			r.Params = t.Params
		}
		tsk := singleTask{
			base: t.base,
			Rule: r,
		}
		uuid, err := tsk.GetUUID()
		if err != nil {
			return nil, err
		}
		tasks[uuid] = tsk
	}

	return tasks, nil
}

type configList []config

func (t *configList) ToTask() (map[string]singleTask, error) {
	var tasks = make(map[string]singleTask)

	for _, c := range *t {
		tsks, err := c.toTask()
		if err != nil {
			return nil, err
		}
		for uuid, tsk := range tsks {
			tasks[uuid] = tsk
		}
	}

	return tasks, nil
}
