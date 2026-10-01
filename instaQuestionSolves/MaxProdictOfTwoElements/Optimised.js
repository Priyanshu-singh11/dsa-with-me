function maxProductOfTwoElementOptimize(arr) {
  let maxProduct = 1
  let largest = arr[0]
  let secondLargest = arr[1]

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] > largest) {
      secondLargest = largest
      largest = arr[i]
    } else if (arr[i] > secondLargest) {
      secondLargest = arr[i]
    }
  }

  maxProduct = secondLargest * largest
  return maxProduct
}

let arr = [3, 7, 2, 9, 5]

let ans = maxProductOfTwoElementOptimize(arr)

console.log(ans)
