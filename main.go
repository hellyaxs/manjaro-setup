package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/spf13/cobra"
    "setup/cli"
)

func execScripts(scriptPath string) {
       
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

    fmt.Println("Script executado com sucesso.")
}



func main() {
    reader := bufio.NewReader(os.Stdin)
    var rootCmd = &cobra.Command{
        Use:   "mycli",
        Short: "My CLI is a tool to demonstrate how to build a CLI in Go",
        Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
        Run: func(cmd *cobra.Command, args []string) {
            // Default action when no subcommands are passed
        text :=`
        ███╗░░░███╗░█████╗░███╗░░██╗░░░░░██╗░█████╗░█ ████╗░░█████╗░  ░██████╗███████╗████████╗██╗░░░██╗██████╗░
        ████╗░████║██╔══██╗████╗░██║░░░░░██║██╔══██╗██╔══██╗██╔══██╗  ██╔════╝██╔════╝╚══██╔══╝██║░░░██║██╔══██╗
        ██╔████╔██║███████║██╔██╗██║░░░░░██║███████║██████╔╝██║░░██║  ╚█████╗░█████╗░░░░░██║░░░██║░░░██║██████╔╝
        ██║╚██╔╝██║██╔══██║██║╚████║██╗░░██║██╔══██║██╔══██╗██║░░██║  ░╚═══██╗██╔══╝░░░░░██║░░░██║░░░██║██╔═══╝░
        ██║░╚═╝░██║██║░░██║██║░╚███║╚█████╔╝██║░░██║██║░░██║╚█████╔╝  ██████╔╝███████╗░░░██║░░░╚██████╔╝██║░░░░░
        ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝░╚════╝░╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░  ╚═════╝░╚══════╝░░░╚═╝░░░░╚═════╝░╚═╝░░░░░`
            fmt.Println(text)
            fmt.Println("\n")
            cli.StartViews()
        },
    }


    rootCmd.Execute()
    for {
        
        fmt.Println("Please enter a command (type 'exit' to quit):")
        fmt.Print("Selecione o tipo de instalação que deseja:\n")
        fmt.Print("1. instalação padrão (manjaro install apps)\n")
        fmt.Print("2. instalação cli (manjaro install cli)\n")
        fmt.Print("3. instalação web/dev (manjaro install web)\n")
        fmt.Print("> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "exit":
            fmt.Println("Exiting...")
            return
        case "hello":
            fmt.Println("Hello! How can I assist you today?")
        case "confirm":
            fmt.Println("Do you want to proceed? (yes/no):")
            fmt.Print("> ")
            confirmation, _ := reader.ReadString('\n')
            confirmation = strings.TrimSpace(confirmation)
            if confirmation == "yes" {
                fmt.Println("Proceeding with the action.")
            } else {
                fmt.Println("Action canceled.")
            }
        case "help":
            fmt.Println("Available commands: hello, confirm, help, exit")
        case "1":
            fmt.Println("preparando setup para [instalação de apps]...")
            //  execScripts("./scripts/manjaro_install_apps.sh")
             return
        case "2":
            fmt.Println("preparando setup para [instalação de cli]...")
            // execScripts("./scripts/manjaro_install_cli.sh")
            return
        case "3":
            fmt.Println("preparando setup para [instalação de web/dev]...")
            // execScripts("./scripts/manjaro_install_web_dev.sh")
            return
        default:
            fmt.Println("Unknown command. Type 'help' to see available commands.")
        }
    }
}
