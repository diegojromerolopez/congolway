build: golgif golsvg golapng randomgol golconv

golgif:
	go build -o bin/golgif cmd/golgif/main.go

golsvg:
	go build -o bin/golsvg cmd/golsvg/main.go

golapng:
	go build -o bin/golapng cmd/golapng/main.go

randomgol:
	go build -o bin/randomgol cmd/randomgol/main.go

golconv:
	go build -o bin/golconv cmd/golconv/main.go

clean:
	rm -rf bin/golgif
	rm -rf bin/golsvg
	rm -rf bin/golapng
	rm -rf bin/golconv
