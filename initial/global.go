package initial

import (
	"github.com/sony/sonyflake"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/util"
	"time"
)

func Global() {
	startTime, _ := time.Parse(time.DateTime, gb.StartTime)
	gb.IdGenerator = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: startTime,
	})
	util.CheckPathAndCreate(gb.VideoAddr)
	util.CheckPathAndCreate(gb.CoverAddr)
	m := map[string]string{
		"0000002066747970":                 ".mp4",
		"0000001c66747970":                 ".mp4",
		"0000001866747970":                 ".mp4",
		"52494646":                         ".avi",
		"3026b2758e66cf11a6d9":             ".wmv",
		"000001BA47000001B3":               ".mpeg",
		"6D6F6F76":                         ".mov",
		"464c5601050000000900":             ".flv",
		"2e524d46000000120001":             ".rmvb",
		"667479703367":                     ".3gb",
		"000001BA":                         ".vob",
		"00000020667479704D34412000000000": ".m4v",
	}
	for k, v := range m {
		gb.FileTypeMap.Store(k, v)
	}
}
