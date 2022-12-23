build-windows:
	GOOS=windows GOARCH=amd64 go build -o /mnt/c/Users/jotar/Desktop/scraper.exe cmd/main.go