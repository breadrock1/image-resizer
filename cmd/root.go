package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"image-resize-service/internal/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:   "./resizer --config=/path/to/config.toml",
	Short: "Launch resizer service",
	Long:  `There is command line interface to launch resizer service.`,
	Run: func(cmd *cobra.Command, _ []string) {
		if configPath, err := cmd.Flags().GetString("config"); err == nil {
			log.Println("Launching resizer service with config: ", configPath)
			return
		}
		log.Fatal("Failed while launching resizer ...")
	},
}

func Execute() *config.Config {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}

	versionFlag, _ := rootCmd.Flags().GetCount("version")
	if versionFlag > 0 {
		PrintVersion()
		os.Exit(0)
	}

	configPath, _ := rootCmd.Flags().GetString("config")
	parsedConfig, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Failed while parsing config file: %s", err)
	}

	return parsedConfig
}

func init() {
	flagSet := rootCmd.Flags()

	flagSet.StringP("config", "c", "./configs/config.toml", "Path to toml config file.")
	flagSet.CountP("version", "V", "Print resizer project version.")
}
