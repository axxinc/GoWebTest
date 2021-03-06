package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 2 {
		t.Error("Wrong id,was expecting 2 but got ", post.Id)
	}
	if post.Content != "GO" {
		t.Error("Wrong content,was expecting 'GO' but got ", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skiping encoding for now")
}

// func TestLongRunningTest(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skiping long-running test in short mode")
// 	}
// 	time.Sleep(time.Second * 10)
// }

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
