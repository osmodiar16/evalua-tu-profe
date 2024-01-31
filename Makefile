SHELL := /usr/bin/zsh
run:
	@templ generate
	@go run cmd/main.go

