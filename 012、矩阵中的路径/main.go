
var board [][]byte

func exist(bd [][]byte, word string) bool {

    board = bd
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(i, j, 0, []byte(word)) {
                return true
            }
        }
    }

    return false
}

func dfs(i, j, k int, chs []byte) bool {
    
    if k == len(chs) {
        return true
    }

    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != chs[k] {
        return false
    }

    temp := board[i][j]
    board[i][j] = 0

    ret :=  dfs(i+1, j, k+1, chs) || dfs(i-1, j, k+1, chs) || dfs(i, j-1, k+1, chs) || dfs(i, j+1, k+1, chs)
    board[i][j] = temp

    return ret
}

