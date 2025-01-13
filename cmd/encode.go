package cmd

import (
    "compress/gzip"
    "encoding/base64"
    "fmt"
    "io"
    "os"

    "github.com/spf13/cobra"
)

var fileToEncode string

// encodeCmd represents the command for compressing + encoding
var encodeCmd = &cobra.Command{
    Use:   "encode",
    Short: "Compresses and encodes the content of a JSON file to base64",
    Long: `Reads the content of a JSON file, compresses it using GZIP, 
and then encodes it using base64, displaying the result to the console.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        if fileToEncode == "" {
            return fmt.Errorf("you must specify the file path with -f or --file")
        }

        // 1. Read file
        data, err := os.ReadFile(fileToEncode)
        if err != nil {
            return fmt.Errorf("error reading file '%s': %w", fileToEncode, err)
        }

        // 2. Compress data in memory using gzip
        compressedData, err := gzipCompress(data)
        if err != nil {
            return fmt.Errorf("error compressing data: %w", err)
        }

        // 3. Encode to base64
        encoded := base64.StdEncoding.EncodeToString(compressedData)

        // 4. Print result to console
        fmt.Println(encoded)
        return nil
    },
}

func init() {
    // Define the -f/--file flag to specify the JSON file to compress
    encodeCmd.Flags().StringVarP(&fileToEncode, "file", "f", "", "Path to the JSON file to be compressed and encoded")
}

// gzipCompress compresses the data with gzip and returns a []byte
func gzipCompress(data []byte) ([]byte, error) {
    // Create a pipe to write/read in memory
    r, w := io.Pipe()

    go func() {
        gw := gzip.NewWriter(w)
        _, err := gw.Write(data)
        gw.Close()
        w.CloseWithError(err)
    }()

    // Read the compressed content
    compressed, err := io.ReadAll(r)
    if err != nil {
        return nil, err
    }
    return compressed, nil
}
