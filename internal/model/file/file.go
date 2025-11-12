package file

import "github.com/DevopsCreator/http-server/internal/model/common"

type Extension string
type Integrity bool // Целостность

type File struct {
	Extension
	common.Path
	Integrity
}
