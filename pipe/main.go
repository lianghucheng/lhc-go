package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	pipe3()
}

func pipe1() {
	cmd0 := exec.Command("echo","My first command echo for Golang.")

	stdout, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: the command can not be startup: %s\n", err)
		return
	}

	output0 := make([]byte, 30)
	n, err := stdout.Read(output0)
	if err != nil {
		fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
		return
	}

	fmt.Printf("the imformage for the pipe1: %s\n", output0[:n])
}

func pipe2() {
	cmd1 := exec.Command("nslookup", "baidu.com")

	stdout1,err := cmd1.StdoutPipe()
	if err !=nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for the command No.1: %s\n",err)
		return
	}

	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}

	var outputBuf bytes.Buffer
	for {
		tempoutput := make([]byte, 5)
		n, err := stdout1.Read(tempoutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
			}
		}
		if n > 0 {
			outputBuf.Write(tempoutput)
		}
	}

	fmt.Printf("%s\n", outputBuf.String())
}

func pipe3() {
	cmd2 := exec.Command("nslookup", "baidu.com")

	stdout2, err := cmd2.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for the command No.2: %s\n", err)
		return
	}

	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}

	outputBuf2 := bufio.NewReader(stdout2)

	var output bytes.Buffer
	for{
		b, _, err := outputBuf2.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			} else {
				fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
				return
			}
		}
		if len(b) > 0 {
			output.Write(b)
		}
	}

	fmt.Printf("The pipe3 result: %s\n", output.String())
}