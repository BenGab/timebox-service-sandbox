package logic

import "time"

type TimestampService struct {
	data *time.Time
}

func (ts *TimestampService) SetTimestamp(time time.Time) {
	t := time
	ts.data = &t
}

func (ts *TimestampService) GetTimeStamp() *time.Time {
	return ts.data
}

func NewtimestampService() TimestampService {
	return TimestampService{data: nil}
}
