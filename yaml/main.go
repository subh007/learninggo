package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	buf, err := os.ReadFile("sample/config1.yaml")
	if err != nil {
		fmt.Printf("err:%+v", err)
		os.Exit(1)
	}

	obj := make(map[string]interface{})
	err = yaml.Unmarshal(buf, obj)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		os.Exit(1)
	}

	for key, val := range obj {
		fmt.Printf("->obj[%v]:%v\n", key, val)
	}

	buf, err = yaml.Marshal(obj)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		os.Exit(1)
	}

	yamlOutput := string(buf)
	fmt.Println(yamlOutput)

	// list all the clusters
	var clusters []interface{}
	var ok bool

	if clusters, ok = obj["clusters"].([]interface{}); !ok {
		fmt.Println("type assertion failed")
		return
	}

	for index, val := range clusters {
		fmt.Printf("index:%d: val:%+v", index, val)
	}

	out, err := yaml.Marshal(clusters)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Println()
	fmt.Println(string(out))

	f, err := os.Create("sample/tmp.yaml")
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(f)
	writer.WriteString(string(out))
	writer.Flush()

	var cluster2 []interface{}
	err = yaml.Unmarshal(out, &cluster2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cluster2...)

}
