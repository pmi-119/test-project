package api

import "file-structure/internal/model"

type SrvI interface {
	Summator(data *model.Data)
}
