package cmd

import (
	"github.com/81120/epidemic-news/service"

	"github.com/spf13/cobra"
)

var cityCmd = &cobra.Command{
	Use:   "city xx",
	Short: "get data for city",
	Long:  "",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city, _ := cmd.Flags().GetString("city")
		if city == "" {
			city = "武汉"
		}
		res := service.FetchDataByCity(city)
		res.DisplayCityData()
	},
}

func init() {
	rootCmd.AddCommand(cityCmd)
	cityCmd.Flags().StringP("city", "c", "武汉", "")
}
