package common

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/pay/context"
	"github.com/seerwo/akspay/util"
)

const(
	QUERY_ORDER_STATUS_URl = "/gateway/partners/%s/orders/%s" //https://pay.akspay.com/api/v1.0/gateway/partners/{partner_code}/orders/{order_id}
)

type ReqOrderStatus struct {
	PartnerCode string `json:"partner_code"`  //必填，商户编码，由4-6位大写字母或数字构成
	OrderId string `json:"order_id"` //必填，商户支付订单号(建议长度8-25位)
}

type ResOrderStatus struct {
	util.CommonPayError
	util.CommonError
	ReturnCode string `json:"return_code"` //执行结果
	ResultCode string `json:"result_code"` //PAYING:等待支付 CREATE_FAIL:订单创建失败 CLOSED:已关闭 PAY_FAIL:支付失败 PAY_SUCCESS:支付成功 PARTIAL_REFUND:部分退款 FULL_REFUND:全额退款
	ChannelOrderId string `json:"channel_order_id"`  //渠道方(微信、支付宝等)订单ID
	OrderId string `json:"order_id"`  //QRtrip订单ID
	PartnerOrderId string `json:"partner_order_id"` //商户订单ID
	TotalFee int `json:"total_fee"`  //订单金额，单位是货币最小面值单位
	RealFee int `json:"real_fee"`  //实际支付金额，单位是货币最小面值单位(目前等于订单金额，为卡券预留)
	Rate float64 `json:"rate"`  //交易时使用的汇率，1JPY=?CNY
	CustomerId string `json:"customer_id"` //客户ID
	PayType string `json:"pay_type"` //(optional) 支付方式(有则展示):对应支付的钱包、或者信用卡、借记卡
	PayTime string `json:"pay_time"` //支付时间（yyyy-MM-dd HH:mm:ss，GMT+8）
	CreateTime string `json:"create_time"` //订单创建时间（最新订单为准）（yyyy-MM-dd HH:mm:ss，GMT+8）
	Currency string `json:"currency"` //币种，通常为JPY
	Channel string `json:"channel"` //支付渠道 Alipay|支付宝、Wechat|微信
	PreOrderXml string `json:"pre_order_xml"` //(Wechat) 下单微信返还xml
	QueryOrderXml string `json:"query_order_xml"` //(Wechat) 查询订单返回xml
}

//Common common struct
type Common struct {
	*context.Context
}

//NewPackage instance
func NewCommon(context *context.Context) *Common {
	common := new(Common)
	common.Context = context
	return common
}

func (o *Common) GetOrderStatus(req ReqOrderStatus, clientReportId string)(res ResOrderStatus, err error) {

	//accessParam, _ := json.Marshal(req)

	var accessUrl string
	if accessUrl, err = o.GetAccessParam(req, util.BASE_WEB_URL + QUERY_ORDER_STATUS_URl, clientReportId); err != nil {
		return
	}
	//fmt.Println(accessUrl)
	//fmt.Println(string(accessParam))
	var response []byte
	if response, err = util.NewHTTPGet(accessUrl, nil); err != nil {
		return
	}
	//println(string(response))
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ReturnCode != "SUCCESS" && res.ErrCode == "INVALID_SIGN" {
		if response, err = util.NewHTTPGet(accessUrl, nil); err != nil {
			return
		}
		if err = json.Unmarshal(response, &res); err != nil {
			return
		}
	}
	if res.ReturnCode != "SUCCESS" {
		err = fmt.Errorf("GetOrderStatus Error , errcode=%s , errmsg=%s", res.ReturnCode, res.ErrMsg)
		return
	}
	return
}