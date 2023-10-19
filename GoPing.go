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
}
