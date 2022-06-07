package gocache

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHTTP(t *testing.T) {
	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			t.Log("[gocache] search key", key)
			if v, ok := fakeDb[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := NewHTTPPool(addr)
	t.Log("gocache is running at", addr)
	t.Fatal(http.ListenAndServe(addr, peers))
}
