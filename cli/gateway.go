package cli

import (
	"fmt"
	"os"
	"os/exec"
)


func ExecScripts(scriptPath string) {
       
    // Verifica se o script existe
    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        fmt.Printf("O script %s não existe.\n", scriptPath)
        return
    }

    // Executa o script
    cmd := exec.Command("bash", scriptPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Erro ao executar o script: %v\n", err)
        return
    }

    fmt.Sprintf("Script executado com sucesso.", scriptPath)
}