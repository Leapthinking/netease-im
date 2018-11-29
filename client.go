package netease

import (
	"github.com/go-resty/resty"
	"github.com/json-iterator/go"
	"strconv"
	"time"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

//ImClient .
type ImClient struct {
	AppKey    string
	AppSecret string
	Nonce     string

	client *resty.Client
}

//CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appkey, appSecret, httpProxy string) *ImClient {
	c := &ImClient{AppKey: appkey, AppSecret: appSecret, Nonce: RandStringBytesMaskImprSrc(64)}
	client := resty.New()
	if len(httpProxy) > 0 {
		client.SetProxy(httpProxy)
	}

	client.SetHeader("Accept", "application/json;charset=utf-8").
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;").
		SetHeader("AppKey", c.AppKey).
		SetHeader("Nonce", c.Nonce).
		SetPreRequestHook(checksumHook(appSecret, c.Nonce)).
		SetTimeout(5 * time.Second)
	c.client = client
	return c
}

// replace this with an noop
func (c *ImClient) setCommonHead(req *resty.Request) {}

func checksumHook(appSecret string, nonce string) func(*resty.Client, *resty.Request) error {
	return func(_ *resty.Client, req *resty.Request) error {
		timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
		req.SetHeader("CurTime", timeStamp)
		req.SetHeader("CheckSum", ShaHashToHexStringFromString(appSecret+nonce+timeStamp))
		return nil
	}
}
