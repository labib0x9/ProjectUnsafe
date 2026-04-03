package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type LabRepository interface {
	Create(lab model.Lab) (model.Lab, error)
	Get(id string) (model.Lab, error)
	List() ([]model.Lab, error)
	Delete(id string) error
	Update(lab model.Lab) error
}

type labRepo struct {
	dbConn *sqlx.DB
}

func NewLabRepo(dbConn *sqlx.DB) LabRepository {
	return &labRepo{dbConn: dbConn}
}

func (l *labRepo) Create(lab model.Lab) (model.Lab, error) {
	query := `
        insert into labs (labname, title, difficulty, category, description, long_description, hints, total_solved, container_id)
        values (:labname, :title, :difficulty, :category, :description, :long_description, :hints, :total_solved, :container_id)
        returning *
    `

	rows, err := l.dbConn.NamedQuery(query, lab)
	if err != nil {
		return model.Lab{}, err
	}
	defer rows.Close()

	var created model.Lab
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return model.Lab{}, err
		}
	}
	return created, nil
}

func (l *labRepo) Get(id string) (model.Lab, error) {
	var lab model.Lab
	query := `select * from labs where labname = $1`
	if err := l.dbConn.Get(&lab, query, id); err != nil {
		return model.Lab{}, err
	}
	return lab, nil
}

func (l *labRepo) List() ([]model.Lab, error) {
	var labs []model.Lab
	query := `select * from labs`
	if err := l.dbConn.Select(&labs, query); err != nil {
		return []model.Lab{}, err
	}
	return labs, nil
}

func (l *labRepo) Delete(id string) error {
	query := `delete from labs where labname = $1`
	if _, err := l.dbConn.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (l *labRepo) Update(lab model.Lab) error {
	return nil
}
