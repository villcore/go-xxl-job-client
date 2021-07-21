package executor

import "villcore.com/common/api"

type ClientExecutorBizImpl struct {
	executor XxlJobSimpleExecutor
}

func (receiver *ClientExecutorBizImpl) Beat() api.ReturnT {
	return api.ReturnT{
		Code:    200,
		Content: nil,
	}
}

func (receiver *ClientExecutorBizImpl) IdleBeat(idleBeatParam api.IdleBeatParam) api.ReturnT {
	if receiver.executor.isJobRunning(idleBeatParam.JobId) {
		return api.ReturnT{
			Code:    500,
			Content: "job goroutine is running or has trigger queue.",
		}
	} else {
		return api.ReturnT{
			Code:    200,
			Content: nil,
		}

		// 03
		// 025243 => 567089
	}
}

func (receiver *ClientExecutorBizImpl) Run(triggerParam api.TriggerParam) api.ReturnT {
	if receiver.executor.runJob(triggerParam) {
		return api.ReturnT{
			Code:    200,
			Content: nil,
		}
	} else {
		return api.ReturnT{
			Code:    500,
			Content: "job goroutine invoke failed",
		}
	}
}

func (receiver *ClientExecutorBizImpl) Kill(killParam api.KillParam) api.ReturnT {
	receiver.executor.removeJob(killParam.JobId)
	return api.ReturnT{
		Code:    200,
		Content: nil,
	}
}

func (receiver *ClientExecutorBizImpl) Log(logParam api.LogParam) api.ReturnT {
	// TODO: just ignore
	return api.ReturnT{
		Code:    200,
		Content: nil,
	}
}

type AdminBizClientImpl struct {
}
