package server

import (
	"github.com/taviti/caldav-go/icalendar"
	"github.com/taviti/caldav-go/icalendar/components"
	"strings"
)

func isInList(summary string, shows []string) bool {
	for _, show := range shows {
		if strings.HasPrefix(summary, show) {
			return true
		}
	}
	return false
}

func GetFilteredIcs(calendar string, shows []string) (string, error) {
	cal := new(components.Calendar)
	icalendar.Unmarshal(calendar, cal)
	events := []*components.Event{}

	for _, event := range cal.Events {
		summary := event.Summary
		if isInList(summary, shows) {
			events = append(events, event)
		}
	}
	cal.Events = events
	ret, err := icalendar.Marshal(cal)
	if err != nil {
		return "", err
	}
	return ret, nil
}
