package time_package

import "time"

func GetTime() int64 {
	return time.Now().Unix()
}
