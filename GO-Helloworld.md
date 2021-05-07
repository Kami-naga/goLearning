# Go Learning
- fast development, fast run and no technical debt!  
- Suitable for big data, microservice,concurrency programming.  
- no "object", no inheritance, no polymorphism, no generic programming, no try/catch...
- have interface, functional programming,   CSP (goroutine + channel)...
## Installation of GO 
- download from https://golang.org/
- after installation, type `go version` in cmd, then we can get the version number (go1.16.3 windows/amd64)
- if necessary, set the proxy by `GOPROXY` variable
- type `go env` to get the env variables
- first set the GO111MODULE to on(it's for dependency management)
  - `go env -w GO111MODULE=on`
    - be careful that it's "on", not "ON",or you'll get some bugs
  - (see here for details https://github.com/golang/go/wiki/Modules)
- get "goimports", it will help us do following 3 things when we save the go file
  - delete unnecessary space
  - add missing imports
  - sort the imports
  - `go get -v golang.org/x/tools/cmd/goimports`
## GO Configuration of IDE (Intellij) & helloworld
- settings - plugins
  - install GO
  - install Filw Watchers
    - for formatting
- new project - Go Modules - Next - set project name - complete
- new- go file - simple application - hello.go
  - add `fmt.Println("hello world")` in main function
  - `go run hello.go`
- Hello World!
  - parameter hints are not so useful in GO, so just close it in Settings-Editor-Inlay Hints-GO
  - add a goimports file watcher(Settings-Tools-File Watcher)
  - module name is in go.mod file (`go mod init MODULENAME`)

## GO dependency
- we use GOPATH，GOVENDOR before manage the dependencies by putting all the dependencies together in GOPATH
- when we need different versions of a lib for 2 projects, it's hard to do if we  use only GOPATH,so we have VENDOR in each project,those different things can be put into the vendor dir, and the same things still put in GOPATH
- above approaches have a strict restriction on path, it's not so convenient, so we have go mod now
    - `go get xxxxxx` to add a dependency,if you need specified version, add `@version` at last
    - `go mod tidy` to delete unused dependency & add imported dependency
    - when migrating from GOPATH&GOVENDOR to go mod, it's also easy,just use `go mod init MODNAME` and then `go mod tidy` to add them to the go mod file, then you can delete everything in the vendor file
## GO output
- use `go build ./...` to check if all the files can be built, but no output file
- use `go install ./...` to get an output, the output files will be in GOPATH(you can get it at `go env`)

## Interaface
- GO is an interface oriented language,no inheritance, no polymorphism, and it is interfaces that implement the functions that inheritance & polymorphism do in other languages. Interfaces in Go are very flexible.
- duck typing:describe the external behavior, not the internal structure
    - duck typing in Go is as flexible as python & c++, and also has a type check as java