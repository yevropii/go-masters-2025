package server

import (
	"context"
	"time"
)

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if stdSrv != nil {
		stdSrv.Shutdown(ctx)
	}
	if fastSrv != nil {
		fastSrv.Shutdown()
	}
}
