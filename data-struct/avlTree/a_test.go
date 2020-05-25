package main

import (
	"fmt"
	"testing"
)

var llTest1 =[]int{10,5,20,3,7,1}
var llTest2 = []int{10,5,20,3,7,4}
var rrTest1 = []int{10,5,20,17,23,21}
var rrTest2 = []int{10,5,20,17,23,26}
var lrTest1 = []int{10,5,20,3,7,8}
var lrTest2 = []int{10,5,20,3,7,6}
var rlTest1 = []int{10,5,20,17,23,16}
var rlTest2 = []int{10,5,20,17,23,18}

func TestShow(t *testing.T) {
	aT := toTestLL(llTest2)
	aT.root.show(0)
}

func TestShow2(t *testing.T) {
	at := toTestLL(llTest2)
	at.show2()
}

/**
1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:1)
|-------------------------------------------1(freq:0,heigh:0)
|-----------------------------7(freq:0,heigh:0)
|---------------20(freq:0,heigh:0)
2、期望输出：
|-5(freq:0,heigh:2)
|---------------3(freq:0,heigh:1)
|-----------------------------1(freq:0,heigh:0)
|---------------10(freq:0,heigh:3)
|-----------------------------7(freq:0,heigh:0)
|-----------------------------20(freq:0,heigh:0)



1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:1)
|-------------------------------------------4(freq:0,heigh:0)
|-----------------------------7(freq:0,heigh:0)
|---------------20(freq:0,heigh:0)
2、期望输出：
|-5(freq:0,heigh:2)
|---------------3(freq:0,heigh:1)
|-----------------------------4(freq:0,heigh:0)
|---------------10(freq:0,heigh:3)
|-----------------------------7(freq:0,heigh:0)
|-----------------------------20(freq:0,heigh:0)
 */
func TestAvlTree_SingleRotateLeft(t *testing.T) {
	at := toTestLL(llTest2)
	fmt.Println("before")
	at.show(at.root)
	at.SingleRotateLeft(&at.root)
	fmt.Println("after")
	at.show(at.root)
}
/**
1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------17(freq:0,heigh:0)
|-----------------------------23(freq:0,heigh:1)
|-------------------------------------------21(freq:0,heigh:0)
2、期望输出：
|-20(freq:0,heigh:2)
|---------------10(freq:0,heigh:3)
|-----------------------------5(freq:0,heigh:0)
|-----------------------------17(freq:0,heigh:0)
|---------------23(freq:0,heigh:1)
|-----------------------------21(freq:0,heigh:0)



1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------17(freq:0,heigh:0)
|-----------------------------23(freq:0,heigh:1)
|-------------------------------------------26(freq:0,heigh:0)
2、期望输出：
|-20(freq:0,heigh:2)
|---------------10(freq:0,heigh:3)
|-----------------------------5(freq:0,heigh:0)
|-----------------------------17(freq:0,heigh:0)
|---------------23(freq:0,heigh:1)
|-----------------------------26(freq:0,heigh:0)
 */
func TestAvlTree_SingleRotateRight(t *testing.T) {
	at := toTestLL(rrTest2)
	fmt.Println("before")
	at.show(at.root)
	at.SingleRotateRight(&at.root)
	fmt.Println("after")
	at.show(at.root)
}
/**
1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:0)
|-----------------------------7(freq:0,heigh:1)
|-------------------------------------------8(freq:0,heigh:0)
|---------------20(freq:0,heigh:0)
2、期望输出：
|-7(freq:0,heigh:1)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:0)
|---------------10(freq:0,heigh:3)
|-----------------------------8(freq:0,heigh:0)
|-----------------------------20(freq:0,heigh:0)



1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:0)
|-----------------------------7(freq:0,heigh:1)
|-------------------------------------------6(freq:0,heigh:0)
|---------------20(freq:0,heigh:0)
2、期望输出：
|-7(freq:0,heigh:1)
|---------------5(freq:0,heigh:2)
|-----------------------------3(freq:0,heigh:0)
|-----------------------------6(freq:0,heigh:0)
|---------------10(freq:0,heigh:3)
|-----------------------------20(freq:0,heigh:0)
*/
func TestAvlTree_DoubleRotateLR(t *testing.T) {
	at := toTestLL(lrTest2)
	fmt.Println("before")
	at.show(at.root)
	at.DoubleRotateLR(&at.root)
	fmt.Println("after")
	at.show(at.root)
}
/**
1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------17(freq:0,heigh:1)
|-------------------------------------------16(freq:0,heigh:0)
|-----------------------------23(freq:0,heigh:0)
2、期望输出：
|-17(freq:0,heigh:1)
|---------------10(freq:0,heigh:3)
|-----------------------------5(freq:0,heigh:0)
|-----------------------------16(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------23(freq:0,heigh:0)



1、输入：
|-10(freq:0,heigh:3)
|---------------5(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------17(freq:0,heigh:1)
|-------------------------------------------18(freq:0,heigh:0)
|-----------------------------23(freq:0,heigh:0)
2、期望输出：
|-17(freq:0,heigh:1)
|---------------10(freq:0,heigh:3)
|-----------------------------5(freq:0,heigh:0)
|---------------20(freq:0,heigh:2)
|-----------------------------18(freq:0,heigh:0)
|-----------------------------23(freq:0,heigh:0)
*/
func TestAvlTree_DoubleRotateRL(t *testing.T) {
	at := toTestLL(rlTest2)
	fmt.Println("before")
	at.show(at.root)


}