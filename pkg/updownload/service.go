package uploaddownload

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"webapp3/pkg/models"
)

type Service struct {
	FolderPath string
	ServeURL   string
}

func NewService(FolderPath string, ServeURL string) *Service {
	return &Service{
		FolderPath: FolderPath,
		ServeURL:   ServeURL,
	}
}

// ---------------------------------
func (s *Service) InitCheck() (bool, error) {
	fmt.Printf("FolderPath=%s ServeURL=%s\n", s.FolderPath, s.ServeURL)
	_, err := os.Stat(s.FolderPath)
	if err == nil {
		fmt.Printf("folder %s exists\n", s.FolderPath)
		return true, nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(s.FolderPath, 0777)
		fmt.Printf("folder %s created\n", s.FolderPath)
		return false, nil
	}
	return false, err
}

// download
func (s *Service) GetStaticFolderContent() ([]*models.FileInfo, error) {
	//res := make([]string, 0)
	fileList := make([]*models.FileInfo, 0)
	folderList := make([]*models.FileInfo, 0)

	files, err := ioutil.ReadDir(s.FolderPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		tmp := &models.FileInfo{Name: file.Name(), IsDir: file.IsDir()}
		if file.IsDir() {
			folderList = append(folderList, tmp)
		} else {
			fileList = append(fileList, tmp)
		}
	}

	res := append(folderList, fileList...)

	return res, nil
}

// upload multi
func (s *Service) SaveMultipleFiles(files []*multipart.FileHeader, i int) error {
	file, err := files[i].Open()
	defer file.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	fname := s.FolderPath + "/" + files[i].Filename
	//fmt.Println(i, " ", fname)
	out, err := os.Create(fname)

	defer out.Close()
	if err != nil {
		log.Println("Unable to create the file for writing. Check your write access privilege")
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(i, " ", files[i].Filename, " <- written to ", fname)
	//
	return nil
}

/*
// upload single (not actual)
func (s *Service) SaveFileFromFormToFolder(file multipart.File, fileheader *multipart.FileHeader) error {
	//=== start
	tempFile, err := os.CreateTemp(s.FolderPath, fmt.Sprintf("*-%s", fileheader.Filename))
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	//
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	fmt.Println("file writteen")
	//=== finish
	return nil
}
*/
