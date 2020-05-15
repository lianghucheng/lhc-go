package main

import "fmt"

func main() {
	//var x, y, n int
	//var xn, yn []int
	//fmt.Scanf("%d%d%d", &x, &y, &n)
	//for i := 0; i < n; i++ {
	//	val := 0
	//	fmt.Scanf("%d", &val)
	//	xn = append(xn, val)
	//	fmt.Scanf("%d",&val)
	//	yn = append(yn, val)
	//}
	//
	//var themap [][]int
	//for i := 0; i <= x; i++ {
	//	var temp []int
	//	for j := 0; j <= y; j++ {
	//		temp = append(temp, 0)
	//	}
	//	themap = append(themap, temp)
	//}
	//
	//for i := 0; i < n; i++ {
	//	themap[xn[i]][yn[i]] = -1
	//}
	//
	//fmt.Printf("%v",roadCount2(themap, []int{0, 0}))
	treeDp()
}

//未通过，障碍物为-1		完全繁琐的想法，时间复杂度极高
func roadCount(themap [][]int) (rt int) {
	x := len(themap)
	if x == 0 {
		return
	}
	y := len(themap[0])
	if y == 0 {
		return
	}

	if themap[0][0] == -1 {
		return
	}
	stack := []int{0, 0}

	for {
		if len(stack) == 0 {
			break
		}
		s := append([]int{},stack[len(stack) - 2 :]...)
		stack = stack[:len(stack) - 2]

		tx, ty := s[0], s[1] + 1

		temp := 0

		if ty < y {
			if themap[tx][ty] == 0 {
				stack = append(stack, tx, ty)
				temp++
			}
			if themap[tx][ty] == -1 {
				if tx + 1 >= x {
					rt--
				} else if themap[tx + 1][ty] == -1{
					rt--
				}
			}
		}
		tx, ty = s[0] + 1, s[1]
		if tx < x {
			if themap[tx][ty] == 0 {
				stack = append(stack, tx, ty)
				temp++
			}
			if themap[tx][ty] == -1 {
				if ty + 1 >= y {
					rt--
				} else if themap[tx][ty + 1] == -1 {
					rt--
				}
			}
		}
		if temp == 2 {
			rt++
		}
	}
	return
}

//未通过，障碍物为-1		不繁琐，时间复杂度正常，但是思想错误
func roadCount1(themap [][]int, branch []int, xc, yc int) (rt int) {
	xl := len(themap)
	if xl == 0 {
		return
	}
	yl := len(themap[0])
	if yl == 0 {
		return
	}

	if themap[0][0] == -1 {
		return
	}

	if xc == xl - 1 && yc == yl - 1 {
		themap[xc][yc] = 1
		rt = 1
		return
	}

	for i := 0; i < len(branch)/2; i++ {
		x := branch[2 * i]
		y := branch[2 * i + 1]

		newBranch := []int{}

		tx, ty := x, y + 1
		//todo:
		if ty < yl {
			if themap[tx][ty] == 0 {
				newBranch = append(newBranch, tx, ty)
			} else if themap[tx][ty] > 0 {
				rt += themap[tx][ty]
			}
		}

		tx, ty = x + 1, y
		//todo:
		if tx < xl {
			if themap[tx][ty] == 0 {
				newBranch = append(newBranch, tx, ty)
			} else if themap[tx][ty] > 0 {
				rt += themap[tx][ty]
			}
		}

		rt += roadCount1(themap, newBranch, x, y) //试试remain，不知道会不会内容超出
	}
	themap[xc][yc] = rt

	return
}

//通过，障碍物为-1
func roadCount2(themap [][]int, branch []int) (rt int) {
	xl := len(themap)
	if xl == 0 {
		return
	}
	yl := len(themap[0])
	if yl == 0 {
		return
	}

	if themap[0][0] == -1 {
		return
	}
	x := branch[0]
	y := branch[1]
	if x == xl - 1 &&  y == yl - 1 {
		themap[x][y] = 1
		rt = 1
		return
	}
	tx, ty := x, y + 1
	//todo:
	if ty < yl {
		if themap[tx][ty] == 0 {
			rt += roadCount2(themap, []int{tx, ty})
		} else if themap[tx][ty] > 0 {
			rt += themap[tx][ty]
		}
	}

	tx, ty = x + 1, y
	//todo:
	if tx < xl {
		if themap[tx][ty] == 0 {
			rt += roadCount2(themap, []int{tx, ty})
		} else if themap[tx][ty] > 0 {
			rt += themap[tx][ty]
		}
	}

	//rt += roadCount2(themap, newBranch, x, y) //试试remain，不知道会不会内容超出

	themap[x][y] = rt

	return
}

//通过，障碍物为-1
func roadCount3(themap [][]int, branch *[]int) (rt int) {
	xl := len(themap)
	if xl == 0 {
		return
	}
	yl := len(themap[0])
	if yl == 0 {
		return
	}

	if themap[0][0] == -1 {
		return
	}
	x := (*branch)[0]
	y := (*branch)[1]
	if x == xl - 1 &&  y == yl - 1 {
		themap[x][y] = 1
		rt = 1
		return
	}
	tx, ty := x, y + 1
	//todo:
	if ty < yl {
		if themap[tx][ty] == 0 {
			(*branch)[0] = tx
			(*branch)[1] = ty
			rt += roadCount3(themap, branch)
		} else if themap[tx][ty] > 0 {
			rt += themap[tx][ty]
		}
	}

	tx, ty = x + 1, y
	//todo:
	if tx < xl {
		if themap[tx][ty] == 0 {
			(*branch)[0] = tx
			(*branch)[1] = ty
			rt += roadCount3(themap, branch)
		} else if themap[tx][ty] > 0 {
			rt += themap[tx][ty]
		}
	}

	//rt += roadCount2(themap, newBranch, x, y) //试试remain，不知道会不会内容超出

	themap[x][y] = rt

	return
}

//通过，障碍物为1，dp
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n ==0{
		return 0
	}
	if m == 1 || n == 1 {
		for i:= range obstacleGrid {
			for j:= range obstacleGrid[i] {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
			}
		}
		return 1
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp:=make([][]int, m)
	for i:=range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		if (dp[i-1][0] == 1 || i == 1) && obstacleGrid[i][0] == 0 {
			dp[i][0]=1
		}else{
			dp[i][0]=0
		}
	}
	for i:=1;i<n;i++{
		if(dp[0][i-1]==1||i==1)&&obstacleGrid[0][i]==0{
			dp[0][i]=1
		}else{
			dp[0][i]=0
		}
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if obstacleGrid[i][j]==0{
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}else{
				dp[i][j]=0
			}
		}
	}
	return dp[m-1][n-1]
}

func treeDp(){
	x,y,n := 0,0,0
	fmt.Scanf("%d%d%d",&x,&y,&n)
	grip := make([][]int, x+1)
	for k:=range grip{
		grip[k] = make([]int, y+1)
	}
	xn,yn :=make([]int, n),make([]int,n)
	for k:=range xn {
		fmt.Scanf("%d",&xn[k])
		fmt.Scanf("%d",&yn[k])
		grip[xn[k]][yn[k]] = -1
	}
	for k:=range grip{
		grip[k][0] = 1
	}
	for k:=range grip[0] {
		grip[0][k] = 1
	}
	for i:=1;i<x+1;i++{
		for j:=1 ;j<y+1;j++ {
			if grip[i][j] == -1 {
				grip[i][j] = 0
			} else {
				grip[i][j] = grip[i-1][j] + grip[i][j - 1]
			}
			//if grip[i-1][j] !=-1 {
			//	grip[i][j] += grip[i-1][j]
			//}
			//if grip[i][j-1] != -1 {
			//	grip[i][j]+=grip[i][j-1]
			//}
		}
	}

	printArr(grip)
}

func printArr(arr [][]int){
	for _,v1:=range arr{
		for _,v2:=range v1 {
			fmt.Printf("%-2v ",v2)
		}
		fmt.Println()
	}

}

//划分子集合
func assignSubOrder() {
	var n int
	fmt.Scanf("%d",&n)
	s := n*(n+1)
	if s%4 !=0 {
		fmt.Println(0)
		return
	}
	s /=4
	dyn := make([]uint64,s+1)
	dyn[0] =1
	for i:=1 ;i<=n ;i++ {
		for j:=s;j>=i;j-- {
			if j-i >=0 {
				dyn[j] += dyn[j-i]
			}
		}
	}

	fmt.Println(dyn[s]/2)
	fmt.Println(dyn)
}