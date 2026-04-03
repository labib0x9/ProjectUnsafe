package model

import "github.com/lib/pq"

type Lab struct {
	Id              int            `db:"id"`
	LabName         string         `json:"labname" db:"labname"`
	Title           string         `json:"title" db:"title"`
	Difficulty      string         `json:"difficulty" db:"difficulty"`
	Category        string         `json:"category" db:"category"`
	Description     string         `json:"description" db:"description"`
	LongDescription string         `json:"long_description" db:"long_description"`
	Hints           pq.StringArray `json:"hints" db:"hints"`
	Completions     int            `json:"completions" db:"total_solved"`
	ContainerId     string         `db:"container_id"`
}
