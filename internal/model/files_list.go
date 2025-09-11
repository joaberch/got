package model

// FilesList lists all the mandatory file/folder
var FilesList = map[string]string{
	"staging.csv":     "File",   //Staging area
	"commits.csv":     "File",   //Tracking file state
	"objects/commits": "Folder", //Contain serialized commits with hash as name
	"objects/trees":   "Folder", //Contain serialized trees with hash as name
	"objects/blobs":   "Folder", //Contain serialized blobs with hash as name
	"head":            "File",   //Store the latest commit hash (HEAD)
}
