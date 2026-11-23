func searchMatrix(arr [][]int, target int) bool {
    //m := len(arr)-1
    n := len(arr[0])
    start := 0
    end := len(arr)-1
    for start<=end{
        med := start+(end-start)/2
        if arr[med][0]<=target && arr[med][n-1]>=target{
            //this is row
            return SearchNumber(arr, target,med)
        }else if arr[med][0]<=target{
            //up side row
            start = med+1
            
        }else{
            //low side row
            end = med-1
            
        
    }
    
    }
    return false
    }


func SearchNumber(arr [][]int,target int,mid int)bool{
    start := 0
    end := len(arr[mid])-1
    for start<=end{
        med := start+(end-start)/2
        if target>arr[mid][med]{
            start = med+1
        }else if target<arr[mid][med]{
            end = med-1
        }else{
            return true 
        }
    }
    return false
}
