package main

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
type Users struct{
	Id int `gorm:"primary_key"`
	Name string
	PassWord string
	RealName string
	Phone string
	Birthday string
	CreateDate string
	WeChat string
	QQ string
	Like string
	SelfText string
}
var Db *sql.DB
func init(){
var err error
Db,err=sql.Open("mysql","root:123456@tcp(00.00.00.00:3306)/test")
   if err!=nil{
	   log.Fatal(err)
	   return
   }

}
