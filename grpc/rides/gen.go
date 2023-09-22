package rides

/* Get the Go gRPC tools by running

	$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

You need the tools to be in your path:

	$ export PATH="$PATH:$(go env GOPATH)/bin"

See more at https://grpc.io/docs/languages/go/quickstart/
*/

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative ./rides.proto

//go:generate grpcurl --plaintext localhost:9292 list
//go:generate grpcurl --plaintext localhost:9292 describe .StartRequest

//go:generate grpcurl --plaintext -d @ localhost:9292 Rides.Start < start.json

//go:generate grpcurl --plaintext -d @ localhost:9292 Rides.End < end.json

/*
The most simple gRPC call is a unary call. This type of call accepts a single request message and returns a single response message.
*/

//go:generate go test -v ./cmd/server/

//go:generate protoc --grpc-gateway_out=./pb --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true ./rides.proto

//go:generate curl -d@end.json http://localhost:8080/Rides/End
