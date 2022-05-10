package dao

import "time"

type News struct {
	Category string `db:"类别"`
	Title    string `db:"标题"`
	Time     time.Time `db:"采集时间"`
	Url      string `db:"网址"`
}
var(
	Category  = make(map[string]string)
)