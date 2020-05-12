build: golgif golsvg randomgol

golgif:
	go build -o bin/golgif cmd/golgif/main.go

golsvg:
	go build -o bin/golsvg cmd/golsvg/main.go

randomgol:
	go build -o bin/randomgol cmd/randomgol/main.go
