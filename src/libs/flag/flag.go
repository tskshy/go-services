package flag

import (
	"fmt"
	"os"
)

/*
 get the value that CMDLINE provided, if not, return the `value`(default)
 cmdline: ./cmd args0 args1 ...
*/
func Parse(key, value string) string {
	var args = os.Args[1:]

	/*parse error*/
	if len(args)%2 != 0 {
		return value
	}

	for i := 0; i < len(args); i += 2 {
		var key_a = args[i]

		var key_1 = fmt.Sprintf("-%s", key)
		var key_2 = fmt.Sprintf("--%s", key)

		if key == key_a || key_1 == key_a || key_2 == key_a {
			return args[i+1]
		}
	}

	return value
}
