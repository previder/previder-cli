package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/previder/previder-go-sdk/client"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	accessToken    string
	customerId     string
	baseUri        string
	outputType     string
	previderClient *client.PreviderClient
)

var rootCmd = &cobra.Command{
	Use:   "previder-cli",
	Short: "Previder CLI is the command line client for the Previder Portal",
	Long: `Previder CLI is the command line client for the Previder Portal.
More information can be found at https://portal.previder.com/api-docs.html or at https://previder.com`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if accessToken == "" {
			if os.Getenv("PREVIDER_TOKEN") != "" {
				accessToken = os.Getenv("PREVIDER_TOKEN")
			} else {
				log.Fatal("No token found")
			}
		}

		previderClient, err = client.New(&client.ClientOptions{Token: accessToken, CustomerId: customerId, BaseUrl: baseUri})
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&accessToken, "token", "t", "", "The Previder access token")
	rootCmd.PersistentFlags().StringVarP(&customerId, "customer", "c", "", "An optional subcustomer id")
	rootCmd.PersistentFlags().StringVarP(&outputType, "output", "o", "json", "Output format [pretty / json]")
	rootCmd.PersistentFlags().StringVarP(&baseUri, "uri", "u", "https://portal.previder.nl/api/", "Optional different URI")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printJson(input any) {
	marshal, _ := json.MarshalIndent(input, "", "  ")
	println(string(marshal))
}

func printTable(headers []string, content []interface{}) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headers)

	for _, row := range content {
		var rowValues []string
		rowVal := reflect.ValueOf(row)
		for _, header := range headers {
			field := rowVal.FieldByName(header)
			if !field.IsValid() {
				rowValues = append(rowValues, "")
				continue
			}

			var value string
			switch field.Kind() {
			case reflect.Bool:
				value = strconv.FormatBool(field.Bool())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(field.Int(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.String:
				value = field.String()
			case reflect.Slice:
				stringSlice, _ := field.Interface().([]string)
				value = strings.Join(stringSlice, ",")
			case reflect.Uint64:
				value = strconv.FormatUint(field.Uint(), 10)
			default:
				value = "Unsupported Type " + field.Kind().String()
			}
			rowValues = append(rowValues, value)
		}
		table.Append(rowValues)
	}
	table.Render()
}
