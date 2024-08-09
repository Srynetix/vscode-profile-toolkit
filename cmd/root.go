package cmd

import (
	"fmt"
	"os"

	"github.com/Srynetix/vscode-profile-toolkit/pkg/archiver"
	"github.com/Srynetix/vscode-profile-toolkit/pkg/extractor"
	"github.com/Srynetix/vscode-profile-toolkit/pkg/parser"
	"github.com/spf13/cobra"
)

var archiveInput string
var archiveOutput string
var extractInput string
var extractOutput string

var rootCmd = &cobra.Command{
	Use:   "vs-prof-tk",
	Short: "A tool to manage .code-profile files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Create a new .code-profile file from a profile directory",
	Run: func(cmd *cobra.Command, args []string) {
		packParser := &parser.ProfilePackParser{}
		packArchiver := &archiver.ProfilePackArchiver{}

		pack := packParser.ParseFolder(archiveInput)
		packArchiver.ArchiveTo(pack, archiveOutput)
	},
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract a .code-profile file to a profile directory",
	Run: func(cmd *cobra.Command, args []string) {
		packParser := &parser.ProfilePackParser{}
		packExtractor := &extractor.ProfilePackExtractor{}

		pack := packParser.ParsePath(extractInput)
		packExtractor.Extract(pack, extractOutput)
	},
}

func init() {
	archiveCmd.PersistentFlags().StringVarP(&archiveInput, "input", "i", "", "input profile directory")
	archiveCmd.MarkPersistentFlagRequired("input")
	archiveCmd.PersistentFlags().StringVarP(&archiveOutput, "output", "o", "", "output path as a .code-profile file")
	archiveCmd.MarkPersistentFlagRequired("output")
	extractCmd.PersistentFlags().StringVarP(&extractInput, "input", "i", "", "input .code-profile file")
	extractCmd.MarkPersistentFlagRequired("input")
	extractCmd.PersistentFlags().StringVarP(&extractOutput, "output", "o", "", "output directory")
	extractCmd.MarkPersistentFlagRequired("output")

	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(extractCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
