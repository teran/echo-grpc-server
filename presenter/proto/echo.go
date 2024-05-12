package proto

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative echo.proto

func NewEchoResponse(s string) (*EchoResponse, error) {
	return &EchoResponse{
		Message: s,
	}, nil
}

func NewEchoRequest(s string) (*EchoRequest, error) {
	return &EchoRequest{
		Message: s,
	}, nil
}
