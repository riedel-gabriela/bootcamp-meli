default: start

start:
	- go run ./goweb/productsApi/cmd/main.go

development:
	- air -c .air.toml