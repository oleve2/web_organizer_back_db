package models

// --------------------------------
// models for analytics
type AnActivNameList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// AnDateVal -> AnNameLabelsValues -> AnListNLV
type AnDateVal struct {
	ActivDate  string `json:"activ_date"`
	ActivValue int    `json:"activ_value"`
}

type AnNameLabelsValues struct {
	ChartNameId int      `json:"chart_name_id"`
	ChartName   string   `json:"chart_name"`
	Labels      []string `json:"labels"`
	Data        []int    `json:"data"`
}

// Всё вмете -----------------------------------
//
type ActivityData struct {
	Label string `json:"label"`
	Data  []int  `json:"data"`
}

type ActivityDataReport struct {
	Labels   []string        `json:"labels"`
	Datasets []*ActivityData `json:"datasets"`
}

type ParamsDates struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}
