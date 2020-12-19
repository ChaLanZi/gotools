package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(sqlCmd)
}
