package repository

import (
	"fmt"

	file "github.com/DevopsCreator/http-server/internal/model/file"
	header "github.com/DevopsCreator/http-server/internal/model/header"
)

var SliceOfFile []file.File

var SliceOfHeader []header.Header

func FileToSlice(file *file.File) {
	SliceOfFile = append(SliceOfFile, *file)
}

func HeaderToSlice(header *header.Header) {
	SliceOfHeader = append(SliceOfHeader, *header)
}