
func searchMatrix(arr [][]int, target int) bool {
    m := len(arr)-1
    n := len(arr[0])-1
    row := 0
    col := n
    for row<=m && col>= 0{
        if arr[row][col]==target{
            return true
        }else if target>arr[row][col]{
            row++
        }else{
            col--
        }
    }
    return false
}

