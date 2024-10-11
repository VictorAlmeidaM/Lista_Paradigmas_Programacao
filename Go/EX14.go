package main

import (
	"fmt"
	"sync"
)

// Configuracao representa a configuração da aplicação
type Configuracao struct {
	// Adicione campos de configuração conforme necessário
	url     string
	timeout int
}

// Variável para armazenar a instância única da Configuracao
var instancia *Configuracao

// Mutex para controlar o acesso à instância
var mu sync.Mutex

// getInstance retorna a instância única da classe Configuracao
func getInstance() *Configuracao {
	mu.Lock()
	defer mu.Unlock()

	// Verifica se a instância já foi criada
	if instancia == nil {
		instancia = &Configuracao{}
	}
	return instancia
}

// Métodos para definir e obter configurações
func (c *Configuracao) SetConfig(url string, timeout int) {
	c.url = url
	c.timeout = timeout
}

func (c *Configuracao) GetConfig() (string, int) {
	return c.url, c.timeout
}

// Exemplo de uso
func main() {
	// Obtendo a instância única da Configuracao
	config1 := getInstance()
	config1.SetConfig("https://example.com", 30)

	// Obtendo a mesma instância
	config2 := getInstance()
	url, timeout := config2.GetConfig()

	// Exibindo as configurações
	fmt.Printf("URL: %s, Timeout: %d\n", url, timeout)

	// Verificando se config1 e config2 são a mesma instância
	fmt.Printf("config1 é config2? %v\n", config1 == config2) // Saída: true
}
