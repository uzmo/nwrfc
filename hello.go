package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/sap/gorfc/gorfc"
)

func abapSystem() gorfc.ConnectionParameter {
	//SXX
	return gorfc.ConnectionParameter{
		Dest:   "SXX",
		Client: "100",
		User:   "RFC_USER",
		Passwd: "123456",
		Lang:   "EN",
		Ashost: "192.168.2.140",
		Sysnr:  "00",
	}
}

func main() {
	//connect SAP System.
	c, err := gorfc.ConnectionFromParams(abapSystem())
	defer c.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//input params.
	params := map[string]interface{}{
		"WEEK": "201847",
	}
	//call function.
	r, err := c.Call("WEEK_GET_FIRST_DAY", params)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	} else {
		fmt.Println(reflect.TypeOf(r["DATE"]))
		fmt.Println(r["DATE"])
	}

}
