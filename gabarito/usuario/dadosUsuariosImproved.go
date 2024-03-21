package usuario

import (
	"fmt"
	"strings"
)

type Usuario struct {
	nomeDoUsuario string
	numeroDaCor   int
	idade         int
}

func MontarFraseDoUsuarioMelhorada(usuario Usuario) string {

	nomedoUsuarioCorrigido := colocarPrimeiraLetraEmMaiusculoMelhorada(usuario.nomeDoUsuario)

	corPreferida := encontrarCorPreferidaMelhorada(usuario.numeroDaCor)

	fraseFormatada := fmt.Sprintf(`Olá! Meu nome é %s, minha idade é %d e minha cor preferida é %s!!`, nomedoUsuarioCorrigido, usuario.idade, corPreferida)

	return fraseFormatada
}

func encontrarCorPreferidaMelhorada(numeroDaCor int) string {
	switch numeroDaCor {
	case 1:
		return "laranja"
	case 2:
		return "roxo"
	case 3:
		return "rosa"
	case 4:
		return "branco"
	case 5:
		return "preto"

	default:
		return "nenhuma"
	}
}

func colocarPrimeiraLetraEmMaiusculoMelhorada(s string) string {
	// A função len em Go é uma função embutida que retorna o número de elementos de um determinado argumento.
	// Ela é amplamente utilizada com vários tipos de dados.
	// Strings -> retorna o número de bytes na string
	// arrays e slices -> retorna o número de elementos presentes
	if len(s) == 0 {
		return ""
	}
	// Converte o primeiro caractere para maiúsculo e o restante para minúsculo.
	// fatiamento: s[:1] pega os bytes da string s do início até, mas não incluindo, o índice 1. resulta no primeiro caractere da string.
	// fatiamento s[1:] pega todos os bytes da string começando no índice 1 até o final da string. Isso exclui o primeiro caractere e retorna o restante da string.
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
