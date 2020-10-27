type Point struct{
	x,y int
}

func movingCount(m int, n int, k int) int {
	
	visited := map[Point]bool{}
	
    dfs(m,n,0,0,k,&visited) //返回的是最长不重复路径的长度
    return len((visited)) //可访问到的位置个数
}

func dfs(m, n, i, j, k int, visited *map[Point]bool ) int {

	point := Point{i,j}

    if i < 0 || i >= m || j < 0 || j>= n || !fit(i,j,k) || (*visited)[point]{
        return 0
	}
	
	(*visited)[point] = true

    up := dfs(m, n, i+1, j, k, visited)
    down := dfs(m, n, i-1, j, k, visited)
    right := dfs(m, n, i, j+1, k, visited)
    left := dfs(m, n, i, j-1, k, visited)
    

    max := right
    if left > right {
        max = left
    }

    if up > max {
        max = up
    }

    if down > max {
        max = down
    }

    return 1 + max

}



func fit(ii,jj,k int) bool {
    sum := 0
    for ii != 0 || jj != 0 {
        sum += ii%10 + jj%10
        if sum > k {
            return false
        }
        ii = ii / 10
        jj = jj / 10
    }

    return true
}