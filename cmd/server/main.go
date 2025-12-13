package main

import (
	"github.com/DevopsCreator/http-server/internal/repository"
	"github.com/DevopsCreator/http-server/internal/service"
)

func main() {

	repository.FileToSlice(service.NewFile())
	repository.HeaderToSlice(service.NewHeader())
}
