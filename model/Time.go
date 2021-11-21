package model

import (
	"strconv"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	// 秒级
	*t = Time(time.Unix(int64(num), 0))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	// 毫秒级
	return ([]byte)(strconv.FormatInt(time.Time(t).UnixNano()/1e6, 10)), nil
}
