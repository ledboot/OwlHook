package notifiers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ledboot/OwlHook/internal/model"
)

// DingTalkNotifier implements the NotifierInterface for DingTalk
type DingTalkNotifier struct {
	Config *model.PlatformConfig
}

// NewDingTalkNotifier creates a new DingTalk notifier instance
func NewDingTalkNotifier(config *model.PlatformConfig) *DingTalkNotifier {
	return &DingTalkNotifier{
		Config: config,
	}
}

type dingTalkMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}

// generateSignature generates the signature for DingTalk webhook
func (n *DingTalkNotifier) generateSignature(timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, n.Config.Secret)
	mac := hmac.New(sha256.New, []byte(n.Config.Secret))
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// Send implements NotifierInterface
func (n *DingTalkNotifier) Send(message *model.WebhookMessage) error {
	text := GenerateAlertMessage(message)

	dingMsg := &dingTalkMessage{
		MsgType: "markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: fmt.Sprintf("Alert %s", message.Status),
			Text:  text,
		},
	}

	payload, err := json.Marshal(dingMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal dingtalk message: %v", err)
	}

	timestamp := time.Now().UnixNano() / 1e6
	signature := n.generateSignature(timestamp)

	// Add signature to webhook URL
	baseURL, err := url.Parse(n.Config.WebhookUrl)
	if err != nil {
		return fmt.Errorf("invalid webhook URL: %v", err)
	}

	q := baseURL.Query()
	q.Set("timestamp", fmt.Sprintf("%d", timestamp))
	q.Set("sign", signature)
	baseURL.RawQuery = q.Encode()

	resp, err := http.Post(baseURL.String(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send dingtalk message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("dingtalk API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
