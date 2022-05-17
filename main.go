package main

import (
	"bufio"
	"errors"
	"fmt"
	bf "github.com/ali-mukhamed-palitaev/brainfuck-interpreter-go/lib"
	"os"
)

// Exists checks whether the file with the given name exists
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please, provide brainfuck script filename")
		return
	}
	var filename = os.Args[1]
	if exists, _ := Exists(filename); !exists {
		fmt.Println("File with the given name does not exist")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	interpreter := bf.NewInterpreter(nil, nil)
	interpreter.AddCommand("%", Remainder)
	interpreter.Execute(reader)
}
