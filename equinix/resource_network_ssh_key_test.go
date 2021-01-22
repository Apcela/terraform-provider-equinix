package equinix

import (
	"testing"

	"github.com/equinix/ne-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestNetworkSSHKey_createFromResourceData(t *testing.T) {
	//given
	expected := ne.SSHPublicKey{
		Name:  "testKey",
		Value: "testKeyValue",
	}
	rawData := map[string]interface{}{
		networkSSHKeySchemaNames["Name"]:  expected.Name,
		networkSSHKeySchemaNames["Value"]: expected.Value,
	}
	d := schema.TestResourceDataRaw(t, createNetworkSSHKeyResourceSchema(), rawData)
	//when
	key := createNetworkSSHKey(d)
	//then
	assert.Equal(t, expected, key, "Created key matches expected result")
}

func TestNetworkSSHKey_updateResourceData(t *testing.T) {
	//given
	input := &ne.SSHPublicKey{
		UUID:  "059c3020-aec5-44ca-816c-235435f16df9",
		Name:  "testKey",
		Value: "testKeyValue",
	}
	d := schema.TestResourceDataRaw(t, createNetworkSSHKeyResourceSchema(), make(map[string]interface{}))
	//when
	err := updateNetworkSSHKeyResource(input, d)
	//then
	assert.Nil(t, err, "Update of resource data does not return error")
	assert.Equal(t, input.UUID, d.Get(networkSSHKeySchemaNames["UUID"]), "UUID matches")
	assert.Equal(t, input.Name, d.Get(networkSSHKeySchemaNames["Name"]), "Name matches")
	assert.Equal(t, input.Value, d.Get(networkSSHKeySchemaNames["Value"]), "Value matches")
}
