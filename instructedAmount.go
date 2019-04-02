// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// InstructedAmount is the InstructedAmount of the wire
type InstructedAmount struct {
	// tag
	tag string
	// CurrencyCode
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker
	// (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as
	Amount string `json:"amount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInstructedAmount returns a new InstructedAmount
func NewInstructedAmount() InstructedAmount {
	ia := InstructedAmount{
		tag: TagInstructedAmount,
	}
	return ia
}

// Parse takes the input string and parses the InstructedAmount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ia *InstructedAmount) Parse(record string) {
	ia.tag = record[:6]
	ia.CurrencyCode = ia.parseStringField(record[6:9])
	ia.Amount = ia.parseStringField(record[9:24])
}

// String writes InstructedAmount
func (ia *InstructedAmount) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(24)
	buf.WriteString(ia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on InstructedAmount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ia *InstructedAmount) Validate() error {
	if err := ia.fieldInclusion(); err != nil {
		return err
	}
	// ToDo: Add CurrencyCode lookup
	if err := ia.isAlphanumeric(ia.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, ia.CurrencyCode)
	}
	if err := ia.isAmount(ia.Amount); err != nil {
		return fieldError("Amount", err, ia.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ia *InstructedAmount) fieldInclusion() error {
	return nil
}
