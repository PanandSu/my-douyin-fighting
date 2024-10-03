package initial

import (
	"github.com/sony/sonyflake"
	gb "my-douyin-fighting/glob"
	"time"
)

func Global() {
	startTime, err := time.Parse(time.DateTime, gb.StartTime)
	if err != nil {
		return
	}
	sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: startTime,
	})
}
