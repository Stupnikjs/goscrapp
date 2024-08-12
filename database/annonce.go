package database

import (
	"database/sql"
	"fmt"

	"github.com/Stupnikjs/goscrapp/data"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) InitTable() error {
	sqlInit := `CREATE TABLE annonces (
    id          VARCHAR(255),
    url         TEXT,
    pubdate     DATETIME,
    ville       VARCHAR(255),
    lieu        VARCHAR(255),
    departement INT,
    description TEXT,
    profession  VARCHAR(255),
    contrat     VARCHAR(255),
    created_at  DATETIME
);`
	fmt.Println(sqlInit)
	return nil
}

func (m *PostgresRepo) InsertAnnonces(data.Annonce) error {

	return nil
}
