package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
	//"github.com/go-chi/chi/v5"
)

// echo
func (s *Server) handleEcho(writer http.ResponseWriter, request *http.Request) {
	msg := fmt.Sprintf("this is echo page")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Println(err)
		return
	}
}

// helpers ----------------------------------------

// write response
func WriteAnswer(dataJSON []byte, writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(dataJSON)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ------------------------------------------------------------------------------------------------------
// Knowbase posts

func (s *Server) handleAllPosts(writer http.ResponseWriter, request *http.Request) {
	posts, err := s.backendSvc.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// single post by id
func (s *Server) handlePostById(writer http.ResponseWriter, request *http.Request) {
	postID := chi.URLParam(request, "post_id")
	//fmt.Println("postID=", postID)
	post, err := s.backendSvc.GetOnePostByID(postID)
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// save post updates to DB
func (s *Server) handleSaveUpdates(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println("entered PUT handleSaveUpdates")
	// request body
	var dataUpdate *models.PostDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	//fmt.Printf("dataUpdate = %+v\n", dataUpdate)
	// save to DB (backend)
	err = s.backendSvc.UpdatePostById(dataUpdate)
	resp := &models.ResponseDTO{Status: "OK"}
	dataJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// insert new post
func (s *Server) handleInsertNewPost(writer http.ResponseWriter, request *http.Request) {
	var dataInsert *models.PostDTO
	err := json.NewDecoder(request.Body).Decode(&dataInsert)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Printf("dataInsert = %+v\n", dataInsert)
	err = s.backendSvc.InsertNewPost(dataInsert)
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

// delete post by id
func (s *Server) handleDeletePostById(writer http.ResponseWriter, request *http.Request) {
	postID := chi.URLParam(request, "post_id")
	fmt.Printf("deleting = %+v\n", postID)
	err := s.backendSvc.DeletePostById(postID)
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
// Activities - names

//
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

//
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

//
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

//
func (s *Server) handleActivitiesNamesDel(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	//fmt.Printf("deleting = %+v\n", delId)
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

// ------------------------------------------------------------------------------------------------------
// Activities - analytics

func (s *Server) handleActiv3(writer http.ResponseWriter, request *http.Request) {
	dateFrom := chi.URLParam(request, "date_from")
	dateTo := chi.URLParam(request, "date_to")

	data, err := s.backendSvc.AnalyticsActiv3(dateFrom, dateTo)
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

func (s *Server) handleActivRepCommon(writer http.ResponseWriter, request *http.Request) {
	/**/
	dateFrom := chi.URLParam(request, "date_from")
	dateTo := chi.URLParam(request, "date_to")

	data, err := s.backendSvc.Activ4CommonGraphByDates(dateFrom, dateTo)
	if err != nil {
		log.Println(err)
		return
	}
	/*fmt.Println(data.Labels)
	for _, v := range data.Datasets {
		fmt.Printf("%+v\n", v)
	}*/

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
	//fmt.Println(df, dt)

	data := &models.ParamsDates{DateFrom: df.String()[:10], DateTo: dt.String()[:10]}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}
