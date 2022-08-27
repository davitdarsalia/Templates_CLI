# Windows - 64-Bit Arch -> 32-Bit Arch
GOOS=windows GOARCH=amd64 go build -o bin/templ.exe main.go
GOOS=windows GOARCH=386 go build -o bin/templ.exe main.go

# OSX - 64-Bit Arch -> 32-Bit Arch
GOOS=darwin GOARCH=amd64 go build -o bin/templ_osx main.go

# Linux - 64-Bit Arch -> 32-Bit Arch
GOOS=linux GOARCH=amd64 go build -o bin/templ main.go
GOOS=linux GOARCH=386 go build -o bin/templ main.go
