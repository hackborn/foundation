package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"

	"github.com/hackborn/foundation/model"
)

func WriteOrganizations(orgs []model.Org) error {
	f, err := os.Create("orgs.txt")
	if err != nil {
		return fmt.Errorf("WriteOrganizations: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(makeOrgString(orgs))
	w.Flush()
	return err
}

func makeOrgString(orgs []model.Org) string {
	var sb strings.Builder
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	for i, org := range orgs {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString(". ")
		sb.WriteString(org.Name)
		sb.WriteString("\n")
		sb.WriteString(ac.FormatMoney(org.Amount))
		sb.WriteString("\n")
		if org.Address != "" {
			sb.WriteString(org.Address)
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
