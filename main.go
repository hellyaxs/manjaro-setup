package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/spf13/cobra"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var rootCmd = &cobra.Command{
        Use:   "mycli",
        Short: "My CLI is a tool to demonstrate how to build a CLI in Go",
        Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
        Run: func(cmd *cobra.Command, args []string) {
            // Default action when no subcommands are passed
            fmt.Println("tesndaod cobrea model")
        },
    }


    for {
        rootCmd.Execute()
        fmt.Println("Please enter a command (type 'exit' to quit):")
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
        default:
            fmt.Println("Unknown command. Type 'help' to see available commands.")
        }
    }
}
