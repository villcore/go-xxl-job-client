package executor

import "villcore.com/common/api"

func init() {
	// 1. register
}

type JobExecutor interface {
}

type XxlJobSimpleExecutor struct {
	// TODO state mark start

	// TODO admin biz api client

	// TODO log clean ignore

	// TODO trigger callback goroutine

	// TODO trigger retry callback

	// TODO inner http server

}

func New(config XxlJobConfig) *XxlJobSimpleExecutor {
	return &XxlJobSimpleExecutor{}
}

func (e *XxlJobSimpleExecutor) AddJobHandler(name string, handler func(param string), init func(), destroy func()) {

}

func (e *XxlJobSimpleExecutor) Start() {

}

func (e *XxlJobSimpleExecutor) Destroy() {

}

func (e *XxlJobSimpleExecutor) isJobRunning(jobId int32) bool {

}

func (e *XxlJobSimpleExecutor) getJob(jobId int32) interface{} {

}

func (e *XxlJobSimpleExecutor) removeJob(id int32) {

}

func (e *XxlJobSimpleExecutor) runJob(param api.TriggerParam) bool {

}
