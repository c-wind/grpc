module github.com/c-wind/grpc

go 1.14

require (
	github.com/golang/protobuf v1.3.5
	golang.org/x/net v0.0.0-20200319234117-63522dbf7eec // indirect
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200319113533-08878b785e9c // indirect
	google.golang.org/grpc v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
