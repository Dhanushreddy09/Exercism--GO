package census

import "fmt"

func Census() {
	address := map[string]string{"street": "oxford street"}
	resident := NewResident("Dhanush", 20, address)

	fmt.Println(resident.HasRequiredInfo())
}
