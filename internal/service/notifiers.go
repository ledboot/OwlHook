// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/ledboot/OwlHook/internal/consts"
	"github.com/ledboot/OwlHook/internal/model"
)

type (
	INotify interface {
		Send(ctx context.Context, platform consts.Platform, message *model.WebhookMessage) error
	}
)

var (
	localNotify INotify
)

func Notify() INotify {
	if localNotify == nil {
		panic("implement not found for interface INotify, forgot register?")
	}
	return localNotify
}

func RegisterNotify(i INotify) {
	localNotify = i
}
