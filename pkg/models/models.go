package models

// --------------------------------
// models for postBase

type PostDTO struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Theme string `json:"theme"`
	Part  string `json:"part"`
}

type ResponseDTO struct {
	Status string `json:status`
}

// --------------------------------
// models for activities

// activ_names
type ActivityDTO struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
	NormId    int    `json:"norm_id"`
}

// activ_normative
type ActivityNormativeDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	NormPeriod  string `json:"norm_period"`
	NormMeasure string `json:"norm_measure"`
	NormValue   int    `json:"norm_value"`
}

// activ_log
type ActivityLogDTO struct {
	Id          int    `json:"id"`
	ActivNameId int    `json:"activ_name_id"`
	ActivNormId int    `json:"activ_norm_id"`
	ActivDate   string `json:"activ_date"`
	ActivValue  int    `json:"activ_value"`
	ActivName   string `json:"activ_name"`
}
