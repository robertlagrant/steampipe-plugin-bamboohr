package bamboohr

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	bamboohr_client "github.com/robertlagrant/bamboohr-client-go"
)

func tableBambooEmployee() *plugin.Table {
	return &plugin.Table{
		Name:        "bamboohr_employee",
		Description: "BambooHR only returns employees your credentials have access to.",
		List: &plugin.ListConfig{
			Hydrate: listEmployee,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Employee's BambooHR ID"},
			{Name: "first_name", Transform: transform.FromField("FirstName"), Type: proto.ColumnType_STRING, Description: "Employee's first name"},
			{Name: "last_name", Transform: transform.FromField("LastName"), Type: proto.ColumnType_STRING, Description: "Employee's last name"},
			{Name: "display_name", Transform: transform.FromField("DisplayName"), Type: proto.ColumnType_STRING, Description: "Employee's full name for display"},
			{Name: "job_title", Transform: transform.FromField("JobTitle"), Type: proto.ColumnType_STRING, Description: "Employee's job title"},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "Employee's department"},
			{Name: "division", Type: proto.ColumnType_STRING, Description: "Employee's division"},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Employee's work location"},
			{Name: "supervisor", Type: proto.ColumnType_STRING, Description: "Employee's supervisor"},
			{Name: "pay_rate", Transform: transform.FromField("PayRate").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's pay rate"},
			{Name: "termination_date", Transform: transform.FromField("TerminationDate").NullIfZero().NullIfEqual("0000-00-00"), Type: proto.ColumnType_STRING, Description: "Employee's termination date"},
			{Name: "employee_number", Transform: transform.FromField("EmployeeNumber").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's employee number"},
			{Name: "hire_date", Transform: transform.FromField("HireDate").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's hire date"},
			{Name: "original_hire_date", Transform: transform.FromField("OriginalHireDate").NullIfZero().NullIfEqual("0000-00-00"), Type: proto.ColumnType_STRING, Description: "Employee's original hire date"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Employee's record status"},
			{Name: "supervisor_id", Transform: transform.FromField("SupervisorID").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's supervisor's BambooHR ID"},
			{Name: "supervisor_employee_number", Transform: transform.FromField("SupervisorEID").NullIfZero(), Type: proto.ColumnType_STRING, Description: "Employee's supervisor's employee number"},
		},
	}
}

func listEmployee(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)
	employeesResponse, err := bamboohr_client.ListEmployees()
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve employees. Reason: %s", err.Error())
	}

	for _, employee := range employeesResponse {
		d.StreamListItem(ctx, employee)
	}

	return nil, nil
}
