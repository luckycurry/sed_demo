package main

import (
	"bufio"
	"example.com/gp-demo-1/gram_demo/demo1/alo/bubblesort"
	"example.com/gp-demo-1/gram_demo/demo1/alo/qsort"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var infile = flag.String("i", "infile", "File contains values for sorting")
var outfile = flag.String("o", "outfile", "File to receive sorted values")
var algorithm = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("failed to open the file")
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("这行太长啦，不是所期待的")
			break
		}

		str := string(line) // 转换字符数组为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("文件创建失败")
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			bubblesort.BubbleSort(values)
		case "qsort":
			qsort.QuickSort(values)
		default:
			fmt.Println("请输入qsot或者bubblesort")
		}
		t2 := time.Now()
		fmt.Println("本次使用了算法为", *algorithm, "花费时间", t2.Sub(t1))
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
	//writeValues(values, *outfile)

}
