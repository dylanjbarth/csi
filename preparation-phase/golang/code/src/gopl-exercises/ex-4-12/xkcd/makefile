test: 
	go test ./...

.PHONY: extract
extract: 
	go run main.go --extract

.PHONY: force-extract
force-extract: 
	go run main.go --extract --force

.PHONY: index
index: 
	go run main.go --index

.PHONY: search
search: 
	go run main.go "barrel"