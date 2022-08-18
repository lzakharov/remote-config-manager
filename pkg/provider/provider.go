package provider

import (
	"context"
)

// Provider is a remote config provider.
type Provider interface {
	ListKeys(ctx context.Context) ([]string, error)
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key, value string) error
}
