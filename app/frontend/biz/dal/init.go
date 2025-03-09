package dal

import (
	"github.com/wwbab/Go_mall/app/frontend/biz/dal/mysql"
	"github.com/wwbab/Go_mall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
