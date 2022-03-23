package main

import (
	"testing"

	"github.com/CobaltSato/mock/unit_test/hey"
)

// # gomod記載のパッケージ名を指定する
// go test github.com/CobaltSato/mock/unit_test -v
// # もしくはカレントディレクトリ指定でもOK
// go test . -v

// サブテスト
// go test github.com/CobaltSato/mock/unit_test -run TestHeyDo/A  -v
// go test github.com/CobaltSato/mock/unit_test -run TestHeyDo/B  -v

func TestHeyDo(t *testing.T) {
	t.Run("A", func(t *testing.T) {
  	want := "こんにちは"
  	got := hey.Do()
  	if got != want {
  	  t.Errorf("want %q, got %q", want, got)
  	}
	})
	t.Run("B", func(t *testing.T) {
  	want := "hello"
  	got := hey.Do()
  	if got != want {
  	  t.Errorf("want %q, got %q", want, got)
  	}
	})
}


