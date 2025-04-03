package notifiers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ledboot/OwlHook/internal/logic/notifiers/lark"
	"github.com/ledboot/OwlHook/internal/model"
)

// LarkNotifier implements the NotifierInterface for Lark (Feishu)
type LarkNotifier struct {
	Config *model.PlatformConfig
}

// NewLarkNotifier creates a new Lark notifier instance
func NewLarkNotifier(config *model.PlatformConfig) *LarkNotifier {
	return &LarkNotifier{
		Config: config,
	}
}

type larkMessage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

// Send implements NotifierInterface
func (n *LarkNotifier) Send(ctx context.Context, payload *model.Payload) error {
	// tmpl, ok := PlatformTemplate[string(consts.PlatformLark)]
	// if ok {
	// 	tmpl, err := template.New(templateName).Funcs(defaultFuncs).Parse(tmpl)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to parse lark template: %v", err)
	// 	}
	// 	tmpl.ExecuteTemplate()
	// } else {

	// }

	msg := lark.NewMsgInteractiveFromPayload(payload)

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal lark message: %v", err)
	}

	resp, err := http.Post(n.Config.WebhookUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send lark message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("lark API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
