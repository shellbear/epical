package epical

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

const (
	CalendarName = "EpiCal"
	Version      = "0.1.4"
)

func ListEvents(epitechToken string) {
	data, err := GetRegisteredEvents(epitechToken)
	if err != nil {
		log.Fatalln(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	if _, err := fmt.Fprintln(writer, "NAME\tSTART\tEND\tROOM"); err != nil {
		log.Fatalln(err)
	}

	if len(data) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, evt := range data {
			rdv, valid := evt.RdvGroupRegistered.(string)
			parts := []string{evt.Start, evt.End}

			if valid {
				parts = strings.Split(rdv, "|")
			}

			if _, err := fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n",
				evt.ActiTitle, parts[0], parts[1], evt.Room.Code); err != nil {
				log.Fatalln(err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		log.Fatalln(err)
	}
}

func ClearEvents(credentialsPath string, deleteFrom time.Time, deleteCalendar bool) {
	svc, err := GetGoogleCalendarService(credentialsPath)
	if err != nil {
		log.Fatalln("Failed to get Google calendar service,", err)
	}

	cal, err := GetGoogleCalendarByName(svc, CalendarName)
	if err != nil {
		log.Fatalln("Failed to get Google calendar ", err)
	}

	if cal != nil {
		events, err := svc.Events.List(cal.Id).Do()
		if err != nil {
			log.Fatalln("Failed to list calendar events", err)
		}

		for _, evt := range events.Items {
			t, err := time.Parse(time.RFC3339, evt.Start.DateTime)
			if err != nil {
				log.Fatalln("Failed to parse calendar event datetime,", err)
			}

			if t.After(deleteFrom) {
				err = svc.Events.Delete(cal.Id, evt.Id).Do()
				if err != nil {
					log.Fatalln("Failed to delete calendar event,", err)
				}
			}
		}

		if deleteCalendar {
			if err = svc.Calendars.Delete(cal.Id).Do(); err != nil {
				log.Fatalln("Failed to delete Google calendar,", err)
			}
		}
	}
}

func SyncCalendar(credentialsPath, token string) {
	data, err := GetRegisteredEvents(token)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Found %d events to synchronize.\n", len(data))
	svc, err := GetGoogleCalendarService(credentialsPath)
	if err != nil {
		log.Fatalln("Failed to get Google calendar service,", err)
	}

	t := time.Now().Truncate(time.Hour * 24)

	ClearEvents(credentialsPath, t, false)
	fmt.Println("Cleared old calendar events.")

	cal, err := GetOrCreateGoogleCalendar(svc, CalendarName)
	if err != nil {
		log.Fatalln("Failed to get Google calendar,", err)
	}

	if len(data) == 0 {
		fmt.Println("There is no upcoming Epitech event.")
	} else {
		for _, c := range data {
			newEvt, err := NewGoogleCalendarEvent(&c)
			if err != nil {
				log.Fatalln("Failed to create Google calendar event,", err)
			}

			evt, err := svc.Events.Insert(cal.Id, newEvt).Do()
			if err != nil {
				log.Fatalln("Failed to create Google calendar event", err)
			}

			log.Println("Created event", evt.Summary)
		}
	}
}
