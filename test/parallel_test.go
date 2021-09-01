package main

import (
	"testing"
	"time"
)

func TestParaller1(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestParaller2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParaller3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
