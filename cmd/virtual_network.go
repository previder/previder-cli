package cmd

import (
	"fmt"
	"github.com/previder/previder-go-sdk/client"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	var virtualNetworkCmd = &cobra.Command{
		Use:   "virtualnetwork",
		Short: "Virtual network commands",
	}
	rootCmd.AddCommand(virtualNetworkCmd)

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Get a list of virtual networks",
		Args:  cobra.NoArgs,
		RunE:  listVirtualNetwork,
	}
	virtualNetworkCmd.AddCommand(cmdList)

	var cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Get a virtual network",
		Args:  cobra.ExactArgs(1),
		RunE:  getVirtualNetwork,
	}
	virtualNetworkCmd.AddCommand(cmdGet)

	var cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Create a virtual network",
		Args:  cobra.NoArgs,
		RunE:  createVirtualNetwork,
	}

	cmdCreate.Flags().String("name", "", "Name of the virtual server")
	cmdCreate.MarkFlagRequired("name")
	cmdCreate.Flags().String("group", "", "Group")
	cmdCreate.Flags().String("type", "VXLAN", "Type of network [VLAN,IAN,VXLAN]")

	virtualNetworkCmd.AddCommand(cmdCreate)

	var cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete a virtual network",
		Args:  cobra.ExactArgs(1),
		RunE:  deleteVirtualNetwork,
	}
	virtualNetworkCmd.AddCommand(cmdDelete)

}

func listVirtualNetwork(cmd *cobra.Command, args []string) error {
	var page client.PageRequest
	page.Size = 100
	page.Page = 0
	page.Sort = "+name"
	page.Query = ""

	_, content, err := previderClient.VirtualNetwork.Page(page)
	if err != nil {
		fmt.Println(err)
	}

	if outputType == "pretty" {
		intContent := make([]interface{}, len(*content))
		for idx, row := range *content {
			intContent[idx] = row
		}
		printTable([]string{"Id", "Name", "Group", "Type", "State"}, intContent)
	} else {
		printJson(content)
	}

	return nil
}

func getVirtualNetwork(cmd *cobra.Command, args []string) error {
	content, err := previderClient.VirtualNetwork.Get(args[0])
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

func createVirtualNetwork(cmd *cobra.Command, args []string) error {
	var err error
	var vm client.VirtualNetworkUpdate
	vm.Name, err = cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	vm.Group, err = cmd.Flags().GetString("group")
	if err != nil {
		return err
	}

	vm.Type, err = cmd.Flags().GetString("type")
	if err != nil {
		return err
	}

	task, err := previderClient.VirtualNetwork.Create(&vm)
	if err != nil {
		return err
	}

	finishedTask, err := previderClient.Task.WaitFor(task.Id, client.DefaultTimeout)
	if err != nil {
		return err
	}

	if outputType == "pretty" {
		fmt.Printf("%+v\n", finishedTask)
	} else {
		printJson(finishedTask)
	}
	return nil
}

func deleteVirtualNetwork(cmd *cobra.Command, args []string) error {
	_, err := previderClient.VirtualNetwork.Delete(args[0])
	if err != nil {
		return err
	}

	return nil
}
