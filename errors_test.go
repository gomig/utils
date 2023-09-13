package utils_test

import (
	"testing"

	"github.com/gomig/utils"
)

func TestTaggedError(t *testing.T) {
	err := utils.TaggedError([]string{"Test", "TagB"}, "test %s", "error")
	if err.Error() != "[Test] [TagB] test error" {
		t.Log(err.Error())
		t.Fatal("failed!")
	}
}

func TestIsErrorOf(t *testing.T) {
	err := utils.TaggedError([]string{"Test", "TagB"}, "test %s", "error")
	if !utils.IsErrorOf("TagB", err) {
		t.Fatal("failed!")
	}
}

func TestHasError(t *testing.T) {
	err := utils.TaggedError([]string{"Test", "TagB"}, "test %s", "error")
	if utils.HasError(err) == false {
		t.Fatal("failed!")
	}
}
