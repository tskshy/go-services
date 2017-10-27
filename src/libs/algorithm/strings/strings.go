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
	for i := 0; i < len(atr); {
		for j < len(apr) {
			var char_i = atr[i]
			var char_j = apr[j]

			fmt.Println(fmt.Sprintf("%2d, %2d, %s, %s", i, j, string(char_i), string(char_j)))

			if char_i != char_j {
				if j <= 0 {
					j = 1
				}
				j = arr[j-1] + 1

				if j == 0 {
					fmt.Println("--")
					i++
				}
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

var MP1 = func(target, pattern string) bool {
	var fun = func(r []rune) []int {
		var next_arr = make([]int, len(r))

		for n := 0; n < len(next_arr); n++ {
			next_arr[n] = -1
		}

		var i int = 0
		var j int = 0

		next_arr[i] = 0
		i++

		fmt.Println("len(r)", len(r))

		for i < len(r) {
			var char_i = r[i]
			var char_j = r[j]

			// 0 0 1 1 2 0 1 2 3 4 5
			fmt.Println(next_arr, i, j)

			if j == 0 && char_i != char_j {
				if next_arr[i] == -1 {
					next_arr[i] = 0
				}

				i++
				j = next_arr[j] //j always 0
				continue
			}

			if j == 0 && char_i == char_j {
				next_arr[i] = 1

				i++
				j++
				continue
			}

			if j != 0 && char_i != char_j {
				if next_arr[i] == -1 {
					//next_arr[i] = 0
				}

				j = next_arr[j]
				continue
			}

			if j != 0 && char_i == char_j {
				if next_arr[i] == -1 {
					next_arr[i] = next_arr[i-1] + 1
				}

				i++
				j++
				continue
			}

		}

		return next_arr
	}

	var atr = []rune(target)
	var apr = []rune(pattern)

	if len(atr) < len(apr) || len(atr) == 0 || len(apr) == 0 {
		return false
	}
	var next_arr = fun(apr)
	fmt.Println(next_arr)
	return false
}

//http://www.cnblogs.com/SYCstudio/p/7194315.html
//http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
