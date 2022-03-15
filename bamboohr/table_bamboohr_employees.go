package bamboohr

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/plugin"

    "github.com/robertlagrant/bamboohr-client-go"
)

func tableBambooEmployee() *plugin.Table {
    return &plugin.Table {
        Name:        "bamboohr_employee",
        Description: "BambooHR only returns employees your credentials have access to.",
        List: &plugin.ListConfig {
            Hydrate: listEmployee,
        },
        Get: &plugin.GetConfig {
            KeyColumns: plugin.SingleColumn("id"),
            Hydrate:    getUser,
        },
        Columns: []*plugin.Column {
            { Name: "displayName", Type: proto.ColumnType_STRING, Description: "Employee's full name for display" },
        },
    }
}

func listEmployee(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    fmt.Println(bamboohr_client.ListEmployees())

    opts := &zendesk.UserListOptions{
        PageOptions: zendesk.PageOptions{
            Page:    1,
            PerPage: 100,
        },
    }
    for true {
        users, page, err := conn.GetUsers(ctx, opts)
        if err != nil {
            return nil, err
        }
        for _, t := range users {
            d.StreamListItem(ctx, t)
        }
        if !page.HasNext() {
            break
        }
        opts.Page++
    }
    return nil, nil
}
