package netease

import (
	"encoding/json"
	"testing"
)

var client = CreateImClient("", "", "")

func TestToken(t *testing.T) {
	user := &ImUser{ID: "test1", Name: "test3", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestRefreshToken(t *testing.T) {
	tk, err := client.RefreshToken("7")
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(tk)
	t.Log(string(b), err)
}

func Benchmark_SyncMap(b *testing.B) {
	CreateImClient("", "", "")
}
