build:
	go build -o bin/gen-cli-app cli/app/main.go

clean:
	trash bin
