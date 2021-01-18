package goft

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once
var onceCron sync.Once
var taskCron *cron.Cron
func init() {
	chlist := getTaskList() //get process list
	go func() {
		for t := range chlist {
			//run process
			doTask(t)
		}
	}()
}
func getCronTask()  *cron.Cron{
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				//回调
				t.callback()
			}
		}()
		t.Exec()
	}()
}
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskExecutor struct {
	f        TaskFunc
	p        []interface{}
	callback func()
}

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (this *TaskExecutor) Exec() {
	this.f(this.p...)
}

func Task(f TaskFunc, callback func(), params ...interface{}) {
	if f == nil {
		return
	}
	//放进channel
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, callback) //增加任务队列
	}()
}
