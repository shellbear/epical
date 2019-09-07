package epical

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EpitechEvent struct {
	Scolaryear       string `json:"scolaryear"`
	Codemodule       string `json:"codemodule"`
	Codeinstance     string `json:"codeinstance"`
	Codeacti         string `json:"codeacti"`
	Codeevent        string `json:"codeevent"`
	Semester         int    `json:"semester"`
	InstanceLocation string `json:"instance_location"`
	Titlemodule      string `json:"titlemodule"`
	ProfInst         []struct {
		Type    string `json:"type"`
		Login   string `json:"login"`
		Title   string `json:"title"`
		Picture string `json:"picture"`
	} `json:"prof_inst"`
	ActiTitle               string      `json:"acti_title"`
	NumEvent                int         `json:"num_event"`
	Start                   string      `json:"start"`
	End                     string      `json:"end"`
	TotalStudentsRegistered int         `json:"total_students_registered"`
	Title                   interface{} `json:"title"`
	TypeTitle               string      `json:"type_title"`
	TypeCode                string      `json:"type_code"`
	IsRdv                   string      `json:"is_rdv"`
	NbHours                 string      `json:"nb_hours"`
	AllowedPlanningStart    string      `json:"allowed_planning_start"`
	AllowedPlanningEnd      string      `json:"allowed_planning_end"`
	NbGroup                 int         `json:"nb_group"`
	NbMaxStudentsProjet     interface{} `json:"nb_max_students_projet"`
	Room                    struct {
		Code  string `json:"code"`
		Type  string `json:"type"`
		Seats int    `json:"seats"`
	} `json:"room"`
	Dates              interface{} `json:"dates"`
	ModuleAvailable    bool        `json:"module_available"`
	ModuleRegistered   bool        `json:"module_registered"`
	Past               bool        `json:"past"`
	AllowRegister      bool        `json:"allow_register"`
	EventRegistered    interface{} `json:"event_registered"`
	Display            string      `json:"display"`
	Project            bool        `json:"project"`
	RdvGroupRegistered interface{} `json:"rdv_group_registered"`
	RdvIndivRegistered interface{} `json:"rdv_indiv_registered"`
	AllowToken         bool        `json:"allow_token"`
	RegisterStudent    bool        `json:"register_student"`
	RegisterProf       bool        `json:"register_prof"`
	RegisterMonth      bool        `json:"register_month"`
	InMoreThanOneMonth bool        `json:"in_more_than_one_month"`
}

const (
	EPITECH_API_QUERY_TIME_FORMAT = "2006-01-02"
	EPITECH_EVENT_TIME_FORMAT     = "2006-01-02 15:04:05"
	EPITECH_BASE_URL              = "https://intra.epitech.eu"
)

// GetCalendar aims to fetch Epitech calendar data
func GetRegisteredEvents(token string) ([]EpitechEvent, error) {
	start := time.Now().Format(EPITECH_API_QUERY_TIME_FORMAT)
	resp, err := http.Get(fmt.Sprintf("%s/auth-%s/planning/load?format=json&start=%s", EPITECH_BASE_URL, token, start))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request to %s failed (%s)", resp.Request.Host, resp.Status)
	}

	var data []EpitechEvent
	var registeredEvents []EpitechEvent

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for _, c := range data {
		status, valid := c.EventRegistered.(bool)
		if !valid || status {
			registeredEvents = append(registeredEvents, c)
		}
	}

	return registeredEvents, nil
}
