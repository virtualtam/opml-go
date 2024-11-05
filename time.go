// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"time"
)

var (
	locationGMT *time.Location = gmtLocation()
)

func gmtLocation() *time.Location {
	location, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}

	return location
}

func formatRFC1123Time(t time.Time) string {
	return t.In(locationGMT).Format(time.RFC1123)
}

func parseRFC1123Time(timeStr string) (time.Time, error) {
	parsed, err := time.ParseInLocation(time.RFC1123, timeStr, locationGMT)
	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}
