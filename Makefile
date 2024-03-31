push:
	@go build -o ./bin/push ./push
	@./bin/push

short-poll:
	@go build -o ./bin/short-poll ./short-poll
	@./bin/short-poll

long-poll:
	@go build -o ./bin/long-poll ./long-poll
	@./bin/long-poll
	
.PHONY: push short-poll long-poll