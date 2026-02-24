package tests

import (
	"context"
	rc "gin/app/queue/rabbitmq/consumer"
	p "gin/app/queue/rabbitmq/producer"
	"gin/common/ctxkey"
	"gin/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRabbitMQPublish(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(context.Background(), ctxkey.TraceIdKey, "test-trace-id")

	if config.GetConfig().Rabbitmq.Enabled {
		rc.NewRabbitmqDemoConsumer()
		pub := p.NewRabbitmqDemoPublisher()
		defer func() {
			require.NoError(t, pub.Close())
		}()

		err := pub.Publish(ctx, []byte(`{"orderId":333}`))
		require.NoError(t, err)
	}
}
