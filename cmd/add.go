package cmd

import (
	"Got/utils"
)

func Add(path string) {
	//Read the file
	contents := utils.StreamFile(path)

	//Get file hash
	hashed := utils.HashContents(contents)

	//Create a blob object, store it in objects with hash as his name

	//Add (the relative path, hash, perm) to staging.csv

}
