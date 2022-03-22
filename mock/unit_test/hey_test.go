package main

import (
	"fmt"
	"github.com/CobaltSato/mock/unit_test/hey"
)


func () Test (t *testing.T) { // 関数名?
  want := "こんにちは"
  go := hey.Do()
  if got != want {
    t.Errorf("want %q, got %q", want, got)
  }
}
