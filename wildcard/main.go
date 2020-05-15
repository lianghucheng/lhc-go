package main

import (
	"fmt"
	"strings"
)

func main() {
	reStr := "o*m"
	tagStr := "shopeemobile.com"

	//fmt.Scanf("%s\n%s", &reStr, &tagStr)
	remain := tagStr

	matchs := strings.Split(reStr, "*") //?   *abc*ef

	rtIndex := []int{}
	rubbish := 0
	flag := -1
	for {
		if flag == 1 {
			break
		}
		firstMatch := -1
		firstLength := 0
		indexContainer := []int{}
		isFirst := true
		for k, v := range matchs {
			index := strings.Index(remain, v)
			if len(v) != 0 && isFirst{
				firstMatch = index + rubbish
				firstLength = len(v)
				isFirst = false
			} else if k == 0 && len(v) == 0 {
				continue
			} else if k == len(matchs) - 1 && len(v) == 0{
				index = len(remain) - 1
			} else if len(v) == 0 {
				continue
			}

			if len(indexContainer) > 0 {
				if index <= indexContainer[len(indexContainer) - 1] && index != -1 {
					tempIndex := index
					tempReamin := remain
					tempRubbish := 0
					for {
						tempReamin = tempReamin[:tempIndex] + tempReamin[tempIndex + len(v):]
						tempRubbish += len(v)

						tempIndex = strings.Index(tempReamin, v)

						if tempIndex + tempRubbish + rubbish > indexContainer[len(indexContainer) - 1] || tempIndex == -1 {
							index = tempIndex + tempRubbish
							break
						}
					}
				}
			}

			if index == -1 && firstMatch - rubbish != -1 {
				remain = remain[:firstMatch - rubbish] + remain[firstMatch - rubbish + firstLength:]
				rubbish += len(matchs[0]) // ? *asd*
				indexContainer = []int{}
				break
			} else if firstMatch - rubbish < 0 {
				flag = 1
				indexContainer = []int{}
				break
			}

			indexContainer = append(indexContainer, index + rubbish)
		}

		if len(indexContainer) > 0 {
			length := indexContainer[len(indexContainer)-1] - indexContainer[0] + len(matchs[len(matchs) - 1])

			rtIndex = append(rtIndex, indexContainer[0], length)

			remain = remain[:firstMatch - rubbish] + remain[firstMatch - rubbish + firstLength:]
			rubbish += firstLength
		}
	}

	rt := [][]int{}

	for i := 0;i < len(rtIndex)/2; i++ {
		if len(matchs[0]) == 0 {
			for j := 0; j < rtIndex[2*i]; j++ {
				rt = append(rt,[]int{j, rtIndex[2*i] - j + rtIndex[2*i + 1]})
			}
		}

		rt = append(rt, []int{rtIndex[2*i],rtIndex[2*i + 1]})
	}

	for i := 0; i < len(rt); i++ {
		for j := 0; j < len(rt) - i - 1; j++ {
			if rt[j][0] > rt[j + 1][0] {
				temp := rt[j]
				rt[j] = rt[j + 1]
				rt[j + 1] = temp
			}
			if rt[j][0] == rt[j + 1][0] && rt[j][1] > rt[j + 1][1] {
				temp := rt[j]
				rt[j] = rt[j + 1]
				rt[j + 1] = temp
			}
		}
	}

	for _,v := range rt{
		fmt.Println(v[0],v[1])
	}
}