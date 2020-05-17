build: golstdout golgif golsvg golapng randomgol golconv

golstdout:
	go build -o bin/golstdout cmd/golstdout/main.go

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

tests:
	go test -v ./...

clean:
	rm -rf bin/golstdout
	rm -rf bin/golgif
	rm -rf bin/golsvg
	rm -rf bin/golapng
	rm -rf bin/golconv
