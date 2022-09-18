package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func SetTable() *tablewriter.Table {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Service", "DETAIL", "RESOURCE", "IP"})

	return table
}
