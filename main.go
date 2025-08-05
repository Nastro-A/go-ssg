package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"git.doganeplatz.eu/nastro/indexhtml"
	"git.doganeplatz.eu/nastro/listener"
	"github.com/joho/godotenv"
)

func ifNotExistsMkDir(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir(dirPath, 0700)
		if err != nil {
			log.Fatalf("Error creating %s directory: %v\n", dirPath, err)
		}
		err = os.Mkdir(strings.Join([]string{dirPath, "tmp"}, "/"), 0700)
		if err != nil {
			log.Fatalln("Error creating tmp folder")
		}
	} else {
		return false
	}
	return true
}

func main() {
	godotenv.Load(".env")

	MDDir := os.Getenv("MD_DIR")
	if MDDir == "" {
		log.Fatalln("MD_DIR not set in environment")
	}
	HTMLDir := os.Getenv("HTML_DIR")
	if HTMLDir == "" {
		log.Fatalln("HTML_DIR not set in environment")
	}
	port := os.Getenv("PORT")
	port = strings.Join([]string{":", port}, "")
	if port == "" {
		log.Fatalln("PORT not set in environment")
	}
	title := os.Getenv("HTML_TITLE")
	if title == "" {
		log.Fatalln("HTML_TITLE not set in environment")
	}
	theme := os.Getenv("THEME")
	if theme == "" {
		log.Fatalln("THEME not set in environment, check https://picocss.com for the available colors")
	}
	statusMDDir := ifNotExistsMkDir(MDDir)
	statusHTMLDir := ifNotExistsMkDir(HTMLDir)

	if statusMDDir || statusHTMLDir {
		fmt.Println("Directory creation succeded")
	}

	indexhtml.IndexHTML(MDDir, HTMLDir, title, theme)
	fmt.Println("index.html creation succeded")

	fmt.Println("Starting watcher...")
	go listener.Listener(MDDir, HTMLDir, title, theme)

	fmt.Printf("Starting web server at port %s\n", port)
	http.Handle("/", http.FileServer(http.Dir(HTMLDir)))
	http.ListenAndServe(port, nil)
}
