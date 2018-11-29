package netease

import (
	"os"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	msg := &TextMessage{Message: "message test 1"}
	err := client.SendTextMessage("1", "169143", msg, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchTextMessage(t *testing.T) {
	msg := &TextMessage{Message: "message test"}
	str, err := client.SendBatchTextMessage("1", []string{"169143"}, msg, nil)
	t.Log(str)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchAttachMessage(t *testing.T) {
	err := client.SendBatchAttachMsg("1", "{'msg':'test'}", []string{"2", "3"}, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestBroadcastMsg(t *testing.T) {
	os.Setenv("GOCACHE", "off")
	t.Log(client.BroadcastMsg("好久不见了呢，我在这里等你哦", "", nil, nil))
}

func TestRecallMsg(t *testing.T) {
	t.Log(client.RecallMessage("456", "time", "from", "to", 7))
}
