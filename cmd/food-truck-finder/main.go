package main

import (
	"fmt"
	"food-truck-finder/pkg/foodtruck"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

const (
	URLFlagName = "url"
	LimitFlagName = "limit"
	SortFlagName = "sort"
	TimeoutFlagName = "timeout"
)

func main() {
	app := &cli.App{
		Name: "food-truck-finder",
		Usage: "Find food trucks in San Francisco that's open right now",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: URLFlagName,
				Usage: "Base URL to query against",
				Value: foodtruck.DefaultBaseURL,
				DefaultText: foodtruck.DefaultBaseURL,
			},
			&cli.UintFlag{
				Name: LimitFlagName,
				Usage: "Number of food trucks to show per page",
				Value: 10,
				DefaultText: "10",
			},
			&cli.BoolFlag{
				Name: SortFlagName,
				Usage: "Boolean value to sort listings",
				Value: true,
				DefaultText: "true",
			},
			&cli.DurationFlag{
				Name: TimeoutFlagName,
				Usage: "Timeout duration in seconds to stop waiting for response",
				Value: time.Second * 10,
				DefaultText: "10",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "find",
				Aliases: []string{"f"},
				Usage: "Starts finding food trucks",
				Subcommands: []*cli.Command{
					{
						Name: "all",
						Aliases: []string{"a"},
						Usage: "returns all listings of food trucks",
						Action: func (c *cli.Context) error {
							fmt.Println("dummy text")
							return nil
						},
					},
					{
						Name: "now",
						Aliases: []string{"c"},
						Usage: "returns listings of fodd trucks that are open now",
						Action: func (ctx *cli.Context) error {
							config := foodtruck.NewConfig(ctx.String(URLFlagName), ctx.Uint(LimitFlagName), ctx.Bool(SortFlagName), ctx.Duration(TimeoutFlagName))
							client := foodtruck.NewClient(config)
							client.GetOpenFoodTrucks(nil)
							// use client to call api and fetch
							return nil
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}