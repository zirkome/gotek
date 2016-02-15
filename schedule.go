package gotek

import (
	"encoding/json"
	"io/ioutil"
)

type Event struct {
	Scolaryear              string
	Codemodule              string
	Codeinstance            string
	Codeacti                string
	Codeevent               string
	Semester                int
	InstanceLocation        string `json:"instance_location"`
	Titlemodule             string `json:"titlemodule"`
	ActiTitle               string `json:"acti_title"`
	Start                   string
	End                     string
	TotalStudentsRegistered int    `json:"total_students_registered"`
	TypeTitle               string `json:"type_title"`
	NbHours                 string `json:"nb_hours"`
	RegisterStudent         bool   `json:"register_student"`
}

func (c *Client) GetSchedule() ([]Event, error) {
	var schedule []Event

	response, err := client.Get(SCHEDULE_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
