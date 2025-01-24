SOURCES=$(shell find . -type f -name "*.go")

all: run_bot
rebuild: clean app

run_bot: .PHONY
	@echo "-- Run Bot"
	@go run cmd/app/main.go

app: $(SOURCES)
	@echo "-- Run Bot"
	@go build -o app ./cmd/app/main.go

clean:
	@echo "-- Clean"
	@rm -f app

.PHONY: