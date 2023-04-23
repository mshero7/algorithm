package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	root := NewNode("")

	success := Insert(root, "tea")
	assert.True(t, success)

	success = Insert(root, "ted")
	assert.True(t, success)

	success = Insert(root, "team")
	assert.True(t, success)

	success = Insert(root, "in")
	assert.True(t, success)

	success = Insert(root, "int")
	assert.True(t, success)
}

func TestAutoCompleteEng(t *testing.T) {
	root := NewNode("")

	success := Insert(root, "tea")
	assert.True(t, success)

	success = Insert(root, "ted")
	assert.True(t, success)

	success = Insert(root, "team")
	assert.True(t, success)

	success = Insert(root, "teaming")
	assert.True(t, success)

	success = Insert(root, "in")
	assert.True(t, success)

	success = Insert(root, "int")
	assert.True(t, success)

	res := AutoComplete(root, "teami")
	assert.Equal(t, "teaming", res)
}

func TestAutoCompleteKo(t *testing.T) {
	root := NewNode("")

	success := Insert(root, "한국")
	assert.True(t, success)

	success = Insert(root, "한국사람")
	assert.True(t, success)

	success = Insert(root, "한국사람 사랑해")
	assert.True(t, success)

	success = Insert(root, "한국사람 사랑해요우")
	assert.True(t, success)

	success = Insert(root, "한국사람과 같이")
	assert.True(t, success)

	res := AutoComplete(root, "한국사람 사랑해")
	assert.Equal(t, "한국사람 사랑해요우", res)
}
