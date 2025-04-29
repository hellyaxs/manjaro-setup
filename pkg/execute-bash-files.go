package pkg

import (
	"fmt"
	"os"
	"os/exec"
    "sync"
)

type Script struct {
    Name string
    Path string
    Status ScriptStatus

}

const  (
    ScriptStatus Success = "success"
    ScriptStatus Failed = "failed"
    ScriptStatus Running = "running"
    ScriptStatus NotFound = "not found"
)

func ExecScripts(script []Script) {
       
    var wg sync.WaitGroup

	for i := range scripts {
		script := &scripts[i] // capturar o ponteiro do elemento correto
		if _, err := os.Stat(script.Path); os.IsNotExist(err) {
			fmt.Printf("O script %s não existe.\n", script.Name)
			continue
		}

		wg.Add(1)
		go func(s *Script) {
			defer wg.Done()

			cmd := exec.Command("bash", s.Path)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Erro ao executar o script %s: %v\n", s.Name, err)
				s.Status = Failed
				return
			}

			s.Status = Success
			fmt.Printf("Script %s executado com sucesso.\n", s.Name)
		}(script)
	}

	wg.Wait()
}