package markdowntohtml

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func RetriveMDFiles(dirpath string) ([]os.DirEntry, error) {
	_, err := os.Stat(dirpath)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	files, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func saveHTML(file []byte, HTMLPath string, fileName string) error {
	fileName = strings.Split(fileName, ".")[0]
	htmlFileName := []string{fileName, "html"}
	fileName = strings.Join(htmlFileName, ".")
	path := []string{HTMLPath, fileName}
	err := os.WriteFile(strings.Join(path, "/"), file, 0700)
	if err != nil {
		return err
	}
	return nil
}

func ConvertSingletoHTMLAndSave(filepath string, savePath string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", filepath, err)
	}
	array := strings.Split(filepath, "/")
	fileName := array[len(array)-1]
	if file == nil {
		log.Fatalf("File empty, Ignoring %s", fileName)
	}

	if strings.Split(fileName, ".")[1] != "md" {
		log.Printf("File %s not a markdown file", fileName)
		return
	}
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(file)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.Render(doc, renderer)

	err = saveHTML(html, savePath, fileName)
	if err != nil {
		log.Fatalf("Error saving file %s : %v", fileName, err)
	}
}
