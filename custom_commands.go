package main

import (
	bf "github.com/ali-mukhamed-palitaev/brainfuck-interpreter-go/lib"
)

// Remainder is a Command implementation. It sets the value of the remainder after division by 2
func Remainder(i *bf.Interpreter) {
	i.SetCurrentValue(i.GetCurrentValue() % 2)
}
