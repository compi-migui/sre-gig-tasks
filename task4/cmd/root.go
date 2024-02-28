package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flags.
var print_lines bool
var print_bytes bool

type results struct {
	filename string
	lines int
	bytes int
}


func ScanFile(path string) (*results, error) {
	ret := results{filename: path, lines: 0, bytes: 0}

	var file *os.File
	var err error
	if path == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(path)
		if err != nil  {
			return nil, err
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	// bufio's ScanLines strips \n and \r\n from the data, leading to
	// incorrect byte counts. So we have to count them ourselves.
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		ret.bytes += 1
		if scanner.Bytes()[0] == '\n' {
			ret.lines += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &ret, nil
}


func PrintResults(result results, print_lines bool, print_bytes bool) {
	if print_lines {
		fmt.Printf("%d ", result.lines)
	}
	if print_bytes {
		fmt.Printf("%d ", result.bytes)
	}

	fmt.Printf("%s", result.filename)
	fmt.Printf("\n")
}

var rootCmd = &cobra.Command{
	Use:   "gowc",
	Short: "wc - print newline and byte counts for each file",
	Long: `Print newline and byte counts for each file.

With no FILE, or when FILE is -, read standard input.

The  options  below  may  be used to select which counts are printed, always in the following order: newline, byte.`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !print_lines && !print_bytes {
			print_lines = true
			print_bytes = true
		}

		if len(args) == 0 {
			args = append(args, "-")
		}

		total := results{filename: "total", lines: 0, bytes: 0}

		for _, path := range args {
			// TODO: if printing to a pipe, print only one total number
			result, err := ScanFile(path)
			if err != nil  {
				// Don't print usage text for internal errors.
				// Setting this here preserves usage printing for errors in
				// command/flag parsing.
				cmd.SilenceUsage = true
				return err
			}
			PrintResults(*result, print_lines, print_bytes)
			total.lines += result.lines
			total.bytes += result.bytes
		}
		if len(args) > 1 {
			PrintResults(total, print_lines, print_bytes)
		}
		return nil
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// Error message is printed by cobra, so we just exit.
		// TODO: return more appropriate error codes?
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&print_bytes, "bytes", "c", false, "print the byte counts")
	rootCmd.PersistentFlags().BoolVarP(&print_lines, "lines", "l", false, "print the line counts")
}
