package cmd

import (
	"IMfourm-go/database/seeders"
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/seed"
	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use: "seed",
	Short: "insert fake data to the database",
	Run: runSeeders,
	Args: cobra.MaximumNArgs(1),
}

func runSeeders(cmd *cobra.Command, args []string){
	seeders.Initialize()
	if len(args) > 0 {
		//有传参的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("seeder not found: " + name)
		}
	} else {
		//默认运行全部
		seed.RunAll()
		console.Success("done seeding.")
	}
}
