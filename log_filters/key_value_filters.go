package logfilters

import "github.com/AliAlhajji/caddylogs/models"

//KeyValueFilter is a function that filters based on key-value pairs
type KeyValueFilter func(log *models.Log, key, value string) bool

func RequestHeaderIs(log *models.Log, headerKey, headerValue string) bool {
	if val, ok := log.Request.Headers[headerKey]; ok {
		return val[0] == headerValue
	}

	return false
}
