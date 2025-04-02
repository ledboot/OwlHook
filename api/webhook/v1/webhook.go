package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/ledboot/OwlHook/internal/consts"
	"github.com/ledboot/OwlHook/internal/model"
)

type WebhookReq struct {
	g.Meta         `path:"/webhook/{platform}" method:"POST" summary:"Webhook" tags:"Webhook"`
	Platform       consts.Platform `v:"enums" path:"platform"`
	WebhookMessage *model.WebhookMessage
}

type WebhookRes struct {
}
