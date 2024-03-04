package app

import (
	"encoding/json"
	"log"
	"net/http"
	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
)

// ------------------------------------------------------------------------------------------------------
// Activities - names
func (s *Server) handleActivitiesNames(writer http.ResponseWriter, request *http.Request) {
	dataResp, err := s.backendSvc.GetActivitiesList()
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(dataResp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesNamesNew(writer http.ResponseWriter, request *http.Request) {
	var dataNewActiv *models.ActivityDTO
	err := json.NewDecoder(request.Body).Decode(&dataNewActiv)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.InsertNewActiv(dataNewActiv) //insert
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesNamesUpd(writer http.ResponseWriter, request *http.Request) {
	var dataUpdActiv *models.ActivityDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdActiv)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.UpdateNewActivById(dataUpdActiv) // update
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesNamesDel(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	err := s.backendSvc.DeleteNewActivById(delId) //delete
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// ------------------------------------------------------------------------------------------------------
// Activities - normatives
func (s *Server) handleActivNormativs(writer http.ResponseWriter, request *http.Request) {
	dataResp, err := s.backendSvc.GetActivNorm()
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(dataResp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivNormativsNew(writer http.ResponseWriter, request *http.Request) {
	var dataNewActivNorm *models.ActivityNormativeDTO
	err := json.NewDecoder(request.Body).Decode(&dataNewActivNorm)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.ActivNormNew(dataNewActivNorm)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivNormativsUpd(writer http.ResponseWriter, request *http.Request) {
	var dataUpdActivNorm *models.ActivityNormativeDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdActivNorm)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.ActivNormUpdate(dataUpdActivNorm)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivNormativsDel(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	err := s.backendSvc.ActivNormDelById(delId)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// ------------------------------------------------------------------------------------------------------
// Activities - logs
func (s *Server) handleActivitiesLogs(writer http.ResponseWriter, request *http.Request) {
	dataResp, err := s.backendSvc.ActivLogsAll() //
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(dataResp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesLogsNew(writer http.ResponseWriter, request *http.Request) {
	var dataNewActivLog *models.ActivityLogDTO
	err := json.NewDecoder(request.Body).Decode(&dataNewActivLog)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.ActivLogsNew(dataNewActivLog)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesLogsUpd(writer http.ResponseWriter, request *http.Request) {
	var dataUpdActivLog *models.ActivityLogDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdActivLog)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.backendSvc.ActivLogsUpdate(dataUpdActivLog)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleActivitiesLogsDel(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	err := s.backendSvc.ActivLogsDelById(delId)
	if err != nil {
		log.Println(err)
		return
	}
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}
