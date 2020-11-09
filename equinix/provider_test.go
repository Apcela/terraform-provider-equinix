package equinix

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"equinix": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv(endpointEnvVar); v == "" {
		t.Fatalf("%s env variable has to be set", endpointEnvVar)
	}
	if v := os.Getenv(clientIDEnvVar); v == "" {
		t.Fatalf("%s env variable has to be set", clientIDEnvVar)
	}
	if v := os.Getenv(clientSecretEnvVar); v == "" {
		t.Fatalf("%s env variable has to be set", clientSecretEnvVar)
	}
}

func nprintf(format string, params map[string]interface{}) string {
	for key, val := range params {
		var strVal string
		switch val.(type) {
		case []string:
			r := regexp.MustCompile(`" "`)
			strVal = r.ReplaceAllString(fmt.Sprintf("%q", val), `", "`)
		default:
			strVal = fmt.Sprintf("%v", val)
		}
		format = strings.Replace(format, "%{"+key+"}", strVal, -1)
	}
	return format
}

func sourceMatchesTargetSchema(t *testing.T, source interface{}, sourceFields []string, target interface{}, targetFields map[string]string) {
	val := reflect.ValueOf(source)
	for _, fName := range sourceFields {
		val := val.FieldByName(fName)
		assert.NotEmptyf(t, val, "Value of a field %v not found", fName)
		var schemaValue interface{}
		switch target.(type) {
		case *schema.ResourceData:
			schemaValue = target.(*schema.ResourceData).Get(targetFields[fName])
		case map[string]interface{}:
			schemaValue = target.(map[string]interface{})[targetFields[fName]]
		default:
			assert.Fail(t, "Target type not supported")
		}
		switch val.Kind() {
		case reflect.String, reflect.Int, reflect.Bool, reflect.Float64:
			assert.Equal(t, val.Interface(), schemaValue, fName+" matches")
		case reflect.Slice:
			assert.ElementsMatch(t, val.Interface().([]string), schemaValue.(*schema.Set).List(), fName+" matches")
		default:
			assert.Failf(t, "Type of field not supported: field %v, type %v", fName, val.Kind())
		}
	}
}

func structToSchemaMap(src interface{}, schema map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	val := reflect.ValueOf(src)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		schemaName, ok := schema[typ.Field(i).Name]
		if !ok {
			continue
		}
		ret[schemaName] = val.Field(i).Interface()
	}
	return ret
}

func randInt(n int) int {
	src := rand.NewSource(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	i := rand.New(src).Intn(n)
	mu.Unlock()
	return i
}

func randString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	result := make([]byte, length)
	set := "abcdefghijklmnopqrstuvwxyz012346789"
	var mu sync.Mutex
	mu.Lock()
	r := rand.New(src)
	for i := 0; i < length; i++ {
		result[i] = set[r.Intn(len(set))]
	}
	mu.Unlock()
	return string(result)
}

func getFromEnv(varName string) (string, error) {
	if v := os.Getenv(varName); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("environmental variable '%s' is not set", varName)
}
