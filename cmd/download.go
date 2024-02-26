package cmd

import "github.com/spf13/cobra"

var downloadCmd = &cobra.Command{
	Use:     "download <playlistname1>",
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"d"},
	Short:   "downloads songs associated with a given spotify playlist",

	Run: func(cmd *cobra.Command, args []string) {

	},
}
