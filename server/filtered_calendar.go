package server

import (
	"fmt"
	"github.com/taviti/caldav-go/icalendar"
	"github.com/taviti/caldav-go/icalendar/components"
)

func GetFilteredIcs(calendar string, shows []string) {
	cal := new(components.Calendar)
	icalendar.Unmarshal(calendar, cal)

	fmt.Println(cal.Events[0])
}
