# 第四节课作业lv3

## 前排提示：

1、每次关闭postman之类软件页面之前请退出（“/quit”）“账号”。

2、初次使用是没有账号的，需要注册再登录才可使用。

3、qqword是qq密码的意思。

4、登录后才可以使用除了“注册”以外的功能。

5、添加好友时没有填写分组”group“，就默认该好友是”myfriend“组

## 功能介绍：

### 1、注册

在请求处输入“/register”

key需要qqname、qqword

### 2、登录

在请求处输入“/login”

key需要qqname、qqword

无法多个账户同时登录，换账号前需要退出已登录的账号

### 3、用令牌获取用户信息

在请求处输入“/user/get”

Header处需要Authorization

### 4、退出账号

在请求处输入“/quit”

不需要任何key

每次关闭接口测试软件之前一定要退出！！！！！！！

### 5、注销账号

在请求处输入“/unsubscribe”

不需要任何key

### 6、添加好友（附带好友分组）

在请求处输入“/add friend”

添加的好友必须是已经注册的账户哦

被添加的用户会默认将主动添加的用户分到”myfriend“组中

key需要friend

key可以添加group

(不添加group就默认该好友是”myfriend“组)

### 7、删除好友

在请求处输入“/delete friend”

key需要friend

### 8、查看所有好友

在请求处输入“/scan friends”

不需要任何key

### 9、查看某分组内的所有好友

在请求处输入“/scan group”

key需要group

### 10、更改好友分组

在请求处输入“/change group”

key需要friend、new group

更改的分组可以是目前没有的新分组

### 11、查找好友

在请求处输入“/search friend”

key需要friend

该功能会显示该好友位于哪个分组
![$}TX% BN``)`7NONBS`79GH](https://user-images.githubusercontent.com/116962163/205443106-89c484a7-55f2-45b7-a3d3-9662a5bd3f6f.png)
