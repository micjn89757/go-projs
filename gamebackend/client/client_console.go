package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 从控制台接收指令和参数
type ClientConsole struct {
	chInput chan *InputParam
}

type InputParam struct {
	Command string		// 输入指令
	Param   []string	// 输入参数
}

func NewClientConsole() *ClientConsole {
	c := &ClientConsole{}
	return c
}

func (c *ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input err ,check your input and  try again !!!")
			continue
		}
		readString = strings.TrimRight(readString, "\n")
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			fmt.Println("input err, check your input and  try again !!! ")
			continue
		}
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		c.chInput <- in
	}
}
