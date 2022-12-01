package model

import "github.com/dgrijalva/jwt-go"

type Use struct {
	QQname string `form:"qqname" json:"qqname" binding:"required"`
	QQword string `form:"qqword" json:"qqword" binding:"required"`
}
type MyClaims struct {
	QQname string `json:"qqname"`
	jwt.StandardClaims
}
type AboutFriend struct {
	Friend string `form:"friend" json:"friend" binding:"required"`
}
type Group struct {
	Friend   string `form:"friend" json:"friend" binding:"required"`
	NewGroup string `form:"new group" json:"new group" binding:"required"`
}
type ScanGroup struct {
	Group string `form:"group" json:"group" binding:"required"`
}
type Users struct {
	Id     int
	QQname string
	QQword string
	Login  string
}
type Friends struct {
	Id     int
	QQname string
	Group  string
	Scan   rune
}
