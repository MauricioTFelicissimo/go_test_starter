package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

// testSayHello testa a função sayHello para garantir que ela retorne a mensagem de saudação correta
// este teste se relaciona diretamente com o TDD ao validar que a função adere ao comportamento esperado
// estabelecido na fase inicial 'vermelha' de escrever o teste
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)
	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

// testOddOrEven testa a funcionalidade de números pares e ímpares na função OddOrEven
// estes testes confirmam a fase 'verde' do TDD onde a função deve passar nessas verificações assertivas,
// seguindo os testes iniciais que falharam
func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

// testCheckhealth valida que o ponto de verificação de saúde retorne o status e a resposta corretos
// este é um exemplo de TDD onde a funcionalidade é verificada contra os requisitos delineados no cenario de teste
func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.Checkhealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}