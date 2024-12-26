package shardedpg

import (
	"context"
	"fmt"
	"github.com/godepo/elephant/internal/sharded"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Pool interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Transactional(ctx context.Context, fn func(ctx context.Context) error) (out error)
}

type Builder interface {
	ShardPicker(pickFn ShardPicker) Builder
	Shard(key byte, shard Pool) Builder
	Go() (*sharded.Sharded, error)
}

type ShardPicker interface {
	Pick(ctx context.Context, key string) byte
}

type builder struct {
	size   uint8
	shards map[byte]Pool
	picker ShardPicker
}

func New(poolSize uint8) Builder {
	return &builder{
		size:   poolSize,
		shards: make(map[byte]Pool, poolSize),
	}
}

func (b *builder) ShardPicker(pickFn ShardPicker) Builder {
	b.picker = pickFn
	return b
}

func (b *builder) Shard(key byte, shard Pool) Builder {
	b.shards[key] = shard
	return b
}

func (b *builder) Go() (*sharded.Sharded, error) {
	if b.size == 0 {
		return nil, fmt.Errorf("shards pool size must be greater than zero")
	}
	if b.picker == nil {
		return nil, fmt.Errorf("no shard picker provided")
	}
	shards := make([]sharded.Pool, 0, b.size)
	for key := range b.size {
		shard, ok := b.shards[key]
		if !ok {
			return nil, fmt.Errorf("shard not found for key %d", key)
		}
		shards = append(shards, shard)
	}
	return sharded.New(shards, b.picker), nil
}
