

type DBRepo interface {
    InitTable() error
    InsertAnnonce (Annonce) error



}