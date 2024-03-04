package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
)

// ------------------------------------------------------------------------------------------------------
// Activities - analytics
func (s *Server) handleCommongraphs(writer http.ResponseWriter, request *http.Request) {
	dateFrom := chi.URLParam(request, "date_from")
	dateTo := chi.URLParam(request, "date_to")

	data, err := s.backendSvc.PrepareCommonGraphs(dateFrom, dateTo) //AnalyticsActiv3
	if err != nil {
		log.Println(err)
		return
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleIndividualGraphs(writer http.ResponseWriter, request *http.Request) {
	dateFrom := chi.URLParam(request, "date_from")
	dateTo := chi.URLParam(request, "date_to")

	data, err := s.backendSvc.PrepareIndivGraphs(dateFrom, dateTo) //Activ4CommonGraphByDates
	if err != nil {
		log.Println(err)
		return
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleAnalyticParams(writer http.ResponseWriter, request *http.Request) {
	d := time.Now()
	currentLocation := d.Location()

	currentYear, currentMonth, _ := d.Date()
	df := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	dt := df.AddDate(0, 1, 0)

	data := &models.ParamsDates{DateFrom: df.String()[:10], DateTo: dt.String()[:10]}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}
