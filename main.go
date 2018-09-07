package main

import (
	"log"
	"encoding/json"
	"strconv"
	"reflect"
  "net/http"
  "github.com/julienschmidt/httprouter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type RequestMsg struct{
    msg string
    data interface{}
    code int
}

func main(){
  //初始化新路由器
  mux:=httprouter.New()
  
  mux.POST("/HandleRequest/:name",handleRequest)

  server:=http.Server{
    Addr:"127.0.0.1:2020",
    Handler:mux,
  }
  server.ListenAndServe()
}
var maps=map[string]interface{}{"GetUserByIdManager":GetUserByIdManager,"AddUserManager":AddUserManager,"GetUserManager":GetUserManager,"ChangeUserByKeyManager":ChangeUserByKeyManager}

func handleRequest(w http.ResponseWriter,r *http.Request,h httprouter.Params){
  r.ParseForm()
   h1:=maps[h.ByName("name")]
   v1:=reflect.ValueOf(h1)
   params:=make([]reflect.Value,2)
   params[0]=reflect.ValueOf(w)
   params[1]=reflect.ValueOf(r)
   rs:=v1.Call(params)
   output:=rs[0].Interface().([]byte)
   w.Header().Set("Content-type","application/json")
   w.Write(output)
   return
}


func AddUserManager(w http.ResponseWriter,r *http.Request) []byte{
  user:=Users{
    Id:1,
    Name:"特殊t",
    PassWord:"123666",
    RealName:"asd",
    Phone:"13524874512",
    Birthday:"2015-02-01",
    CreateDate:"2017-05-02",
    WeChat:"wwww",
    QQ :"qq",
    Like:"not bug is run go",
    SelfText:"asdasd",
  }
  msg:=RequestMsg{}
  
    _,err:=AddUser(user)
    if err!=nil{
    msg=RequestMsg{
      msg:"创建成功",
      data:user,
      code:0,
    }
  }
  
  output,err:=json.MarshalIndent(&msg,"","\t\t")

  return output
}

func GetUserByIdManager(w http.ResponseWriter,r *http.Request) []byte{
  r.ParseForm()
  user:=Users{}
  msg:=make(map[string]interface{})
  id,err:=strconv.Atoi(r.FormValue("id"))
  if err!=nil{
    // log.Fatal(err)
    msg["msg"]="查询失败"
    msg["data"]=user
    msg["code"]=-1
    

  }else{

    user,err=GetUserById(id)
    if err!=nil{
      msg["msg"]="查询失败"
      msg["data"]=""
      msg["code"]=-1
    }else{
      msg["msg"]="查询成功"
      msg["data"]=user
      msg["code"]=0
    }
  }
  if err!=nil{
      msg["msg"]="查询失败"
      msg["data"]=""
      msg["code"]=-1
  }
  log.Println(msg)
  output,err:=json.MarshalIndent(&msg,"","\t\t")
  return output
}

func GetUserManager(w http.ResponseWriter,r *http.Request) []byte{
   uList:=GetUser()
   output,err:=json.MarshalIndent(&uList,"","\t\t")
   if err!=nil{
     log.Fatal(err)
   }
   return output
}

func ChangeUserByKeyManager(w http.ResponseWriter,r *http.Request) []byte{
 num,err:=ChangeUserByKey("test")
 if err!=nil{
   log.Fatal(err)
 }
 output,err:=json.MarshalIndent(&num,"","\t\t")
 return output
}