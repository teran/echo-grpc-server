package proto

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative remote_echo.proto

func NewRemoteEchoResponse(s string) (*RemoteEchoResponse, error) {
	return &RemoteEchoResponse{
		Message: s,
	}, nil
}

func NewRemoteEchoRequest(remote, message string) (*RemoteEchoRequest, error) {
	return &RemoteEchoRequest{
		Remote:  remote,
		Message: message,
	}, nil
}
