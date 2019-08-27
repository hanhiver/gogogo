// dup2 打印输入中多次出现的行的个数和文本。
// 从stdin或者制定的文件列表读取。

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	fmt.Printf("Main: i am in.\n")
    counts := make(map[string]int)
    files := os.Args[1:]

	fmt.Printf("%d files.\n", len(files))

    if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
        for _, arg := range files {
			fmt.Printf("Handle the first file %s.\n", arg)
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
			fmt.Printf("File %s openned.\n", arg)
            countLines(f, counts)
			fmt.Printf("File %s will be closed.\n", arg)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        fmt.Println(input.Text())
        counts[input.Text()]++
    }
}


