package model

type Lab struct {
	Id              string   `json:"id"`
	Title           string   `json:"title"`
	Difficulty      string   `json:"difficulty"`
	Category        string   `json:"category"`
	Description     string   `json:"description"`
	LongDescription string   `json:"longDescription"`
	Hints           []string `json:"hints"`
	Completions     int      `json:"completions"`
	EstimatedTime   string   `json:"estimatedTime"`
	Tags            []string `json:"tags"`
}

var LabList []Lab
