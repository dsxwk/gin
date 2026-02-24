package tests

import (
	"context"
	"gin/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisPubSub(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})

	err := config.GetRedisCache().Redis().WithContext(ctx).Subscribe("testRedisChan", func(channel, payload string) {
		assert.Equal(t, "testRedisChan", channel)
		assert.Contains(t, payload, "test")
		close(done)
	})
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	err = config.GetRedisCache().Redis().Publish("testRedisChan", map[string]interface{}{
		"test": "ok",
	})
	require.NoError(t, err)

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting pubsub")
	}
}
