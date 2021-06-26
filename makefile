run:
	go run cmd/web/*.go

clean:
	rm dist/*

dist/server: cmd/web/*.go
	go build -o dist/server cmd/web/*.go

serve: dist/server
	dist/server

