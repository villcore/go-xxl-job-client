package api

type ReturnT struct {
	Code    int32
	Msg     string
	Content interface{}
}

type IdleBeatParam struct {
	JobId int32
}

type TriggerParam struct {
	JobId                 int32
	ExecutorHandler       string
	ExecutorBlockStrategy string
	ExecutorTimeout       int64

	LogId        int64
	LongDateTime int64

	GlueType       string
	GlueSource     string
	GlueUpdateTime string

	BroadcastIndex int32
	BroadcastTotal int32
}

type KillParam struct {
	JobId int32
}

type LogParam struct {
	LogId       int64
	LogDateTime int64
	FromLineNum int32
}
