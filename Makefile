build:
	@mkdir -p ./bin
	@go build -o bin/reltrend

run: build
	@./bin/reltrend

clean:
	@rm -rf bin/

setup:
	go get github.com/charmbracelet/lipgloss
	go mod vendor
	go mod tidy