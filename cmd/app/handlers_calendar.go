package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp3/pkg/models"

	"github.com/go-chi/chi"
)

// ------------------------------------------------------------------------------------------------------
// CALENDAR

func (s *Server) handleCalendarGridByMonth(writer http.ResponseWriter, request *http.Request) {
	var dataYearMonth *models.CalendarYearMonthDTO
	err := json.NewDecoder(request.Body).Decode(&dataYearMonth)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println("YearMonth=", dataYearMonth.YearMonth)

	clItems := s.backendGormServ.Calend_GetGrid(dataYearMonth.YearMonth)

	dataJSON, err := json.Marshal(clItems)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleCalendarItemsAll(writer http.ResponseWriter, request *http.Request) {
	var dataYearMonth *models.CalendarYearMonthDTO
	err := json.NewDecoder(request.Body).Decode(&dataYearMonth)
	if err != nil {
		log.Println(err)
		return
	}

	calendItems, err := s.backendGormServ.CalendItemAll(dataYearMonth.YearMonth)
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(calendItems)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleCalendarItemsNofiltered(writer http.ResponseWriter, request *http.Request) {
	calendItems, err := s.backendGormServ.CalendItemsAllNoFilter()
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(calendItems)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleCalendarInsertOne(writer http.ResponseWriter, request *http.Request) {
	var dataInsert *models.CalendarData
	err := json.NewDecoder(request.Body).Decode(&dataInsert)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendGormServ.CalendItemInsertOne(dataInsert)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer([]byte("insert done"), writer)
}

func (s *Server) handleCalendarUpdateOne(writer http.ResponseWriter, request *http.Request) {
	var dataUpdate *models.CalendarData
	err := json.NewDecoder(request.Body).Decode(&dataUpdate)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.backendGormServ.CalendItemUpdateOne(dataUpdate)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer([]byte("update done"), writer)
}

func (s *Server) handleCalendarDeleteOne(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	delIdInt, err := strconv.Atoi(delId)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Println("delId=", delIdInt)
	err = s.backendGormServ.CalendItemDeleteOne(delIdInt)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer([]byte("delete done"), writer)
}
