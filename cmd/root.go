package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// RootCmd is the root command of the application.
var RootCmd = &cobra.Command{
    Use:   "ziptool",
    Short: "ziptool is a tool for compressing and decompressing data with GZIP and base64",
    Long:  `ziptool allows you to compress (encode) and decompress (decode) JSON content using GZIP + base64.`,
    // You could define a Run function here if you want the root command to do something by default,
    // but typically, actions are left to the subcommands.
    Run: func(cmd *cobra.Command, args []string) {
        // Default message if no subcommands are called
        fmt.Println("ziptool: Use 'encode' or 'decode'. Run 'ziptool --help' for more information.")
    },
}

func init() {
    // Add the subcommands
    RootCmd.AddCommand(encodeCmd)
    RootCmd.AddCommand(decodeCmd)
}
