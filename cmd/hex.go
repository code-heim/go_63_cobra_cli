/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generates Hex numbers",
	Long: `Provided length, it generates hex numbers.
For example:
securerandom hex -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		hex, err := generateHexString(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(hex)
		}
	},
}

func generateHexString(length int) (string, error) {
	// Calculate the number of bytes needed (each byte is two hex characters)
	byteLength := (length + 1) / 2
	bytes := make([]byte, byteLength)

	// Generate random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Convert to hex and truncate to the desired length
	hexString := hex.EncodeToString(bytes)[:length]
	return hexString, nil
}

func init() {
	rootCmd.AddCommand(hexCmd)

	hexCmd.Flags().IntP("length", "l", 4, "Length of hex")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
