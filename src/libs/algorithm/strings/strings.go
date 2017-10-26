package strings

import "fmt"

/*字符串与模式匹配 常见算法*/

/*检查target中是否包含pattern*/
var BruteForce = func(target, pattern string) bool {
	var atr = []rune(target)
	var apr = []rune(pattern)

	if len(atr) < len(apr) || len(atr) == 0 || len(apr) == 0 {
		return false
	}

	for i := 0; i < len(atr)-len(apr)+1; i++ {
		for j, char_p := range apr {
			if char_p != atr[i+j] {
				break
			}

			if j == len(apr)-1 {
				return true
			}
		}
	}

	return false
}

var MP = func(target, pattern string) bool {
	var _lambda = func(r []rune) []int {
		var a = make([]int, len(r))

		for i := 0; i < len(r); {
			if i == 0 {
				a[i] = -1
				i++
				continue
			}

			for _, char := range r {
				if char == r[i] {
					a[i] = a[i-1] + 1
					i++
				} else {
					a[i] = -1
					break
				}
			}

			i++
		}
		return a
	}

	var atr = []rune(target)
	var apr = []rune(pattern)

	if len(atr) < len(apr) || len(atr) == 0 || len(apr) == 0 {
		return false
	}

	var arr = _lambda(apr)
	fmt.Println(arr)

	var j = 0
	for i := 0; i < len(atr)-len(apr)+1; {
		for j < len(apr) {
			fmt.Print(i, j, " ")
			var char_i = atr[i]
			var char_j = apr[j]

			fmt.Println(string(char_i), string(char_j))
			if char_i != char_j {
				j = arr[j] + 1
				i++
				break
			}

			if char_i == char_j && j < len(apr)-1 {
				j++
				i++
				continue
			}

			return true
		}
	}

	return false
}
