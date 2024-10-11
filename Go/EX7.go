package main

import (
	"fmt"
)

type Professor struct {
	nome    string
	escolas []*Escola // Lista para armazenar as escolas que o professor leciona
}

func NewProfessor(nome string) *Professor {
	return &Professor{nome: nome, escolas: []*Escola{}}
}

func (p *Professor) adicionarEscola(escola *Escola) {
	// Verifica se a escola já está na lista
	for _, e := range p.escolas {
		if e == escola {
			return
		}
	}
	p.escolas = append(p.escolas, escola)
	escola.adicionarProfessor(p) // Adiciona o professor à escola também
}

func (p *Professor) String() string {
	return p.nome
}

type Escola struct {
	nome        string
	professores []*Professor // Lista para armazenar os professores da escola
}

func NewEscola(nome string) *Escola {
	return &Escola{nome: nome, professores: []*Professor{}}
}

func (e *Escola) adicionarProfessor(professor *Professor) {
	// Verifica se o professor já está na lista
	for _, p := range e.professores {
		if p == professor {
			return
		}
	}
	e.professores = append(e.professores, professor)
}

func (e *Escola) String() string {
	return e.nome
}

func (e *Escola) listarProfessores() []string {
	nomesProfessores := make([]string, len(e.professores))
	for i, professor := range e.professores {
		nomesProfessores[i] = professor.String()
	}
	return nomesProfessores
}

func main() {
	escola1 := NewEscola("Escola A")
	escola2 := NewEscola("Escola B")

	professor1 := NewProfessor("Roberto")
	professor2 := NewProfessor("Fernanda")

	// Associando professores às escolas
	professor1.adicionarEscola(escola1)
	professor1.adicionarEscola(escola2)
	professor2.adicionarEscola(escola1)

	// Exibindo os professores de cada escola
	fmt.Printf("Professores da %s: ", escola1)
	for i, nome := range escola1.listarProfessores() {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(nome)
	}
	fmt.Println()

	fmt.Printf("Professores da %s: ", escola2)
	for i, nome := range escola2.listarProfessores() {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(nome)
	}
	fmt.Println()
}
