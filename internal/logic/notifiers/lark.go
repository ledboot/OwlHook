package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
func (n *LarkNotifier) Send(message *model.WebhookMessage) error {
	text := GenerateAlertMessage(message)

	larkMsg := &larkMessage{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: text,
		},
	}

	payload, err := json.Marshal(larkMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal lark message: %v", err)
	}

	resp, err := http.Post(n.Config.WebhookUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send lark message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("lark API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
