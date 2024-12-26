package sharded

import (
	"context"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardPickFn_Pick(t *testing.T) {
	t.Run("should be able to be able", func(t *testing.T) {
		fakeByte := faker.New().UInt8()
		picker := Picker(func(ctx context.Context, key string) byte {
			return fakeByte
		})
		actual := picker.Pick(context.Background(), faker.New().RandomStringWithLength(10))
		assert.Equal(t, fakeByte, actual)
	})
}
