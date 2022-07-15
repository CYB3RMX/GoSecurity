package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

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

func getHashes(flname string) {
	var hashes []string
	coll := colly.NewCollector()
	coll.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "MD5") && strings.Contains(e.Attr("href"), "direction") == false {
			hashes = append(hashes, strings.Split(e.Attr("href"), "=")[1])
		}
	})
	coll.Visit("http://vxvault.net/ViriList.php")
	createFile(flname, hashes)
}

func getIPAddrs(flname string) {
	var ipaddrs []string
	coll := colly.NewCollector()
	coll.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "IP=") {
			ipaddrs = append(ipaddrs, strings.Split(e.Attr("href"), "=")[1])
		}
	})
	coll.Visit("http://vxvault.net/ViriList.php")
	createFile(flname, ipaddrs)
}

func main() {
	var (
		dtype      string
		outputfile string
	)
	var banner = `
	 _   _      _   _           _     
	| | | |    | | | |         | |    
	| | | |_  _| |_| | __ _ ___| |__  
	| | | \ \/ /  _  |/ _  / __|  _ \
	\ \_/ />  <| | | | (_| \__ \ | | |
	 \___//_/\_\_| |_/\__,_|___/_| |_|

			> Simple IoC crawler
		
		@CYB3RMX https://github.com/CYB3RMX
	`
	fmt.Println(banner)
	flag.StringVar(&dtype, "dtype", "hash", "Specify data type (hash, ipaddr).")
	flag.StringVar(&outputfile, "output", "output.txt", "Specify output file.")
	flag.Parse()
	switch dtype {
	case "hash":
		getHashes(outputfile)
	case "ipaddr":
		getIPAddrs(outputfile)
	default:
		fmt.Println("Wrong choice :(")
	}
}
