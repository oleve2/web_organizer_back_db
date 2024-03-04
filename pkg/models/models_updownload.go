package models

type FileInfo struct {
	Name  string `json:"file_name"`
	IsDir bool   `json:"file_isdir"`
}

type FilesListDTO struct {
	FilesList []*FileInfo `json:"files_list"`
	ServeURL  string      `json:"serve_url"`
}

type FilesUpdateItemDTO struct {
	FileNewName  string    `json:"file_new_name"`
	OriginalFile *FileInfo `json:"file_original"`
}

type FilesUpdateItemResponseDTO struct {
	File     *FileInfo `json:"file_original"`
	Status   bool      `json:"status"`
	ErrorStr string    `json:"error_str"`
}
