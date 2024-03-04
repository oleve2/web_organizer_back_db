package app

import (
	"encoding/json"
	"log"
	"net/http"
	models "webapp3/pkg/models"

	"github.com/go-chi/chi"
)

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
