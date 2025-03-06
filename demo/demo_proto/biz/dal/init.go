package dal

import (
	"github.com/wwbab/Go_mall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
