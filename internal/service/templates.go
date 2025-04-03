// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/ledboot/OwlHook/internal/model"
)

type (
	ITemplate interface {
		GetTemplate(platform string) (*model.SafeTemplate, error)
	}
)

var (
	localTemplate ITemplate
)

func Template() ITemplate {
	if localTemplate == nil {
		panic("implement not found for interface ITemplate, forgot register?")
	}
	return localTemplate
}

func RegisterTemplate(i ITemplate) {
	localTemplate = i
}
