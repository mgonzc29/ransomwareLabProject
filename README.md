# Práctica académica: simulación de ransomware

---

> ## AVISO: EJERCICIO PURAMENTE ACADÉMICO
>
> **Este repositorio es un ejercicio académico / de laboratorio de ciberseguridad. NO es un ransomware real, no busca causar daño y NO debe emplearse con fines maliciosos.**
>
> - Es una **demostración didáctica** para estudiar técnicas de malware en un entorno aislado.
> - **No ejecutar** fuera de una máquina de pruebas propia o expresamente autorizada.
> - **No utilizar** contra equipos de terceros, sistemas de producción ni datos reales.
> - **No redistribuir** como herramienta de ataque.
>
> Cualquier uso indebido del código es responsabilidad exclusiva de quien lo ejecuta.

---

## ¿Qué contiene?

Proyecto en **Go** (módulo `ransomware`, Go 1.27) que implementa un **cliente simulado** con las piezas típicas de un ransomware de laboratorio: conexión a un servidor de comando y control (C2), cifrado de archivos, descifrado de "vacuna", borrado de copias shadow, keylogger opcional y consola remota.

Todo el comportamiento depende de un servidor C2 de prueba; la dirección está fija en el código (`10.130.72.98:9090`). Si no hay conexión, el programa no ejecuta cambios en los archivos.

### Funcionalidades (con fines didácticos)

- **Cifrado AES-256** de archivos `.jpg`, `.jpeg`, `.png`, `.gif`, `.docx`, `.xlsx`, `.pptx` y `.pdf` bajo `C:\`. La clave se deriva con SHA-256 y se usa AES en modo ECB con padding PKCS#7 (forma insegura, elegida solo para ilustrar el concepto).
- **Descifrado (`vacuna`)**: revierte exactamente el cifrado anterior con la misma clave.
- **Borrado de copias shadow** mediante `vssadmin delete shadows /all /quiet`, para mostrar una técnica habitual de los ransomware reales.
- **Keylogger opcional** basado en `GetAsyncKeyState` de `user32.dll` (sin binarios externos).
- **Autoborrado opcional** del ejecutable mediante PowerShell.
- **Consola remota** sobre TCP: acepta `ransom`, `vacuna`, `keys` y, para cualquier otro comando, lo pasa a `cmd.exe /C`.

### Por qué NO es un ransomware real

- No hay cifrado asimétrico: la clave viaja en texto plano desde el C2 y sirve tanto para cifrar como para descifrar.
- Incluye un comando explícito para **descifrar sin pago** (`vacuna`).
- No tiene persistencia, propagación, evasión avanzada, ni mecanismo de rescate/extorsión.
- Está limitado a un entorno de pruebas Windows y a una lista de extensiones concreta.

### Estructura de archivos

| Archivo | Contenido |
| --- | --- |
| `main.go` | Punto de entrada: conexión con el C2, arranque de las partes opcionales y bucle de comandos. |
| `ransomUtils/ataques.go` | Funciones `Atacar` (cifrar) y `Curar` (descifrar) los archivos de la lista. |
| `ransomUtils/cifrado.go` | Derivación de clave, AES-256, padding PKCS#7 y cifrado/descifrado de archivos. |
| `ransomUtils/utils.go` | Listado de copias de sombra, búsqueda de archivos por extensión y autoborrado. |
| `ransomUtils/keylogger.go` | Captura de pulsaciones de teclado (parte opcional). |
| `ransomUtils/reverseCon.go` | Conexión TCP al C2 y procesamiento de los comandos recibidos. |
| `go.mod` | Definición del módulo Go. |

## Requisitos

- Go 1.27.
- Windows: el código usa `user32.dll`, `vssadmin`, `cmd.exe`, `powershell` y rutas `C:\`.

## Uso responsable

Este proyecto solo debe compilarse y ejecutarse en una **máquina aislada de laboratorio**, con fines de aprendizaje e investigación, y bajo tu propia responsabilidad. Para probarlo hace falta un servidor TCP de prueba en la dirección indicada que envíe una clave terminada en salto de línea y acepte los comandos anteriores.

El objetivo del repositorio es **comprender y analizar** estas técnicas para mejorar la defensa, no disponer de una herramienta usable contra sistemas reales.
