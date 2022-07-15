package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Domz struct {
	Domain []string
}

type RespData struct {
	Hashdata []HashData
}

type HashData struct {
	Md5    string `json:"md5"`
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
}

func createFile(flname string, datawrite []string) {
	fl, err := os.Create(flname)
	if err != nil {
		panic(err)
	}
	defer fl.Close()

	for _, data := range datawrite {
		fl.WriteString(data + "\n")
	}
	fmt.Printf("Data saved into: %s\n", flname)
}

func gatherHashes(apikey string) []string {
	var (
		respData RespData
		hasharr  []string
	)
	// Make HTTP Get request to API
	newurl := "https://malshare.com/api.php?api_key=" + apikey + "&action=getlist"
	resp, err := http.Get(newurl)
	if err != nil {
		panic(err)
	}

	// Parsing incoming data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parsing JSON data and append to array
	json.Unmarshal(body, &respData.Hashdata)
	for i := 0; i < len(respData.Hashdata); i++ {
		hasharr = append(hasharr, respData.Hashdata[i].Md5)
	}
	return hasharr
}

func gatherDomains(apikey string) []string {
	var (
		ddz  Domz
		darr []string
	)
	newurl := "https://malshare.com/api.php?api_key=" + apikey + "&action=getsources"
	resp, err := http.Get(newurl)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &ddz.Domain)
	for _, data := range ddz.Domain {
		darr = append(darr, data)
	}
	return darr
}

func main() {
	var (
		apikey     string
		dtype      string
		outputfile string
	)
	var banner = `
	___  ___      _ _____                          
	|  \/  |     | /  ___|                         
	| .  . | __ _| \  --.  ___  _   _ _ __ ___ ___ 
	| |\/| |/ _  | | --. \/ _ \| | | |  __/ __/ _ \
	| |  | | (_| | /\__/ / (_) | |_| | | | (_|  __/
	\_|  |_/\__,_|_\____/ \___/ \__,_|_|  \___\___|

			> Simple MalShare crawler tool.

			@CYB3RMX https://github.com/CYB3RMX
	`
	fmt.Println(banner)
	flag.StringVar(&apikey, "apikey", "foo", "Specify Malshare API key.")
	flag.StringVar(&dtype, "dtype", "hash", "Specify data type (hash, domain).")
	flag.StringVar(&outputfile, "output", "output.txt", "Specify output file.")
	flag.Parse()
	switch dtype {
	case "hash":
		hdat := gatherHashes(apikey)
		createFile(outputfile, hdat)
	case "domain":
		ddat := gatherDomains(apikey)
		createFile(outputfile, ddat)
	default:
		fmt.Println("Wrong argument :(")
	}
}
