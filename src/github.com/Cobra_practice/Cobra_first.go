package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	//var echoTimes int

	var cmdAdd = &cobra.Command{
		Use:   "add [filename to add]",
		Short: "Add a file to IPFS",
		Long:  `add is for adding a file to the IPFS system`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//add code for adding file to IPFS
			fmt.Println("Filename: " + strings.Join(args, " "))
		},
	}

	var cmdUpload = &cobra.Command{
		Use:   "upload [JWT to upload]",
		Short: "upload file to Blockchain",
		Long:  `upload the file,filehash and the owner has as a JWT to the blockchain`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//add code to upload to blockchain
			fmt.Println("JWT: " + strings.Join(args, " "))
		},
	}

	var cmdGet = &cobra.Command{
		Use:   "get [filename to get]",
		Short: "get a file from the blockchain",
		Long:  `get the file from the blockchain with appropriate permissions`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//add code to get a file from blockchain
			fmt.Println("Filename: " + strings.Join(args, " "))
		},
	}

	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "dswac"}
	rootCmd.AddCommand(cmdAdd, cmdUpload, cmdGet)
	rootCmd.Execute()
	//cmdPrint.Execute()
}
