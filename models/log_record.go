package models

import (
	"fmt"
	"net/http"
	"time"
)

type Log struct {
	Level     string  `json:"level"`
	Timestamp float64 `json:"ts"`
	Message   string  `json:"msg"`
	Logger    string  `json:"logger"`
	Request   Request `json:"request"`
	CommonLog string  `json:"common_log"`
	Status    int     `json:"status"`
}

func (l *Log) GetTime() time.Time {
	return time.Unix(int64(l.Timestamp), 0)
}

func (l *Log) StatusText() string {
	return http.StatusText(l.Status)
}

func (l *Log) StatusCodeText() string {
	return fmt.Sprintf("%d %s", l.Status, l.StatusText())
}
