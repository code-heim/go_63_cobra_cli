/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates UUID",
	Long: `Generates UUID. For example:

securerandom uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID := uuid.New()
		fmt.Println(newUUID.String())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uuidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
