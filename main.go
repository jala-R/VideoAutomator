package main

import (
	"encoding/xml"
	"fmt"
	"os"

	types "jalar.me/xml/Types"
)

const XMLHEADER = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE xmeml>
`

func main() {
	a := types.Xmeml{}
	// var b any = a

	// var b int = 10
	// fmt.Println(reflect.ValueOf(b).Kind() == reflect.Ptr)
	// fmt.Println(reflect.ValueOf(&b).Kind())
	types.Constructor(&a)
	temp := &types.Label{
		Label2: "Forest",
	}
	types.Constructor(temp)

	a.Sequence.Label = *temp

	xmlData, _ := xml.MarshalIndent(a, "", " ")

	file, err := os.Create("./output1.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file.Write([]byte(XMLHEADER))
	file.Write(xmlData)
	file.Close()
}
