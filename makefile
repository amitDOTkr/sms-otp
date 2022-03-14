docker:
	CGO_ENABLED=0 go build main.go && \
	docker build -t amitdotkr/sms-otp:dev . && \
	docker push amitdotkr/sms-otp:dev && \
	rm main

gen:
	protoc -I=. --go_out=. ./src/protos/*.proto

pb:
	protoc --go_out=. --go-grpc_out=. ./src/protos/*.proto