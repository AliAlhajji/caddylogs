package caddylogs

import (
	"bufio"
	"encoding/json"
	"os"

	logfilters "github.com/AliAlhajji/caddylogs/log_filters"
	"github.com/AliAlhajji/caddylogs/models"
)

type AccessLogs struct {
	logs []*models.Log
}

//New loads the logs from the file logsPath. The logs must be in Caddy's access logs JSON format
func New(logsPath string) (*AccessLogs, error) {
	repo := &AccessLogs{}

	err := repo.loadLogs(logsPath)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

//GetLogs returns the current logs. You can use this method after applying the filters to get the filtered logs
func (r *AccessLogs) GetLogs() []*models.Log {
	return r.logs
}

//Filter filters the logs based on the selected filter
func (r *AccessLogs) Filter(filter logfilters.FilterFunc) *AccessLogs {
	var filterLog []*models.Log
	for _, log := range r.logs {
		if filter(log) {
			filterLog = append(filterLog, log)
		}
	}

	r.logs = filterLog

	return r
}

//StringFilter filters the logs based on the selected filter and the str that is passed to the filter
func (r *AccessLogs) StringFilter(filter logfilters.StringFilterFunc, str string) *AccessLogs {
	var filterLog []*models.Log
	for _, log := range r.logs {
		if filter(log, str) {
			filterLog = append(filterLog, log)
		}
	}

	r.logs = filterLog

	return r
}

//IntFilter filters the logs based on the selected filter and the int value passed to the filter
func (r *AccessLogs) IntFilter(filter logfilters.IntFilters, n int) *AccessLogs {
	var filterLog []*models.Log
	for _, log := range r.logs {
		if filter(log, n) {
			filterLog = append(filterLog, log)
		}
	}

	r.logs = filterLog

	return r
}

func (r *AccessLogs) KeyValueFilter(filter logfilters.KeyValueFilter, key, value string) *AccessLogs {
	var filterLog []*models.Log
	for _, log := range r.logs {
		if filter(log, key, value) {
			filterLog = append(filterLog, log)
		}
	}

	r.logs = filterLog

	return r
}

//First returns the first n records of the current logs
func (r *AccessLogs) First(n int) *AccessLogs {
	if n > len(r.logs) {
		n = len(r.logs)
	}

	r.logs = r.logs[0:n]
	return r
}

//Last returns the last n records of the current logs
func (r *AccessLogs) Last(n int) *AccessLogs {
	if n > len(r.logs) {
		n = len(r.logs)
	}

	r.logs = r.logs[len(r.logs)-n:]

	return r
}

//Reverse reverses the logs array
func (r *AccessLogs) Reverse() *AccessLogs {
	for i, j := 0, len(r.logs)-1; i < j; i, j = i+1, j-1 {
		r.logs[i], r.logs[j] = r.logs[j], r.logs[i]
	}

	return r
}

func (r *AccessLogs) loadLogs(logsPath string) error {
	file, err := os.Open(logsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var logs []*models.Log

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		logText := scanner.Text()

		log, err := marshallLog([]byte(logText))
		if err != nil {
			return err
		}

		logs = append(logs, log)
	}

	r.logs = logs
	return nil
}

func marshallLog(rawLog json.RawMessage) (*models.Log, error) {
	var log *models.Log

	err := json.Unmarshal(rawLog, &log)
	if err != nil {
		return nil, err
	}

	return log, nil
}
