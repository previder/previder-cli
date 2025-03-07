package cmd

import (
	"fmt"
	"github.com/previder/previder-go-sdk/client"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	var customerCmd = &cobra.Command{
		Use:   "customer",
		Short: "Customer commands",
	}
	rootCmd.AddCommand(customerCmd)

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Get a list of customers",
		Args:  cobra.NoArgs,
		RunE:  listCustomer,
	}
	customerCmd.AddCommand(cmdList)

	var cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Get a customer",
		Args:  cobra.ExactArgs(1),
		RunE:  getCustomer,
	}
	customerCmd.AddCommand(cmdGet)

	var cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Create a customer",
		RunE:  createCustomer,
	}

	cmdCreate.Flags().StringP("name", "", "", "Customer name")
	cmdCreate.Flags().StringP("accountName", "", "", "Account Name used to log in")
	cmdCreate.Flags().StringP("address", "", "", "The street name")
	cmdCreate.Flags().StringP("addressNumber", "", "", "House number")
	cmdCreate.Flags().StringP("addressSuffix", "", "", "Address suffix")
	cmdCreate.Flags().StringP("zipCode", "", "", "Address zipcode")
	cmdCreate.Flags().StringP("city", "", "", "City of residence")
	cmdCreate.Flags().StringP("countryCode", "", "", "Country code")
	cmdCreate.Flags().StringP("language", "", "", "language")
	cmdCreate.Flags().StringP("purchaseOrderNumber", "", "", "purchaseOrderNumber")
	cmdCreate.Flags().StringP("cocNumber", "", "", "cocNumber")
	cmdCreate.Flags().BoolP("partner", "", false, "partner")
	cmdCreate.Flags().BoolP("hidingPrices", "", false, "hidingPrices")
	cmdCreate.Flags().BoolP("invoiceToPartner", "", false, "invoiceToPartner")
	customerCmd.AddCommand(cmdCreate)

	var cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete a Customer",
		Args:  cobra.ExactArgs(1),
		RunE:  deleteCustomer,
	}
	customerCmd.AddCommand(cmdDelete)

}

func deleteCustomer(cmd *cobra.Command, args []string) error {
	err := previderClient.Customer.Delete(args[0])
	if err != nil {
		return err
	}
	return nil
}

func listCustomer(cmd *cobra.Command, args []string) error {
	var page client.PageRequest
	page.Size = 100
	page.Page = 0
	page.Sort = "+name"
	page.Query = ""
	_, content, err := previderClient.Customer.Page(page)
	if err != nil {
		fmt.Println(err)
	}
	if outputType == "pretty" {
		intContent := make([]interface{}, len(*content))
		for idx, row := range *content {
			intContent[idx] = row
		}
		printTable([]string{"Id", "Name", "City", "OcfId", "NfaId", "VerificationStatus"}, intContent)
	} else {
		printJson(content)
	}
	return nil
}

func getCustomer(cmd *cobra.Command, args []string) error {
	content, err := previderClient.Customer.Get(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	if outputType == "pretty" {
		fmt.Printf("%+v\n", content)
	} else {
		printJson(content)
	}
	return nil
}

func createCustomer(cmd *cobra.Command, args []string) error {
	var customerCreate client.CustomerCreate
	var err error

	customerCreate.Name, err = cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	customerCreate.AccountName, err = cmd.Flags().GetString("accountName")
	if err != nil {
		return err
	}
	customerCreate.Address, err = cmd.Flags().GetString("address")
	if err != nil {
		return err
	}
	customerCreate.AddressNumber, err = cmd.Flags().GetString("addressNumber")
	if err != nil {
		return err
	}
	customerCreate.AddressSuffix, err = cmd.Flags().GetString("addressSuffix")
	if err != nil {
		return err
	}
	customerCreate.Zipcode, err = cmd.Flags().GetString("zipCode")
	if err != nil {
		return err
	}
	customerCreate.City, err = cmd.Flags().GetString("city")
	if err != nil {
		return err
	}
	customerCreate.CountryCode, err = cmd.Flags().GetString("countryCode")
	if err != nil {
		return err
	}
	customerCreate.CountryCode, err = cmd.Flags().GetString("language")
	if err != nil {
		return err
	}
	customerCreate.CountryCode, err = cmd.Flags().GetString("purchaseOrderNumber")
	if err != nil {
		return err
	}
	customerCreate.CountryCode, err = cmd.Flags().GetString("cocNumber")
	if err != nil {
		return err
	}
	customerCreate.Partner, err = cmd.Flags().GetBool("partner")
	if err != nil {
		return err
	}
	customerCreate.HidingPrices, err = cmd.Flags().GetBool("hidingPrices")
	if err != nil {
		return err
	}
	customerCreate.InvoiceToPartner, err = cmd.Flags().GetBool("invoiceToPartner")
	if err != nil {
		return err
	}
	createdUser, err := previderClient.Customer.Create(customerCreate)
	if err != nil {
		return err
	}

	if outputType == "pretty" {
		fmt.Printf("%+v\n", createdUser)
	} else {
		printJson(createdUser)
	}
	return nil
}
