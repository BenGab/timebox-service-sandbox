package logic_test

import (
	"testing"
	"time"

	"github.com/bengab/timebox-service/src/logic"
)

func TestSetTimeStamp_NotNil(t *testing.T) {
	sut := logic.NewtimestampService()
	sut.SetTimestamp(time.Now())

	if sut.GetTimeStamp() == nil {
		t.Error("Timestamp is still nil")
	}
}

func TestGetSetTimeStamp_Expected(t *testing.T) {
	sut := logic.NewtimestampService()
	time := time.Now()
	sut.SetTimestamp(time)

	if *sut.GetTimeStamp() != time {
		t.Error("GettimeStamp()")
	}
}

func GetTimeStamp_Nil(t *testing.T) {
	sut := logic.NewtimestampService()

	if sut.GetTimeStamp() != nil {
		t.Error("timestamp not nil")
	}
}
