package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
)

// ------------------------------------------------------------------------------------------------------
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
