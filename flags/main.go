package main

import (
	"flag"
	"fmt"
	"os"
)

//  file system create:
//    token
//    region
//    name

//  file system grant:
//    token
//    region
//    file system id
//    addr

//  file system revoke:
//    token
//    region
//    file system id
//    addr

//  file system list:
//    token
//    region
//    file system id

// file system show:
//    token
//    region
//    file system id

//  file system delete:
//    token
//    region
//    file system id

//  file system update:
//    token
//    region
//    file system id
//    name

var token = flag.String("token", "", "Security Token")
var region = flag.String("region", "iad", "CFS Region")
var name = flag.String("name", "cfsdrive", "File System Name")
var fsid = flag.String("fsid", "", "File System ID")
var addr = flag.String("addr", "127.0.0.1", "Client Network Address")

var regions = []string{"iad", "aio"}

// FileSystem ...
type FileSystem struct {
	ID     string   `json:"id"`
	AcctID string   `json:"acctid"`
	Name   string   `json:"name"`
	Addrs  []string `json:"addrs"`
}

func main() {
	flag.Parse()
	var fs FileSystem
	fmt.Println(os.Args)

	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}

	if *token == "" {
		fmt.Printf("%s\n", "Token is required")
		printUsage()
		os.Exit(1)
	}

	if !validRegion(*region) {
		fmt.Printf("%s %s\n", "Invalid Region. The valid regions are", regions)
		printUsage()
		// os.Exit(1)
	}

	fs.ID = *fsid
	fs.AcctID = "1-1-1-1"
	fs.Name = *name
	fs.Addrs = []string{*addr}

	fmt.Printf("Token:\t\t\t%s\n", *token)
	fmt.Printf("Region:\t\t\t%s\n", *region)
	fmt.Printf("Name:\t\t\t%s\n", *name)
	fmt.Printf("FSID:\t\t\t%s\n", *fsid)
	fmt.Printf("Addr:\t\t\t%s\n", *addr)

	fmt.Printf("Args[1]:\t\t\t%s\n", flag.Args())

	fmt.Printf("%v\n\n", fs)
}

func printUsage() {
	fmt.Printf("Application: flagtest\n")
}

func validRegion(r string) bool {
	for _, v := range regions {
		if v == r {
			return true
		}
	}
	return false
}
