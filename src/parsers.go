package src

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

func ParseFile(filepath string, info fs.FileInfo) (content string, fp string) {

	fmt.Printf("Parsing %v , size: %v\n", filepath, info.Size())
	switch extension := path.Ext(filepath); strings.ToLower(extension) {
	case ".xml", ".xhtml":
		//Open file
		file, fileErr := os.Open(filepath)
		if fileErr != nil {
			fmt.Printf("Error: could open file: %v \n", fileErr)
		}
		defer file.Close()
		content = parseXmlDoc(file)
	case ".txt", ".md":
		b, fileErr := os.ReadFile(filepath)
		if fileErr != nil {
			fmt.Printf("Error: could open file: %v \n", fileErr)
		}
		content = string(b)
	default:
		fmt.Printf("File extension %v not yet supported \n", extension)
	}

	fp = filepath
	return strings.TrimSpace(content), fp
}

func parseXmlDoc(file *os.File) (content string) {
	// decode xml
	decoder := xml.NewDecoder(file)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error getting token: %v\n", err)
			break
		}

		if charData, ok := token.(xml.CharData); ok {
			// process as text. How do I read the text data?
			s := string([]byte(charData))
			content += s
		}

	}

	return (content)
}
