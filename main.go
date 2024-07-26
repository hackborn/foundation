package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	orgs, year, err := LoadOrganizations("Donations.csv")
	if err != nil {
		panic(err)
	}
	// Sort because why not
	sort.Slice(orgs, func(i, j int) bool {
		return strings.ToLower(orgs[i].Name) < strings.ToLower(orgs[j].Name)
	})

	err = WriteOrganizations(orgs)
	if err != nil {
		panic(err)
	}
	sum := 0.0
	for _, org := range orgs {
		sum += org.Amount
	}
	fmt.Println("year", year, "sum", sum)
}
