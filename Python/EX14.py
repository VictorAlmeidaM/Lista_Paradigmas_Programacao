class Configuracao:
    _instancia = None  # Atributo de classe para armazenar a instância única

    def __new__(cls, *args, **kwargs):
        if cls._instancia is None:  # Verifica se a instância já foi criada
            cls._instancia = super(Configuracao, cls).__new__(cls)
            # Aqui você pode inicializar qualquer atributo da instância
            cls._instancia.configuracao = {}
        return cls._instancia

    def set_config(self, chave, valor):
        self.configuracao[chave] = valor

    def get_config(self, chave):
        return self.configuracao.get(chave, None)

# Exemplo de uso
if __name__ == "__main__":
    # Criando instâncias da classe Configuracao
    config1 = Configuracao()
    config2 = Configuracao()

    # Definindo configurações
    config1.set_config("url", "https://example.com")
    config1.set_config("timeout", 30)

    # Obtendo configurações usando a segunda instância
    print(f"URL: {config2.get_config('url')}")  # Saída: URL: https://example.com
    print(f"Timeout: {config2.get_config('timeout')}")  # Saída: Timeout: 30

    # Verificando se ambas as instâncias são a mesma
    print(f"config1 é config2? {config1 is config2}")  # Saída: config1 é config2? True
