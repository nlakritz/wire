// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// InputMessageAccountabilityData (IMAD) {1520}
type InputMessageAccountabilityData struct {
	// tag
	tag string
	// InputCycleDate CCYYMMDD
	InputCycleDate string `json:"inputCycleDate"`
	// InputSource
	InputSource string `json:"inputSource"`
	// InputSequenceNumber
	InputSequenceNumber string `json:"inputSequenceNumber"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewInputMessageAccountabilityData returns a new InputMessageAccountabilityData
func NewInputMessageAccountabilityData() InputMessageAccountabilityData {
	imad := InputMessageAccountabilityData{
		tag: TagInputMessageAccountabilityData,
	}
	return imad
}

// Parse takes the input string and parses the InputMessageAccountabilityData values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (imad *InputMessageAccountabilityData) Parse(record string) {
	imad.tag = record[:6]
	imad.InputCycleDate = imad.validateDate(record[6:14])
	imad.InputSource = imad.parseStringField(record[14:22])
	imad.InputSequenceNumber = imad.parseStringField(record[22:28])
}

// String writes InputMessageAccountabilityData
func (imad *InputMessageAccountabilityData) String() string {
	var buf strings.Builder
	buf.Grow(22)
	buf.WriteString(imad.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on InputMessageAccountabilityData and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (imad *InputMessageAccountabilityData) Validate() error {
	if err := imad.fieldInclusion(); err != nil {
		return err
	}
	if err := imad.isAlphanumeric(imad.InputSource); err != nil {
		return fieldError("InputSource", err, imad.InputSource)
	}
	if err := imad.isAlphanumeric(imad.InputSequenceNumber); err != nil {
		return fieldError("InputSequenceNumber", err, imad.InputSequenceNumber)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (imad *InputMessageAccountabilityData) fieldInclusion() error {
	if imad.InputCycleDate == "" {
		return fieldError("InputCycleDate", ErrFieldRequired, imad.InputCycleDate)
	}
	if imad.InputSource == "" {
		return fieldError("InputSource", ErrFieldRequired, imad.InputSource)
	}
	if imad.InputSequenceNumber == "" {
		return fieldError("InputSequenceNumber", ErrFieldRequired, imad.InputSequenceNumber)
	}
	return nil
}
