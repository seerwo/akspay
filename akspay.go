package akspay

import (
	"github.com/seerwo/akspay/cache"
	"github.com/seerwo/akspay/erp"
	config_erp "github.com/seerwo/akspay/erp/config"
	"github.com/seerwo/akspay/pay"
	config_pay "github.com/seerwo/akspay/pay/config"
	"github.com/seerwo/akspay/prize"
	config_prize "github.com/seerwo/akspay/prize/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	//Log as JSON instead of the fefault ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	//Output to stdout instead of the default stderr
	//Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//Only log the warning seerity or above.
	log.SetLevel(log.DebugLevel)
}

//Ark struct
type Akspay struct {
	cache cache.Cache
}

//NewArk init
func NewAkspay() *Akspay {
	return &Akspay{}
}

//SetCache set cache
func(a *Akspay) SetCache(cache cache.Cache) {
	a.cache = cache
}

func (a *Akspay) GetPay(cfg *config_pay.Config) *pay.Pay{
	return pay.NewPay(cfg)
}

func (a *Akspay) GetErp(cfg *config_erp.Config) *erp.Erp{
	return erp.NewErp(cfg)
}

func (a *Akspay) GetPrize(cfg *config_prize.Config) *prize.Prize{
	return prize.NewPrize(cfg)
}