package logic

import "time"

type timestampService struct {
	data *time.Time
}

func (ts *timestampService) SetTimestamp(time time.Time) {
	t := time
	ts.data = &t
}

func (ts *timestampService) GetTimeStamp() *time.Time {
	return ts.data
}

func NewtimestampService() timestampService {
	return timestampService{data: nil}
}
