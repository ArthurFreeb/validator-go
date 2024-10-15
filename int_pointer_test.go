package schema_test

import (
	"testing"

	schema "github.com/Jamess-Lucass/validator-go"
	"github.com/stretchr/testify/assert"
)

func TestIntPointer_Type(t *testing.T) {
	s := schema.IntPointer()

	assert.True(t, s.Parse(newIntPointer(0)).IsValid())
	assert.True(t, s.Parse(newIntPointer(int(uint(10)))).IsValid())

	assert.False(t, s.Parse("123").IsValid())
	assert.False(t, s.Parse(nil).IsValid())
	assert.False(t, s.Parse(map[string]int{
		"one": 1,
		"two": 2,
	}).IsValid())
	assert.False(t, s.Parse([]int{1, 2, 3}).IsValid())
	assert.False(t, s.Parse(0.015).IsValid())

	assert.False(t, s.Parse(uint(10)).IsValid())
	assert.False(t, s.Parse(uint8(10)).IsValid())
	assert.False(t, s.Parse(uint16(10)).IsValid())
	assert.False(t, s.Parse(uint32(10)).IsValid())
	assert.False(t, s.Parse(uint64(10)).IsValid())

	assert.False(t, s.Parse(int8(10)).IsValid())
	assert.False(t, s.Parse(int16(10)).IsValid())
	assert.False(t, s.Parse(int32(10)).IsValid())
	assert.False(t, s.Parse(int64(10)).IsValid())
}

func TestIntPointer_Lt(t *testing.T) {
	s := schema.IntPointer().Lt(5)
	assert.True(t, s.Parse(newIntPointer(4)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-5)).IsValid())

	assert.False(t, s.Parse(newIntPointer(5)).IsValid())
	assert.False(t, s.Parse(newIntPointer(500)).IsValid())
}

func TestIntPointer_Lte(t *testing.T) {
	s := schema.IntPointer().Lte(5)

	assert.True(t, s.Parse(newIntPointer(4)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-5)).IsValid())
	assert.True(t, s.Parse(newIntPointer(5)).IsValid())

	assert.False(t, s.Parse(newIntPointer(6)).IsValid())
	assert.False(t, s.Parse(newIntPointer(500)).IsValid())
}

func TestIntPointer_Gt(t *testing.T) {
	s := schema.IntPointer().Gt(5)

	assert.True(t, s.Parse(newIntPointer(6)).IsValid())
	assert.True(t, s.Parse(newIntPointer(500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(5)).IsValid())
	assert.False(t, s.Parse(newIntPointer(-5)).IsValid())
}

func TestIntPointer_Gte(t *testing.T) {
	s := schema.IntPointer().Gte(5)

	assert.True(t, s.Parse(newIntPointer(5)).IsValid())
	assert.True(t, s.Parse(newIntPointer(6)).IsValid())
	assert.True(t, s.Parse(newIntPointer(500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(4)).IsValid())
	assert.False(t, s.Parse(newIntPointer(-5)).IsValid())
}

func TestIntPointer_Positive(t *testing.T) {
	s := schema.IntPointer().Positive()
	assert.True(t, s.Parse(newIntPointer(1)).IsValid())
	assert.True(t, s.Parse(newIntPointer(6)).IsValid())
	assert.True(t, s.Parse(newIntPointer(500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(0)).IsValid())
	assert.False(t, s.Parse(newIntPointer(-5)).IsValid())
}

func TestIntPointer_Nonnegative(t *testing.T) {
	s := schema.IntPointer().Nonnegative()

	assert.True(t, s.Parse(newIntPointer(0)).IsValid())
	assert.True(t, s.Parse(newIntPointer(6)).IsValid())
	assert.True(t, s.Parse(newIntPointer(500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(-5)).IsValid())
}

func TestIntPointer_Negative(t *testing.T) {
	s := schema.IntPointer().Negative()

	assert.True(t, s.Parse(newIntPointer(-1)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(0)).IsValid())
	assert.False(t, s.Parse(newIntPointer(5)).IsValid())
}

func TestIntPointer_Nonpositive(t *testing.T) {
	s := schema.IntPointer().Nonpositive()

	assert.True(t, s.Parse(newIntPointer(0)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-500)).IsValid())

	assert.False(t, s.Parse(newIntPointer(1)).IsValid())
	assert.False(t, s.Parse(newIntPointer(5)).IsValid())
}

func TestIntPointer_MultipleOf(t *testing.T) {
	s := schema.IntPointer().MultipleOf(5)

	assert.True(t, s.Parse(newIntPointer(0)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-5)).IsValid())
	assert.True(t, s.Parse(newIntPointer(-25)).IsValid())
	assert.True(t, s.Parse(newIntPointer(50)).IsValid())

	assert.False(t, s.Parse(newIntPointer(1)).IsValid())
	assert.False(t, s.Parse(newIntPointer(3)).IsValid())
}

func newIntPointer(value int) *int {
	return &value
}
