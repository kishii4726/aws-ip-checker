package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-ip-checker",
	Short: "Checks if the security group and WEF and ALB contain the specified IP address",
	Long: `Checks if the security group and WAF and ALB contain the specified IP address.
You can specify IP addresses as arguments or read a csv file.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
