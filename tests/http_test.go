package tests

import (
	"context"
	"gin/common/ctxkey"
	"gin/common/errcode"
	"gin/pkg"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "msg": "pong", "data": nil})
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, ctxkey.TraceIdKey, "test-trace-id")

	resp, err := pkg.HttpRequestJson[errcode.SuccessResponse](
		ctx,
		"GET",
		ts.URL+"/ping",
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "pong", resp.Msg)
}
