package credential

//AccessTokenHandle AccessToken interface
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
	GetAccessParam(req interface{}, url, clientReportId string)(accessParam string, err error)
	GetPrizeParam(tel string)(accessParam string, err error)
}