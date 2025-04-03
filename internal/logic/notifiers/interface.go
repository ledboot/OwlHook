package notifiers

import (
	"context"

	"github.com/ledboot/OwlHook/internal/model"
)

// NotifierInterface defines the interface for different chat platform notifiers
type NotifierInterface interface {
	Send(ctx context.Context, payload *model.Payload) error
}
