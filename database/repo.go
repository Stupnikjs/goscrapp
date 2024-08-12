package database

type DBRepo interface {
    InitTable() error
    InsertAnnonce (data.Annonce) error

}