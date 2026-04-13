package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type Pattern struct {
	Flags    string   `json:"flags"`
	Regexp   string   `json:"regexp"`
	Pattern  string   `json:"pattern"`
	Patterns []string `json:"patterns"`
}

var (
	listFlag bool
	allFlag  bool
)

func init() {
	flag.BoolVar(&listFlag, "list", false, "Show all available patterns in your .gf folder")
	flag.BoolVar(&allFlag, "all", false, "Run every pattern you have against the input")

	flag.Usage = func() {
		fmt.Printf("\033[35m")
		fmt.Println("VANDEX-GF - The Ultimate Recon Filtering Tool")
		fmt.Println("Author: Vandex (Mohamed Magdy)")
		fmt.Printf("\033[0m")
		fmt.Println("\nUsage:")
		fmt.Println("  cat targets.txt | vandex-gf [pattern_names] [flags]")
		fmt.Println("\nMain Flags:")
		fmt.Println("  -list    : List available patterns")
		fmt.Println("  -all     : Run all patterns at once (Multi-threaded)")
		fmt.Println("  -h       : Show this help menu")
		fmt.Println("\nExamples:")
		fmt.Println("  cat urls.txt | vandex-gf xss sqli")
		fmt.Println("  cat urls.txt | vandex-gf -all")
		fmt.Println("  cat urls.txt | vandex-gf lfi")
	}
}

func main() {
	home, _ := os.UserHomeDir()
	gfPath := filepath.Join(home, ".gf")

	flag.Parse()
	targets := flag.Args()

	if listFlag {
		files, err := ioutil.ReadDir(gfPath)
		if err != nil {
			fmt.Printf("\033[31m[-] Error: Could not read .gf folder at %s\033[0m\n", gfPath)
			return
		}
		fmt.Println("[*] Available patterns:")
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				fmt.Println("  -", strings.TrimSuffix(f.Name(), ".json"))
			}
		}
		return
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && arg != "-list" && arg != "-all" && arg != "-h" {
			clean := strings.TrimLeft(arg, "-")
			fmt.Printf("\033[31m[!] Error: You used '%s'. Correct usage is '%s' (without the dash).\033[0m\n", arg, clean)
			return
		}
	}

	if allFlag {
		files, _ := ioutil.ReadDir(gfPath)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				targets = append(targets, strings.TrimSuffix(f.Name(), ".json"))
			}
		}
	}

	if len(targets) == 0 {
		flag.Usage()
		return
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			processPattern(t, gfPath, lines)
		}(target)
	}
	wg.Wait()
}

func processPattern(target string, gfPath string, lines []string) {
	patternFile := filepath.Join(gfPath, target+".json")

	data, err := ioutil.ReadFile(patternFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31m[-] Pattern not found: %s\033[0m\n", target)
		return
	}

	var p Pattern
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31m[-] Malformed JSON: %s\033[0m\n", target)
		return
	}

	finalRegex := p.Regexp
	if finalRegex == "" {
		finalRegex = p.Pattern
	}
	if finalRegex == "" && len(p.Patterns) > 0 {
		finalRegex = "(" + strings.Join(p.Patterns, "|") + ")"
	}

	if finalRegex == "" {
		fmt.Fprintf(os.Stderr, "\033[31m[-] No regex found in: %s\033[0m\n", target)
		return
	}

	if strings.Contains(p.Flags, "i") && !strings.HasPrefix(finalRegex, "(?i)") {
		finalRegex = "(?i)" + finalRegex
	}

	re, err := regexp.Compile(finalRegex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31m[-] Invalid Regex in: %s\033[0m\n", target)
		return
	}

	outputFileName := target + ".txt"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	found := false

	for _, line := range lines {
		if re.MatchString(line) {
			writer.WriteString(line + "\n")
			found = true
		}
	}
	writer.Flush()

	if found {
		fmt.Printf("\033[32m[+] Finished filtering for %-10s -> Saved to %s\033[0m\n", target, outputFileName)
	} else {
		os.Remove(outputFileName)
		fmt.Printf("\033[33m[!] Finished filtering for %-10s (No matches)\033[0m\n", target)
	}
}
