package serialization

import (
	"fmt"
	"testing"
)

type DemoStruct struct {
	Name string
}

func TestMarshal(t *testing.T) {
	fmt.Println(Marshal(DemoStruct{
		Name: "123",
	}))
	fmt.Println(Marshal(&DemoStruct{
		Name: "123",
	}))
	var arr []DemoStruct
	fmt.Println(Marshal(arr))
	fmt.Println(json.MarshalToString(arr))
}

func TestUnmarshalFromString(t *testing.T) {
	unmarshalString := "[]"
	var arr []DemoStruct
	UnmarshalFromString(unmarshalString, &arr)

}
