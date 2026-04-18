package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type PdfRepository interface {
	GetProfile(id int) (model.User, error)
}

type pdfRepo struct {
	db *sqlx.DB
}

func NewPdfRepository(db *sqlx.DB) PdfRepository {
	return &userRepo{db: db}
}

func (r *pdfRepo) GetProfile(id int) () {}