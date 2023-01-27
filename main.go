package main

import (
	"crypto/x509"
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sys/cpu"
)

func main() {
	fmt.Println("go version:", runtime.Version())
	fmt.Println("cpu.S390X.HasVX:", cpu.S390X.HasVX)

	roots := x509.NewCertPool()
	data, err := os.ReadFile("/ca-bundle.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if ok := roots.AppendCertsFromPEM(data); ok {
		fmt.Println("all good")
	}
}
