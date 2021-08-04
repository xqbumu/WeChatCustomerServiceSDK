package WeChatCustomerServiceSDK

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
)

const (
	//获取会话状态
	serviceStateGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get?access_token=%s"
	// 变更会话状态
	serviceStateTransAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans?access_token=%s"
)

// ServiceStateGetOptions 获取会话状态请求参数
type ServiceStateGetOptions struct {
	OpenKFID string `json:"open_kfid"`							// 客服帐号ID
	ExternalUserID string `json:"external_userid"`				// 微信客户的external_userid
}

// ServiceStateGetSchema 获取会话状态响应内容
type ServiceStateGetSchema struct {
	BaseModel
	ServiceState int `json:"service_state"`						// 当前的会话状态，状态定义参考概述中的表格
	ServiceUserID string `json:"service_userid"`				// 接待人员的userid，仅当state=3时有效
}

// ServiceStateGet 获取会话状态
//0	未处理	新会话接入。可选择：1.直接用API自动回复消息。2.放进待接入池等待接待人员接待。3.指定接待人员进行接待
//1	由智能助手接待	可使用API回复消息。可选择转入待接入池或者指定接待人员处理。
//2	待接入池排队中	在待接入池中排队等待接待人员接入。可选择转为指定人员接待
//3	由人工接待	人工接待中。可选择结束会话
//4	已结束	会话已经结束。不允许变更会话状态，等待用户重新发起咨询
func (r *Client) ServiceStateGet(options ServiceStateGetOptions) (info ServiceStateGetSchema, err error) {
	target := fmt.Sprintf(serviceStateGetAddr, r.accessToken)
	r.recordUpdate(target)

	data, err := util.HttpPost(target, options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// ServiceStateTransOptions 变更会话状态请求参数
type ServiceStateTransOptions struct {
	OpenKFID string `json:"open_kfid"`							// 客服帐号ID
	ExternalUserID string `json:"external_userid"`				// 微信客户的external_userid
	ServiceState int `json:"service_state"`						// 变更的目标状态，状态定义和所允许的变更可参考概述中的流程图和表格
	ServicerUserID string `json:"servicer_userid"`				// 接待人员的userid，当state=3时要求必填
}

// ServiceStateTrans 变更会话状态
func (r *Client) ServiceStateTrans(options ServiceStateTransOptions) (info BaseModel, err error) {
	target := fmt.Sprintf(serviceStateTransAddr, r.accessToken)
	r.recordUpdate(target)

	data, err := util.HttpPost(target, options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}