package gotech

type Board struct {
	Projets	[]Project
	Modules	[]Module
	Infos	Infos
	Current	Current
}

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

type Schedule struct {
	ActiveLog	string `json:"active_log"`
	CreditsMin	string `json:"credits_min"`
	CreditsNorm	string `json:"credits_norm"`
	CreditsObj	string `json:"credits_obj"`
	Achieved	int
	Failed		int
	Inprogress	int
}

type Event struct {
	Scolaryear	 string
	Codemodule	 string
	Codeinstance	 string
	Codeacti	 string
	Codeevent	 string
	Semester	 int
	InstanceLocation string `json:"instance_location"`
	Titlemodule	 string `json:"titlemodule"`
	ActiTitle	 string `json:"acti_title"`
	Start		 string
	End		 string
	TotalStudentsRegistered string `json:"total_students_registered"`
	TypeTitle	 string `json:"type_title"`
	NbHours		 string `json:"nb_hours"`
	RegisterStudent	 bool `json:"register_student"`
}
