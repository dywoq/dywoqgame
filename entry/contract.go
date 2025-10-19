package entry

import (
	"fmt"
	"reflect"

	"github.com/dywoq/dywoqgame/resource"
)

type contract struct {
	kind   resource.Kind
	fields []contractField
}

type contractField struct {
	name  string
	tType string
}

var (
	contracts []*contract
)

func contractsInit() {
	contracts = nil

	// sprite contracts initialization
	spriteFields := []contractField{
		{"name", "string"},
		{"path", "string"},
		{"width", "int"},
		{"height", "int"},
		{"raw", "*ebiten.Image"},
	}
	contractSprite := newContract(resource.Sprite, spriteFields)
	contracts = append(contracts, contractSprite)

	// object contracts initialization
	objectFields := []contractField{
		{"name", "string"},
		{"has_sprite", "bool"},
		{"sprite", "map[string]interface{}"},
	}
	objectContract := newContract(resource.Object, objectFields)
	contracts = append(contracts, objectContract)

	if contractHasDoubles() {
		panic("entry: contracts can't share the same kind")
	}
}

func newContract(kind resource.Kind, field []contractField) *contract {
	return &contract{kind, field}
}

func contractHasDoubles() bool {
	kinds := make(map[resource.Kind]bool)
	for _, c := range contracts {
		if kinds[c.kind] {
			return true
		}
		kinds[c.kind] = true
	}
	return false
}

func contractCheck(r resource.Resource) error {
	kind := r.Kind()
	for _, contract := range contracts {
		if contract.kind == kind {
			t := r.Table()
			for key, value := range t {
				for _, field := range contract.fields {
					if field.name == key {
						got := reflect.TypeOf(value).String()
						if got != field.tType {
							return fmt.Errorf("resource")
						}
					}
				}
			}
			return nil
		}
	}
	return fmt.Errorf("entry: didn't find matching contract for the resource with \"%v\"", r.Kind)
}
