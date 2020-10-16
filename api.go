package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func getList(c *gin.Context) {

	root := GetEnvVar("MAPS_PATH")

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if path != root {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(files)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = c.Writer.Write(data)
}

func getMap(c *gin.Context) {

	var err error

	params := c.Request.URL.Query()
	nameParam, ok := params["name"]
	if !ok || len(nameParam) == 0 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	name := nameParam[0]

	var file *os.File
	file, err = getFile(name)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	defer file.Close()

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	_, err = file.Read(FileHeader)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := file.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=" + name)
	c.Writer.Header().Set("Content-Type", FileContentType)
	c.Writer.Header().Set("Content-Length", FileSize)

	_, err = file.Seek(0, 0)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getFile(filename string) (*os.File, error) {

	file, err := os.Open(GetEnvVar("MAPS_PATH") + string(os.PathSeparator) + filename)
	if err != nil {
		//File not found, send 404
		return nil, err
	}

	return file, nil
}
