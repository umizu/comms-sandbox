push:
	@go build -o ./bin/push ./push
	@./bin/push

short-poll:
	@go build -o ./bin/short-poll ./short-poll
	@./bin/short-poll

long-poll:
	@go build -o ./bin/long-poll ./long-poll
	@./bin/long-poll
	
sse:
	@go build -o ./bin/sse ./server-sent-events
	@./bin/sse

ws:
	@go build -o ./bin/ws ./ws
	@./bin/ws

.PHONY: push short-poll long-poll ws
