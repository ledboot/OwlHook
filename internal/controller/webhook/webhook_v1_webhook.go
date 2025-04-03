package webhook

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/ledboot/OwlHook/api/webhook/v1"
	"github.com/ledboot/OwlHook/internal/model"
	"github.com/ledboot/OwlHook/internal/service"
)

func (c *ControllerV1) Webhook(ctx context.Context, req *v1.WebhookReq) (res *v1.WebhookRes, err error) {
	bytes := g.RequestFromCtx(ctx).GetBody()
	g.Log().Infof(ctx, "webhook request received: %v", string(bytes))
	safeTemplate, err := service.Template().GetTemplate(string(req.Platform))
	if err != nil {
		return nil, err
	}
	payload, err := model.ToPayload(safeTemplate, req.WebhookMessage)
	if err != nil {
		return nil, err
	}

	if err := service.Notify().Send(ctx, req.Platform, payload); err != nil {
		return nil, err
	}
	return
}
