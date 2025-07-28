default: start

start:
	- go run ./goweb/productsApi/cmd/main.go

development:
	- air -c .air.toml

cover:
	- go test ./... -cover -coverprofile=coverage.out

cover-html:
	- go tool cover -html=coverage.out -o coverage.html