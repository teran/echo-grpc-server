package proto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEchoRequest(t *testing.T) {
	r := require.New(t)

	req, err := NewEchoRequest("blah")
	r.NoError(err)
	r.Equal(&EchoRequest{
		Message: "blah",
	}, req)
}

func TestEchoResponse(t *testing.T) {
	r := require.New(t)

	resp, err := NewEchoResponse("blah")
	r.NoError(err)
	r.Equal(&EchoResponse{
		Message: "blah",
	}, resp)
}
