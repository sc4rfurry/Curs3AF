package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/Lu1sDV/wafme0w/pkg/wafme0w"
	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var info = color.Bold + color.Yellow + "[info] " + color.Reset
var error_msg = color.Bold + color.Red + "[error] " + color.Reset

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func banner() {
	var author string = "sc4rfurry"
	var version string = "1.0.0"
	var go_version string = "1.19"
	var github string = "https://github.com/sc4rfurry"
	var description string = "Tool to check if a website is protected by a WAF(HTTP/HTTPS)."
	banner := figure.NewColorFigure("Curs3AF", "", "purple", true)
	banner.Print()
	fmt.Printf("\n %vDescription: %v%v", color.Bold+color.Green, color.Reset, description)
	println("\n")
	println("\n"+color.Ize(color.Bold+color.Yellow, "\tAuthor: \t"), author)
	println(color.Ize(color.Bold+color.Yellow, "\tVersion: \t"), version)
	println(color.Ize(color.Bold+color.Yellow, "\tGo Version: \t"), go_version)
	println(color.Ize(color.Bold+color.Yellow, "\tGithub: \t"), github)
	println(color.Ize(color.Bold+color.Blue, "===================================================================================================\n"))
}

func help() {
	var helper string = `
		-- Help for Curs3AF --
	
	usage: ./main -u/--url <domain>
---------------------------------------------------------
	Description:
		- Tool to check if a website is protected by a WAF.
			
			* This tool uses wafme0w, which is a fork of wafw00f.
	
	Installation:
		- sudo apt-get update && sudo apt-get golang
		- git clone https://github.com/sc4rfurry/Curs3AF.git
		- cd Curs3AF
		- go get .
		- go build

	Binary:
		- You can download the PreCompiled binaries from the releases page.
`
	println(helper)
	os.Exit(0)
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	var targets []byte
	args := os.Args[1:]
	if len(args) == 0 {
		banner()
		help()
	}
	if args[0] == "-h" || args[0] == "--help" {
		banner()
		help()
	}
	banner()
	if args[0] == "-u" || args[0] == "--url" {
		domain := args[1]
		if domain == "" {
			log.Fatalf("Invalid Domain/IP: %v", domain)
		} else {
			targets = []byte("http://" + domain + "\n" + "https://" + domain + "\n")

		}

	} else {
		log.Fatalf("Invalid Argument: %v", args[0])
	}
	targetsReader := bytes.NewReader(targets)
	// Download FingerPrints
	if _, err := os.Stat("waf-fingerprints.json"); os.IsNotExist(err) {
		println(info + "FingerPrints not found")
		fmt.Printf("%v[+]%v Downloading FingerPrints... \n", Green, Reset)
		err := downloadFile("waf-fingerprints.json", "https://raw.githubusercontent.com/Lu1sDV/wafme0w/main/cmd/wafme0w/resources/waf-fingerprints.json")
		if err != nil {
			log.Fatalf("%v error downloading fingerprints: %v", error_msg, err)
		}
	}
	// Open FingerPrints
	fingerPrintsFile, err := os.Open("waf-fingerprints.json")
	if err != nil {
		log.Fatalf("error opening fingerprints file: %v", err)
	}
	defer fingerPrintsFile.Close()

	// Run WAF Detection
	opts := &wafme0w.Options{
		Inputs:       targetsReader,
		FingerPrints: fingerPrintsFile,
		Silent:       true,
		Concurrency:  50,
		FastMode:     true,
	}
	println(info + "Starting WAF Detection on " + color.Bold + color.Green + args[1] + color.Reset + "\n")
	runner := wafme0w.NewRunner(opts)
	result, err := runner.Scan()
	if err != nil {
		log.Fatalf("error running scan: %v", err)
	}

	// Print Results
	for _, target := range result {
		fmt.Printf("%v[!]%v %v%v%v is protected by %v%v%v\n", Purple, Reset, Yellow, target.Target, Reset, Cyan, target.FingerPrint, Reset)
	}
	println("\n")
}
