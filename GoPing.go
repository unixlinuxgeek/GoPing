// This is a console application for quickly checking the connection with a remote host

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	pth, err := exec.LookPath("ping")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println("ping installed in:", pth)

	if len(os.Args) == 2 {
		out, err := exec.Command("ping", "-c", "3", os.Args[1]).Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		fmt.Println(string(out))
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", "incorrect number of arguments")
	}
}
