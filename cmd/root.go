package cmd

import (
	"fmt"
	"github.com/mp091689/gtingo"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var info bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gtingen",
	Short: "GTIN generator/calculator",
	Long: `The "gtingen" is a application that can calculate check digit for different GTIN formats.
Also it can generate random GTINs with specified quantity and specified format(default 13).
Available GTIN formats: 8(GTIN-8), 12(GTIN-12), 13(GTIN-13), 14(GTIN-14).

Examples:
$ gtingen --calculate 1234567 #calculate check digit
$ gtingen -c 1234567 #shorthand of calculate flag
$ gtingen --generate 10 #generate ten random GTIN codes(default format GTIN-13)
$ gtingen -g 10 #shorthand of generate flag
$ gtingen --generate 10 --format 12 #generate ten random GTIN codes with GTIN-12 format
$ gtingen -g 10 -f 12 #shorthand of previous command

All commands can be used with --info or -i flag to show detailed information about GTIN:
$ gtingen -i -c 1234567 #calculate check digit with detailed info
$ gtingen -i -g 10 -f 12 #generate ten random GTIN codes with detailed info

It possible to combine calculation and random generation GTINs:
$ gtingen -i -c 123456789012 -g 10 -f 14 #calculate check digit and generate ten GTINs with detailed info
`,
	Run: func(cmd *cobra.Command, args []string) {
		calc, _ := cmd.Flags().GetUint64("calc")
		if calc != 0 {
			fmt.Println("Calculated result:")
			g := gtingo.Calculate(calc)
			if info {
				format, err := g.GetFormat()
				if err != nil {
					log.Fatalln(err.Error())
				}
				country, err := g.GetCountry()
				if err != nil {
					log.Fatalln(err.Error())
				}
				fmt.Println(g, format, country.Name())
			} else {
				s := string(g)
				fmt.Println(s[len(s)-1:])
			}
		}

		formatInt, _ := cmd.Flags().GetInt("format")
		gen, _ := cmd.Flags().GetInt("generate")
		if gen <= 100 && gen != 0 {
			fmt.Println("Generated GTINs:")
			for i := 0; i < gen; i++ {
				newGtin := gtingo.Generate(formatInt)
				var format string
				var country gtingo.Country
				if info {
					format, _ = newGtin.GetFormat()
					country, _ = newGtin.GetCountry()
				}
				fmt.Println(string(newGtin), format, country.Name())
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Uint64P(
		"calculate",
		"c",
		0,
		"Calculate check digit",
	)
	rootCmd.Flags().IntP(
		"generate",
		"g",
		0,
		"Generate random GTINs with specified quantity. Max value: 100",
	)
	rootCmd.Flags().IntP(
		"format",
		"f",
		13,
		"Specify format to generate GTIN. Available types: 8, 12, 13, 14",
	)
	rootCmd.Flags().BoolVarP(
		&info,
		"info",
		"i",
		false,
		"Show GTIN info",
	)
}
