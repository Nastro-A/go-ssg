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

func ConvertSingletoHTMLAndSave(filepath string, savePath string, theme string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", filepath, err)
	}
	array := strings.Split(filepath, "/")
	fileName := array[len(array)-1]
	if file == nil {
		log.Fatalf("File empty, Ignoring %s", fileName)
	}
	fileNameWOExt := strings.Split(fileName, ".")[0]
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
	html = []byte(`
	<!DOCTYPE html>
	<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="color-scheme" content="light dark" />
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.` + theme + `.min.css"
    />
    <title>` + fileNameWOExt + `</title>
  </head>
  <header>
      <nav>
        <ul>
          <li></li>
          <li><strong>` + fileNameWOExt + `</strong></li>
        </ul>
        <ul>
          <li><a href="index.html">Home</a></li>
          <li></li>
        </ul>
      </nav>
    </header>
  <body>
	<main style="padding: 15px;">` + string(html) + `</main>
	</body>
  <footer>
      <hr />
      <p style="text-align: center;">Created with <a href="https://github.com/Nastro-A/go-ssg">go-ssg</a> and <a href="https://picocss.com/">pico</a></p>
    </footer>
</html>
	`)
	err = saveHTML(html, savePath, fileName)
	if err != nil {
		log.Fatalf("Error saving file %s : %v", fileName, err)
	}
}
