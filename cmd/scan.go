package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/phoneinfoga/pkg/scanners"
	"github.com/sundowndev/phoneinfoga/pkg/utils"
)

func init() {
	// Register command
	rootCmd.AddCommand(scanCmd)

	// Register flags
	scanCmd.PersistentFlags().StringVarP(&number, "number", "n", "", "The phone number to scan (E164 or international format)")
	scanCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Text file containing a list of phone numbers to scan (one per line)")
	scanCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output to save scan results")
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a phone number",
	Run: func(cmd *cobra.Command, args []string) {
		n := utils.FormatNumber(number)

		utils.LoggerService.Infoln("Scanning phone number", n)

		scanners.ScanCLI(n)

		utils.LoggerService.Infoln("Job finished.")
	},
}
