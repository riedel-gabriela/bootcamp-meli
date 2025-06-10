default: start

start:
	- go run ./goweb/aula10get/cmd/main.go

development:
	- air -c .air.toml