package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func SetTable() *tablewriter.Table {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SERVICE", "DETAIL", "RESOURCE", "ID,ARN", "IP"})

	return table
}
