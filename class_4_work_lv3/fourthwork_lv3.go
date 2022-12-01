package main

import (
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/api"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
