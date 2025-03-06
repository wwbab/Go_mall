package main

import (
	"github.com/joho/godotenv"
	"github.com/wwbab/Go_mall/demo/demo_proto/biz/dal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	/* //CRUD
	//mysql.DB.Create(&model.User{Email: "demo@example.com",Password: "nihao"}) //创建记录
	mysql.DB.Model(&model.User{}).Where("email = ?","demo@example.com").Update("Password","dada")	//更新
	//查询（First单行，Find多行）
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?","demo@example.com").First(&row)

	fmt.Printf("row: %+v\n",row)
	//软删除,当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录
	mysql.DB.Where("email = ?","demo@example.com").Delete(&model.User{}) */

	//永久删除
	//mysql.DB.Unscoped().Where("email = ?","demo@example.com").Delete(&model.User{})
}

