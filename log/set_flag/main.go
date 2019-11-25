package main

import (
	"log"
	"time"
)

//todo:还有时间格式转换的使用
func main(){
	log.SetFlags(log.Lshortfile|log.Ltime)

	t:=time.Now().UTC().Format("20060102 15:04:05 -07")
	log.Println(t)

	t1,_:=time.Parse("20060102 15:04:05 -07","20190918 22:05:05 +08")

	log.Println(t1.Unix())

	t2,_:=time.Parse("20060102 15:04:05 -07","20190918 22:05:05 +00")

	log.Println(t2.Unix())

	t3:=time.Unix(t1.Unix(),0).Local().Format("20060102 15:04:05 -07")

	log.Println(t3)

	lo,_:=time.LoadLocation("Local")
	log.Println(lo)
}
