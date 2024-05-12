package proto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoteEchoRequest(t *testing.T) {
	r := require.New(t)

	req, err := NewRemoteEchoRequest("test", "blah")
	r.NoError(err)
	r.Equal(&RemoteEchoRequest{
		Remote:  "test",
		Message: "blah",
	}, req)
}

func TestRemoteEchoResponse(t *testing.T) {
	r := require.New(t)

	resp, err := NewRemoteEchoResponse("blah")
	r.NoError(err)
	r.Equal(&RemoteEchoResponse{
		Message: "blah",
	}, resp)
}
