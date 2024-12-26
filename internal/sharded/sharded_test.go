package sharded

import (
	"context"
	"errors"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFailedRow(t *testing.T) {
	t.Run("should be able to scan err", func(t *testing.T) {
		expErr := errors.New(faker.New().RandomStringWithLength(10))
		row := failedRow{
			err: expErr,
		}
		err := row.Scan()
		assert.Equal(t, expErr, err)
	})
}

func TestNew(t *testing.T) {
	t.Run("should be able to be able", func(t *testing.T) {
		mockPool := []Pool{NewMockPool(t), NewMockPool(t), NewMockPool(t)}
		mockPicker := NewMockshardPicker(t)
		sharded := New(mockPool, mockPicker)
		require.NotNil(t, sharded)
		assert.Equal(t, mockPool, sharded.shards)
		assert.Equal(t, mockPicker, sharded.shardPicker)
	})
}

func TestSharded_Begin(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(ArrangeContext).
			Then(AssertTxIsNil, AssertErrorAs(errCantPickShardID))

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.Begin(fxt.State.ctx)
	})
	t.Run("should be able to get shard by id and begin", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeTx,
			).
			When(ActBegin).
			Then(AssertTxAsExpected, AssertNoError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.Begin(fxt.State.ctx)
	})
	t.Run("should be able to get shard by id and begin and return error", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeExpectError,
			).
			When(ActBeginFailed).
			Then(AssertTxIsNil, AssertExpectedError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.Begin(fxt.State.ctx)
	})
	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeTx,
			).
			When(ActGotShardID, ActBegin).
			Then(AssertTxAsExpected, AssertNoError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.Begin(fxt.State.ctx)
	})
	t.Run("should be able to get shard by key and return err", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeExpectError,
			).
			When(ActGotShardID, ActBeginFailed).
			Then(AssertTxIsNil, AssertExpectedError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.Begin(fxt.State.ctx)
	})
}

func TestSharded_BeginTx(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(ArrangeContext).
			Then(AssertTxIsNil, AssertErrorAs(errCantPickShardID))

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
	t.Run("should be able to get shard by id and begin", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeTx, ArrangeTxOptions,
			).
			When(ActBeginTx).
			Then(AssertTxAsExpected, AssertNoError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
	t.Run("should be able to get shard by id and begin and return error", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeTxOptions,
				ArrangeExpectError,
			).
			When(ActBeginTxFailed).
			Then(AssertTxIsNil, AssertExpectedError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeTx, ArrangeTxOptions,
			).
			When(ActGotShardID, ActBeginTx).
			Then(AssertTxAsExpected, AssertNoError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
	t.Run("should be able to get shard by key and return err", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeTxOptions,
				ArrangeExpectError,
			).
			When(ActGotShardID, ActBeginTxFailed).
			Then(AssertTxIsNil, AssertExpectedError)

		fxt.State.Result.Tx, fxt.State.Result.Error = fxt.SUT.BeginTx(fxt.State.ctx, fxt.State.Expect.TxOptions)
	})
}

func TestSharded_Query(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext,
				ArrangeQuery, ArrangeArgs,
			).
			Then(AssertErrorAs(errCantPickShardID))

		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.SUT.Query(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by id", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs, ArrangeRows,
			).
			When(ActQuery).
			Then(AssertRows, AssertNoError)

		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.SUT.Query(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by id and return error", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs,
				ArrangeExpectError,
			).
			When(ActQueryFailed).
			Then(AssertExpectedError)

		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.SUT.Query(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs, ArrangeRows,
			).
			When(ActGotShardID, ActQuery).
			Then(AssertRows, AssertNoError)

		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.SUT.Query(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by key and return err", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs,
				ArrangeExpectError,
			).
			When(ActGotShardID, ActQueryFailed).
			Then(AssertExpectedError)

		fxt.State.Result.Rows, fxt.State.Result.Error =
			fxt.SUT.Query(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
}

func TestSharded_QueryRow(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext,
				ArrangeQuery, ArrangeArgs,
			).
			Then(AssertErrorAs(errCantPickShardID))

		fxt.State.Result.Row =
			fxt.SUT.QueryRow(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
		fxt.State.Result.Error = fxt.State.Result.Row.Scan()
	})
	t.Run("should be able to get shard by id", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs, ArrangeRow,
			).
			When(ActQueryRow).
			Then(AssertRow)

		fxt.State.Result.Row =
			fxt.SUT.QueryRow(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs, ArrangeRow,
			).
			When(ActGotShardID, ActQueryRow).
			Then(AssertRow)

		fxt.State.Result.Row =
			fxt.SUT.QueryRow(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
}

func TestSharded_Exec(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext,
				ArrangeQuery, ArrangeArgs,
			).
			Then(AssertErrorAs(errCantPickShardID))

		_, fxt.State.Result.Error =
			fxt.SUT.Exec(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by id", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs,
			).
			When(ActExec).
			Then(AssertNoError)

		_, fxt.State.Result.Error =
			fxt.SUT.Exec(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by id and return error", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
				ArrangeQuery, ArrangeArgs,
				ArrangeExpectError,
			).
			When(ActExecFailed).
			Then(AssertExpectedError)

		_, fxt.State.Result.Error =
			fxt.SUT.Exec(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs,
			).
			When(ActGotShardID, ActExec).
			Then(AssertNoError)

		_, fxt.State.Result.Error =
			fxt.SUT.Exec(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
	t.Run("should be able to get shard by key and return err", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
				ArrangeQuery, ArrangeArgs,
				ArrangeExpectError,
			).
			When(ActGotShardID, ActExecFailed).
			Then(AssertExpectedError)

		_, fxt.State.Result.Error =
			fxt.SUT.Exec(
				fxt.State.ctx,
				fxt.State.Expect.Query,
				fxt.State.Expect.Args...,
			)
	})
}

func TestSharded_Transactional(t *testing.T) {
	t.Run("should be able to return error if could not pick shard", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(ArrangeContext).
			Then(AssertErrorAs(errCantPickShardID))

		fxt.State.Result.Error =
			fxt.SUT.Transactional(
				fxt.State.ctx,
				func(ctx context.Context) error {
					return nil
				},
			)
	})
	t.Run("should be able to get shard by id", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardID,
			).
			When(ActTransactional).
			Then(AssertNoError)

		fxt.State.Result.Error =
			fxt.SUT.Transactional(
				fxt.State.ctx,
				func(ctx context.Context) error {
					return nil
				},
			)
	})

	t.Run("should be able to get shard by key", func(t *testing.T) {
		fxt := newTestCase(t)
		fxt.
			Given(
				ArrangeContext, ArrangeContextShardingKey,
			).
			When(ActGotShardID, ActTransactional).
			Then(AssertNoError)

		fxt.State.Result.Error =
			fxt.SUT.Transactional(
				fxt.State.ctx,
				func(ctx context.Context) error {
					return nil
				},
			)
	})
}
