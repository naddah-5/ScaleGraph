package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Init() {
	
}

func ReadEnvironment() ([]string, error) {
	env, err := os.Open(string(".env"))
	if err != nil {
		fmt.Println("could not open file ", env)
	}
	defer env.Close()

	buf := bufio.NewReader(env)
	scanner := bufio.NewScanner(buf)
	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		linePair := strings.Split(line, "=")
		if len(linePair) > 2 {
			return nil, errors.New("incorrect .env configuration, can only assign one value per line")
		}
		content = append(content, )
	}
	return content, nil
}
