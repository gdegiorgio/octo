package version

import (
	"os"
	"github.com/spf13/cobra"
)

func NewVersionCmd() (*cobra.Command){
	return &cobra.Command{
		Use: "version",
		Short: "Show current Octo CLI version",
		Long: "Show current Octo CLI version",
		Run: func(cmd *cobra.Command, args []string){
			version := os.Getenv("OCTO_CLI_VERSION")
			if version == "" {
				version = "Unknown Version"
			}
			cmd.Println(version)
		},
	}
}
