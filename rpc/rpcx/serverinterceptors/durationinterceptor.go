package serverinterceptors

import (
	"context"
	"encoding/json"
	"fmt"
	logx2 "github.com/sunmi-OS/gocore/rpc/rpcx/logx"
	xlog2 "github.com/sunmi-OS/gocore/utils/xlog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var serverSlowThreshold = (time.Millisecond * 500).Milliseconds()

// UnaryStatInterceptor
// 链路用时打印 and Panic拦截打印
func UnaryStatInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer handleCrash(func(r interface{}) {
			err = toPanicError(r)
		})
		startTime := time.Now()
		defer func() {
			duration := time.Since(startTime)
			//logDuration(ctx, info.FullMethod, req, duration)
			if duration.Milliseconds() > serverSlowThreshold {
				xlog2.Zap().Warnf("rpc-sever-slow: method: {%s}, duration: %dms, req: {%+v}", info.FullMethod, duration.Milliseconds(), req)
				//logx.LoggerObj.Info("rpc-sever-slow", map[string]string{"addr": addr, "method": method, "content": string(content), "duration": fmt.Sprintf("%d", duration/time.Millisecond)})
				return
			}

		}()
		return handler(ctx, req)
	}
}

// @desc 打印请求时间
// @auth liuguoqiang 2020-06-11
// @param
// @return
func logDuration(ctx context.Context, method string, req interface{}, duration time.Duration) {
	var addr string
	client, ok := peer.FromContext(ctx)
	if ok {
		addr = client.Addr.String()
	}
	content, err := json.Marshal(req)
	if err != nil {
		logx2.LoggerObj.Error("rpc-sever-fail", map[string]string{"addr": addr, "method": method, "content": err.Error(), "duration": fmt.Sprintf("%d", duration/time.Millisecond)})
		return
	}
	if duration.Milliseconds() > serverSlowThreshold {
		logx2.LoggerObj.Info("rpc-sever-slow", map[string]string{"addr": addr, "method": method, "content": string(content), "duration": fmt.Sprintf("%d", duration/time.Millisecond)})
		return
	}
	//logx.LoggerObj.Info("rpc-sever-call", map[string]string{"addr": addr, "method": method, "content": string(content), "duration": fmt.Sprintf("%d", duration/time.Millisecond)})
}
