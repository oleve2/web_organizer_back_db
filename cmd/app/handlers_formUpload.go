package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	models "webapp3/pkg/models"
)

// ------------------------------------------------------------------------------------------------------
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

// ------------------------------------------------------------------------------------------------------
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

func (s *Server) handleFilesUpdateItem(writer http.ResponseWriter, request *http.Request) {
	// request body
	var dataUpdate *models.FilesUpdateItemDTO
	err := json.NewDecoder(request.Body).Decode(&dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Printf("dataUpdate %+v \n", dataUpdate)

	//
	err = s.upDownSvc.Files_UpdateDeleteItem("update", dataUpdate)
	if err != nil {
		log.Println(err)
		//http.Error(writer, err.Error(), http.StatusBadRequest)
		//return
		data := &models.FilesUpdateItemResponseDTO{
			File:     dataUpdate.OriginalFile,
			Status:   false,
			ErrorStr: err.Error(),
		}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}
		WriteAnswer2(dataJSON, http.StatusBadRequest, writer)
		return
	}

	data := &models.FilesUpdateItemResponseDTO{
		File:     dataUpdate.OriginalFile,
		Status:   true,
		ErrorStr: "",
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}

func (s *Server) handleFilesDeleteItem(writer http.ResponseWriter, request *http.Request) {
	// request body
	var dataDelete *models.FilesUpdateItemDTO
	err := json.NewDecoder(request.Body).Decode(&dataDelete)
	if err != nil {
		log.Println(err)
		//http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Printf("dataDelete %+v \n", dataDelete)

	//
	err = s.upDownSvc.Files_UpdateDeleteItem("delete", dataDelete)
	if err != nil {
		log.Println(err)
		//http.Error(writer, err.Error(), http.StatusBadRequest)
		//return
		data := &models.FilesUpdateItemResponseDTO{
			File:     dataDelete.OriginalFile,
			Status:   false,
			ErrorStr: err.Error(),
		}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}
		WriteAnswer2(dataJSON, http.StatusBadRequest, writer)
		return
	}

	data := &models.FilesUpdateItemResponseDTO{
		File:     dataDelete.OriginalFile,
		Status:   true,
		ErrorStr: "",
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	WriteAnswer(dataJSON, writer)
}
