package notifiers

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/ledboot/OwlHook/internal/consts"
	"github.com/ledboot/OwlHook/internal/model"
	"github.com/ledboot/OwlHook/internal/service"
)

type sNotify struct {
	Provider map[consts.Platform]NotifierInterface
}

func Init(ctx context.Context) {
	s := &sNotify{}
	notifyCfgMap := map[consts.Platform]*model.PlatformConfig{}

	if err := g.Cfg().MustGet(ctx, "webhook").Scan(&notifyCfgMap); err != nil {
		g.Log().Fatal(ctx, "failed to scan webhook config: %v", err)
	}

	s.Provider = make(map[consts.Platform]NotifierInterface)

	for notifyType, cfg := range notifyCfgMap {
		if !cfg.Enabled {
			continue
		}
		switch notifyType {
		case consts.PlatformLark:
			s.Provider[notifyType] = NewLarkNotifier(cfg)
		case consts.PlatformDingTalk:
			s.Provider[notifyType] = NewDingTalkNotifier(cfg)
		case consts.PlatformWeCom:
			s.Provider[notifyType] = NewWeComNotifier(cfg)
		}
	}
	service.RegisterNotify(s)
}

func (s *sNotify) Send(ctx context.Context, platform consts.Platform, payload *model.Payload) error {
	if notifier, ok := s.Provider[platform]; ok {
		return notifier.Send(ctx, payload)
	}
	return fmt.Errorf("notifier not found: %s", platform)
}
