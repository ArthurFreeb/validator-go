package schema_test

import (
	"testing"

	schema "github.com/Jamess-Lucass/validator-go"
	"github.com/stretchr/testify/assert"
)

func TestStringPointer_Type(t *testing.T) {
	s := schema.StringPointer()

	assert.True(t, s.Parse(newStringPointer("string")).IsValid())

	assert.False(t, s.Parse(123).IsValid())
	assert.False(t, s.Parse(nil).IsValid())
	assert.False(t, s.Parse(map[string]int{
		"one": 1,
		"two": 2,
	}).IsValid())
	assert.False(t, s.Parse([]int{1, 2, 3}).IsValid())
	assert.False(t, s.Parse(0).IsValid())
}

func TestStringPointer_Path(t *testing.T) {
	s := schema.StringPointer().Min(4).Parse("123")

	assert.Len(t, s.Errors, 1)
	assert.Equal(t, "", s.Errors[0].Path)
}

func TestStringPointer_Max(t *testing.T) {
	s := schema.StringPointer().Max(5)

	assert.True(t, s.Parse(newStringPointer("12345")).IsValid())
	assert.True(t, s.Parse(newStringPointer("1234")).IsValid())

	assert.False(t, s.Parse(newStringPointer("123456")).IsValid())
}

func TestStringPointer_Min(t *testing.T) {
	s := schema.StringPointer().Min(5)

	assert.True(t, s.Parse(newStringPointer("12345")).IsValid())
	assert.True(t, s.Parse(newStringPointer("123456")).IsValid())

	assert.False(t, s.Parse(newStringPointer("1234")).IsValid())
}

func TestStringPointer_Length(t *testing.T) {
	s := schema.StringPointer().Length(5)

	assert.True(t, s.Parse(newStringPointer("12345")).IsValid())

	assert.False(t, s.Parse(newStringPointer("123456")).IsValid())
	assert.False(t, s.Parse(newStringPointer("1234")).IsValid())
}

func TestStringPointer_Url(t *testing.T) {
	s := schema.StringPointer().Url()

	assert.True(t, s.Parse(newStringPointer("http://google.com")).IsValid())
	assert.True(t, s.Parse(newStringPointer("https://google.com/asdf?asdf=ljk3lk4&asdf=234#asdf")).IsValid())

	assert.False(t, s.Parse(newStringPointer("asdf")).IsValid())
	assert.False(t, s.Parse(newStringPointer("https:/")).IsValid())
	assert.False(t, s.Parse(newStringPointer("https")).IsValid())
	assert.False(t, s.Parse(newStringPointer("asdfj@lkjsdf.com")).IsValid())
}

func TestStringPointer_Includes(t *testing.T) {
	s := schema.StringPointer().Includes("test")

	assert.True(t, s.Parse(newStringPointer("X_test_X")).IsValid())
	assert.True(t, s.Parse(newStringPointer("test")).IsValid())

	assert.False(t, s.Parse(newStringPointer("Test")).IsValid())
	assert.False(t, s.Parse(newStringPointer("X_Test_X")).IsValid())
	assert.False(t, s.Parse(newStringPointer("TEST")).IsValid())
	assert.False(t, s.Parse(newStringPointer("3t3est")).IsValid())
}

func TestStringPointer_StartsWith(t *testing.T) {
	s := schema.StringPointer().StartsWith("test")

	assert.True(t, s.Parse(newStringPointer("test_X")).IsValid())
	assert.True(t, s.Parse(newStringPointer("test")).IsValid())

	assert.False(t, s.Parse(newStringPointer("Test")).IsValid())
	assert.False(t, s.Parse(newStringPointer("Test_X")).IsValid())
	assert.False(t, s.Parse(newStringPointer("TEST")).IsValid())
	assert.False(t, s.Parse(newStringPointer("teslt3")).IsValid())
}

func TestStringPointer_EndsWith(t *testing.T) {
	s := schema.StringPointer().EndsWith("test")

	assert.True(t, s.Parse(newStringPointer("X_test")).IsValid())
	assert.True(t, s.Parse(newStringPointer("test")).IsValid())

	assert.False(t, s.Parse(newStringPointer("Test")).IsValid())
	assert.False(t, s.Parse(newStringPointer("X_Test")).IsValid())
	assert.False(t, s.Parse(newStringPointer("TEST")).IsValid())
	assert.False(t, s.Parse(newStringPointer("3tes3t")).IsValid())
}

func newStringPointer(s string) *string {
	return &s
}
