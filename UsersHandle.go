package main

import (
	"log"
)


/* 通过ID查询用户
*/
func GetUserById(id int) (users Users,err error){
  users=Users{}
  err=Db.QueryRow("select id,name from users where id=?",id).Scan(&users.Id,&users.Name)
  return
}


/* 查询所有用户
*/
func GetUser() []Users{
var user Users
var userList []Users
rows,_:=Db.Query("select id,name from users")
for rows.Next(){
	err:=rows.Scan(&user.Id,&user.Name)
	if err!=nil{
		log.Fatal(err)
		return nil
	}
    userList=append(userList,user)
}
return userList
}

/*修改指定用户
*/
func ChangeUserByKey(Key string) (changenum int64,err error){

   sql:="update users set real_name=? where id = 1"
   stmt,err:=Db.Prepare(sql)
   if err!=nil{
	   return
   }
   defer stmt.Close()
   res,err:=stmt.Exec(Key)
   changenum,err=res.RowsAffected()
   return
}

/* 创建用户
*/
func AddUser(user Users) (id int64,err error){
stmt, err := Db.Prepare(`insert into users(name,pass_word,real_name,phone,birthday,create_date,we_chat,qq,like,self_text) values(?,?,?,?,?,?,?,?,?,?)`)
    if err!=nil{
      return
	}
	defer stmt.Close()
    res, err := stmt.Exec(user.Name,user.PassWord,user.RealName,user.Phone,user.Birthday,user.CreateDate,user.WeChat,user.QQ,user.Like,user.SelfText)
      
	id, err = res.LastInsertId()
	if err != nil{
		return
	}

 return
}
