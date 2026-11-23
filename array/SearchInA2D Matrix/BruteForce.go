func searchMatrix(arr [][]int, target int) bool {
    for i:=0;i<len(arr);i++{
        for j:=0;j<len(arr[i]);j++{
            if arr[i][j]==target{
                return true
            }
        }
    }
    return false
}
