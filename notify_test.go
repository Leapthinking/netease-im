package netease

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCheckSum(t *testing.T) {
	os.Setenv("GOCACHE", "off")
	cli := CreateImClient("b2c60dbed0ae2d3c48e6c85664836dc9", "1ed04f7d7085", "")

	body := []byte(`{}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := Md5HashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", ShaHashToHexStringFromString(cli.appSecret+md5+curTime))
	t.Log("checksum:", cli.appSecret+md5+curTime, "checksum-encoded:", ShaHashToHexStringFromString(cli.appSecret+md5+curTime))

	bd, err := cli.GetEventNotification(req)
	t.Log(string(bd), err)
}
