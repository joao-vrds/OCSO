package main

// Importa as bibliotecas do go //
import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func listarUsuarios() {
	// Executa o comando para listar usuários no Windows //
	cmd := exec.Command("wmic", "useraccount", "get", "name")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao listar usuários: %v\n", err)
		return
	}

	// Obtém a saída do comando e divide em linhas //
	resultado := out.String()
	linhas := strings.Split(resultado, "\n")

	// Imprime cada nome de usuário, pulando a primeira linha //
	for i, linha := range linhas {
		if i > 0 {
			fmt.Println(linha)
		}
	}
}

func main() {
	listarUsuarios()
}
