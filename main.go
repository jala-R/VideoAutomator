package main

import (
	"encoding/xml"
	"fmt"

	types "jalar.me/xml/Types"
)

func main() {

	temp := types.Samplecharacteristics{}

	fmt.Println("starting")
	var temp = Person{}

	temp.Name = []Name{
		{
			FirstName: "1.1",
			LastName:  "1.2",
		},
		{
			FirstName: "2.1",
			LastName:  "2.2",
		},
	}

	res, err := xml.MarshalIndent(temp, "", "  ")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", (res))
}
