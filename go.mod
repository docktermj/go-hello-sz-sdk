module github.com/docktermj/go-hello-sz-sdk

go 1.19

require (
	github.com/senzing/g2-sdk-go v0.2.5
	github.com/senzing/g2-sdk-go-grpc v0.0.0-20221215185305-4eeb7d9c3f13
	github.com/senzing/g2-sdk-proto/go v0.0.0-20230104150250-c9b2ad067374
	github.com/senzing/go-helpers v0.1.0
	google.golang.org/grpc v1.51.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/senzing/go-logging v1.1.1 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/senzing/g2-sdk-go-grpc v0.0.0-20221215185305-4eeb7d9c3f13 => /home/senzing/senzing.git/g2-sdk-go-grpc
