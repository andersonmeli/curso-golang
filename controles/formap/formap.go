package main

import "fmt"

func main() {
	wallet := map[string]bool{"and": true, "jo": false}
	walletProfiles := []string{"and", "jo"}
	for _, walletProfile := range walletProfiles {
		if value, found := wallet[walletProfile]; found {
			fmt.Println(walletProfile)
			fmt.Println(value)
		}
	}
}
