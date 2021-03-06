package commands

import (
	"bytes"

	"github.com/jaysonesmith/stardew-crops/formatter"
	"github.com/jaysonesmith/stardew-crops/processors"
	"github.com/jaysonesmith/stardew-crops/utils"
	"github.com/spf13/cobra"
)

// infoCmd returns data based on the input crop
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information on a specific crop",
	Long:  `Specify a crop name to get more information on it i.e. 'stardew-crops info garlic'`,
	Run:   Info,
}

func Info(cmd *cobra.Command, args []string) {
	userFlags := utils.UserSetFlags(cmd)

	b := bytes.NewBuffer([]byte{})
	out, err := processors.Info(args...)
	if err != nil {
		b = formatter.Format(err.Error(), "error")
	} else {
		b = formatter.Format(out, userFlags["format"])
	}

	print(b)
}

func init() {
	StardewCrops.AddCommand(infoCmd)
}
