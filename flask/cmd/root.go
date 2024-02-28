/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"swag.org/http-client/client"
	"swag.org/http-client/client/operations"
)

var pallet string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "http-client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		param := operations.NewGetColorsPaletteParams()
		param.SetPalette(pallet)

		cfg := client.DefaultTransportConfig()
		cfg.Host = "localhost:5000"
		demo := client.NewHTTPClientWithConfig(nil, cfg)

		resp, err := demo.Operations.GetColorsPalette(param)

		if err != nil {
			fmt.Printf("error %+v", err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendRow(table.Row{"colour", resp.Payload.Color})
		t.Render()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.subh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVar(&pallet, "pallet", "all", "name of pallet")
}
