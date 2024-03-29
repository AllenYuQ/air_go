package test

import (
	"fmt"
	"task5/models"
	"testing"
)
//进行单元测试时，需要将setting/setting.go的init() 方法中
//Cfg, err = ini.Load("conf/app.ini") 改为 Cfg, err = ini.Load("../conf/app.ini")
func TestDB(t *testing.T)  {
	//测试数据库连接是否成功 username 用户名 password 密码
	isExist := models.CheckUser("admin", "admin123")
	fmt.Println(isExist)
	//测试增加用户
	data := make(map[string]interface{})
	data["id"] = 3
	data["username"] = "test"
	data["realname"] = "testrealname"
	data["password"] = "test123"
	data["email"] = "2969141711@qq.com"
	data["phone"] = "18816209310"
	data["sex"] = 1
	data["roleId"] = 1
	canAdd := models.AddUser(data)
	fmt.Println(canAdd)
	//列出所有的用户
	var users [] models.User
	users = models.ListUsers()
	for _,user := range users {
		fmt.Println(user)
	}
}

func TestMlmet(t *testing.T)  {
	mlmets := models.ListMlmets()
	fmt.Println(len(mlmets))
	for i, mlmet := range(mlmets) {
		fmt.Println(i)
		fmt.Println(mlmet)
	}
	fmt.Println(mlmets)
}