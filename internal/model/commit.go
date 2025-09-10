package model

type Commit struct {
	TreeHash   string
	ParentHash string
	Author     string
	Message    string
	Timestamp  int64
}
