package usuario

import (
	"fmt"
	"unicode"
)

// variavel acessível em todas as funções desse arquivo
var FraseDoUsuario string

// Função que inicia em maiusculo é global e pode ser chamada fora do pacote
func MontarFraseDoUsuario(nomeDoUsuario string, numeroDaCor, idade int) {
	// variavel acessível apenas nessa função/escopo
	nomedoUsuarioCorrigido := colocarPrimeiraLetraEmMaiusculo(nomeDoUsuario)

	corPreferida := encontrarCorPreferida(numeroDaCor)

	fraseFormatada := fmt.Sprintf(
		`Olá! Meu nome é %s, minha idade é %d e minha cor preferida é %s!!`,
		nomedoUsuarioCorrigido,
		idade,
		corPreferida)

	FraseDoUsuario = fraseFormatada
}

// Função que inicia em minusculo não pode ser chamada fora do pacote
func encontrarCorPreferida(numeroDaCor int) string {
	if numeroDaCor == 1 {
		return "laranja"
	} else if numeroDaCor == 2 {
		return "roxo"
	} else if numeroDaCor == 3 {
		return "rosa"
	} else if numeroDaCor == 4 {
		return "branco"
	} else if numeroDaCor == 5 {
		return "preto"
	} else {
		return "nenhuma"
	}
}

func colocarPrimeiraLetraEmMaiusculo(nome string) string {
	// A função len em Go é uma função embutida que retorna o número de elementos de um determinado argumento.
	// Ela é amplamente utilizada com vários tipos de dados.
	// Strings -> retorna o número de bytes na string
	// arrays e slices -> retorna o número de elementos presentes
	if len(nome) == 0 {
		return ""
	}

	// Convertendo a string para uma slice de runas.
	// Isso é necessário para lidar corretamente com caracteres multibyte.
	// arrays são usados para guardar múltiplos valores do mesmo time em apenas uma variável
	// slice é igual um array, mas sem tamanho fixado
	runas := []rune(nome)

	// Converte apenas a primeira runa para maiúsculo.
	runas[0] = unicode.ToUpper(runas[0])

	// A partir da segunda runa, converte todas para minúsculo.
	for i := 1; i < len(runas); i++ {
		runas[i] = unicode.ToLower(runas[i])
	}

	// Converte a slice de runas de volta para string e retorna.
	return string(runas)
}
