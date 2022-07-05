package cmd

import (
	"IMfourm-go/database/migrations"
	"IMfourm-go/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use: "migrate",
	Short: "Run database migration",
	//所有migrate下的子命令都会执行以下代码
}

var CmdMigrateUp = &cobra.Command{
	Use: "up",
	Short: "Run unmigrated migrations",
	Run: runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use: "down",
	Aliases: []string{"rollback"},
	Short: "Reverse the up command",
	Run: runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use: "reset",
	Short: "Rollback all database migrations",
	Run: runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use: "refresh",
	Short: "Reset and re-run all migrations",
	Run: runRefresh,
}

var CmdMigrateFresh = &cobra.Command{
	Use: "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run: runFresh,
}

func init()  {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateReset,
		CmdMigrateRefresh,
		CmdMigrateFresh,
		)
}
func migrator() *migrate.Migrator{
	//注册database/migrations下的所有迁移文件
	migrations.Initialize()
	//初始化migrator
	return migrate.NewMigrator()
}
func runUp(cmd *cobra.Command, ags []string)  {
	migrator().Up()
}
func runDown(cmd *cobra.Command, args []string)  {
	migrator().Rollback()
}

func runReset(cmd *cobra.Command, args []string)  {
	migrator().Reset()
}
func runRefresh(cmd *cobra.Command, args []string)  {
	migrator().Refresh()
}
func runFresh(cmd *cobra.Command, args []string)  {
	migrator().Fresh()

}