install:
	@echo "-----------------------"
	@cd assetsbuild && npm install
	go mod download
	go get -d github.com/go-bindata/go-bindata/...
	@echo "-----------------------"

build-assets:
	@echo "-----------------------"
	@cd assetsbuild && npx tailwindcss -i ./css/input.css -o ../public/assets/css/index.css
	go-bindata ./public/assets/... 
	@echo "-----------------------"

run-local:
	ENV=local go run main.go bindata.go

run-prod:
	ENV=local go run main.go bindata.go

build:
	go build -o build main.go bindata.go

generate_mocks:
	go generate ./...

run-test:
	go test ./...