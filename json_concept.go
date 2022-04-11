package main

import "fmt"

func main() {
	//{{
	//	'allowProvisioning': True,
	//	'parameterOverrides': {
	//	'DeviceLocation': 'Seattle'
	//}
	//}

	provisioning_result := map[string]bool{
		"allowProvisioning": true,
	}
	fmt.Println(provisioning_result)
}
