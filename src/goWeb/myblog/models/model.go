package models

import (
	"fmt"
	"goWeb/myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0是正常状态，1 是删除
	Createtime int64
}

type Album struct {
	Id 			int
	Filepath 	string
	Filename 	string
	Status 		int 
	Createtime 	int64
}

// ----------------数据库操作--------------------
// 插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	fmt.Println("row", row)
	id := 0
	row.Scan(&id)
	fmt.Println("row2", &row)
	fmt.Println("row2", id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

// 插入图片
func InsertAlbum(album Album) (int64,error) {
	return utils.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",album.Filepath,album.Filename,album.Status,album.Createtime)
}

//查询图片
func FindAllAlbums()([]Album,error){
	rows,err :=utils.QueryDB("selfct id,filepath,filename,status,createtime,form album")
	if err !=nil{
		return nil,err
	}
	var albums []Album
	for rows.Next() {
		id :=0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id,&filepath,&filename,&status,&createtime)
		album := Album{id,filepath,filename,status,createtime}
		albums =append(albums,album)
	}
	return albums,nil
}