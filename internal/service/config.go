// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/ledboot/OwlHook/internal/model"
)

type (
	IConfig interface {
		GetWebhookConfig(ctx context.Context) (webhookConfig map[string]*model.PlatformConfig, err error)
	}
)

var (
	localConfig IConfig
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}
