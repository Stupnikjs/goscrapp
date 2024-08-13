package database

import (
	"context"
	"database/sql"

	"github.com/Stupnikjs/goscrapp/data"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) InitTable() error {
	sqlInit := `CREATE TABLE IF NOT EXISTS annonces (
    id          VARCHAR(255),
    url         TEXT,
    pubdate     TEXT,
    ville       VARCHAR(255),
    lieu        VARCHAR(255),
    departement INT,
    description TEXT,
    profession  VARCHAR(255),
    contrat     VARCHAR(255),
    created_at  TEXT
);`
	_, err := m.DB.Exec(sqlInit)
	return err
}

func (m *PostgresRepo) InsertAnnonce(annonce data.Annonce) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sqlInsert := `
    INSERT INTO annonces (id, url, pubdate, ville, lieu, departement, description, profession, contrat, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
    `
	_, err := m.DB.ExecContext(ctx, sqlInsert, annonce.Id, annonce.Url, annonce.PubDate, annonce.Ville, annonce.Lieu, annonce.Departement, annonce.Description, annonce.Profession, annonce.Contrat, annonce.Created_at)
	return err
}
