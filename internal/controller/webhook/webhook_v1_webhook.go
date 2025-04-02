package webhook

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/ledboot/OwlHook/api/webhook/v1"
	"github.com/ledboot/OwlHook/internal/service"
)

func (c *ControllerV1) Webhook(ctx context.Context, req *v1.WebhookReq) (res *v1.WebhookRes, err error) {
	g.Log().Infof(ctx, "webhook request received: %v", req)
	if err := service.Notify().Send(ctx, req.Platform, req.WebhookMessage); err != nil {
		return nil, err
	}
	return
}
