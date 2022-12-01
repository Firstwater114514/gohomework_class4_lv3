package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/api/middleware"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/dao"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/model"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/utils"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.Use{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后进行注册")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := c.PostForm("qqname")
	flag := dao.SelectUser(qqname)
	if !flag {
		utils.RespFail(c, "用户名已被使用")
		return
	}
	qqword := c.PostForm("qqword")
	ok := dao.AddUser(qqname, qqword)
	if ok == "return1" {
		utils.RespFail(c, "insert failed")
		return
	} else if ok == "return2" {
		utils.RespFail(c, "create table failed")
		return
	}
	utils.RespSuccess(c, "注册成功!快去登录吧~")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.Use{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后登录")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := c.PostForm("qqname")
	qqword := c.PostForm("qqword")
	flag := dao.SelectUser(qqname)
	if flag {
		utils.RespFail(c, "该用户不存在")
		return
	}
	selectPassword := dao.SelectQQwordFromQQname(qqname)
	if selectPassword == "" {
		utils.RespFail(c, "scan failed")
		return
	} else if selectPassword != qqword {
		utils.RespFail(c, "密码错误")
		return
	}
	dao.Login(qqname)
	claim := model.MyClaims{
		QQname: qqname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "sqy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.LoginSuccess(c, tokenString, "注销账号请使用unsubscribe", "加好友请用add friend", "删除好友请用delete friend", "浏览全部好友请用scan friends", "浏览某组所有好友请用scan group", "改变好友的分组请用change group", "查找好友请用search friend", "记得关闭窗口前一定要退出账号哦!!!!!用quit")
}

func getUsernameFromToken(c *gin.Context) {
	qqname, _ := c.Get("qqname")
	utils.RespSuccess(c, qqname.(string))
}

func quit(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "没有登录的账户")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	dao.Quit()
	utils.RespSuccess(c, "成功退出账号")
	return
}
func unsubscribe(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := dao.WhoLogin()
	dao.Unsubscribe(qqname)
	utils.RespSuccess(c, "注销账户成功")
}
func addFriend(c *gin.Context) {
	if err := c.ShouldBind(&model.AboutFriend{}); err != nil {
		utils.RespFail(c, "请输入对方的qq昵称")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	friend := c.PostForm("friend")
	qqname := dao.WhoLogin()
	if friend == qqname {
		utils.RespFail(c, "无法添加自己为好友哦~")
		return
	}
	flag1 := dao.SelectUser(friend)
	if flag1 {
		utils.RespFail(c, "该用户不存在")
		return
	}
	flag2 := dao.CheckNewFriend(qqname, friend)
	if !flag2 {
		utils.RespFail(c, "请勿重复添加好友")
	}
	group := c.PostForm("group")
	ok := dao.AddFriend(qqname, friend, group)
	if ok == "return" {
		utils.RespFail(c, "insert failed")
		return
	}
	utils.RespSuccess(c, "添加好友成功！")
}
func deleteFriend(c *gin.Context) {
	if err := c.ShouldBind(&model.AboutFriend{}); err != nil {
		utils.RespFail(c, "请输入对方的qq昵称")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	friend := c.PostForm("friend")
	qqname := dao.WhoLogin()
	if qqname == friend {
		utils.RespFail(c, "无法删除自己哦，但是您可以选择unsubscribe来注销账号")
		return
	}
	find := dao.FindFriend(qqname, friend)
	if find == "return" {
		utils.RespFail(c, "该用户不存在或不是您的好友")
		return
	} else if find == "unsubscribe" {
		utils.RespFail(c, "很抱歉，该好友将账号注销了")
		return
	}
	dao.DeleteFriend(qqname, friend)
	utils.RespSuccess(c, "删除成功")
}
func scanFriends(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := dao.WhoLogin()
	f1 := dao.ScanFriends(qqname)
	if f1 == nil {
		utils.RespFail(c, "scan failed")
		return
	}
	utils.AllFriends(c, f1)
}
func scanGroup(c *gin.Context) {
	if err := c.ShouldBind(&model.ScanGroup{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := dao.WhoLogin()
	group := c.PostForm("group")
	flag := dao.CheckGroup(qqname, group)
	if !flag {
		utils.RespFail(c, "该组不存在")
		return
	}
	friends := dao.ScanGroup(qqname, group)
	if friends == nil {
		utils.RespFail(c, "scan or query failed")
	}
	utils.AllFriends(c, friends)
}
func changeGroup(c *gin.Context) {
	if err := c.ShouldBind(&model.Group{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := dao.WhoLogin()
	friend := c.PostForm("friend")
	if qqname == friend {
		utils.RespFail(c, "不可以对自己操作哦~")
		return
	}
	find := dao.FindFriend(qqname, friend)
	if find == "return" {
		utils.RespFail(c, "该用户不存在或不是您的好友")
		return
	} else if find == "unsubscribe" {
		utils.RespFail(c, "很抱歉，该好友将账号注销了")
		return
	}
	newGroup := c.PostForm("new group")
	flag := dao.ChangGroup(qqname, friend, newGroup)
	if !flag {
		utils.RespFail(c, "update failed")
	}
	utils.RespSuccess(c, "更改成功！")
}
func searchFriend(c *gin.Context) {
	if err := c.ShouldBind(&model.AboutFriend{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	qqname := dao.WhoLogin()
	friend := c.PostForm("friend")
	if qqname == friend {
		utils.RespFail(c, "不可以对自己操作哦~")
		return
	}
	find := dao.FindFriend(qqname, friend)
	if find == "return" {
		utils.RespFail(c, "该用户不存在或不是您的好友")
		return
	} else if find == "unsubscribe" {
		utils.RespFail(c, "很抱歉，该好友将账号注销了")
		return
	}
	group := dao.SearchFriend(qqname, friend)
	if group == "" {
		utils.RespFail(c, "scan failed")
	}
	sen := friend + "在" + group + "组"
	utils.RespSuccess(c, sen)
}
