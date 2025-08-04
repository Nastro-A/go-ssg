module git.doganeplatz.eu/nastro/listener

go 1.24.5

replace git.doganeplatz.eu/nastro/markdowntohtml => ../markdowntohtml

replace git.doganeplatz.eu/nastro/indexhtml => ../indexhtml

require git.doganeplatz.eu/nastro/markdowntohtml v0.0.0

require git.doganeplatz.eu/nastro/indexhtml v0.0.0

require github.com/fsnotify/fsnotify v1.9.0

require (
	github.com/gomarkdown/markdown v0.0.0-20250731182530-5d03d1963446 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
