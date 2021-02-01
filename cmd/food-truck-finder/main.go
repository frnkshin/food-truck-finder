// Package main provides entry point and command line program for food-truck-finder.
package main

import (
	"fmt"
	"food-truck-finder/pkg/foodtruck"
	"github.com/SebastiaanKlippert/go-soda"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	URLFlagName    = "url"
	LimitFlagName  = "limit"
	SortFlagName   = "asc"
	FormatFlagName = "format"
	TokenFlagName  = "token"
	OpenNowStatement = `dayofweekstr='%s' AND start24<='%s' AND end24>'%s'`
	SodasHHMMFormat = "%d:%d"
	SortColumn = "applicant"
)


// Main runs the main program. It uses cli library to process all parameters given
// to the application at runtime, and as well as sets default values for params.
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
				Usage:       "Boolean value to sort listings by applicant (name). true = asc, false = desc",
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
							hhmm := getCurrentHHMM(SodasHHMMFormat)
							requestMetadata := foodtruck.NewRequestBuilder(client.BaseURL(), client.Token()).SetFormat(client.Format()).SetWhere(fmt.Sprintf(OpenNowStatement, time.Now().Weekday(), hhmm, hhmm)).SetOrder(SortColumn, soda.DirDesc)
							if client.Sorted() {
								requestMetadata = requestMetadata.SetOrder(SortColumn, soda.DirAsc)
							}
							client.GetFoodtrucksPaginated(*requestMetadata, handleShowAllView)
							return nil
						},
					},
					{
						Name:    "now",
						Aliases: []string{"c"},
						Usage:   "returns listings of fodd trucks that are open now",
						Action: func(ctx *cli.Context) error {
							client := foodtruck.NewClient(ctx.String(URLFlagName), ctx.Uint(LimitFlagName), ctx.Bool(SortFlagName), ctx.String(FormatFlagName), ctx.String(TokenFlagName))
							hhmm := getCurrentHHMM(SodasHHMMFormat)
							requestMetadata := foodtruck.NewRequestBuilder(client.BaseURL(), client.Token()).SetFormat(client.Format()).SetWhere(fmt.Sprintf(OpenNowStatement, time.Now().Weekday(), hhmm, hhmm)).SetOrder(SortColumn, soda.DirDesc)
							if client.Sorted() {
								requestMetadata = requestMetadata.SetOrder(SortColumn, soda.DirAsc)
							}
							client.GetFoodtrucksPaginated(*requestMetadata, handlePaginatedView)
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
