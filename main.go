package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "setup/cli"
)

func main() {
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

}
