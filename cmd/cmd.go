package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/anqiansong/sqlgen/internal/gen/flags"
)

var dsn string
var filename []string
var table []string

var rootCmd = &cobra.Command{
	Use:   "sqlgen",
	Short: "A cli for mysql generator",
}
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "Generate SQL model",
	Run: func(cmd *cobra.Command, args []string) {
		flags.Run(dsn, filename, table, flags.SQL)
	},
}

var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "Generate gorm model",
	Run: func(cmd *cobra.Command, args []string) {
		flags.Run(dsn, filename, table, flags.GORM)
	},
}

var xormCmd = &cobra.Command{
	Use:   "xorm",
	Short: "Generate xorm model",
	Run: func(cmd *cobra.Command, args []string) {
		flags.Run(dsn, filename, table, flags.XORM)
	},
}

var sqlxCmd = &cobra.Command{
	Use:   "sqlx",
	Short: "Generate sqlx model",
	Run: func(cmd *cobra.Command, args []string) {
		flags.Run(dsn, filename, table, flags.SQLX)
	},
}

func init() {
	// flags init
	var persistentFlags = rootCmd.PersistentFlags()
	persistentFlags.StringVarP(&dsn, "dsn", "d", "", "Mysql address")
	persistentFlags.StringSliceVarP(&table, "table", "t", []string{"*"}, "Patterns of table name")
	persistentFlags.StringSliceVarP(&filename, "filename", "f", []string{"*.sql"}, "Patterns of SQL filename")

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
