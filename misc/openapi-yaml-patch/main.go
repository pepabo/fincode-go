package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k1LoW/yrep"
)

const openAPIPath = "spec/fincode-openapi.yml"

func main() {
	if err := _main(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func _main() error {
	b, err := os.ReadFile(openAPIPath)
	if err != nil {
		return err
	}
	funcs := []yrep.ReplaceFunc{
		renameToQuery,
		fixNumberValidation,
		removeUnnecessaryAllOf,
		fixInvalidPayType,
		fixInvalidAllOf,
		appendRequiredToRequestBody,
	}

	applied, err := yrep.Apply(b, funcs...)
	if err != nil {
		return err
	}

	if err := os.WriteFile(openAPIPath, applied, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// renameToQuery rename `name: クエリ` to `name: Query`
func renameToQuery(in any) (any, error) {
	item, ok := in.(yaml.MapItem)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	if item.Key == "name" && item.Value == "クエリ" {
		item.Value = "Query"
		return item, nil
	}
	return in, yrep.ErrNotReplaced
}

// fixNumberValidation fix number/integer validation
func fixNumberValidation(in any) (any, error) {
	v, ok := in.(yaml.MapSlice)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	number := false
	minLength := false
	maxLength := false
	for _, item := range v {
		if item.Key == "type" && (item.Value == "number" || item.Value == "integer") {
			number = true
		}
		if item.Key == "minLength" {
			minLength = true
		}
		if item.Key == "maxLength" {
			maxLength = true
		}
	}
	if !number || (!minLength && !maxLength) {
		return in, yrep.ErrNotReplaced
	}
	for i, item := range v {
		if item.Key == "minLength" {
			v[i].Key = "minimum"
		}
		if item.Key == "maxLength" {
			v[i].Key = "maximum"
		}
	}

	return v, nil
}

// removeUnnecessaryAllOf remove unnecessary allOf
func removeUnnecessaryAllOf(in any) (any, error) {
	v, ok := in.(yaml.MapSlice)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	for _, item := range v {
		if item.Key == "allOf" {
			vv, ok := item.Value.([]any)
			if !ok {
				return nil, fmt.Errorf("unexpected type %v: %T", item.Value, item.Value)
			}
			if len(vv) != 1 {
				continue
			}
			return vv[0], nil
		}
	}
	return in, yrep.ErrNotReplaced
}

// fixInvalidPayType fix invalid pay_type
func fixInvalidPayType(in any) (any, error) {
	v, ok := in.(yaml.MapSlice)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	found := false
	for i, item := range v {
		if item.Key == "$ref" && strings.HasSuffix(item.Value.(string), "/pay_type") {
			item.Value = strings.ReplaceAll(item.Value.(string), "/pay_type", "/PayType")
			v[i] = item
			found = true
		}
	}
	if !found {
		return in, yrep.ErrNotReplaced
	}
	return v, nil
}

// fixInvalidAllOf fix invalid allOf
//
// allOf:
// - $ref: "#/components/schemas/CardBrand"
// - nullable: true
//
// to
//
// $ref: "#/components/schemas/CardBrand"
// nullable: true

func fixInvalidAllOf(in any) (any, error) {
	v, ok := in.(yaml.MapSlice)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	var replaced yaml.MapSlice
	for _, item := range v {
		if item.Key != "allOf" {
			replaced = append(replaced, item)
			continue
		}
		vv, ok := item.Value.([]any)
		if !ok {
			return nil, fmt.Errorf("unexpected type %v: %T", item.Value, item.Value)
		}
		oneRef := true
		typeCount := 0
		refCount := 0
		for _, vvv := range vv {
			vvvv, ok := vvv.(yaml.MapSlice)
			if !ok {
				return nil, fmt.Errorf("unexpected type %v: %T", vvv, vvv)
			}
			for _, item := range vvvv {
				if item.Key == "$ref" {
					refCount++
					break
				}
				if item.Key == "type" {
					typeCount++
					break
				}
			}
		}
		if refCount != 1 {
			oneRef = false
		}
		if !oneRef || len(vv) == (typeCount+refCount) {
			replaced = append(replaced, item)
			continue
		}
		for _, vvv := range vv {
			vvvv, ok := vvv.(yaml.MapSlice)
			if !ok {
				return nil, fmt.Errorf("unexpected type %v: %T", vvv, vvv)
			}
			for _, item := range vvvv {
				replaced = append(replaced, item)
			}
		}
	}

	return replaced, nil
}

func appendRequiredToRequestBody(in any) (any, error) {
	item, ok := in.(yaml.MapItem)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	if item.Key != "requestBody" {
		return in, yrep.ErrNotReplaced
	}
	v, ok := item.Value.(yaml.MapSlice)
	if !ok {
		return in, yrep.ErrNotReplaced
	}
	hasRequired := false
	for _, item := range v {
		if item.Key == "required" {
			hasRequired = true
			break
		}
	}
	if hasRequired {
		return in, yrep.ErrNotReplaced
	}
	v = append(v, yaml.MapItem{
		Key:   "required",
		Value: true,
	})
	item.Value = v
	return item, nil
}
