package common

import (
	"net"
	"strconv"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port, nil
}

func Atoi32(str string) int32 {
	r, _ := strconv.Atoi(str)
	return int32(r)
}
