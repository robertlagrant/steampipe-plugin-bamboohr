package bamboohr

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	bamboohr_client "github.com/robertlagrant/bamboohr-client-go"
)

func tableBambooSalary() *plugin.Table {
	return &plugin.Table{
		Name:        "bamboohr_salary",
		Description: "BambooHR only returns employees your credentials have access to.",
		List: &plugin.ListConfig{
			Hydrate: listEmployee,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Employee's BambooHR ID"},
			{Name: "pay_rate", Transform: transform.FromField("PayRate").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's pay rate - raw field from BambooHR"},
			{Name: "salary", Transform: transform.FromField("PayRateParsedSalary").NullIfEqual(""), Type: proto.ColumnType_STRING, Description: "Employee's salary - parsed from pay_rate field or NULL"},
			{Name: "currency", Transform: transform.FromField("PayRateParsedCurrency").NullIfEqual(""), Type: proto.ColumnType_STRING, Description: "Employee's salary currency - parsed from pay_rate field or NULL"},
			// {Name: "frequency", Transform: transform.FromField("PayRate").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's salary currency - parsed from pay_rate field or NULL"},
		},
	}
}

func listSalary(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := makeConfig()
	employeesResponse, err := bamboohr_client.ListEmployees(*config)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve employees. Reason: %s", err.Error())
	}

	for _, employee := range employeesResponse {
		d.StreamListItem(ctx, employee)
	}

	return nil, nil
}
