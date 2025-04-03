package templates

import (
	"context"
	"fmt"

	_ "embed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/ledboot/OwlHook/internal/model"
	"github.com/ledboot/OwlHook/internal/service"
)

//go:embed tmpl/lark.tmpl
var defaultTmplLark string

//go:embed tmpl/dingtalk.tmpl
var defaultTmplDingTalk string

//go:embed tmpl/wecom.tmpl
var defaultTmplWecom string

var platformTemplate = map[string]string{
	"lark":     defaultTmplLark,
	"dingtalk": defaultTmplDingTalk,
	"wecom":    defaultTmplWecom,
}

type sTemplate struct {
	promMsgTemplatesMap map[string]*model.SafeTemplate
}

func init() {
	s := &sTemplate{
		promMsgTemplatesMap: make(map[string]*model.SafeTemplate),
	}

	for k, v := range platformTemplate {
		t := &model.SafeTemplate{}
		if err := t.UpdateTemplate(v); err != nil {
			g.Log().Fatalf(context.Background(), "UpdateTemplate for (%s) failed, err: %s", k, err)
		}
		s.promMsgTemplatesMap[k] = t
	}
	service.RegisterTemplate(s)
}

func (s *sTemplate) GetTemplate(platform string) (*model.SafeTemplate, error) {
	t, ok := s.promMsgTemplatesMap[platform]
	if !ok {
		return nil, fmt.Errorf("template for (%s) not found", platform)
	}
	return t, nil
}

// // AlertLevel returns the severity level of the alert
// func AlertLevel(labels map[string]string) string {
// 	if severity, ok := labels["severity"]; ok {
// 		return strings.ToUpper(severity)
// 	}
// 	return "UNKNOWN"
// }

// // FormatTime formats time to a readable string
// func FormatTime(t time.Time) string {
// 	return t.Format("2006-01-02 15:04:05")
// }

// // FormatDuration formats the duration between two times
// func FormatDuration(start, end time.Time) string {
// 	duration := end.Sub(start)
// 	return duration.String()
// }

// // GenerateAlertMessage generates a common alert message format
// func GenerateAlertMessage(msg *model.WebhookMessage) string {
// 	var builder strings.Builder

// 	// Write header
// 	builder.WriteString(fmt.Sprintf("🔔 **Alert %s**\n\n", msg.Status))

// 	// Write summary
// 	if summary, ok := msg.CommonAnnotations["summary"]; ok {
// 		builder.WriteString(fmt.Sprintf("**Summary:** %s\n", summary))
// 	}

// 	// Write description
// 	if description, ok := msg.CommonAnnotations["description"]; ok {
// 		builder.WriteString(fmt.Sprintf("**Description:** %s\n", description))
// 	}

// 	// Write alerts details
// 	builder.WriteString(fmt.Sprintf("\n**Number of Alerts:** %d", len(msg.Alerts)))
// 	if msg.TruncatedAlerts > 0 {
// 		builder.WriteString(fmt.Sprintf(" (truncated: %d)", msg.TruncatedAlerts))
// 	}
// 	builder.WriteString("\n\n")

// 	// Write each alert
// 	for i, alert := range msg.Alerts {
// 		builder.WriteString(fmt.Sprintf("**Alert %d:**\n", i+1))
// 		builder.WriteString(fmt.Sprintf("- Status: %s\n", alert.Status))
// 		builder.WriteString(fmt.Sprintf("- Severity: %s\n", AlertLevel(alert.Labels)))
// 		builder.WriteString(fmt.Sprintf("- Start: %s\n", FormatTime(alert.StartsAt)))
// 		if !alert.EndsAt.IsZero() {
// 			builder.WriteString(fmt.Sprintf("- End: %s\n", FormatTime(alert.EndsAt)))
// 			builder.WriteString(fmt.Sprintf("- Duration: %s\n", FormatDuration(alert.StartsAt, alert.EndsAt)))
// 		}
// 		builder.WriteString("\n")
// 	}

// 	// Write footer
// 	if msg.ExternalURL != "" {
// 		builder.WriteString(fmt.Sprintf("🔗 [View in Alertmanager](%s)\n", msg.ExternalURL))
// 	}

// 	return builder.String()
// }
