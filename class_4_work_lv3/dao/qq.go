package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/model"
	"log"
)

var db *sql.DB

func AddUser(qqname, qqword string) string {
	us := model.Users{
		QQname: qqname,
		QQword: qqword,
	}
	sqlStrA := "insert into qquser(qqname,qqword) values (?,?)"
	_, errA := db.Exec(sqlStrA, us.QQname, us.QQword)
	if errA != nil {
		fmt.Printf("insert failed, err:%v\n", errA)
		return "return1"
	}
	sen1 := "CREATE TABLE `"
	sen2 := "`( `id` BIGINT(20) NOT NULL AUTO_INCREMENT,`friend` VARCHAR(20) DEFAULT 'nobody',`kind`VARCHAR(10) DEFAULT 'myfriend',PRIMARY KEY(`id`))ENGINE=InnoDB AUTO_INCREMENT=1 CHARSET=utf8mb4;"
	sqlStrF := sen1 + qqname + sen2
	_, errF := db.Exec(sqlStrF)
	if errF != nil {
		fmt.Printf("create table failed,err:%v\n", errF)
		return "return2"
	}
	return "ok"
}

func SelectUser(qqname string) bool {
	id := 1
	sqlStr := "select qqname from qqUser where id >= ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	for rows.Next() {
		var us model.Users
		err := rows.Scan(&us.QQname)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return false
		}
		if qqname == us.QQname {
			return false
		}
	}
	return true
}

func IfLogin() string {
	id := 1
	sqlStr := "select login from qqUser where id >= ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return "return1"
	}
	defer rows.Close()
	for rows.Next() {
		var us model.Users
		err := rows.Scan(&us.Login)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return "return2"
		}
		if us.Login == "yes" {
			return "yes"
		}
	}
	return "no"
}

func SelectQQwordFromQQname(qqname string) string {
	sqlStr := "select qqword from qquser where qqname=?"
	var us model.Users
	err := db.QueryRow(sqlStr, qqname).Scan(&us.QQword)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return us.QQword
}
func Login(qqname string) {
	var us model.Users
	us.Login = "yes"
	us.QQname = qqname
	sqlStr := "update qquser set login=? where qqname=?"
	_, err := db.Exec(sqlStr, us.Login, us.QQname)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func WhoLogin() string {
	login := "yes"
	sqlStr := "select qqname from qquser where login=?"
	var us model.Users
	err := db.QueryRow(sqlStr, login).Scan(&us.QQname)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return "return"
	}
	return us.QQname
}
func Quit() {
	var us1, us2 model.Users
	us2.Login = "yes"
	us1.Login = "no"
	sqlStr := "update qquser set login=? where login=?"
	_, err := db.Exec(sqlStr, us1.Login, us2.Login)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func Unsubscribe(qqname string) {
	sqlStr1 := "delete from qquser where qqname=?"
	_, err1 := db.Exec(sqlStr1, qqname)
	if err1 != nil {
		fmt.Printf("delete failed, err:%v\n", err1)
		return
	}
	sen1 := "drop table"
	sqlStr2 := sen1 + " " + qqname + ";"
	_, err2 := db.Exec(sqlStr2)
	if err2 != nil {
		fmt.Printf("delete failed, err:%v\n", err2)
		return
	}
}
func InitDB() {
	var err error
	dsn := "root:sqy040213@tcp(127.0.0.1:3306)/qq"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}
func AddFriend(qqname, friend, group string) string {
	sen1 := "insert into"
	sen2 := "(friend,kind) values (?,?)"
	sqlStr1 := sen1 + " " + qqname + sen2
	fmt.Println(sqlStr1)
	_, err1 := db.Exec(sqlStr1, friend, group)
	if err1 != nil {
		fmt.Printf("insert failed, err:%v\n", err1)
		return "return"
	}
	myfriend := "myfriend"
	sqlStr2 := sen1 + " " + friend + sen2
	_, err2 := db.Exec(sqlStr2, qqname, myfriend)
	if err2 != nil {
		fmt.Printf("insert failed, err:%v\n", err2)
		return "return"
	}
	return "ok"
}
func FindFriend(qqname, friend string) string {
	sen1 := "select friend from"
	sen2 := " where friend=?"
	sqlStr1 := sen1 + " " + qqname + sen2
	var fr model.Friends
	err1 := db.QueryRow(sqlStr1, friend).Scan(&fr.QQname)
	if err1 != nil {
		fmt.Printf("scan failed, err:%v\n", err1)
		return "return"
	}
	sqlStr2 := "select qqname from qquser where qqname=?"
	var us model.Users
	err2 := db.QueryRow(sqlStr2, friend).Scan(&us.QQname)
	if err2 != nil {
		fmt.Printf("scan failed, err:%v\n", err2)
		sen3 := "delete from"
		sen4 := " where friend=?"
		sqlStr3 := sen3 + " " + qqname + sen4
		_, err3 := db.Exec(sqlStr3, friend)
		if err3 != nil {
			fmt.Printf("delete failed, err:%v\n", err3)
			return "return"
		}
		return "unsubscribe"
	}
	return "ok"
}
func DeleteFriend(qqname, friend string) {
	sen1 := "delete from"
	sen2 := " where friend=?"
	sqlStr1 := sen1 + " " + qqname + sen2
	_, err1 := db.Exec(sqlStr1, friend)
	if err1 != nil {
		fmt.Printf("delete failed, err:%v\n", err1)
		return
	}
	sqlStr2 := sen1 + " " + friend + sen2
	_, err2 := db.Exec(sqlStr2, qqname)
	if err2 != nil {
		fmt.Printf("delete failed, err:%v\n", err2)
		return
	}
}
func ScanFriends(qqname string) []string {
	friends := make([]string, 1)
	friends[0] = "所有好友"
	sen1 := "select friend from"
	sen2 := " where id>=1"
	sqlStr := sen1 + " " + qqname + sen2
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var fr model.Friends
		err := rows.Scan(&fr.QQname)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		sqlStr1 := "select qqname from qquser where qqname=?"
		var us model.Users
		err1 := db.QueryRow(sqlStr1, fr.QQname).Scan(&us.QQname)
		if err1 != nil {
			fmt.Printf("scan failed, err:%v\n", err1)
			sen3 := "delete from"
			sen4 := " where friend=?"
			sqlStr3 := sen3 + " " + qqname + sen4
			_, err3 := db.Exec(sqlStr3, fr.QQname)
			if err3 != nil {
				fmt.Printf("delete failed, err:%v\n", err3)
				return nil
			}
			continue
		}
		if fr.QQname != "" {
			friends = append(friends, fr.QQname)
		}
	}
	return friends
}
func ScanGroup(qqname, group string) []string {
	friends := make([]string, 1)
	friends[0] = group + "组："
	sen1 := "select friend,kind from"
	sen2 := " where id>=1"
	sqlStr := sen1 + " " + qqname + sen2
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var fr model.Friends
		err := rows.Scan(&fr.QQname, &fr.Group)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		sqlStr1 := "select qqname from qquser where qqname=?"
		var us model.Users
		err1 := db.QueryRow(sqlStr1, fr.QQname).Scan(&us.QQname)
		if err1 != nil {
			fmt.Printf("scan failed, err:%v\n", err1)
			sen3 := "delete from"
			sen4 := " where friend=?"
			sqlStr3 := sen3 + " " + qqname + sen4
			_, err3 := db.Exec(sqlStr3, fr.QQname)
			if err3 != nil {
				fmt.Printf("delete failed, err:%v\n", err3)
				return nil
			}
			continue
		}
		if fr.QQname != "" && fr.Group == group {
			friends = append(friends, fr.QQname)
		}
	}
	return friends
}
func ChangGroup(qqname, friend, newGroup string) bool {
	sen1 := "update"
	sen2 := " set kind=? where friend=?"
	sqlStr := sen1 + " " + qqname + sen2
	_, err := db.Exec(sqlStr, newGroup, friend)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return false
	}
	return true
}
func SearchFriend(qqname, friend string) string {
	sen1 := "select kind from"
	sen2 := " where friend=?"
	sqlStr := sen1 + " " + qqname + sen2
	var fr model.Friends
	err := db.QueryRow(sqlStr, friend).Scan(&fr.Group)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return fr.Group
}
func CheckNewFriend(qqname, friend string) bool {
	sen1 := "select friend from"
	sen2 := " where friend=?"
	sqlStr := sen1 + " " + qqname + sen2
	var fr model.Friends
	err := db.QueryRow(sqlStr, friend).Scan(&fr.QQname)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return true
	}
	return false
}
func CheckGroup(qqname, group string) bool {
	sen1 := "select kind from"
	sen2 := " where id >= 1"
	sqlStr := sen1 + " " + qqname + sen2
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var fr model.Friends
		err := rows.Scan(&fr.Group)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return false
		}
		if group == fr.Group {
			return true
		}
	}
	return false
}
