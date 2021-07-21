package api

type AdminBiz interface {
	Beat() ReturnT
	IdleBeat(param *IdleBeatParam) ReturnT
	Run(param *TriggerParam) ReturnT
	Kill(param *KillParam) ReturnT
	Log(param *LogParam) ReturnT
}
