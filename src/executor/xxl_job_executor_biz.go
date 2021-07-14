package executor

type ClientExecutorBiz interface {
}

type ReturnT struct {
	code    int32
	msg     string
	content interface{}
}

type IdleBeatParam struct {
	jobId int32
}

type TriggerParam struct {
	jobId                 int32
	executorHandler       string
	executorBlockStrategy string
	executorTimeout       int64

	logId        int64
	longDateTime int64

	glueType       string
	glueSource     string
	glueUpdateTime string

	broadcastIndex int32
	broadcastTotal int32
}

type KillParam struct {
	jobId int32
}

type LogParam struct {
	logId       int64
	logDateTime int64
	fromLineNum int32
}

type ClientExecutorBizImpl struct {

	executor XxlJobSimpleExecutor
}

func (receiver *ClientExecutorBizImpl) beat() ReturnT {
	return ReturnT{
		code:    200,
		content: nil,
	}
}

func (receiver *ClientExecutorBizImpl) idleBeat(idleBeatParam IdleBeatParam) ReturnT {
	if receiver.executor.isJobRunning(idleBeatParam.jobId) {
		return ReturnT{
			code: 500,
			content: "job goroutine is running or has trigger queue.",
		}
	} else {
		return ReturnT{
			code: 200,
			content: nil,
		}
	}
}

func (receiver *ClientExecutorBizImpl) run(triggerParam TriggerParam) ReturnT {
	receiver.executor.runJob(triggerParam)
}

func (receiver *ClientExecutorBizImpl) kill(killParam KillParam) ReturnT {
	receiver.executor.removeJob(killParam.jobId)
	return ReturnT{
		code:    200,
		content: nil,
	}
}

func (receiver *ClientExecutorBizImpl) log(logParam LogParam) ReturnT {
	// TODO: just ignore
	return ReturnT{
		code: 200,
		content: nil,
	}
}
