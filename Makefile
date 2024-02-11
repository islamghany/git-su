

.PHONY: run
run:
	@go run main.go
	
.PHONY: build/win-x64
build/win-x64:
	@echo "Building..."
	GOOS=windows GOARCH=amd64 go build -ldflags="-s" -o bin/win-x64/ ./...
	@echo "Done."

.PHONY: build/win-x86
build/win-x86:
	@echo "Building..."
	GOOS=windows GOARCH=386 go build -ldflags="-s" -o bin/win-x86/ ./...
	@echo "Done."

.PHONY: build/linux-x64
build/linux-x64:
	@echo "Building..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o	bin/linux-x64/ ./...
	@echo "Done."

.PHONY: build/linux-x86
build/linux-x86:
	@echo "Building..."
	GOOS=linux GOARCH=386 go build -ldflags="-s" -o bin/linux-x86/ ./...
	@echo "Done."

.PHONY: build
build: build/win-x64 build/win-x86 build/linux-x64 build/linux-x86
	@echo "Building..."
	@echo "Done."
	
.PHONY: install/linux-x64
install/linux-x64:
	@echo "Installing..."
	@sudo cp bin/linux-x64/* /usr/local/bin/
	@echo "Done."

.PHONY: install/linux-x86
install/linux-x86:
	@echo "Installing..."
	@sudo cp bin/linux-x86/* /usr/local/bin/
	@echo "Done."


.PHONY: install/win-x64
install/win-x64:
	@echo "Installing..."
	@cp bin/win-x64/* /usr/local/bin/
	@echo "Done."

.PHONY: install/win-x86
install/win-x86:
	@echo "Installing..."
	@cp bin/win-x86/* /usr/local/bin/
	@echo "Done."


.PHONY: clear
clear:
	@echo "Clearing..."
	@rm -rf bin/*
	@echo "Done."
	