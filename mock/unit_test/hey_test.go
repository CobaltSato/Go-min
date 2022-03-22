package main

import (
  "testing"
	"github.com/CobaltSato/mock/unit_test/hey"
)

// $ go test github.com/CobaltSato/mock/unit_test -v # gomod記載のパッケージ名を指定する
// $ go test . -v # もしくはカレントディレクトリ指定でもOK


func TestHeyDo(t *testing.T) {
  want := "こんにちは"
  got := hey.Do()
  if got != want {
    t.Errorf("want %q, got %q", want, got)
  }
}


