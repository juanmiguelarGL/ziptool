package cmd

import (
    "compress/gzip"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "os"

    "github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var prettyPrint bool

// decodeCmd represents the command for decoding + decompressing
var decodeCmd = &cobra.Command{
    Use:   "decode",
    Short: "Decodes from base64 and decompresses GZIP data from an input file",
    Long: `Reads a file containing base64-encoded data, decodes it, and then 
decompresses it using GZIP to retrieve the original JSON. If -o/--output is 
provided, the result is saved to the specified file; otherwise, it is printed 
to the console. Use --pretty to beautify JSON output in the console.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // 1. Validate input file flag
        if inputFile == "" {
            return fmt.Errorf("you must specify an input file with -i or --input")
        }

        // 2. Read base64-encoded data from input file
        base64Content, err := os.ReadFile(inputFile)
        if err != nil {
            return fmt.Errorf("error reading input file '%s': %w", inputFile, err)
        }

        // 3. Decode from base64
        compressedData, err := base64.StdEncoding.DecodeString(string(base64Content))
        if err != nil {
            return fmt.Errorf("error decoding base64 string: %w", err)
        }

        // 4. Decompress GZIP
        jsonData, err := gzipDecompress(compressedData)
        if err != nil {
            return fmt.Errorf("error decompressing data: %w", err)
        }

        // 5. Save to file or print to console
        if outputFile != "" {
            // Save to output file
            if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
                return fmt.Errorf("error saving result to '%s': %w", outputFile, err)
            }
            fmt.Printf("JSON content has been saved to '%s'\n", outputFile)
        } else {
            // Print to console
            if prettyPrint {
                // Attempt to pretty-print the JSON
                var tmp interface{}
                if err := json.Unmarshal(jsonData, &tmp); err != nil {
                    return fmt.Errorf("error parsing JSON for pretty print: %w", err)
                }
                prettyData, err := json.MarshalIndent(tmp, "", "  ")
                if err != nil {
                    return fmt.Errorf("error marshaling JSON for pretty print: %w", err)
                }
                fmt.Println(string(prettyData))
            } else {
                // Standard (raw) output
                fmt.Println(string(jsonData))
            }
        }

        return nil
    },
}

func init() {
    decodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input file containing base64-encoded data")
    decodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to the file where the decompressed JSON will be saved")
    decodeCmd.Flags().BoolVar(&prettyPrint, "pretty", false, "Beautify JSON output in the console (no effect if -o is used)")
}

// gzipDecompress decompresses GZIP data and returns the original bytes
func gzipDecompress(compressed []byte) ([]byte, error) {
    r, w := io.Pipe()
    go func() {
        _, _ = w.Write(compressed)
        _ = w.Close()
    }()

    gr, err := gzip.NewReader(r)
    if err != nil {
        return nil, err
    }
    defer gr.Close()

    return io.ReadAll(gr)
}
