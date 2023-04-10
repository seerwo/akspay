package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/pay/context"
	"github.com/seerwo/akspay/util"
)

const(
	CREATE_SDK_ORDER_URL = "/customs/partners/%s/declare/report/%s"
	SEARCH_SDK_ORDER_URL = "/customs/partners/%s/declare/query/%s"
	REDECLARE_ORDER_URL = "/customs/partners/%s/redeclare/report/%s"
	MODIFY_CUSTOM_URL = "/customs/partners/%s/customs_declare_modify/reports/%s"
)
type SubOrders []SubJson

func (i SubOrders) MarshalJSON() ([]byte, error) {
	if len(i) == 0 {
		return json.Marshal([]interface{}{})
	} else {
		var tempValue []interface{} // 这里需要重新定义一个变量，再序列化，否则会造成循环调用
		for _, item := range i {
			tempValue = append(tempValue, item)
		}
		return json.Marshal(tempValue)
	}
}

type ReqCreateOrder struct {
	OrderId string `json:"order_id"` //order_id 	String 必填，商户支付订单号(建议长度8-25位)，要求同一商户唯一
	Custom string `json:"custom"` //custom 	String 必填，海关编号 * 渠道海关编号
	MchCustomId string `json:"mch_custom_id"` //mch_custom_id 	String 必填，商户在海关备案的编号
	MchCustomName string `json:"mch_custom_name"` //mch_custom_name 	String 必填，商户海关备案名称
	SubOrder SubOrders `json:"sub_order"` //sub_order 	JSON[] 子订单(拆单)
}

type SubJson struct {
	SubOrderNo string `json:"sub_order_no"` //sub_order_no 	String 商户子订单号(建议长度8-25位)
	FeeType string `json:"fee_type"` //fee_type 	String 币种代码 默认值: CNY 允许值: CNY
	OrderFee float64 `json:"order_fee"` //order_fee 	Double 子订单金额
	TransportFee float64 `json:"transport_fee"` //transport_fee 	Double 子订单物流金额
	ProductFee float64 `json:"product_fee"` //product_fee 	Double 子订单商品金额
}

type ResCreateOrder struct {
	util.CommonPayError
	util.CommonError
	ReportId string `json:"report_id"` //report_id 	String QRtrip海关单号
	ClientReportId string `json:"client_report_id"` //client_report_id 	String 商户申请报关单号
	ReportStatus string `json:"report_status"` //report_status 	String 报关单状态: PROCCESSING,SUBMITED,FAILED,SUCCESS
	Channel string `json:"channel"` //channel 	String 支付渠道
	Custom string `json:"custom"` //custom 	String 海关编号
	MchCustomNo string `json:"mch_custom_no"` //mch_custom_no 	String 商户在海关备案的编号
	MchCustomName string `json:"mch_custom_name"` //mch_custom_name 	String 商户海关备案名称
	OrderId string `json:"order_id"` //order_id 	String QRtrip订单号
	TransactionId string `json:"transaction_id"` //transaction_id 	String 支付渠道订单号
	OrderCurrency string `json:"order_currency"` //order_currency 	String 币种
	OrderAmount float64 `json:"order_amount"` //order_amount 	Double 订单金额
	ReportTime string `json:"report_time"` //report_time 	String 报关时间
	CreationDate string `json:"creation_date"` //creation_date 	String 报关单创建时间
	LastUpdateDate string `json:"last_update_date"` //last_update_date 	String 更新时间

	ClientOrderId string `json:"client_order_id"`
	ReportAmount float64 `json:"report_amount"`
	ReportCurrency string `json:"report_currency"`
	ChannelOrderId string `json:"channel_order_id"`
	PartnerOrderId string `json:"partner_order_id"`
	IdentityCheck string `json:"identity_check"`
	RequestXml string `json:"request_xml"`
	RequestCode string `json:"request_code"`
	SubOrders []SubOrder `json:"sub_orders"`
}

type SubOrder struct {
	ChannelSubOrderId string `json:"channel_sub_order_id"`
	SubOrderNo string `json:"sub_order_no"`
	OrderFee float64 `json:"order_fee"`
	ProductFee float64 `json:"product_fee"`
	ReportStatus string `json:"report_status"`
	TransportFee float64 `json:"transport_fee"`
}

type ReqSearchOrder struct {

}

type ResSearchOrder struct {
	util.CommonError
	ReportStatus string `json:"report_status"` //report_status 	String 报关单状态: PROCCESSING,SUBMITED,FAILED,SUCCESS
	ReportId string `json:"report_id"` //report_id 	String QRtrip海关单号
	ClientReportId string `json:"client_report_id"` //client_report_id 	String 商户申请报关单号
	OrderId string `json:"order_id"` //order_id 	String QRtrip订单号
	PartnerOrderId string `json:"partner_order_id"` //partner_order_id 	String 商户订单ID
}

type ReqRedeclareOrder struct {

}

type ResRedeclareOrder struct {
	ResCreateOrder
}

type ReqModifyCustom struct {
	PartnerReportId string `json:"partner_report_id"` //partner_report_id 	String 必填，商户申请报关单号，要求同一商户唯一
	Customs string `json:"customs"` //customs 	String 必填，海关编号 * 渠道海关编号
	MchCustomId string `json:"mch_customs_id"` //mch_customs_id 	String 必填，商户在海关备案的编号
	MchCustomsName string `json:"mch_customs_name"` //mch_customs_name 	String 必填，商户海关备案名称
	ConsumerEncryptId string `json:"consumer_encrypt_id"` //consumer_encrypt_id 	String 消费者信息加密key id，见获取加密密钥
	ConsumerEncryptInfo string `json:"consumer_encrypt_info"` //consumer_encrypt_info 	String 加密后的消费者信息，使用[获取加密密钥]获取的公钥对消费者信息组成的json字符串进行RSA加密，以base64格式传输；报关消费者信息和付款人不一致时需要提交。
}

type ConsumerInfo struct {
	ConsumerName string `json:"consumer_name"` //consumer_name 	String 消费者姓名
	ConsumerIdNo string `json:"consumer_id_no"` //consumer_id_no 	String消费者身份证号
}

type ResModifyCustom struct {
	util.CommonPayError
	util.CommonError
	Reports []Report `json:"reports"` //reports 	Array 报关单信息列表
}

type Report struct {
	ReportId string `json:"report_id"` //report_id 	String QRtrip海关单号
	PartnerReportId string `json:"partner_report_id"` //partner_report_id 	String 商户申请报关单号
	Status string `json:"status"` //partner_report_id 	String 商户申请报关单号
	Channel string `json:"channel"` //channel 	String 报关渠道
	CannelReportId string `json:"cannel_report_id"` //channel_report_id 	String 渠道报关单号
	Customs string `json:"customs"` //customs 	String 海关编号
	MchCustomsNo string `json:"mch_customs_no"` //mch_customs_no 	String 商户在海关备案的编号
	MchCustomsName string `json:"mch_customs_name"` //mch_customs_name 	String 商户海关备案名称
	OrderId string `json:"order_id"` //order_id 	String QRtrip订单号
	TransactionId string `json:"transaction_id"` //transaction_id 	String  支付渠道订单号
	OrderCurrency string `json:"order_currency"` //order_currency 	String 币种，CNY
	OrderAmount int64 `json:"order_amount"` //order_amount 	Integer 订单金额，单位是人民币分
	CreationDate string `json:"creation_date"` //creation_date 	String 报关单创建时间，2021-01-01 01:00:00 格式
	LastUpdateDate string `json:"last_update_date"` //last_update_date 	String 渠道更新时间，2021-01-01 01:00:00 格式
	VerifyDepartment string `json:"verify_department"` //verify_department 	String 校验单位
	VerifyDepartmentTradeId string `json:"verify_department_trade_id"` //verify_department_trade_id 	String 校验号
	ErrorCode string `json:"error_code"` //error_code 	String 错误代码
	ErrorMsg string `json:"error_msg"` //error_msg 	String 错误返回的信息描述
}

//Order order struct
type Order struct {
	*context.Context
}

//NewPackage instance
func NewOrder(context *context.Context) *Order {
	order := new(Order)
	order.Context = context
	return order
}

func (o *Order) GetCreateOrderReq(req ReqCreateOrder, clientReportId string)(res ResCreateOrder, err error) {

	accessParam, _ := json.Marshal(req)

	var accessUrl string
	if accessUrl, err = o.GetAccessParam(req, util.BASE_WEB_URL + CREATE_SDK_ORDER_URL, clientReportId); err != nil {
		return
	}
	//fmt.Println(accessUrl)
	//fmt.Println(string(accessParam))
	var response []byte
	if response, err = util.NewHTTPPut(accessUrl, string(accessParam)); err != nil {
		return
	}
	//println(string(response))
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode != "SUCCESS" && res.ErrMsg != "customs declaration of order already exists" && res.ErrMsg != "customs declaration already exists"{
		err = fmt.Errorf("GetCreateOrderReq Error , errcode=%s , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

func (o *Order) GetRedeclareOrderReq(req ReqRedeclareOrder, clientReportId string)(res ResRedeclareOrder, err error) {

	accessParam, _ := json.Marshal(req)

	var accessUrl string
	if accessUrl, err = o.GetAccessParam(req, util.BASE_WEB_URL + REDECLARE_ORDER_URL, clientReportId); err != nil {
		return
	}
	//fmt.Println(accessUrl)
	//fmt.Println(string(accessParam))

	var response []byte
	if response, err = util.NewHTTPPut(accessUrl, string(accessParam)); err != nil {
		return
	}
	//println(string(response))
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode == "INVALID_SIGN"{
		if _, err = o.GetRedeclareOrderReq(req, clientReportId); err != nil {
			return
		}
	}
	if res.ErrCode != "SUCCESS" && res.ErrMsg != "customs declaration of order already exists"{
		err = fmt.Errorf("GetRedeclareOrderReq Error , errcode=%s , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

func (o *Order) GetModifyCustomReq(req ReqModifyCustom, clientReportId string)(res ResModifyCustom, err error) {

	accessParam, _ := json.Marshal(req)

	var accessUrl string
	if accessUrl, err = o.GetAccessParam(req, util.BASE_WEB_URL + MODIFY_CUSTOM_URL, clientReportId); err != nil {
		return
	}
	//fmt.Println(accessUrl)
	//fmt.Println(string(accessParam))

	var response []byte
	if response, err = util.NewHTTPPut(accessUrl, string(accessParam)); err != nil {
		return
	}
	println(string(response))
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	//if res.ErrCode == "INVALID_SIGN"{
	//	if _, err = o.GetRedeclareOrderReq(req, clientReportId); err != nil {
	//		return
	//	}
	//}
	if res.ErrCode != "SUCCESS" && res.ErrMsg != "customs declaration of order already exists"{
		err = fmt.Errorf("GetModifyCustomReq Error , errcode=%s , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}