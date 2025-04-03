package lark

import (
	"github.com/ledboot/OwlHook/internal/consts"
	"github.com/ledboot/OwlHook/internal/logic/notifiers/lark/card"
	"github.com/ledboot/OwlHook/internal/model"
)

type Msg struct {
	MsgType consts.LarkMessageType `json:"msg_type"`

	Content *Content `json:"content,omitempty"`

	Card *Card `json:"card,omitempty"`
}

type Content struct {
	Text string `json:"text,omitempty"`
}

type Card struct {
	Config   *CardConfig       `json:"config,omitempty"`
	Header   *CardHeader       `json:"header,omitempty"`
	CardLink *card.MultiURL    `json:"card_link,omitempty"`
	Elements []card.CardModule `json:"elements"` // 最多可堆叠 50 个模块
}

// CardConfig 卡片配置
type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode,omitempty"` // 2021/03/22 之后，此字段废弃，所有卡片均升级为自适应屏幕宽度的宽版卡片
	EnableForward  bool `json:"enable_forward"`             // 是否允许卡片被转发，默认 false
}

type CardHeader struct {
	Title    *card.Text `json:"title"`              // 卡片标题内容, text 对象（仅支持 "plain_text")
	Template string     `json:"template,omitempty"` // 控制标题背景颜色, https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN
}

func NewMsgInteractive(card *Card) *Msg {
	return &Msg{
		MsgType: consts.LarkMessageTypeInteractive,
		Card:    card,
	}
}

func NewMsgInteractiveFromPayload(payload *model.Payload) *Msg {
	card := &Card{
		Elements: []card.CardModule{
			&card.ModuleDiv{
				Tag: "div",
				Text: &card.Text{
					Content: payload.Markdown,
					Tag:     "lark_md",
				},
			},
		},
		Header: &CardHeader{
			Title: &card.Text{
				Content: payload.Title,
				Tag:     "plain_text",
			},
		},
	}
	switch payload.AlertStatus {
	case consts.AlertStatusFiring:
		card.Header.Template = "red"

	case consts.AlertStatusResolved:
		card.Header.Template = "green"
	}

	return NewMsgInteractive(card)
}
