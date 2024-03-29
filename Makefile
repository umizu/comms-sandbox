push:
	@go build -o ./bin/push ./push
	@./bin/push
	
.PHONY: push