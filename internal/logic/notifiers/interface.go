package notifiers

import "github.com/ledboot/OwlHook/internal/model"

// NotifierInterface defines the interface for different chat platform notifiers
type NotifierInterface interface {
	Send(message *model.WebhookMessage) error
}
