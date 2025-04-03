package consts

type Platform string

const (
	PlatformLark     Platform = "lark"
	PlatformDingTalk Platform = "dingtalk"
	PlatformWeCom    Platform = "wecom"
)

type LarkMessageType string

const (
	LarkMessageTypeText        LarkMessageType = "text"
	LarkMessageTypeInteractive LarkMessageType = "interactive"
)

type AlertStatus string

const (
	AlertStatusFiring   AlertStatus = "firing"
	AlertStatusResolved AlertStatus = "resolved"
)
