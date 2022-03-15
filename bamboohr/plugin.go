package bamboohr

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name:             "steampipe-plugin-bamboohr",
        DefaultTransform: transform.FromGo().NullIfZero(),
        TableMap: map[string]*plugin.Table{
            "bamboohr_employees":        	tableBambooEmployee(),
            // "zendesk_group":       	 		tableZendeskGroup(),
            // "zendesk_organization":	 		tableZendeskOrganization(),
            // "zendesk_search":      	 		tableZendeskSearch(),
            // "zendesk_ticket":      	 		tableZendeskTicket(),
            // "zendesk_ticket_audit":	 		tableZendeskTicketAudit(),
            // "zendesk_trigger":     	 		tableZendeskTrigger(),
            // "zendesk_user":         		tableZendeskUser(),
        },
    }
    return p
}
