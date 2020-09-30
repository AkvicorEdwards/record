package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadUInt32() (val uint32) {
	_, _ = fmt.Scanf("%d\n", &val)
	return val
}

func ReadInt() (val int) {
	_, _ = fmt.Scanf("%d\n", &val)
	return val
}

func ReadLine() (str string) {
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	var err error
	for {
		str, err = inputReader.ReadString('\n')
		if err != nil || len(str) == 0{
			continue
		}
		if str == "\n" || str == " " {
			continue
		}
		break
	}

	for str[len(str)-1] == '\n' {
		str = str[:len(str)-1]
	}

	return strings.TrimSpace(str)
}
