package logfilters

import "github.com/AliAlhajji/caddylogs/models"

type IntFilters func(*models.Log, int) bool

func StatusCodeIs(log *models.Log, statusCode int) bool {
	return log.Status == statusCode
}
