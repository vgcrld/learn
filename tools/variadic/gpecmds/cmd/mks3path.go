/*
	Turn this: "/share/process/True_Value/input/exetrn.20191211.182014.GMT__755922__.914156-459137866.oracle.gz"

	Into this: "s3://etl-x2gpe.galileosuite.com/True_Value/2019/12/11/exetrn.20191211.182014.GMT__755922__.914156-459137866.oracle.gz"
*/
package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// mks3pathCmd represents the mks3path command
var mks3pathCmd = &cobra.Command{
	Use:     "mks3path",
	Short:   "Change a legacy file name to an s3 path",
	Example: "mks3path /share/prod/process/atsgroup/input/gvic1.20191101.124000.GMT__334234_.uuid-123-455.oracle.gz",
	Run: func(cmd *cobra.Command, args []string) {
		spew.Dump(cmd.Example)
	},
}

func init() {
	rootCmd.AddCommand(mks3pathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mks3pathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mks3pathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
