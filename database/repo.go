package database

import "github.com/Stupnikjs/goscrapp/data"

type DBRepo interface {
	InitTable() error
	InsertAnnonce(data.Annonce) error
	DropTable() error
}
