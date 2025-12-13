package service

import (
	file "github.com/DevopsCreator/http-server/internal/model/file"
	header "github.com/DevopsCreator/http-server/internal/model/header"
)

func NewFile() *file.File {
	return &file.File{
		Name:            "",
		Type:            "",
		LastTimeChanges: 0,
		ReRead:          false,
		CacheTime:       0,
	}
}

func NewHeader() *header.Header {
	return &header.Header{
		Name: "",
	}
}
