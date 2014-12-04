package util

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

// Must be of type int, otherwise panic
func MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Join(ins []int, sep string) string {
	strSlice := make([]string, len(ins))
	for i, one := range ins {
		strSlice[i] = strconv.Itoa(one)
	}
	return strings.Join(strSlice, sep)
}

func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
