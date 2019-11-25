package main

import (
	"fmt"
	"sync"
)

func main(){
	a:=GetInstance()
	a.Name="a"
	b:=GetInstance()
	b.Name="b"
	fmt.Println(&a.Name,a)
	fmt.Println(&b.Name,b)
	fmt.Printf("%p %T\n",a,b)
	fmt.Printf("%p %T\n",a,b)
}

var instance *single
var lock sync.Mutex

type single struct{
	Name string
}

func GetInstance()*single{
	if instance==nil{
		lock.Lock()
		defer lock.Unlock()
		instance=&single{}
	}
	return instance
}