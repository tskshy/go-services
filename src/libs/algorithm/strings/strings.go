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

var MP = func(target, pattern string, use_kmp bool) (found int) {
	var fun = func(r []rune) []int {
		var next = make([]int, len(r)+1)

		var i int = 0
		var j int = -1

		next[0] = -1

		for i < len(r) {
			if j == -1 || r[i] == r[j] {
				i++
				j++

				/*
				 KMP(1977年)是MP(1970年)算法的改进，`use_kmp`分支是它们之间的区别
				*/
				if use_kmp {
					if i == len(r) || r[i] != r[j] {
						next[i] = j
					} else {
						next[i] = next[j]
					}
				} else {
					next[i] = j
				}
			} else {
				j = next[j]
			}
		}

		return next
	}

	var atr = []rune(target)
	var apr = []rune(pattern)

	if len(atr) < len(apr) || len(atr) == 0 || len(apr) == 0 {
		return 0
	}
	var next = fun(apr)
	fmt.Println(next)

	var i int = 0
	var j int = 0

	for i < len(atr) {
		if j == -1 || atr[i] == apr[j] {
			i++
			j++
		} else {
			j = next[j]
		}

		if j >= len(apr) {
			/**TODO del*/
			var info = fmt.Sprintf("[%d %d)", i-j, i)
			fmt.Println("found", info)
			/***/

			j = next[j]
			found++
		}
	}

	return found
}

//http://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
