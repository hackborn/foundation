package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hackborn/foundation/model"
)

func LoadOrganizations(fn string) ([]model.Org, string, error) {
	p := filepath.Join("data", fn)
	f, err := os.Open(p)
	if err != nil {
		return nil, "", fmt.Errorf("LoadOrganizations: %w", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	var orgs []model.Org
	var ri *readIndexes = nil
	for {
		r, err := r.Read()
		if err == io.EOF {
			year := ""
			if ri != nil {
				year = ri.ForYear
			}
			return orgs, year, nil
		}
		if err != nil {
			return nil, "", fmt.Errorf("LoadOrganizations: %w", err)
		}
		if ri == nil {
			ri = makeReadIndexes(r)
		} else {
			if org, ok := makeOrg(r, *ri); ok {
				orgs = append(orgs, org)
			}
		}
	}
}

func makeOrg(r []string, ri readIndexes) (model.Org, bool) {
	var org model.Org
	amt, _ := strings.CutPrefix(r[ri.Amount], "$")
	if amt == "" {
		return org, false
	}
	amt = strings.ReplaceAll(amt, ",", "")
	if f, err := strconv.ParseFloat(amt, 64); err == nil && f > 0.0 {
		org.Amount = f
	} else {
		return org, false
	}

	org.Name = r[ri.Org]
	org.Category = r[ri.Cat]
	org.Address = r[ri.Address]
	if org.Name == "" || org.Name == "subtotal" || org.Name == "total" || org.Name == "target" {
		return org, false
	}
	return org, true
}

// makeReadIndexes converts the title row into indexes for each column
func makeReadIndexes(r []string) *readIndexes {
	var ri readIndexes
	for i, n := range r {
		// Category,Organization,,2023,2022,2021,,Address,Tax ID
		switch n {
		case "Category":
			ri.Cat = i
		case "Organization":
			ri.Org = i
		case "Address":
			ri.Address = i
		default:
			ri.Amount, ri.ForYear = makeReadAmountYear(ri, n, i)
		}
	}
	if ri.isValid() {
		return &ri
	}
	return nil
}

func makeReadAmountYear(ri readIndexes, s string, idx int) (int, string) {
	if len(s) != 4 {
		return ri.Amount, ri.ForYear
	}
	i, err := strconv.Atoi(s)
	if err != nil || i < 2000 || i >= 3000 {
		return ri.Amount, ri.ForYear
	}
	if ri.ForYear == "" {
		return idx, s
	}
	prevI, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if i > prevI {
		return i, s
	}
	return ri.Amount, ri.ForYear
}

type readIndexes struct {
	Org     int
	Cat     int
	Address int
	Amount  int
	ForYear string
}

func (ri readIndexes) isValid() bool {
	m := make(map[int]struct{})
	m[ri.Org] = struct{}{}
	m[ri.Cat] = struct{}{}
	m[ri.Address] = struct{}{}
	m[ri.Amount] = struct{}{}
	return len(m) == 4 && ri.ForYear != ""
}
