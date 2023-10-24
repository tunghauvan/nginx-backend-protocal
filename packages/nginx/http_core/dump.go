package httpcore

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unicode"

	"github.com/tufanbarisyildirim/gonginx"
)

// Translate location context to nginx config using gonginx
func (c *LocationContext) ToNginx() string {
	// Init Empty Directives
	directives := &Directives{}

	// Add http_access to directives
	directives.AddHttpAccessContext(c.HttpAccessContext)

	// Add CoreProps to directives
	directives.AddCoreProps(reflect.ValueOf(c.CoreProps))

	// Add error_page to directives
	directives.AddErrorPageContext(c.ErrorPageContext)

	location_directive := gonginx.Directive{
		Name:       "location",
		Parameters: c.Location,
		Block: &gonginx.Block{
			Directives: directives.Directives,
		},
	}

	return gonginx.DumpDirective(&location_directive, gonginx.IndentedStyle)
}

// Dump server context to nginx config
func (c *ServerContext) ToNginx() string {
	// Init Empty Directives
	directives := &Directives{}

	// Add server_name directive
	if len(c.ServerNames) > 0 {
		server_name_directive := gonginx.Directive{
			Name:       "server_name",
			Parameters: c.ServerNames,
		}

		directives.AddDirective(&server_name_directive)
	}

	// Add listen directive multiple times
	if len(c.Listens) > 0 {
		for _, listen := range c.Listens {
			listen_directive := gonginx.Directive{
				Name:       "listen",
				Parameters: []string{listen},
			}

			directives.AddDirective(&listen_directive)
		}
	}

	// Add http_access to directives
	directives.AddHttpAccessContext(c.HttpAccessContext)

	// Add CoreProps to directives
	directives.AddCoreProps(reflect.ValueOf(c.CoreProps))

	// Add error_page to directives
	directives.AddErrorPageContext(c.ErrorPageContext)

	server_directive := gonginx.Directive{
		Name: "server",
		Block: &gonginx.Block{
			Directives: directives.Directives,
		},
	}

	return gonginx.DumpDirective(&server_directive, gonginx.IndentedStyle)
}

func intSliceToString(intSlice []int) string {
	var result string
	for _, i := range intSlice {
		result += fmt.Sprintf("%d ", i)
	}
	return result
}

func toLowerSnakeCase(s string) string {
	var result string
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i > 0 {
				result += "_"
			}
			result += string(unicode.ToLower(c))
		} else {
			result += string(c)
		}
	}
	return result
}

// Dump location context to json string
func (c *LocationContext) Dump() string {
	// Convert the LocationContext object to a JSON string
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Print the resulting JSON string
	fmt.Println(string(jsonString))
	return string(jsonString)
}