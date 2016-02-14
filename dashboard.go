package gotek

import (
	"encoding/json"
	"io/ioutil"
)

type Project struct {
	Title           string
	TitleLink       string
	TimelineStart   string
	TimelineEnd     string
	TimelineBarre   string
	DateInscription string
	IDActivite      string
	SoutenanceName  string
	SoutenanceLink  string
	SoutenanceDate  string
	SoutenanceSalle string
}

type Module struct {
	Title           string
	TitleLink       string `json:"title_link"`
	TimelineStart   string `json:"timeline_start"`
	TimelineEnd     string `json:"timeline_end"`
	TimelineBarre   string `json:"timeline_barre"`
	DateInscription string `json:"date_inscription"`
}

type Infos struct {
	ID            string
	Login         string
	Title         string
	Email         string
	InternalEmail string `json:"internal_email"`
	Lastname      string
	Firstname     string
	Picture       string
	Promo         int
	Semester      int
	UID           int
	GID           int
	Location      string
}

type Current struct {
	ActiveLog   string `json:"active_log"`
	CreditsMin  string `json:"credits_min"`
	CreditsNorm string `json:"credits_norm"`
	CreditsObj  string `json:"credits_obj"`
}

type History struct {
	Title      string
	Content    string
	Date       string
	ID         string
	Visible    string
	IDActivite string `json:"id_activite"`
	Class      string
}

type Board struct {
	Projets []Project
	Modules []Module
}

type Dashboard struct {
	IP      string
	Board   Board
	History []History
	Infos   Infos
	Current []Current
}

func (c *Client) GetDashboard() (*Dashboard, error) {
	var dashboard Dashboard

	response, err := client.Get(DASHBOARD_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &dashboard)
	if err != nil {
		return nil, err
	}
	return &dashboard, nil
}
