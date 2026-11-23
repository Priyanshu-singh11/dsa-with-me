func searchMatrix(arr [][]int, target int) bool {
    for i:=0;i<len(arr);i++{
        if arr[0][i]==target || arr[1][i]==target || arr[2][i] == target{
            return true
        }
    }
    return false
}
