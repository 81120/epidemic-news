package cmd

import (
	"github.com/81120/epidemic-news/service"

	"github.com/spf13/cobra"
)

var areaCmd = &cobra.Command{
	Use:   "area xx",
	Short: "area",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		area, _ := cmd.Flags().GetString("area")
		if area == "" {
			area = "湖北"
		}
		res := service.FetchDataByArea(area)
		res.DisplayCityList()
	},
}

func init() {
	rootCmd.AddCommand(areaCmd)
	areaCmd.Flags().StringP("area", "a", "湖北", "")
}
