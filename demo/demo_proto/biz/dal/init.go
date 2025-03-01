package dal

import (
	"github.com/wwbab/Go_mall/demo/demo_proto/biz/dal/mysql"
	"github.com/wwbab/Go_mall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
