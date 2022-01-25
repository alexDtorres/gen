build:
	go build -o bin/gen-app app/main.go

clean:
	trash bin
	trash testdata
