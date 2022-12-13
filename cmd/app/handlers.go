package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
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
	// request body
	var dataUpdate *models.PostDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
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

// static files handler
func (s *Server) handleFilesList(writer http.ResponseWriter, request *http.Request) {
	files, err := s.upDownSvc.GetStaticFolderContent()
	if err != nil {
		log.Println(err)
		return
	}

	data := &models.FilesListDTO{
		FilesList: files,
		ServeURL:  s.upDownSvc.ServeURL,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// form upload
func (s *Server) handleFormUpload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(200 * 1024)

	/*
		for k, v := range request.Form {
			fmt.Println(k, v)
		}
		formVal1 := request.FormValue("val1")
		formVal2 := request.FormValue("val2")
		fmt.Printf("formVal1=%s  formVal2=%s\n", formVal1, formVal2)
	*/

	// save multiple files
	// https://socketloop.com/tutorials/upload-multiple-files-golang
	files := request.MultipartForm.File["dataFile"]

	for i, _ := range files { // loop through the files one by one
		err := s.upDownSvc.SaveMultipleFiles(files, i)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// save one file
	/*
		file, fileheader, err := request.FormFile("dataFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploaded File: %+v %d\n", fileheader.Filename, fileheader.Size)
		defer file.Close()
		err = s.upDownSvc.SaveFileFromFormToFolder(file, fileheader)
	*/

	// response
	data := "form recieved"
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

// TAGS
func (s *Server) handleTagsAll(writer http.ResponseWriter, request *http.Request) {
	tags, err := s.backendGormServ.TagsAll()
	if err != nil {
		log.Println(err)
		return
	}
	dataJSON, err := json.Marshal(tags)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleTagsInsertOne(writer http.ResponseWriter, request *http.Request) {
	var dataInsert *models.Tags
	err := json.NewDecoder(request.Body).Decode(&dataInsert)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = s.backendGormServ.TagsInsertOne(dataInsert)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	WriteAnswer([]byte("insert done"), writer)
}

func (s *Server) handleTagsUpdateOne(writer http.ResponseWriter, request *http.Request) {
	var dataUpdate *models.Tags
	err := json.NewDecoder(request.Body).Decode(&dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = s.backendGormServ.TagsUpdateOne(dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	WriteAnswer([]byte("update done"), writer)
}

func (s *Server) handleTagsDeleteOne(writer http.ResponseWriter, request *http.Request) {
	delId := chi.URLParam(request, "del_id")
	delIdInt, err := strconv.Atoi(delId)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Println("delId=", delId)
	err = s.backendGormServ.TagsDeleteOne(delIdInt)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	WriteAnswer([]byte("delete done"), writer)
}
