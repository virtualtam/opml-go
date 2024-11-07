// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	cases := []struct {
		tname   string
		dateStr string
		want    time.Time
	}{
		{
			tname:   "RFC 1123",
			dateStr: "Mon, 27 Feb 2006 12:09:48 GMT",
			want:    mustDecodeRFC1123Time("Mon, 27 Feb 2006 12:09:48 GMT"),
		},
		{
			tname:   "Newsblur",
			dateStr: "2024-11-07 20:18:01.109756",
			want:    mustDecodeRFC1123Time("Thu, 07 Nov 2024 20:18:01.109756 GMT"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			got, err := decodeTime(tc.dateStr)
			if err != nil {
				t.Fatalf("failed to parse date: %q", err)
			}

			if !got.Equal(tc.want) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func mustDecodeRFC1123Time(dateStr string) time.Time {
	parsed, err := time.ParseInLocation(time.RFC1123, dateStr, locationGMT)
	if err != nil {
		panic(err)
	}

	return parsed
}
