package main

import (
	"go_local_search_engine/src"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//todo
	//add command to interactivey search
	//modify model.go to update index.json based on file chandes in dir
	//store index in sqlite or leave as json

	indexDir()

}

func indexDir() {

	m := src.NewModel()

	dir := "./docs.gl"
	// dir := "./test_docs"

	_ = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Error couldnt walk %v \n", err)
			}
			content, fp := src.ParseFile(path, info)
			if len(strings.Trim(content, " ")) == 0 {
				return nil
			}
			t := src.NewTokenizer(content, fp)
			m.NewDoc(t)

			return nil
		})

	m.SaveAllDocuments()

	// t:=src.NewTokenizerQuery("linear interpolation")
	t := src.NewTokenizerQuery("hyperbolic cosine")
	// t:=src.NewTokenizerQuery("bvec4")
	m.Query_terms(t)

}
