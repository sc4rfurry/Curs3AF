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
	var version string = "1.0.1"
	var go_version string = "1.18.1 or higher"
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
	println(color.Bold + color.Green + "\n\t\t\t~ Help Menu ~" + color.Reset)
	println("\n\t" + color.Bold + color.Cyan + "Usage: " + color.Reset + "./main -u/--url <domain>")
	println(color.Bold + color.Gray + "____________________________________________________________" + color.Reset)
	println("\n\t" + color.Bold + color.Green + "Options: " + color.Reset)
	println("\t\t" + color.Bold + color.Yellow + "-u/--url " + color.Reset + "\tDomain to scan")
	println("\t\t" + color.Bold + color.Yellow + "-f/--file " + color.Reset + "\tFile with domains to scan (Each domain in a new line)")
	println("\t\t" + color.Bold + color.Yellow + "-g/--generic " + color.Reset + "\tScan for mostly used WAFs")
	println("\t\t" + color.Bold + color.Yellow + "-h/--help " + color.Reset + "\tShow this help menu" + "\n")
	os.Exit(0)
}

func downloadFile(filename string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func scanFile(filename string, isGeneric bool) {
	// Open FingerPrints
	fingerPrintsFile, err := os.Open("waf-fingerprints.json")
	if err != nil {
		log.Fatalf("error opening fingerprints file: %v", err)
	}
	defer fingerPrintsFile.Close()

	// Open Targets File
	targetsFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening targets file: %v", err)
	}
	defer targetsFile.Close()

	// Run WAF Detection
	if isGeneric {
		opts := &wafme0w.Options{
			Inputs:           targetsFile,
			FingerPrints:     fingerPrintsFile,
			Concurrency:      50,
			FastMode:         true,
			ExcludeGeneric:   false,
			ListWAFS:         false,
			Silent:           true,
			NoColors:         false,
			SuppressWarnings: false,
		}
		println(info + "Starting WAF Detection on " + color.Bold + color.Green + filename + color.Reset)
		println(info + "Running in Generic Mode (Scan for mostly used Wafs)" + "\n")
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
	} else {

		opts := &wafme0w.Options{
			Inputs:           targetsFile,
			FingerPrints:     fingerPrintsFile,
			Concurrency:      50,
			FastMode:         false,
			ExcludeGeneric:   false,
			ListWAFS:         false,
			Silent:           true,
			NoColors:         false,
			SuppressWarnings: false,
		}
		println(info + "Starting WAF Detection on " + color.Bold + color.Green + filename + color.Reset)
		println(info + "Running in Normal Mode (Scan for all 153 Wafs)" + "-" + " Could take time to scan" + "\n")
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
	os.Exit(0)
}

func main() {
	var targets []byte
	var isGeneric bool
	var filename string
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-g" || arg == "--generic" {
			isGeneric = true
		} else {
			isGeneric = false
		}
	}
	if len(args) == 0 {
		help()
	}
	if args[0] == "-h" || args[0] == "--help" {
		help()
	}
	banner()
	if args[0] == "-f" || args[0] == "--file" {
		if args[1] != "" {
			filename = args[1]
		} else {
			log.Fatalf("Invalid or Null Filename: %v", filename)
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %v", filename)
		}
		scanFile(filename, isGeneric)
	}
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
	if isGeneric {
		opts := &wafme0w.Options{
			Inputs:           targetsReader,
			FingerPrints:     fingerPrintsFile,
			Concurrency:      50,
			FastMode:         true,
			ExcludeGeneric:   false,
			ListWAFS:         false,
			Silent:           true,
			NoColors:         false,
			SuppressWarnings: false,
		}
		println(info + "Starting WAF Detection on " + color.Bold + color.Green + args[1] + color.Reset)
		println(info + "Running in Generic Mode (Scan for mostly used Wafs)" + "\n")
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
	} else {

		opts := &wafme0w.Options{
			Inputs:           targetsReader,
			FingerPrints:     fingerPrintsFile,
			Concurrency:      50,
			FastMode:         false,
			ExcludeGeneric:   false,
			ListWAFS:         false,
			Silent:           true,
			NoColors:         false,
			SuppressWarnings: false,
		}
		println(info + "Starting WAF Detection on " + color.Bold + color.Green + args[1] + color.Reset)
		println(info + "Running in Normal Mode (Scan for all 153 Wafs)" + "-" + " Could take time to scan" + "\n")
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
}
