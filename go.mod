module git.doganeplatz.eu/nastro/go-ssg

go 1.24.5

replace git.doganeplatz.eu/nastro/go-ssg/utils/listener => ./utils/listener

replace git.doganeplatz.eu/nastro/go-ssg/utils/markdowntohtml => ./utils/markdowntohtml

replace git.doganeplatz.eu/nastro/go-ssg/utils/indexhtml => ./utils/indexhtml

require git.doganeplatz.eu/nastro/go-ssg/utils/listener v0.0.0

require (
	git.doganeplatz.eu/nastro/go-ssg/utils/indexhtml v0.0.0
	github.com/joho/godotenv v1.5.1
)

require (
	git.doganeplatz.eu/nastro/go-ssg/utils/markdowntohtml v0.0.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gomarkdown/markdown v0.0.0-20250731182530-5d03d1963446 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
