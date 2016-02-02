package gotek

import (
	"io/ioutil"
	"encoding/json"
)

type Project struct {
	Title		string
	TitleLink	string
	TimelineStart	string
	TimelineEnd	string
	TimelineBarre	string
	DateInscription	string
	IdActivite	string
	SoutenanceName	string
	SoutenanceLink	string
	SoutenanceDate	string
	SoutenanceSalle	string
}

type Module struct {
	title			string
	title_link		string
	timeline_start		string
	timeline_end		string
	timeline_barre		string
	date_inscription	string
}

type Infos struct {
	Id		string
	Login		string
	Title		string
	Email		string
	InternalEmail	string `json:"internal_email"`
	Lastname	string
	Firstname	string
	Picture		string
	Promo		int
	Semester	int
	Uid		int
	Gid		int
	Location	string
}

type Current struct {
	ActiveLog	string `json:"active_log"`
	CreditsMin	string `json:"credits_min"`
	CreditsNorm	string `json:"credits_norm"`
	CreditsObj	string `json:"credits_obj"`
	Achieved	int
	Failed		int
	Inprogress	int
}

type Board struct {
	Projets	[]Project
	Modules	[]Module
	Infos	Infos
	Current	Current
}

func (c *Client) GetDashboard() (*Board, error) {
	var dashboard Board

	response, err := client.Get(DASHBOARD_ENDPOINT)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(contents, &dashboard)
		if err != nil {
			return nil, err
		}
	}
	return &dashboard, nil
}
