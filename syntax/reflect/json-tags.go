package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

type VMFilterSpec struct {
	Filter_folders        []string `json:"filter.folders,omitempty"`
	Filter_power_states   []string `json:"filter.power_states,omitempty"`
	Filter_hosts          []string `json:"filter.hosts,omitempty"`
	Filter_names          []string `json:"filter.names,omitempty"`
	Filter_clusters       []string `json:"filter.clusters,omitempty"`
	Filter_datacenters    []string `json:"filter.datacenters,omitempty"`
	Filter_resource_pools []string `json:"filter.resource_pools,omitempty"`
	Filter_vms            []string `json:"filter.vms,omitempty"`
}

type FolderFilterSpec struct {
	Filter_folders        []string `json:"filter.folders,omitempty"`
	Filter_parent_folders []string `json:"filter.parent_folders,omitempty"`
	Filter_types          []string `json:"filter.types,omitempty"`
	Filter_names          []string `json:"filter.names,omitempty"`
	Filter_datacenters    []string `json:"filter.datacenters,omitempty"`
}

/*
func Spec2Params(spec interface{}) (string, error) {
	val := reflect.ValueOf(spec).Elem()

	params := []string{}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		pname := typeField.Tag.Get("json")

		for _, n := range valueField.Interface().([]string) {
			params = append(params, fmt.Sprintf("%s=%v", pname, n))
		}
	}

	return strings.Join(params, "&"), nil
}
*/

func Spec2Params(spec interface{}) (string, error) {

	t := reflect.TypeOf(spec)
	v := reflect.ValueOf(spec)

	params := []string{}

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)
		jsontag := typeField.Tag.Get("json")
		tags := strings.Split(jsontag, ",")
		pname := tags[0]

		for _, n := range valueField.Interface().([]string) {
			params = append(params, fmt.Sprintf("%s=%v", pname, n))
		}
	}

	return strings.Join(params, "&"), nil
}

func Si2pec2Params(spec interface{}) (string, error) {

	fmt.Println(reflect.TypeOf(spec))
	fmt.Println(reflect.ValueOf(spec))
	v := reflect.ValueOf(spec)
	t := reflect.TypeOf(spec)
	for n := 0; n < v.NumField(); n++ {
		vi := v.Field(n)
		ti := t.Field(n)
		fmt.Println(vi)
		fmt.Println(vi.Type())
		fmt.Printf("tag: %s\n", ti.Tag.Get("json"))
		name := ti.Tag.Get("json")
		arr := strings.Split(name, ",")
		fmt.Printf("name: %s\n", arr[0])
		//field := reflect.TypeOf(spec).Elem().Field(n)
		//fmt.Printf("tag1: %s\n", field.Tag)
	}
	return "string", nil
}

func main1() {

	var folder VMFilterSpec
	folder.Filter_folders = []string{"1", "2", "3"}
	folder.Filter_power_states = []string{"on", "off"}
	folder.Filter_names = []string{"zongzw", "andrew_zong"}

	params, err := Spec2Params(folder)

	if err == nil {
		fmt.Printf("%v\n", params)
	}

	fmt.Printf("sepc: %v", folder)
	bytes, err := json.Marshal(folder)
	if err != nil {
		fmt.Printf("failed to bytes")
		return
	}

	fmt.Printf("json: %s\n", string(bytes))
	return

	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()

	var spec VMFilterSpec
	spec.Filter_folders = append(spec.Filter_folders, "a")
	spec.Filter_clusters = append(spec.Filter_clusters, "b")
	spec.Filter_clusters = append(spec.Filter_clusters, "c")
	val := reflect.ValueOf(&spec).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		fmt.Printf("%v: %v\n", typeField.Name, valueField.Len())
		fmt.Printf("value of valueField: %s\n", valueField.Interface().([]string))

	}

}
