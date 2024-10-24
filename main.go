package main

import (
	"encoding/xml"
	"fmt"

	types "jalar.me/xml/Types"
)

func main() {

	temp := types.Samplecharacteristics{}

	types.Constructor(temp)

	res, err := xml.MarshalIndent(temp, "", "  ")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", (res))
}
