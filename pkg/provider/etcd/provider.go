package etcd

import (
	"context"
	"fmt"
	"log"

	"github.com/lzakharov/remote-config-manager/pkg/provider"
	"go.etcd.io/etcd/client/v3"
)

// Provider is a remote config provider.
type Provider struct {
	client *clientv3.Client
}

// NewProvider creates a new remote config provider.
func NewProvider(client *clientv3.Client) *Provider {
	return &Provider{client: client}
}

// ListKeys lists keys.
func (p *Provider) ListKeys(ctx context.Context) ([]string, error) {
	resp, err := p.client.Get(ctx, "",
		clientv3.WithPrefix(),
		clientv3.WithKeysOnly())
	if err != nil {
		log.Printf("get keys: %v", err)
		return nil, fmt.Errorf("get keys: %w", provider.ErrInternal)
	}

	keys := make([]string, len(resp.Kvs))
	for i, k := range resp.Kvs {
		keys[i] = string(k.Key)
	}

	return keys, nil
}

// Get gets value by key.
func (p *Provider) Get(ctx context.Context, key string) (string, error) {
	resp, err := p.client.Get(ctx, key)
	if err != nil {
		log.Printf("get %q: %v", key, err)
		return "", fmt.Errorf("get: %w", provider.ErrInternal)
	}

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("get: %w", provider.ErrNotFound)
	}

	return string(resp.Kvs[0].Value), nil
}

// Put puts value by key.
func (p *Provider) Put(ctx context.Context, key, value string) error {
	_, err := p.client.Put(ctx, key, value)
	if err != nil {
		log.Printf("put %q with %q: %v", key, value, err)
		return fmt.Errorf("put: %w", provider.ErrInternal)
	}

	return nil
}
