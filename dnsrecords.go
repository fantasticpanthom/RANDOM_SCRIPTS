package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"runtime"
)


var Reset   = "\033[0m"
var Red     = "\033[31m"
var Green   = "\033[32m"
var Yellow  = "\033[33m"
var Blue    = "\033[34m"
var Purple  = "\033[35m"
var Cyan    = "\033[36m"
var Gray    = "\033[37m"
var White   = "\033[97m"


func init() {
	if runtime.GOOS == "windows" {
		Reset   = ""
		Red     = ""
		Green   = ""
		Yellow  = ""
		Blue    = ""
		Purple  = ""
		Cyan    = ""
		Gray    = ""
		White   = ""
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		domain := scan.Text()
		ip, _ := net.LookupIP(domain)
		cname, _ := net.LookupCNAME(domain)
		ns, _ := net.LookupNS(domain)
		mx, _ := net.LookupMX(domain)
		txtrecords, _ := net.LookupTXT(domain)
		fmt.Println(Red + "Domain : " +  domain + Reset)
		fmt.Println(Cyan+"IPv4 and IPv6 addresses"+Reset)
		for _, i := range ip {
			fmt.Println(i)
		}
		fmt.Println(Purple+"CNAMEs"+Reset)
		for _,j := range cname {
			fmt.Println(j)
		}
		fmt.Println(Blue+"Name Server"+Reset)
		for _,k := range ns {
			fmt.Println(k)
		}
		fmt.Println(Green+"MX"+Reset)
		for _,l := range mx {
			fmt.Println(l)
		}
		fmt.Println(Yellow+"text record(SPF info)"+Reset)
		for _,m := range txtrecords {
			fmt.Println(m)
		}
	}
}
