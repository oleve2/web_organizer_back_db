package models

type FilesListDTO struct {
	FilesList []string `json:"files_list"`
	ServeURL  string   `json:"serve_url"`
}
