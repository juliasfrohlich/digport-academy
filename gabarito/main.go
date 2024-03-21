// package main: Um pacote executável que gera um arquivo binário quando compilado.
// sinaliza ao compilador de Go que o pacote deve ser compilado como um programa autônomo.
package main

// libs
import (
	"example/Digport-Academy/gabarito/usuario"
	"fmt"
)

// Função main: A função main serve como o ponto de entrada para a execução de um programa Go.
// Quando um programa Go é executado, o processo começa na função main do pacote main.
// Esta função não recebe nenhum argumento e não retorna nenhum valor. Ela inicia a execução do programa e quando termina, o programa é encerrado.
func main() {
	idade := 24
	nome := "Julia"
	fmt.Println(usuario.FraseDoUsuario)
	usuario.MontarFraseDoUsuario(nome, 2, idade)

}
