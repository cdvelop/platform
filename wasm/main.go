package main

import "syscall/js"

func main() {
	js.Global().Get("console").Call("log", "¡Hi 3 Go y WebAssembly!")
	// fmt.Scanln() // para que el programa no se cierre inmediatamente
	// medical.Hello()
	select {}

}
