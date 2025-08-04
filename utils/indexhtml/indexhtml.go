package indexhtml

import (
	"html/template"
	"log"
	"os"
	"strings"
)

// helped by chat-gpt-4o
func IndexHTML(dirPath string, title string, theme string) {
	type Link struct {
		Href  string
		Label string
	}

	type IndexData struct {
		Title string
		Files []Link
		Theme string
	}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error in reading %s directory: %v", dirPath, err)
	}
	filesWOIndex := []Link{}
	for i := 0; i < len(files); i++ {
		fileName := files[i].Name()
		if fileName != "index.html" && !files[i].IsDir() {
			file := Link{
				Href:  fileName,
				Label: strings.Split(fileName, ".")[0],
			}
			filesWOIndex = append(filesWOIndex, file)
		}
	}

	data := IndexData{
		Title: title,
		Files: filesWOIndex,
		Theme: theme,
	}

	dirPath = strings.Join([]string{dirPath, "index.html"}, "/")
	f, err := os.Create(dirPath)
	if err != nil {
		log.Fatalln("Error creating index.html")
	}

	const temp = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="color-scheme" content="light dark" />
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.{{.Theme}}.min.css"
    />
    <title>{{.Title}}</title>
  </head>
  <header>
      <nav>
        <ul>
          <li></li>
          <li><strong>{{.Title}}</strong></li>
        </ul>
        <ul>
          <li><a href="index.html">Home</a></li>
          <li></li>
        </ul>
      </nav>
    </header>
  <body>
    <ul>
      {{range .Files}}
      <li><a href="{{.Href}}">{{.Label}}</a></li>
      {{end}}
    </ul>
  </body>
  <footer>
      <hr />
      <p style="text-align: center;">Created with <a href="http://www.github.com/Nastro-A/go-ssg">go-ssg</a> and <a href="https://picocss.com/">pico</a></p>
    </footer>
</html>`

	t := template.Must(template.New("index").Parse(temp))
	err = t.Execute(f, data)
	if err != nil {
		log.Printf("Error in creating index.html")
	}
}

func aboutHTML(dirpath string, title string, theme string) {
}
