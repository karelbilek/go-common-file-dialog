package main

import (
	"log"

	"github.com/karelbilek/go-common-file-dialog/cfd"
	"github.com/karelbilek/go-common-file-dialog/cfdutil"
)

func main() {
	results, err := cfdutil.ShowOpenMultipleFilesDialog(cfd.DialogConfig{
		Title: "Open Multiple Files",
		Role:  "OpenFilesExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.txt",
		DefaultExtension:        "txt",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file(s): %s\n", results)
}
