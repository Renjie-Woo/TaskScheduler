package model

import (
	"github.com/Renjie-Woo/TaskScheduler/pkg/tool"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"path/filepath"
)

type taskScheduler struct {
	scheduler  *cron.Cron            // 调度器
	configCtrl configController      // 配置控制器
	tasks      map[string]singleTask // 当前执行中的任务列表
	logger     *zap.SugaredLogger
	workDir    string
}

func NewTaskScheduler(workDir string, configCtrl configController, logger *zap.SugaredLogger) *taskScheduler {
	return &taskScheduler{
		scheduler:  cron.New(cron.WithSeconds()),
		configCtrl: configCtrl,
		tasks:      make(map[string]singleTask),
		logger:     logger,
		workDir:    workDir,
	}
}

func (t *taskScheduler) removeTaskByID(tsk singleTask) {
	t.scheduler.Remove(tsk.id)
	delete(t.tasks, tsk.uuid)

	t.logger.Infof("[DELET TASK] remove task %s[%s] success", tsk.Name, tsk.Rule)
}

func (t *taskScheduler) RemoveTask(tsk singleTask) {
	if old, ok := t.tasks[tsk.uuid]; ok {
		t.removeTaskByID(old)
	}
}

func (t *taskScheduler) addTaskDirectly(tsk singleTask) {
	id, err := t.scheduler.AddFunc(tsk.Rule.Rule, func() {
		t.logger.Infof("[EXEC] begin to execute task %s[%s]", tsk.Name, tsk.Rule)
		var result, err = tool.RunPythonScript(filepath.Join(t.workDir, tsk.ScriptPath), tsk.Rule.Params)
		if err != nil {
			t.logger.Errorf("[EXEC FATAL] add task %s[%s] with err %v", tsk.Name, tsk.Rule, err)
			return
		}
		t.logger.Infof("[EXEC SUCCESS] execute task %s[%s] with result %s", tsk.Name, tsk.Rule, string(result))
	})
	if err != nil {
		t.logger.Errorf("[ADD TASK] add task %s[%s] with error %v", tsk.Name, tsk.Rule, err)
		return
	}
	tsk.SetTaskID(id)

	t.tasks[tsk.uuid] = tsk

	t.logger.Infof("[ADD TASK] add task %s[%s] success", tsk.Name, tsk.Rule)
}

func (t *taskScheduler) AddTask(tsk singleTask) {
	t.RemoveTask(tsk)

	t.addTaskDirectly(tsk)
}

func (t *taskScheduler) Run() {
	t.logger.Info(">>> start scheduler")

	go t.RefreshConfig()

	t.scheduler.Run()
}

func (t *taskScheduler) updateTasks(new map[string]singleTask) {
	var old = t.tasks
	var delCnt, addCnt = 0, 0
	for uuid, tsk := range old {
		if _, ok := new[uuid]; !ok {
			t.removeTaskByID(tsk)
			delCnt += 1
		}
	}

	for uuid, tsk := range new {
		if _, ok := old[uuid]; !ok {
			t.addTaskDirectly(tsk)
			addCnt += 1
		}
	}

	t.logger.Infof("[UPDATE] total %d tasks changed that %d been deleted and %v been added", delCnt+addCnt, delCnt, addCnt)
}

func (t *taskScheduler) refreshConfig() {
	t.logger.Info("[REFRESH] refresh configuration")
	var tasks, err = t.configCtrl.Refresh()
	if err != nil {
		t.logger.Errorf("[REFRESH FATAL] refresh with error %v", err)
	} else {
		t.logger.Info("[REFRESH SUCCESS] refresh success, begin to update tasks")
		t.updateTasks(tasks)
	}
}

func (t *taskScheduler) RefreshConfig() {
	t.logger.Info("[REFRESH] start refresh process")
	t.refreshConfig()

	for range t.configCtrl.GetTicker().C {
		t.refreshConfig()
	}
}
