package model

type PlatformConfig struct {
	WebhookUrl string `json:"webhookUrl"`
	Secret     string `json:"secret"`
	Enabled    bool   `json:"enabled"`
}
