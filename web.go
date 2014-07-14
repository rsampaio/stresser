package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for t := range ticker.C {
			go func() {
				fmt.Println(t)
				addrs, _ := net.InterfaceAddrs()
				ip, _, _ := net.ParseCIDR(addrs[1].String())
				cmd := exec.Command(
					"./.bin/ab", "-c", "1000", "-n", "100000", "http://"+ip.String()+":"+os.Getenv("PORT")+"/")
				cmd.Env = []string{"LD_LIBRARY_PATH=./.bin/"}
				out, err := cmd.Output()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(string(out))
			}()
		}
	}()

	http.HandleFunc("/", hello)
	fmt.Println("listening..." + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}
