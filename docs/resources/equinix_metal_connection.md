---
subcategory: "Metal"
---

# equinix_metal_connection (Resource)

Use this resource to request of create an Interconnection from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

## Example Usage

```hcl
resource "equinix_metal_connection" "test" {
    name            = "My Interconnection"
    organization_id = local.my_organization_id
    project_id      = local.my_project_id
    metro           = "sv"
    redundancy      = "redundant"
    type            = "shared"
}
```

## Argument Reference

* `name` - (Required) Name of the connection resource
* `organization_id` - (Required) ID of the organization responsible for the connection
* `redundancy` - (Required) Connection redundancy - redundant or primary
* `type` - (Required) Connection type - dedicated or shared
* `mode` - Mode for connections in IBX facilities with the dedicated type - standard or tunnel
* `project_id` - (Optional) ID of the project where the connection is scoped to, must be set for shared connection
* `metro` - (Optional) Metro where the connection will be created
* `facility` - (Optional) Facility where the connection will be created
* `description` - (Optional) Description for the connection resource
* `tags` - (Optional) String list of tags

## Attributes Reference

* `status` - Status of the connection resource
* `token` - Fabric Token from the [Equinix Fabric Portal](https://fabric.equinix.com/dashboard)
* `speed` - Port speed in bits per second
* `ports` - List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the [equinix_metal_connection datasource](../data-sources/equinix_metal_connection.md).
 