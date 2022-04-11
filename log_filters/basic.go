package logfilters

import "github.com/AliAlhajji/caddylogs/models"

type FilterFunc func(*models.Log) bool

//InfoLogs returns only the logs with level "info"
func InfoLogs(log *models.Log) bool {
	return log.Level == models.LogLevel_Info
}

var Info = func(log *models.Log) bool {
	return true
}

//ErrorLogs returns only the logs with level "error"
func ErrorLogs(log *models.Log) bool {
	return log.Level == models.LogLevel_Error
}
