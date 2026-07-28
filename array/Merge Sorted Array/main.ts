function MergeSortedArray(
  num1: number[],
  m: number,
  num2: number[],
  n: number
): number[] {

  let i = m - 1;
  let j = n - 1;
  let idx = num1.length - 1;

  while (i >= 0 && j >= 0) {
    if (num1[i] > num2[j]) {
      num1[idx] = num1[i];
      i--;
    } else {
      num1[idx] = num2[j];
      j--;
    }
    idx--;
  }

  while (j >= 0) {
    num1[idx] = num2[j];
    j--;
    idx--;
  }

  return num1;
}

const num1 = [1, 2, 3, 0, 0, 0];
const num2 = [2, 5, 6];

const m = 3;
const n = 3;

const result: number[] = MergeSortedArray(num1, m, num2, n);

console.log(result);
