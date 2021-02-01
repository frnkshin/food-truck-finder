package main

import (
	"fmt"
	"food-truck-finder/pkg/foodtruck"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"github.com/eiannone/keyboard"
	tm "github.com/buger/goterm"
)

const (
	URLFlagName    = "url"
	LimitFlagName  = "limit"
	SortFlagName   = "sort"
	FormatFlagName = "format"
	TokenFlagName  = "token"
)

func main() {
	app := &cli.App{
		Name:  "food-truck-finder",
		Usage: "Find food trucks in San Francisco that's open right now",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        URLFlagName,
				Usage:       "Base URL to query against",
				Value:       foodtruck.DefaultBaseURL,
				DefaultText: foodtruck.DefaultBaseURL,
			},
			&cli.UintFlag{
				Name:        LimitFlagName,
				Usage:       "Number of food trucks to show per page",
				Value:       foodtruck.DefaultPageLimit,
				DefaultText: string(foodtruck.DefaultPageLimit),
			},
			&cli.BoolFlag{
				Name:        SortFlagName,
				Usage:       "Boolean value to sort listings",
				Value:       foodtruck.DefaultSorted,
				DefaultText: strconv.FormatBool(foodtruck.DefaultSorted),
			},
			&cli.StringFlag{
				Name:        FormatFlagName,
				Usage:       "Type of response format from Food Trucks API",
				Value:       foodtruck.DefaultFormat,
				DefaultText: foodtruck.DefaultFormat,
			},
			&cli.StringFlag{
				Name:        TokenFlagName,
				Usage:       "API Token for Food Trucks API",
				Value:       foodtruck.DefaultToken,
				DefaultText: foodtruck.DefaultToken,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "find",
				Aliases: []string{"f"},
				Usage:   "Starts finding food trucks",
				Subcommands: []*cli.Command{
					{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "returns all listings of food trucks",
						Action: func(ctx *cli.Context) error {
							client := foodtruck.NewClient(ctx.String(URLFlagName), ctx.Uint(LimitFlagName), ctx.Bool(SortFlagName), ctx.String(FormatFlagName), ctx.String(TokenFlagName))
							client.GetFoodtrucks(func(trucks *foodtruck.FoodTrucks) {
								for _, truck := range *trucks {
									fmt.Println(truck.String())
								}
							})
							return nil
						},
					},
					{
						Name:    "now",
						Aliases: []string{"c"},
						Usage:   "returns listings of fodd trucks that are open now",
						Action: func(ctx *cli.Context) error {
							client := foodtruck.NewClient(ctx.String(URLFlagName), ctx.Uint(LimitFlagName), ctx.Bool(SortFlagName), ctx.String(FormatFlagName), ctx.String(TokenFlagName))
							client.GetFoodtrucksPaginated(func(trucks *foodtruck.FoodTrucks) {
								tm.Clear()
								for _, truck := range *trucks {
									fmt.Println(truck.String())
									tm.Flush()
								}
								fmt.Println("Press any key to show next results\nPress Q to quit")
								char, _, err := keyboard.GetSingleKey()
								if (err != nil) {
									panic(err)
								}
								switch char {
									case 'q', 'Q':
										fmt.Println("Pressed Q -- Quitting program")
										os.Exit(0)
										break
									default:
										tm.Clear()
										return
								}
							})
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
