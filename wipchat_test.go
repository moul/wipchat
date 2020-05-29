package wipchat

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientUnauthenticated(t *testing.T) {
	client := New("")
	ctx := context.Background()

	{
		ret, err := client.QueryViewer(ctx)
		require.Nil(t, ret)
		require.Equal(t, err, ErrTokenRequired)
	}

	{
		ret, err := client.MutateCreateTodo(ctx, "", nil, nil)
		require.Nil(t, ret)
		require.Equal(t, err, ErrTokenRequired)
	}

	{
		ret, err := client.uploadAttachment(ctx, Attachment{})
		require.Nil(t, ret)
		require.Equal(t, err, ErrTokenRequired)
	}
}
