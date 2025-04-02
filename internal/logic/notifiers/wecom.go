package notifiers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ledboot/OwlHook/internal/model"
)

// WeComNotifier implements the NotifierInterface for WeCom (Enterprise WeChat)
type WeComNotifier struct {
	Config *model.PlatformConfig
}

// NewWeComNotifier creates a new WeCom notifier instance
func NewWeComNotifier(config *model.PlatformConfig) *WeComNotifier {
	return &WeComNotifier{
		Config: config,
	}
}

type weComMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// Send implements NotifierInterface
func (n *WeComNotifier) Send(message *model.WebhookMessage) error {
	text := GenerateAlertMessage(message)

	wecomMsg := &weComMessage{
		MsgType: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: text,
		},
	}

	payload, err := json.Marshal(wecomMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal wecom message: %v", err)
	}

	resp, err := http.Post(n.Config.WebhookUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send wecom message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wecom API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
