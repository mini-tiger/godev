// Copyright 2016 - 2019 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX files. Support reads and writes XLSX file generated by
// Microsoft Excel™ 2007 and later. Support save file without losing original
// charts of XLSX. This library needs Go version 1.10 or later.

package excelize

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataValidation(t *testing.T) {
	resultFile := filepath.Join("test", "TestDataValidation.xlsx")

	f := NewFile()

	dvRange := NewDataValidation(true)
	dvRange.Sqref = "A1:B2"
	dvRange.SetRange(10, 20, DataValidationTypeWhole, DataValidationOperatorBetween)
	dvRange.SetError(DataValidationErrorStyleStop, "error title", "error body")
	dvRange.SetError(DataValidationErrorStyleWarning, "error title", "error body")
	dvRange.SetError(DataValidationErrorStyleInformation, "error title", "error body")
	f.AddDataValidation("Sheet1", dvRange)
	if !assert.NoError(t, f.SaveAs(resultFile)) {
		t.FailNow()
	}

	dvRange = NewDataValidation(true)
	dvRange.Sqref = "A3:B4"
	dvRange.SetRange(10, 20, DataValidationTypeWhole, DataValidationOperatorGreaterThan)
	dvRange.SetInput("input title", "input body")
	f.AddDataValidation("Sheet1", dvRange)
	if !assert.NoError(t, f.SaveAs(resultFile)) {
		t.FailNow()
	}

	dvRange = NewDataValidation(true)
	dvRange.Sqref = "A5:B6"
	dvRange.SetDropList([]string{"1", "2", "3"})
	f.AddDataValidation("Sheet1", dvRange)
	if !assert.NoError(t, f.SaveAs(resultFile)) {
		t.FailNow()
	}
}

func TestDataValidationError(t *testing.T) {
	resultFile := filepath.Join("test", "TestDataValidationError.xlsx")

	f := NewFile()
	f.SetCellStr("Sheet1", "E1", "E1")
	f.SetCellStr("Sheet1", "E2", "E2")
	f.SetCellStr("Sheet1", "E3", "E3")

	dvRange := NewDataValidation(true)
	dvRange.SetSqref("A7:B8")
	dvRange.SetSqref("A7:B8")
	dvRange.SetSqrefDropList("$E$1:$E$3", true)

	err := dvRange.SetSqrefDropList("$E$1:$E$3", false)
	assert.EqualError(t, err, "cross-sheet sqref cell are not supported")

	f.AddDataValidation("Sheet1", dvRange)
	if !assert.NoError(t, f.SaveAs(resultFile)) {
		t.FailNow()
	}

	dvRange = NewDataValidation(true)
	err = dvRange.SetDropList(make([]string, 258))
	if dvRange.Formula1 != "" {
		t.Errorf("data validation error. Formula1 must be empty!")
		return
	}
	assert.EqualError(t, err, "data validation must be 0-255 characters")
	dvRange.SetRange(10, 20, DataValidationTypeWhole, DataValidationOperatorGreaterThan)
	dvRange.SetSqref("A9:B10")

	f.AddDataValidation("Sheet1", dvRange)
	if !assert.NoError(t, f.SaveAs(resultFile)) {
		t.FailNow()
	}
}
