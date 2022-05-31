package task3

import "io"

func PrintString(w io.Writer, arr ...interface{}) {
	for _, el := range arr {
		if val, ok := el.(string); ok {
			w.Write([]byte(val))
		}
	}
}