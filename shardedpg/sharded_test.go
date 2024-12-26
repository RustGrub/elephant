package shardedpg

import (
	"context"
	"fmt"
	"github.com/godepo/elephant/internal/sharded"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should be able to error if poolSize is 0", func(t *testing.T) {
		pool, err := New(0).Go()
		assert.Nil(t, pool)
		assert.ErrorContains(t, err, "shards pool size must be greater than zero")
	})
	t.Run("should be able to error if nil shard picker provided", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeNilShardPicker,
			).
			Then(AssertErrorAs(fmt.Errorf("no shard picker provided")))
		fxt.State.Result.ShardedPool, fxt.State.Result.Error = fxt.SUT.ShardPicker(fxt.State.shardPicker).Go()
	})
	t.Run("should be able to error if no shard picker provided", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Then(AssertErrorAs(fmt.Errorf("no shard picker provided")))
		fxt.State.Result.ShardedPool, fxt.State.Result.Error = fxt.SUT.ShardPicker(fxt.State.shardPicker).Go()
	})
	t.Run("should be able to error if one of shards not provided", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			Then(
				AssertErrorAs(fmt.Errorf("shard not found for key 2")),
				AssertNilShardedPool,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				ShardPicker(fxt.State.shardPicker).
				Go()
	})
	t.Run("should be able to error if one of shards not provided", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			Then(
				AssertErrorAs(fmt.Errorf("shard not found for key 1")),
				AssertNilShardedPool,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				ShardPicker(fxt.State.shardPicker).
				Go()
	})
	t.Run("should be able to error if one of shards not provided", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			Then(
				AssertErrorAs(fmt.Errorf("shard not found for key 0")),
				AssertNilShardedPool,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				ShardPicker(fxt.State.shardPicker).
				Go()
	})
	t.Run("should be able to query from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs, ArrangeRows,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActQuery,
			).
			Then(
				AssertNoError,
				AssertRows,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Query(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)
	})
	t.Run("should be able to query from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs, ArrangeRows,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActQuery).
			Then(
				AssertNoError,
				AssertRows,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Query(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)
	})
	t.Run("should be able to queryRow from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs, ArrangeRow,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActQueryRow,
			).
			Then(
				AssertNoError,
				AssertRow,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		fxt.State.Result.Row =
			fxt.State.Result.ShardedPool.QueryRow(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)

	})
	t.Run("should be able to queryRow from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs, ArrangeRow,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActQueryRow).
			Then(
				AssertNoError,
				AssertRow,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Row =
			fxt.State.Result.ShardedPool.QueryRow(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)
	})
	t.Run("should be able to exec from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActExec,
			).
			Then(
				AssertNoError,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		_, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Exec(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)

	})
	t.Run("should be able to exec from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActExec).
			Then(
				AssertNoError,
				AssertRow,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		_, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Exec(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)
	})
	t.Run("should be able to begin from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeTx,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActBegin,
			).
			Then(
				AssertNoError,
				AssertTxAsExpected,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Tx, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Begin(fxt.State.ctx)

	})
	t.Run("should be able to begin from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeTx,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActBegin).
			Then(
				AssertNoError,
				AssertRow,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Tx, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Begin(fxt.State.ctx)
	})
	t.Run("should be able to beginTx from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeTx, ArrangeTxOptions,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActBeginTx,
			).
			Then(
				AssertNoError,
				AssertTxAsExpected,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Tx, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)

	})
	t.Run("should be able to beginTx from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeTx, ArrangeTxOptions,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActBeginTx).
			Then(
				AssertNoError,
				AssertRow,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Tx, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
	t.Run("should be able to Transactional from shard by sharding key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(
				ActPickShard,
				ActTransactional,
			).
			Then(
				AssertNoError,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Transactional(fxt.State.ctx, func(ctx context.Context) error {
				return nil
			})

	})
	t.Run("should be able to Transactional from shard by sharding id in context", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeShardPicker(fxt.Deps.shardPicker),
			).
			When(ActTransactional).
			Then(
				AssertNoError,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Transactional(fxt.State.ctx, func(ctx context.Context) error {
				return nil
			})
	})
	t.Run("should be able to be able by sharding key With sharded.Picker", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs, ArrangeRows,
				ArrangeShardPicker(sharded.Picker(func(ctx context.Context, key string) byte {
					return fxt.State.shardID
				})),
			).
			When(
				ActQuery,
			).
			Then(
				AssertNoError,
				AssertRows,
			)
		fxt.State.Result.ShardedPool, fxt.State.Result.Error =
			fxt.SUT.
				Shard(0, fxt.State.shards[0]).
				Shard(1, fxt.State.shards[1]).
				Shard(2, fxt.State.shards[2]).
				ShardPicker(fxt.State.shardPicker).
				Go()
		require.NoError(t, fxt.State.Result.Error)
		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.State.Result.ShardedPool.Query(fxt.State.ctx, fxt.State.Expect.Query, fxt.State.Expect.Args...)
	})
}
