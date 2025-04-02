package notifiers

import (
	"fmt"
	"strings"
	"time"

	"github.com/ledboot/OwlHook/internal/model"
)

// AlertLevel returns the severity level of the alert
func AlertLevel(labels map[string]string) string {
	if severity, ok := labels["severity"]; ok {
		return strings.ToUpper(severity)
	}
	return "UNKNOWN"
}

// FormatTime formats time to a readable string
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDuration formats the duration between two times
func FormatDuration(start, end time.Time) string {
	duration := end.Sub(start)
	return duration.String()
}

// GenerateAlertMessage generates a common alert message format
func GenerateAlertMessage(msg *model.WebhookMessage) string {
	var builder strings.Builder

	// Write header
	builder.WriteString(fmt.Sprintf("🔔 **Alert %s**\n\n", msg.Status))

	// Write summary
	if summary, ok := msg.CommonAnnotations["summary"]; ok {
		builder.WriteString(fmt.Sprintf("**Summary:** %s\n", summary))
	}

	// Write description
	if description, ok := msg.CommonAnnotations["description"]; ok {
		builder.WriteString(fmt.Sprintf("**Description:** %s\n", description))
	}

	// Write alerts details
	builder.WriteString(fmt.Sprintf("\n**Number of Alerts:** %d", len(msg.Alerts)))
	if msg.TruncatedAlerts > 0 {
		builder.WriteString(fmt.Sprintf(" (truncated: %d)", msg.TruncatedAlerts))
	}
	builder.WriteString("\n\n")

	// Write each alert
	for i, alert := range msg.Alerts {
		builder.WriteString(fmt.Sprintf("**Alert %d:**\n", i+1))
		builder.WriteString(fmt.Sprintf("- Status: %s\n", alert.Status))
		builder.WriteString(fmt.Sprintf("- Severity: %s\n", AlertLevel(alert.Labels)))
		builder.WriteString(fmt.Sprintf("- Start: %s\n", FormatTime(alert.StartsAt)))
		if !alert.EndsAt.IsZero() {
			builder.WriteString(fmt.Sprintf("- End: %s\n", FormatTime(alert.EndsAt)))
			builder.WriteString(fmt.Sprintf("- Duration: %s\n", FormatDuration(alert.StartsAt, alert.EndsAt)))
		}
		builder.WriteString("\n")
	}

	// Write footer
	if msg.ExternalURL != "" {
		builder.WriteString(fmt.Sprintf("🔗 [View in Alertmanager](%s)\n", msg.ExternalURL))
	}

	return builder.String()
}
