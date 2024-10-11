class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []  # Lista para armazenar as escolas que o professor leciona

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)  # Adiciona o professor à escola também

    def __str__(self):
        return self.nome


class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []  # Lista para armazenar os professores da escola

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)

    def __str__(self):
        return self.nome

    def listar_professores(self):
        return [str(professor) for professor in self.professores]


# Exemplo de uso
if __name__ == "__main__":
    escola1 = Escola("Escola A")
    escola2 = Escola("Escola B")

    professor1 = Professor("Roberto")
    professor2 = Professor("Fernanda")

    # Associando professores às escolas
    professor1.adicionar_escola(escola1)
    professor1.adicionar_escola(escola2)
    professor2.adicionar_escola(escola1)

    # Exibindo os professores de cada escola
    print(f"Professores da {escola1}: {escola1.listar_professores()}")
    print(f"Professores da {escola2}: {escola2.listar_professores()}")
