package time

import (
	"campiagn-slip/config"
	"time"
)

func InBKK() time.Time {
	conf := config.GetConfig()
	now := time.Now().Add(time.Duration(conf.IncreaseTime) * time.Hour)
	return now
}
