package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func downloadFile(url string, output string) {
	// Doin http GET request
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.ContentLength > 264 {
		// Handling output file
		fl, _ := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
		defer fl.Close()

		// Implementing progress bar
		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			"Downloading...",
		)
		io.Copy(io.MultiWriter(fl, bar), resp.Body)
	} else {
		fmt.Println("[!] Looks like we have 404.")
	}
}

func main() {
	// Banner
	var banner = `
	 ____                            _____                      _                 _           
	|  _ \                          |  __ \                    | |               | |          
	| |_) | __ _ ______ _  __ _ _ __| |  | | _____      ___ __ | | ___   __ _  __| | ___ _ __ 
	|  _ < / _  |_  / _  |/ _  |  __| |  | |/ _ \ \ /\ / /  _ \| |/ _ \ / _  |/ _  |/ _ \  __|
	| |_) | (_| |/ / (_| | (_| | |  | |__| | (_) \ V  V /| | | | | (_) | (_| | (_| |  __/ |   
	|____/ \__,_/___\__,_|\__,_|_|  |_____/ \___/ \_/\_/ |_| |_|_|\___/ \__,_|\__,_|\___|_|

		> Simple malware sample fetch tool.

		@CYB3RMX https://github.com/CYB3RMX
	`
	fmt.Println(banner)

	// Specify program arguments
	var date string
	flag.StringVar(&date, "date", "foo", "Specify a date.")
	flag.Parse()

	// Parsing urls
	newurl := "https://datalake.abuse.ch/malware-bazaar/daily/" + date + ".zip"
	newout := "malware-bazaar-" + date + ".zip"

	// Download content
	downloadFile(newurl, newout)

	// Create a directory to store data
	direct, _ := exists("./Malwarebazaar_data")
	if direct {
		news := "./Malwarebazaar_data/" + newout
		os.Rename(newout, news)
	} else {
		os.Mkdir("Malwarebazaar_data", os.ModePerm)
		news := "./Malwarebazaar_data/" + newout
		os.Rename(newout, news)
	}
}
