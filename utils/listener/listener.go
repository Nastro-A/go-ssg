package listener

import (
	"fmt"
	"log"
	"os"
	"strings"

	"git.doganeplatz.eu/nastro/indexhtml"
	"git.doganeplatz.eu/nastro/markdowntohtml"
	"github.com/fsnotify/fsnotify"
)

func deleteFile(filePath string, savePath string) {
	fileNameArray := strings.Split(filePath, "/")
	fileNameWExt := fileNameArray[len(fileNameArray)-1]
	fileName := strings.Split(fileNameWExt, ".")[0]
	htmlName := strings.Join([]string{fileName, "html"}, ".")
	array := [2]string{savePath, htmlName}
	savePath = strings.Join(array[:], "/")
	err := os.Remove(savePath)
	if err != nil {
		log.Printf("Error removing file %s : %v\n", fileName, err)
	}
}

func Listener(inputPath string, savePath string, title string, theme string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error in creating the watcher: %v\n", err)
	}
	defer watcher.Close()

	err = watcher.Add(inputPath)
	if err != nil {
		log.Fatalf("Error in adding the directory %s to the watcher : %v\n", inputPath, err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Create) {
				filePath := event.Name
				fmt.Printf("New file created: %s\n", filePath)
				markdowntohtml.ConvertSingletoHTMLAndSave(filePath, savePath)
				indexhtml.IndexHTML(savePath, title, theme)
			}
			if event.Has(fsnotify.Rename) {
				filePath := event.Name
				fmt.Println(filePath)
				deleteFile(filePath, savePath)
				fmt.Printf("Removed file: %s\n", filePath)
				indexhtml.IndexHTML(savePath, title, theme)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
