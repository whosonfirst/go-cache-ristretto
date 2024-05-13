package ristretto

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/whosonfirst/go-cache"
	"github.com/whosonfirst/go-ioutil"
)

func TestCacheMemcache(t *testing.T) {

	ctx := context.Background()

	c, err := cache.NewCache(ctx, "ristretto://")

	if err != nil {
		t.Fatalf("Failed to create new cache, %v", err)
	}

	sr := strings.NewReader("world")
	r, err := ioutil.NewReadSeekCloser(sr)

	if err != nil {
		t.Fatalf("Failed to create readseekcloser, %v", err)
	}

	_, err = c.Set(ctx, "hello", r)

	if err != nil {
		t.Fatalf("Failed to set key, %v", err)
	}

	r, err = c.Get(ctx, "hello")

	if err != nil {
		t.Fatalf("Failed to get key, %v", err)
	}

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read body, %v", err)
	}

	if string(body) != "world" {
		t.Fatalf("Unexpected value: '%s'", string(body))
	}

	err = c.Unset(ctx, "hello")

	if err != nil {
		t.Fatalf("Failed to unset key, %v", err)
	}
}
