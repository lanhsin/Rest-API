package schema

type CustomVersion struct {
	Version string `json:"Version"`
	Commit  string `json:"Commit"`
	Date    string `json:"Date"`
}