package equinix

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/packethost/packngo"
)

func connectionPortSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the connection port resource",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the connection port resource",
			},
			"role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role - primary or secondary",
			},
			"speed": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port speed in bits per second",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Port status",
			},
			"link_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Port link status",
			},
			"virtual_circuit_ids": {
				Computed:    true,
				Type:        schema.TypeList,
				Elem:        schema.TypeString,
				Description: "List of IDs of virtual circuits attached to this port",
			},
		},
	}
}

func dataSourceMetalConnection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMetalConnectionRead,

		Schema: map[string]*schema.Schema{
			"connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the connection to lookup",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the connection resource",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the connection resource",
			},
			"tags": {
				Type:        schema.TypeList,
				Description: "Tags attached to the connection",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"organization_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of organization to which the connection belongs",
			},
			"project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of project to which the connection belongs",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the connection resource",
			},
			"redundancy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Connection redundancy - reduntant or primary",
			},
			"facility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Slug of a facility to which the connection belongs",
			},
			"metro": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Slug of a metro to which the connection belongs",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fabric Token for the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Connection type - dedicated or shared",
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mode for connections in IBX facilities with the dedicated type - standard or tunnel",
			},
			"speed": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port speed in bits per second",
			},
			"ports": {
				Type:        schema.TypeList,
				Elem:        connectionPortSchema(),
				Computed:    true,
				Description: "List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)",
			},
		},
	}
}

func getConnectionPorts(cps []packngo.ConnectionPort) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0, 1)

	for _, p := range cps {
		vcIDs := []string{}
		for _, vc := range p.VirtualCircuits {
			vcIDs = append(vcIDs, vc.ID)
		}
		connPort := map[string]interface{}{
			"name":                p.Name,
			"id":                  p.ID,
			"role":                string(p.Role),
			"speed":               p.Speed,
			"status":              p.Status,
			"link_status":         p.LinkStatus,
			"virtual_circuit_ids": vcIDs,
		}
		ret = append(ret, connPort)
	}
	return ret
}

func dataSourceMetalConnectionRead(d *schema.ResourceData, meta interface{}) error {
	connId := d.Get("connection_id").(string)
	d.SetId(connId)
	return resourceMetalConnectionRead(d, meta)
}
