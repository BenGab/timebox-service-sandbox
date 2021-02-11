package logic

type TimestampService struct {
	data *string
}

func (ts *TimestampService) SetTimestamp(time string) {
	ts.data = &time
}

func (ts *TimestampService) GetTimeStamp() *string {
	return ts.data
}

func NewtimestampService() TimestampService {
	return TimestampService{data: nil}
}
