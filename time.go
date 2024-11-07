// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"time"
)

var (
	locationGMT *time.Location = gmtLocation()

	timeFormatDateTimeMicro string = "2006-01-02 15:04:05.000000"
)

func gmtLocation() *time.Location {
	location, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}

	return location
}

func encodeRFC1123Time(t time.Time) string {
	return t.In(locationGMT).Format(time.RFC1123)
}

func decodeTime(timeStr string) (time.Time, error) {
	parsed, err := time.ParseInLocation(time.RFC1123, timeStr, locationGMT)
	if err == nil {
		return parsed, nil
	}

	parsed, err = time.ParseInLocation(timeFormatDateTimeMicro, timeStr, locationGMT)
	if err == nil {
		return parsed, nil
	}

	return time.Time{}, err
}
