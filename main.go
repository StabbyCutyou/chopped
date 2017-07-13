package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"time"
)

var pantry = []string{
	"strings",
	"math",
	"errors",
	"time",
	"testing",
	"flag",
}

func main() {
	fmt.Println("Tonight! On Chopped: Boston Go edition...")
	time.Sleep(2 * time.Second)
	fmt.Println("Our contests will be charged with making a completely new library or application...")
	fmt.Println("Using only THESE 5 packages!")
	time.Sleep(1 * time.Second)
	p := loadPackages()
	size := big.NewInt(int64(len(p)))
	for i := 0; i < 5; i++ {
		alex, _ := rand.Int(rand.Reader, size)
		fmt.Println(p[alex.Int64()])
		time.Sleep(1 * time.Second)
	}
	fmt.Println("And of course, each contestant has access to our staple pantry")
	for _, item := range pantry {
		fmt.Println(item)
		time.Sleep(1 * time.Second)
	}
}

func loadPackages() []string {
	cmd := exec.Command("go", "list", "std")
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(b), "\n")
	finalList := make([]string, 0)
	for _, p := range s {
		if strings.Contains(p, "vendor") ||
			strings.Contains(p, "testing") ||
			strings.Contains(p, "math") ||
			strings.Contains(p, "errors") ||
			strings.Contains(p, "flag") ||
			strings.Contains(p, "time") {
			continue
		}
		finalList = append(finalList, p)
	}
	return finalList
}
