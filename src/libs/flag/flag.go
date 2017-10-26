package flag

import (
	"fmt"
	"os"
	"strings"
)

/*
 get the value that CMDLINE provided, if not, return the `value`(default)
 cmdline: ./cmd args0 args1 ...
*/
func Parse(key, value string) string {
	if key == "" {
		return value
	}

	var args = os.Args[1:]

	for i := 0; i < len(args); i += 1 {
		var kv = strings.Split(args[i], "=")
		if len(kv) != 2 && kv[0] != "" && kv[1] != "" {
			continue
		}

		var key_a = fmt.Sprintf("-%s", key)
		var key_b = fmt.Sprintf("--%s", key)

		if kv[0] == key || kv[0] == key_a || kv[0] == key_b {
			return kv[1]
		}

	}

	return value
}
