package indexhtml

import (
	"bufio"
	"cmp"
	"html/template"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Link struct {
	Href     string
	Label    string
	Date     string
	Snippet1 string
	Snippet2 string
	Snippet3 string
}

type IndexData struct {
	Title string
	Files []Link
	Theme string
}

func temlpateToHTML(temp string, data IndexData, outputPath string) {
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("Error creating index.html")
	}

	t := template.Must(template.New("index").Parse(temp))
	err = t.Execute(f, data)
	if err != nil {
		log.Printf("Error in creating index.html")
	}
}

func extractSnippet(mdPath string, fileName string) []string {
	fileName = strings.Join([]string{fileName, "md"}, ".")
	f, err := os.Open(strings.Join([]string{mdPath, fileName}, "/"))
	if err != nil {
		log.Fatalf("Error opening %v for extracting snippet: %v", f.Name(), err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	snippet := []string{}
	for i := 0; i < 3; i++ {
		scanner.Scan()
		snippet = append(snippet, scanner.Text())
	}

	if scanner.Err() != nil {
		return []string{"", "", ""}
	}

	return snippet
}

// helped by chat-gpt-4o
func IndexHTML(mdPath string, dirPath string, title string, theme string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading %s directory: %v", dirPath, err)
	}
	filesWOIndex := []Link{}
	for i := 0; i < len(files); i++ {
		fileName := files[i].Name()
		if fileName != "index.html" && !files[i].IsDir() {
			snippet := extractSnippet(mdPath, strings.Split(fileName, ".")[0])
			fileInfo, err := files[i].Info()
			if err != nil {
				log.Fatalf("Error obtaining informations from file %s : %v\n", fileName, err)
			}
			file := Link{
				Href:  fileName,
				Label: strings.Split(fileName, ".")[0],
				Date:  fileInfo.ModTime().Local().Format(time.RFC822),

				Snippet1: string(snippet[0]),
				Snippet2: string(snippet[1]),
				Snippet3: string(snippet[2]),
			}
			filesWOIndex = append(filesWOIndex, file)
		}
		slices.SortFunc(filesWOIndex, func(a, b Link) int {
			return cmp.Compare(b.Date, a.Date)
		})
	}

	data := IndexData{
		Title: title,
		Files: filesWOIndex,
		Theme: theme,
	}

	dirPath = strings.Join([]string{dirPath, "index.html"}, "/")

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
      <li><a href="{{.Href}}">{{.Label}} </a>{{.Date}}</li>
			<p>{{.Snippet1}}<br />{{.Snippet2}}<br />{{.Snippet3}}</p>
      {{end}}
    </ul>
  </body>
  <footer>
      <hr />
      <p style="text-align: center;">Created with <a href="https://github.com/Nastro-A/go-ssg">go-ssg</a> and <a href="https://picocss.com/">pico</a></p>
    </footer>
</html>`

	temlpateToHTML(temp, data, dirPath)
}
