build:
	npm run --prefix ./frontend build 

run:
	npm run --prefix ./frontend build
	go run ./cmd/server serve