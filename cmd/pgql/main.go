package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/trace"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	Version = "0.1.0"
)

func main() {
	rootCmd := &cobra.Command{Use: "pgql", SilenceUsage: true}
	rootCmd.PersistentFlags().StringP("file", "f", "", "specify an alternate config file (default: pgql.yaml)")

	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}

		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the pgql version number",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer trace.StartRegion(cmd.Context(), "version").End()
		fmt.Printf("%s\n", Version)
		return nil
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty pgql.yaml settings file",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init called")
		return nil
	},
}

var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate source code from SQL",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("generate called")
		return nil
	},
}

func getConfigPath(f *pflag.Flag) (string, string) {
	if f != nil && f.Changed {
		file := f.Value.String()
		if file == "" {
			fmt.Println("error parsing config: file argument is empty")
			os.Exit(1)
		}

		abspath, err := filepath.Abs(file)
		if err != nil {
			fmt.Printf("error parsing config: absolute file path lookup failed: %s\n", err)
			os.Exit(1)
		}

		return filepath.Dir(abspath), filepath.Base(abspath)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error parsing pgql.json: file does not exist")
		os.Exit(1)
	}

	return wd, ""
}
