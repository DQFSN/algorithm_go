
func maxValue(grid [][]int) int {

    if len(grid) == 0 {
        return 0
    }

    r,c := len(grid)-1,len(grid[0])-1

    for i:=1;i<=c;i++{
        grid[0][i] = grid[0][i-1] + grid[0][i]
    }
    for i:=1;i<=r;i++{
        grid[i][0] = grid[i-1][0] + grid[i][0]
    }

    for i:=1;i<=r;i++ {
        for j:=1;j<=c;j++ {
            left := grid[i][j-1]
            up := grid[i-1][j]

            if up > left {
                grid[i][j] = up + grid[i][j]
            }else {
                grid[i][j] = left + grid[i][j]
            }
        }
    }

    return grid[r][c]

    // return findMaxValue(len(grid)-1,len(grid[0])-1,0,0,grid)

    // return getMaxVaule(len(grid)-1, len(grid[0])-1,grid)
}

func getMaxVaule(i,j int, grid [][]int) int {

    if i == 0 && j == 0 {
        return grid[0][0]
    }

    if i == 0 {
        return grid[i][j] + getMaxVaule(i,j-1,grid)
    }

    if j == 0 {
        return grid[i][j] + getMaxVaule(i-1,j,grid)
    }

    left := grid[i][j] + getMaxVaule(i,j-1,grid)
    up := grid[i][j] + getMaxVaule(i-1,j,grid)

    if left > up {
        return left
    }else {
        return  up
    }

}

func findMaxValue(r,c,i,j int, grid [][]int) int {

    // if i > r || j > c  {
    //     return 0
    // }

    if i == r && j == c {
        return grid[i][j]
    } 

    if i == r {
        return grid[i][j] + findMaxValue(r,c,i,j+1,grid)
    }

    if j == c {
        return grid[i][j] + findMaxValue(r,c,i+1,j,grid)
    }

    cur := grid[i][j]
    right := findMaxValue(r,c,i,j+1,grid)
    down := findMaxValue(r,c,i+1,j,grid)

    if right > down {
        return cur + right
    }
    return down + cur

}