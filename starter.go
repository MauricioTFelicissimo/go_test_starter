package starter

import (
	"fmt"
	"math"
	"net/http"
)

// sayHello gera uma mensagem de saudação para um nome dado, esta função foi desenvolvida seguindo a abordagem TDD
// 1. escrever um teste que falha (fase vermelha)
// 2. implementar a função para fazer o teste passar (fase verde)
// 3. refatorar se necessário garantindo que o teste ainda passa (fase de refatoração)
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// oddOrEven determina se um número é impar ou par usando a metodologia TDD
// o teste inicial é escrito para falhar para um número de entrada, esperando a saida correta em string,
// a função é então implementada para passar este teste e casos extremos adicionais sao considerados
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

// checkhealth verifica a saúde da aplicação e é projetada usando principios de TDD
// um teste é inicialmente criado para esperar uma certa resposta HTTP, que falha
// a função é então elaborada para satisfazer este teste, garantindo que o ponto de verificação de saúde do aplicativo se comporte conforme o esperado
func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}