package bamboohr

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"

	bamboohr_client "github.com/robertlagrant/bamboohr-client-go"
)

func tableBambooEmployee() *plugin.Table {
	return &plugin.Table{
		Name:        "bamboohr_employee",
		Description: "BambooHR only returns employees your credentials have access to.",
		List: &plugin.ListConfig{
			Hydrate: listEmployee,
		},
		// Get: &plugin.GetConfig {
		//     KeyColumns: plugin.SingleColumn("id"),
		//     Hydrate:    getUser,
		// },
		Columns: []*plugin.Column{
			{Name: "firstName", Type: proto.ColumnType_STRING, Description: "Employee's first name"},
			{Name: "lastName", Type: proto.ColumnType_STRING, Description: "Employee's last name"},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: "Employee's full name for display"},
			{Name: "jobTitle", Type: proto.ColumnType_STRING, Description: "Employee's job title"},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "Employee's department"},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Employee's work location"},
			{Name: "supervisor", Type: proto.ColumnType_STRING, Description: "Employee's supervisor"},
		},
	}
}

func listEmployee(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	employees, _ := bamboohr_client.ListEmployees()
	for _, employee := range employees {

		d.StreamListItem(ctx, employee)
	}

	return nil, nil
}
