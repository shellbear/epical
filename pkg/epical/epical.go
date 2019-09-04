package epical

import (
	"fmt"
	"log"
)

const (
	CALENDAR_NAME = "EpiCal"
	VERSION       = "0.1"
)

func ListEvents(epitechToken string) {
	data, err := GetRegisteredEvents(epitechToken)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, evt := range data {
			fmt.Printf("%v (%v-%v)\n", evt.ActiTitle, evt.Start, evt.End)
		}
	}
}

func ClearEvents(credentialsPath string) {
	svc, err := GetGoogleCalendarService(credentialsPath)
	if err != nil {
		log.Fatal(err)
	}

	cal, err := GetGoogleCalendarByName(svc, CALENDAR_NAME)
	if err != nil {
		log.Fatal(err)
	}

	if cal != nil {
		events, err := svc.Events.List(cal.Id).Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, evt := range events.Items {
			err = svc.Events.Delete(cal.Id, evt.Id).Do()
			if err != nil {
				log.Fatal(err)
			}
		}

		err = svc.Calendars.Delete(cal.Id).Do()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SyncCalendar(credentialsPath, token string) {
	data, err := GetRegisteredEvents(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d events to synchronize.\n", len(data))

	svc, err := GetGoogleCalendarService(credentialsPath)
	if err != nil {
		log.Fatal(err)
	}

	ClearEvents(credentialsPath)
	fmt.Println("Cleared old calendar events.")

	cal, err := GetOrCreateGoogleCalendar(svc, CALENDAR_NAME)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		fmt.Println("There is no upcoming Epitech event.")
	} else {
		for _, c := range data {
			newEvt, err := NewGoogleCalendarEvent(&c)
			if err != nil {
				log.Fatal(err)
			}

			evt, err := svc.Events.Insert(cal.Id, newEvt).Do()
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Created event", evt.Summary)
		}
	}
}
