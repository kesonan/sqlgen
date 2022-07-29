package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/anqiansong/sqlgen/internal/gen/flags"
)

var arg flags.RunArg

var rootCmd = &cobra.Command{
	Use:   "sqlgen",
	Short: "A cli for mysql generator",
}
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "Generate SQL model",
	Run: func(cmd *cobra.Command, args []string) {
		arg.Mode = flags.SQL
		flags.Run(arg)
	},
}

var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "Generate gorm model",
	Run: func(cmd *cobra.Command, args []string) {
		arg.Mode = flags.GORM
		flags.Run(arg)
	},
}

var xormCmd = &cobra.Command{
	Use:   "xorm",
	Short: "Generate xorm model",
	Run: func(cmd *cobra.Command, args []string) {
		arg.Mode = flags.XORM
		flags.Run(arg)
	},
}

var sqlxCmd = &cobra.Command{
	Use:   "sqlx",
	Short: "Generate sqlx model",
	Run: func(cmd *cobra.Command, args []string) {
		arg.Mode = flags.SQLX
		flags.Run(arg)
	},
}

func init() {
	// flags init
	var persistentFlags = rootCmd.PersistentFlags()
	persistentFlags.StringVarP(&arg.DSN, "dsn", "d", "", "Mysql address")
	persistentFlags.StringSliceVarP(&arg.Table, "table", "t", []string{"*"}, "Patterns of table name")
	persistentFlags.StringSliceVarP(&arg.Filename, "filename", "f", []string{"*.sql"}, "Patterns of SQL filename")
	persistentFlags.StringVarP(&arg.Output, "output", "o", ".", "The output directory")

	// sub commands init
	rootCmd.AddCommand(sqlCmd)
	rootCmd.AddCommand(gormCmd)
	rootCmd.AddCommand(xormCmd)
	rootCmd.AddCommand(sqlxCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}