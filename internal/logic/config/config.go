package config

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/ledboot/OwlHook/internal/model"
	"github.com/ledboot/OwlHook/internal/service"
)

type sConfig struct{}

func init() {
	service.RegisterConfig(&sConfig{})
}

func (s *sConfig) GetWebhookConfig(ctx context.Context) (webhookConfig map[string]*model.PlatformConfig, err error) {
	webhookConfig = make(map[string]*model.PlatformConfig)

	if err := g.Cfg().MustGet(ctx, "webhook").Scan(&webhookConfig); err != nil {
		return nil, err
	}

	return
}
