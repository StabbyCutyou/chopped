package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
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
	"io",
	"context",
	"fmt",
}

var xmap = map[string]string{
	"crypto":  "golang.org/x/crypto",
	"image":   "golang.org/x/image",
	"net":     "golang.org/x/net",
	"syscall": "golang.org/x/sys",
	"text":    "golang.org/x/text",
	"tools":   "golang.org/x/tools",
}

var xorphans = []string{
	"golang.org/x/blog",
	"golang.org/x/exp",
	"golang.org/x/mobile",
	"golang.org/x/talks",
}

var rejects = append(pantry, "vendor", "internal")

func main() {
	c := flag.Int("c", 1, "Determines if this is in bulk mode or not. Defaults to 1, which disables bulk mode")
	e := flag.Bool("e", false, "Easy mode: Only uses the upper-most packages when determining baskets")
	flag.Parse()
	if *c == 1 {
		runShowMode(*e)
		os.Exit(0)
	}
	runBulkMode(*c, *e)
}

func runShowMode(easy bool) {
	fmt.Println("Tonight! On Chopped: Boston Go edition...")
	time.Sleep(2 * time.Second)
	fmt.Println("Our contests will be charged with making a completely new library or application...")
	fmt.Println("Using only THESE 5 packages!")
	time.Sleep(1 * time.Second)
	p := loadPackages(easy)
	size := big.NewInt(int64(len(p)))
	list := make([]string, 0)
	for i := 0; i < 5; i++ {
		alex, _ := rand.Int(rand.Reader, size)
		pick := p[alex.Int64()]
		if contains(list, pick) {
			i--
			continue
		}
		list = append(list, pick)
		fmt.Println(pick)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("And of course, each contestant has access to our staple pantry")
	for _, item := range pantry {
		fmt.Println(item)
		time.Sleep(1 * time.Second)
	}
}

func contains(l []string, p string) bool {
	for _, s := range l {
		if s == p {
			return true
		}
	}
	return false
}

func pickSeasoning(list []string) string {
	seasons := make([]string, 0)
	// collect all possible seasonins for known packages
	for _, l := range list {
		if s, ok := xmap[l]; ok {
			seasons = append(seasons, s)
		}
	}
	// if none found, use an orphan
	if len(seasons) == 0 {
		// return rand orphan
		size := big.NewInt(int64(len(xorphans)))
		// return rand season
		mark, _ := rand.Int(rand.Reader, size)
		return xorphans[mark.Int64()]
	}
	size := big.NewInt(int64(len(seasons)))
	// return rand season
	aron, _ := rand.Int(rand.Reader, size)
	return seasons[aron.Int64()]
}

func runBulkMode(c int, easy bool) {
	for k := 0; k < c; k++ {
		p := loadPackages(easy)
		size := big.NewInt(int64(len(p)))
		fmt.Printf("------------------Basket %d---------------\n", k)
		list := make([]string, 0)
		for i := 0; i < 5; i++ {
			alex, _ := rand.Int(rand.Reader, size)
			pick := p[alex.Int64()]
			if contains(list, pick) {
				i--
				continue
			}
			list = append(list, pick)
			fmt.Println(pick)
		}
		fmt.Println(">>> Seasoning X <<<")
		fmt.Println(pickSeasoning(list))
		fmt.Println(">>> Staple Pantry <<<")
		for _, item := range pantry {
			fmt.Println(item)
		}
		fmt.Println("------------------------------------------")
	}

}

func isRejected(p string) bool {
	if p == "" {
		return true
	}
	for _, r := range rejects {
		if strings.Contains(p, r) {
			return true
		}
	}
	return false
}

func loadPackages(easy bool) []string {
	cmd := exec.Command("go", "list", "std")
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(b), "\n")
	finalList := make([]string, 0)
	for _, p := range s {
		if isRejected(p) {
			continue
		}
		if easy {
			p = strings.Split(p, "/")[0]
		}
		if !contains(finalList, p) {
			finalList = append(finalList, p)
		}
	}
	return finalList
}
