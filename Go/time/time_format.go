package main

import (
	"fmt"
	"time"
)

var YearDateTimeFormat = "20060102150405"

func main() {
	fmt.Println(CompareExperimentID("20200108184533"))
	fmt.Println(parseTimeDuration("2019-12-25T08:36:28+08:00", "2019-12-25T12:36:28+08:00").Seconds())
}

func parseTimeDuration(start, end string) time.Duration {
	stamp, _ := time.ParseInLocation(time.RFC3339, start, time.Local)
	estamp, _ := time.ParseInLocation(time.RFC3339, end, time.Local)
	return estamp.Sub(stamp)
}

func CompareExperimentID(eid string) bool {
	eTime, _ := time.ParseInLocation(YearDateTimeFormat, eid, time.Local)
	return eTime.Before(time.Now().Add(-time.Hour))
}
