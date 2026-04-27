package repo

import (
	"github.com/jmoiron/sqlx"
)

type PdfRepository interface {
	GetProfile(id int)
}

type pdfRepo struct {
	db *sqlx.DB
}

func NewPdfRepository(db *sqlx.DB) PdfRepository {
	return &pdfRepo{db: db}
}

func (r *pdfRepo) GetProfile(id int) {}
