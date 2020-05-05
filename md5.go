package main

import (
    "crypto/md5"
    "fmt"
    "bufio"
    "os"

)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan(){
	domain := scan.Text()
    data := []byte(domain)
    fmt.Printf(domain+":")
    fmt.Printf("%x\n", md5.Sum(data))
}}
