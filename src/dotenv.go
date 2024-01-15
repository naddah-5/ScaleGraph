package main

import (
	"bufio"
	"fmt"
	"os"
)

func Init() {
	env, err := os.Open(string(".env"))
	if err != nil {
		fmt.Println("could not open file ", env)
	}
	defer env.Close()

	buf := bufio.NewReader(env)
	bufio.NewScanner(buf)
}
