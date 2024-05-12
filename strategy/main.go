package main

import "fmt"

// Definição do tipo de função para a estratégia
type Strategy func()

// Estratégias concretas
func StrategyA() {
	fmt.Println("Executando a estratégia A")
}

func StrategyB() {
	fmt.Println("Executando a estratégia B")
}

func StrategyC() {
	fmt.Println("Executando a estratégia C")
}

// Contexto
type Context struct {
	strategy Strategy
}

// Método para definir a estratégia atual
func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

// Método para executar a estratégia atual
func (c *Context) ExecuteStrategy() {
	if c.strategy != nil {
		c.strategy()
	} else {
		fmt.Println("Nenhuma estratégia selecionada")
	}
}

func main() {
	context := Context{}

	// Definindo e executando estratégias
	context.SetStrategy(StrategyA)
	context.ExecuteStrategy()

	context.SetStrategy(StrategyB)
	context.ExecuteStrategy()

	context.SetStrategy(StrategyC)
	context.ExecuteStrategy()
}
