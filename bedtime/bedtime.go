package bedtime

import (
	"fmt"
	"time"
)

func CalculateBedtimeNow(timezone string) (times []string, err error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= 7; i++ {
		times = append(times, time.Now().In(loc).Add((time.Duration(90*i)*time.Minute)+15*time.Minute).Format("15:04"))
	}

	return times, nil
}

func CalculateBedTime(bedtime string, timezone string) (times []string, err error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	parsedBedtime, err := time.ParseInLocation("15:04", bedtime, loc)
	fmt.Println(parsedBedtime.Format("15:04"))
	if err != nil {
		return nil, err
	}
	for i := 1; i <= 7; i++ {
		times = append(times, parsedBedtime.Add(-((time.Duration(90*i) * time.Minute) + 15*time.Minute)).Format("15:04"))
	}

	return times, nil
}
