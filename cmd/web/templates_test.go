package main

import (
	"testing"
	"time"

	"snippetbox.marouane.com/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 9, 2, 1, 56, 0, 0, time.UTC),
			want: "02 Sep 2024 at 01:56",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		}, {
			name: "CET",
			tm:   time.Date(2024, 9, 2, 1, 56, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "02 Sep 2024 at 00:56",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hd := humanDate(test.tm)
			assert.Equal(t, hd, test.want)
		})
	}

}
