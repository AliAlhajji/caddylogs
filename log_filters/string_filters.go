package logfilters

import (
	"strings"

	"github.com/AliAlhajji/caddylogs/models"
)

//StringFilterFunc is a function that filters the logs based on a string value
type StringFilterFunc func(*models.Log, string) bool

//UrlContains selects the logs whose URL contains a certain string
func UrlContains(log *models.Log, substring string) bool {
	return strings.Contains(log.Request.URL(), substring)
}

//RefererContains selects the logs whose Referer contains a certain string
func RefererContains(log *models.Log, substring string) bool {
	return strings.Contains(log.Request.Referer(), substring)
}

//LoggerIs selects the logs whose logger matches the passed logger name
func LoggerIs(log *models.Log, logger string) bool {
	return log.Logger == logger
}
