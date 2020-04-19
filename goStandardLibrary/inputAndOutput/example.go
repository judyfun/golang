package inputAndOutput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**

 */
func ReadAtDemo() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s，%d\n", p, n)
}

func WriteAtDemo() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	file.WriteString("1234567890")
	n, err := file.WriteAt([]byte("中文网a1"), 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}

func ReadFromDemo() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}
