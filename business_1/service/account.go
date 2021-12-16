// @Description
// @Author  renjingneng
// @CreateTime  2021/12/14 17:55
package service

import (
	"encoding/json"
	"time"

	"github.com/renjingneng/goapp/business_1/database/mysql"
	"github.com/renjingneng/goapp/business_1/database/redis"
)

type AccountService struct {
	OpenDatabase *mysql.Open
	OpenRedis    *redis.Open
}

func NewAccountService() *AccountService {
	res := &AccountService{
		OpenDatabase: mysql.NewOpen(),
		OpenRedis:    redis.NewOpen(),
	}
	return res
}

func (sev *AccountService) GetInfoById(id string) map[string]interface{} {
	cacheKey := "account_" + id
	cacheVal, _ := sev.OpenRedis.Get(cacheKey)
	res := make(map[string]interface{})
	if cacheVal == "" {
		sev.OpenDatabase.SetTablename("account")
		res = sev.OpenDatabase.FetchRowShort(map[string]string{"id": id})
		cacheValByte, _ := json.Marshal(&res)
		sev.OpenRedis.Set(cacheKey, cacheValByte, 1000*time.Second)
	} else {
		json.Unmarshal([]byte(cacheVal), &res)
	}
	return res
}
