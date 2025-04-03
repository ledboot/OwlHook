package model

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ledboot/OwlHook/internal/consts"
)

// WebhookMessage represents the webhook message from Alertmanager
type WebhookMessage struct {
	Version           string            `json:"version"`
	GroupKey          map[string]string `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`

	// extra fields added by us
	MessageAt time.Time `json:"messageAt"` // the time the webhook message was received
}

// Alert represents a single alert from Alertmanager
type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

func RenderTmpl(tmpl *SafeTemplate, tmplName string, data interface{}) (string, error) {
	tmplClone, err := tmpl.Clone()
	if err != nil {
		return "", fmt.Errorf("clone template failed, err: %v", err)
	}
	var buf bytes.Buffer
	if err := tmplClone.ExecuteTemplate(&buf, tmplName, data); err != nil {
		return "", fmt.Errorf("execute template failed, err: %v", err)
	}
	return buf.String(), nil
}

func ToPayload(tmpl *SafeTemplate, alertData *WebhookMessage) (*Payload, error) {
	alertData.MessageAt = time.Now()
	payload := &Payload{
		AlertStatus: consts.AlertStatus(alertData.Status),
	}

	title, err := RenderTmpl(tmpl, "prom.title", alertData)
	if err != nil {
		return nil, err
	}
	payload.Title = title

	text, err := RenderTmpl(tmpl, "prom.text", alertData)
	if err != nil {
		return nil, err
	}
	payload.Text = text

	markdown, err := RenderTmpl(tmpl, "prom.markdown", alertData)
	if err != nil {
		return nil, err
	}
	payload.Markdown = markdown

	return payload, nil
}
