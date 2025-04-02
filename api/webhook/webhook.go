// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package webhook

import (
	"context"

	"github.com/ledboot/OwlHook/api/webhook/v1"
)

type IWebhookV1 interface {
	Webhook(ctx context.Context, req *v1.WebhookReq) (res *v1.WebhookRes, err error)
}
