let str1 = "status";
let str2 = "sutats";

function checkValidPalindrome(str) {
  let left = 0;
  let right = str.length - 1;

  while (left < right) {
    if (str[left] !== str[right]) {

      return palidrome(str, left + 1, right) || palidrome(str, left, right - 1);
    }
    left++;
    right--;
  }
  return true;
}

function palidrome(str, l, r) {
  while (l < r) {
    if (str[l] !== str[r]) return false; // Fixed 's' to 'str'
    l++;
    r--;
  }
  return true;
}

console.log(checkValidPalindrome(str1)); 
console.log(checkValidPalindrome(str2));
