package app

import (
	"fmt"
	"log"
	"net/http"
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

// write responses
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

func WriteAnswer2(dataJSON []byte, status int, writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	//writer.WriteHeader(status)
	_, err := writer.Write(dataJSON)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
