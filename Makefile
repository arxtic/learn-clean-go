install:
	go mod download

dev:
	air -c .air.toml

build:
	go build -o ./dist/main ./app

start:
	chmod +x ./dist/main
	./dist/main