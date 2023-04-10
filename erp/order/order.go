package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/akspay/erp/context"
	"github.com/seerwo/akspay/util"
	"time"
)
const(
	CREATE_CUSTOM_ORDER_URl = "/b2c/custom"
)

type ReqOrder struct {
	Id string `json:"id"`
	AgentId string `json:"agent_id"`
	WareId string `json:"ware_id"`
	AgentName string `json:"agent_name"`
	WareName string `json:"ware_name"`
	OrderDetails []interface{} `json:"order_details"`

	OutId string	`json:"out_id"`
	ConfirmTime time.Time `json:"-"`
	PayTime time.Time `json:"pay_time"`
	Status string `json:"status"`
	PackageId string `json:"package_id"`
	ExpressCompanyCode string `json:"express_company_code"`
	ExpressCompanyName string `json:"express_company_name"`
	ExpressNo string `json:"express_no"`
	ReceiverName string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	Province string `json:"province"`
	City string `json:"city"`
	District string `json:"distribute"`
	TotalNetWeight int `json:"total_net_weight"`
	InternationalExpressNo string `json:"internation_express_no"`
	DeliveryTimePreference string `json:"delivery_time_preference"`
	OrderDeclaredAmount string `json:"order_declared_amount"`
	PaintMarker string `json:"paint_marker"`
	ExpressExtend1 string `json:"express_extend_1"`
	ExpressExtend2 string `json:"express_extend_2"`


	PayAmount float64 `json:"pay_amount"`
	FromType string `json:"from_type"`
	AdTime string `json:"ad_time"`
	//AdUser string `orm:"size(40)" json:"ad_user"`

	ReceiverAddress string `json:"receiver_address"`
	Mobile string `json:"mobile"`
	SellerDiscount float64 `json:"seller_discount"`
	PayType string `json:"pay_type"`
	Versand float64 `json:"versand"`
	Subventionen float64 `json:"subventionen"`
	Customs string `json:"customs"`
	ShippingPrice float64 `json:"shipping_price"`
	UsagingPrice float64 `json:"usaging_price"`
	SkuIds string `json:"sku_ids"`
	YonyouVbillcode string `json:"yonyou_vbillcode"`
	YonyouCsaleorderid string `json:"yonyou_csaleorderid"`
	SendStatus string `json:"send_status"`	//发送用友 和 平台回传标志 0 null 待发  1 已发送 -1取消
	SendCustoms string `json:"send_customs"`     //发送海关 及 物流仓库标志 0 null 待发  1 已发送 -1取消
	StatusCustoms string `json:"status_customs"` // 1 支付 3 订单 7 物流
	Wid string `json:"wid"`
	TradeId string `json:"trade_id"`
	ReceiverZip string `json:"receiver_zip"`
	IdCardNo string `json:"id_card_no"`

	ChannelTrxNo string `json:"channel_trax_no"`
	InteractId string `json:"interact_id"`
	LogisticsNo string `json:"logistics_no"`
	PayOrderId string `json:"pay_order_id"`
	OrderNumber string `json:"order_number"`
	UserName string `json:"user_name"`
	PlatDiscount float64 `json:"plat_discount"`
	PayId string `json:"pay_id"`
	StatusLogist string `json:"status_logist"`
	SalesChannel string `json:"sales_channel"`
}

type ResOrder struct {
	util.CommonErpError
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

func (o *Order) GetCreateOrderReq(req ReqOrder)(res ResOrder, err error) {

	accessParam, _ := json.Marshal(req)

	var response []byte
	if response, err = util.NewHTTPPost(util.BASE_ERP_URL + CREATE_CUSTOM_ORDER_URl, string(accessParam)); err != nil {
		return
	}

	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("GetCreateOrderReq Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
