package num2words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitByThousand(t *testing.T) {
	nums := []int64{123, 456, 789}
	assert.Equal(t, nums, splitByThousand(123456789))
}
