package WeChatCustomerServiceSDK

import (
	"reflect"
	"strings"
)

// Error 错误
type Error string

const (
	// SDKInitFailed 错误码：50001
	SDKInitFailed Error = "SDK初始化失败"
	// SDKCacheUnavailable 错误码：50002
	SDKCacheUnavailable Error = "缓存无效"
	// SDKUnknownError 错误码：50003
	SDKUnknownError Error = "未知错误"
	// SDKInvalidCredential 错误码：40001
	SDKInvalidCredential Error = "不合法的secret参数"
	// SDKInvalidImageSize 错误码：40009
	SDKInvalidImageSize Error = "无效的图片大小"
	// SDKInvalidCorpID 错误码：40013
	SDKInvalidCorpID Error = "无效的 CorpID"
	// SDKAccessTokenInvalid 错误码：40014
	SDKAccessTokenInvalid Error = "AccessToken 无效"
	// SDKMediaIDExceedMinLength 错误码：40058
	SDKMediaIDExceedMinLength Error = "media_id 小于最小长度 1"
	// SDKAccessTokenMissing 错误码：41001
	SDKAccessTokenMissing Error = "缺少AccessToken参数"
	// SDKAccessTokenExpired 错误码：42001
	SDKAccessTokenExpired Error = "AccessToken 已过期"
	// SDKApiFreqOutOfLimit 错误码：45009
	SDKApiFreqOutOfLimit Error = "接口请求次数超频"
	// SDKApiForbidden 错误码：48002
	SDKApiForbidden Error = "API 禁止调用"
	// SDKInvalidOpenKFID 错误码：95000
	SDKInvalidOpenKFID Error = "无效的 open_kfid"
	// SDKOpenKFIDNotExist 错误码：95004
	SDKOpenKFIDNotExist Error = "open_kfid 不存在"
	// SDKWeWorkAlready 错误码：95011
	SDKWeWorkAlready Error = "已在企业微信使用微信客服"
)

//输出错误信息
func (r Error) Error() string {
	return reflect.ValueOf(r).String()
}

// NewSDKErr 初始化SDK实例错误信息
func NewSDKErr(code int, msgList ...string) Error {
	switch code {
	case 50001:
		return SDKInitFailed
	case 50002:
		return SDKCacheUnavailable
	case 40001:
		return SDKInvalidCredential
	case 40009:
		return SDKInvalidImageSize
	case 41001:
		return SDKAccessTokenMissing
	case 42001:
		return SDKAccessTokenExpired
	case 40013:
		return SDKInvalidCorpID
	case 40014:
		return SDKAccessTokenInvalid
	case 40058:
		return SDKMediaIDExceedMinLength
	case 45009:
		return SDKApiFreqOutOfLimit
	case 48002:
		return SDKApiForbidden
	case 95000:
		return SDKInvalidOpenKFID
	case 95004:
		return SDKOpenKFIDNotExist
	case 95011:
		return SDKWeWorkAlready
	default:
		//返回未知的自定义错误
		if len(msgList) > 0 {
			return Error(strings.Join(msgList, ","))
		}
		return SDKUnknownError
	}
}
