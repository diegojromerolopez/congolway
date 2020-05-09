build: golgif randomgol

golgif:
	go build -o bin/golgif cmd/golgif/main.go

randomgol:
	go build -o bin/randomgol cmd/randomgol/main.go
