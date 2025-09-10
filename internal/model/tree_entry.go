package model

type TreeEntry struct {
	Name string
	Mode string //file or folder ...
	Type string //blob or tree
	Hash string
}
