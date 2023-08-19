package models

type FileInfo struct {
	Name  string `json:"file_name"`
	IsDir bool   `json:"file_isdir"`
}

type FilesListDTO struct {
	FilesList []*FileInfo `json:"files_list"`
	ServeURL  string      `json:"serve_url"`
}
