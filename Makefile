push:
	@go build -o ./bin/push ./push
	@./bin/push

short-poll:
	@go build -o ./bin/short-poll ./short-poll
	@./bin/short-poll
	
.PHONY: push short-poll