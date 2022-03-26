package bamboohr

import (
	"context"

	"github.com/hashicorp/go-hclog"
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Employee's BambooHR ID"},
			{Name: "firstName", Type: proto.ColumnType_STRING, Description: "Employee's first name"},
			{Name: "lastName", Type: proto.ColumnType_STRING, Description: "Employee's last name"},
			{Name: "displayName", Type: proto.ColumnType_STRING, Description: "Employee's full name for display"},
			{Name: "jobTitle", Type: proto.ColumnType_STRING, Description: "Employee's job title"},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "Employee's department"},
			{Name: "division", Type: proto.ColumnType_STRING, Description: "Employee's division"},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Employee's work location"},
			{Name: "supervisor", Type: proto.ColumnType_STRING, Description: "Employee's supervisor"},
			{Name: "payRate", Type: proto.ColumnType_STRING, Description: "Employee's pay rate"},
			{Name: "terminationDate", Type: proto.ColumnType_STRING, Description: "Employee's termination date"},
			{Name: "employeeNumber", Type: proto.ColumnType_STRING, Description: "Employee's employee number"},
			{Name: "hireDate", Type: proto.ColumnType_STRING, Description: "Employee's hire date"},
			{Name: "originalHireDate", Type: proto.ColumnType_STRING, Description: "Employee's original hire date"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Employee's record status"},
			{Name: "supervisorId", Type: proto.ColumnType_STRING, Description: "Employee's supervisor's BambooHR ID"},
			{Name: "supervisorEId", Type: proto.ColumnType_STRING, Description: "Employee's supervisor's employee number"},
		},
	}
}

func listEmployee(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	employeesResponse, _ := bamboohr_client.ListEmployees()
	// logger.Warn("Log message and a variable", employeesResponse)
	logger.Warn("Log message and a variable", "employeeResponse", hclog.Fmt("%#v", employeesResponse))

	for _, employee := range employeesResponse {
		logger.Info("Log message and a variable", employee)

		d.StreamListItem(ctx, employee)
	}
	return nil, nil
}
