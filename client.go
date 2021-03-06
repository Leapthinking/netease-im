package netease

import (
	"gopkg.in/resty.v1"
	"github.com/json-iterator/go"
	"strconv"
	"time"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

//ImClient .
type ImClient struct {
	appKey    string
	appSecret string
	nonce     string

	client *resty.Client
}

type ClientOpt interface {
	Apply(*resty.Client)
}

//CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appkey string, appSecret string, opts ...ClientOpt) *ImClient {
	c := &ImClient{appKey: appkey, appSecret: appSecret, nonce: RandStringBytesMaskImprSrc(64)}
	client := resty.New()

	client.SetHeader("Accept", "application/json;charset=utf-8").
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;").
		SetHeader("AppKey", c.appKey).
		SetHeader("Nonce", c.nonce).
		SetPreRequestHook(checksumHook(appSecret, c.nonce)).
		SetTimeout(5 * time.Second)
	for _, o := range opts {
		o.Apply(client)
	}
	c.client = client
	return c
}

func checksumHook(appSecret string, nonce string) func(*resty.Client, *resty.Request) error {
	return func(_ *resty.Client, req *resty.Request) error {
		timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
		req.SetHeader("CurTime", timeStamp)
		req.SetHeader("CheckSum", ShaHashToHexStringFromString(appSecret+nonce+timeStamp))
		return nil
	}
}
