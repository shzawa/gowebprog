package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// $ gotest -v -short -parallel 3
// -parallel 3: 最大3つのテストケースを並列実行
