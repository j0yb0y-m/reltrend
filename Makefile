build:
	@mkdir -p ./bin
	@go build -o bin/reltrend

run: build
	@./bin/reltrend

clean:
	@rm -rf bin/