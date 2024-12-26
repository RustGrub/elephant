package sharded

import "context"

type Picker func(ctx context.Context, key string) byte

func (s Picker) Pick(ctx context.Context, key string) byte {
	return (s)(ctx, key)
}
