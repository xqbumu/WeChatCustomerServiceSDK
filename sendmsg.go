package WeChatCustomerServiceSDK

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
)

const (
	//发送消息
	sendMsgAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=%s"
)

// SendMsgSchema 发送消息响应内容
type SendMsgSchema struct {
	BaseModel
	MsgID string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节, 字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// SendMsg 获取消息
func (r *Client) SendMsg(options interface{}) (info SendMsgSchema, err error) {
	data, err := util.HttpPost(fmt.Sprintf(sendMsgAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
