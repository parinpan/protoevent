package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8089")

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	conn.Write([]byte(`WADOOOOOOOWWWWWW`))
}
