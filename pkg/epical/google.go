package epical

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const (
	GOOGLE_CALENDAR_API_TIME_FORMAT = "2006-01-02T15:04:05"
)

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	if err != nil {
		log.Fatal(err)
	}

	return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetGoogleCalendarByName(service *calendar.Service, name string) (*calendar.Calendar, error) {
	var pageToken string
	var calendarList *calendar.CalendarList
	var err error

	for {
		if pageToken == "" {
			calendarList, err = service.CalendarList.List().Do()
		} else {
			calendarList, err = service.CalendarList.List().PageToken(pageToken).Do()
		}

		if err != nil {
			return nil, err
		}

		for _, cal := range calendarList.Items {
			if cal.Summary == name {
				googleCal, err := service.Calendars.Get(cal.Id).Do()
				if err != nil {
					return nil, err
				}

				return googleCal, nil
			}
		}

		pageToken = calendarList.NextPageToken
		if pageToken == "" {
			break
		}
	}

	return nil, nil
}

func GetOrCreateGoogleCalendar(service *calendar.Service, name string) (*calendar.Calendar, error) {
	cal, err := GetGoogleCalendarByName(service, name)
	if err != nil {
		return nil, err
	}

	if cal != nil {
		return cal, nil
	}

	cal = &calendar.Calendar{
		Summary:     CALENDAR_NAME,
		Description: "https://github.com/shellbear/epical",
		TimeZone:    "Europe/Paris",
	}

	cal, err = service.Calendars.Insert(cal).Do()
	if err != nil {
		return nil, err
	}

	return cal, nil
}

func NewGoogleCalendarEvent(event *EpitechEvent) (*calendar.Event, error) {
	var start, end time.Time
	var err error

	if err != nil {
		return nil, err
	}

	rdv, valid := event.RdvGroupRegistered.(string)
	if valid {
		parts := strings.Split(rdv, "|")

		start, err = time.Parse(EPITECH_EVENT_TIME_FORMAT, parts[0])
		if err != nil {
			return nil, err
		}

		end, err = time.Parse(EPITECH_EVENT_TIME_FORMAT, parts[1])
		if err != nil {
			return nil, err
		}
	} else {
		start, err = time.Parse(EPITECH_EVENT_TIME_FORMAT, event.Start)
		if err != nil {
			return nil, err
		}

		end, err = time.Parse(EPITECH_EVENT_TIME_FORMAT, event.End)
		if err != nil {
			return nil, err
		}
	}

	newEvent := &calendar.Event{
		Summary: event.ActiTitle,
		Description: fmt.Sprintf("%s/module/%s/%s/%s/%s", EPITECH_BASE_URL,
			event.Scolaryear, event.Codemodule, event.Codeinstance, event.Codeacti),
		Start: &calendar.EventDateTime{
			DateTime: start.Format(GOOGLE_CALENDAR_API_TIME_FORMAT),
			TimeZone: "Europe/Paris",
		},
		End: &calendar.EventDateTime{
			DateTime: end.Format(GOOGLE_CALENDAR_API_TIME_FORMAT),
			TimeZone: "Europe/Paris",
		},
		Location: event.Room.Code,
	}

	return newEvent, nil
}

func GetGoogleCalendarService(credentialsPath string) (*calendar.Service, error) {
	b, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)
	calendarService, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))

	return calendarService, err
}
