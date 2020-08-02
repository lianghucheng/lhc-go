package main

import (
	"fmt"
	"time"
	"github.com/yeka/zip"
	"log"
		"sync"
	"io/ioutil"
)

var (
	charMap = `123456789-=qwertyuiop[]\asdfghjkl;'zxcvbnm,./!@#$%^&*()_+QWERTYUIOP{}|ASDFGHJKL:"ZXCVBNM<>?~`+"`0"
	length = 3
	pass = make([]uint8, length)
	findFlag = 0
	wg = &sync.WaitGroup{}
)
func main()  {
	//fmt.Println(isAnswer([]uint8("bq1")))
	wg.Add(3)
	div := len(charMap) / 3

	start1 := make([]int, length)
	end1 := make([]int, length)
	end1[0] = div
	go powerDecode(start1, end1,1)
	start2 := make([]int, length)
	start2[0] = div
	end2 := make([]int, length)
	end2[0] = 2*div
	go powerDecode(start2, end2,2)
	start3 := make([]int, length)
	start3[0] = 2*div
	end3 := make([]int, length)
	go powerDecode(start3, end3,3)
	wg.Wait()
	fmt.Println(string(pass))
}

func powerDecode(start, end []int, i int ) {
	wgg := make(chan int)
	ret := make([]uint8, length)
	initFlag := 0
	t1 := time.Now().Unix()
	c := 0
	for {
		c++
		//fmt.Println("协程",i)
		if findFlag == 1 {
			fmt.Println("协程",i,"结束")
			break
		}
		if isBreak(start, end) && initFlag == 1 {
			break
		}
		initFlag = 1
		ret = getPass(start)
		addCnt(start)
		go func(result []uint8) {
			if isAnswer(result) {
				pass = result
				//设置标志标识结束
				fmt.Println("设置标志标识结束", string(result))
				t2 := time.Now().Unix()
				fmt.Println(t2 - t1)
				findFlag = 1
			}
			wgg <- 1
		}(ret)
	}
	fmt.Println("协程",i, start)
	for {
		c -= <- wgg
		if c == 0 {
			break
		}
	}
}
//只要不是闭环，就不用考虑首尾一样
func isBreak(cnt, end []int) bool {
	for k := range cnt {
		if cnt[k] != end[k] {
			return false
		}
	}
	return true
}

func getPass(cnt []int) []uint8 {
	ret := make([]uint8, length)
	for i:=0;i<length;i++ {
		ret[i] = charMap[cnt[i]]
	}
	return ret
}

func addCnt(cnt []int) {
	cnt[length-1]++
	for i:=length-1;i>=0;i--{
		if cnt[i] == len(charMap) {
			cnt[i] = 0
			if i - 1 >=0 {
				cnt[i-1]++
			}
		}
	}
}

func isAnswer(ret []uint8) bool {
	if findFlag == 1 {
		fmt.Println("协程","结束")
		return false
	}
	r, err := zip.OpenReader("./test.zip")
	defer r.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, f := range r.File {
		if f.IsEncrypted(){
			file := f
			file.SetPassword(string(ret))
			r, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			_, err = ioutil.ReadAll(r)
			r.Close()
			if err == nil {
				return true
			}
			break
		}
	}
	return false
}