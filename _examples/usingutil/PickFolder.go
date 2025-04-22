package main

import (
	"log"

	"github.com/karelbilek/go-common-file-dialog/cfd"
	"github.com/karelbilek/go-common-file-dialog/cfdutil"
)

func main() {
	result, err := cfdutil.ShowPickFolderDialog(cfd.DialogConfig{
		Title:  "Pick Folder",
		Role:   "PickFolderExample",
		Folder: "C:\\",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen folder: %s\n", result)
}
