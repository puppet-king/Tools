// Copyright 2024 The wangkai. ALL rights reserved.

/*
Package test
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
	"tools/internal/custom_print"
	"tools/tools"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name     string
		wantType string
	}{
		{"php.Time", "int"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := php.Time(); fmt.Sprint(reflect.TypeOf(got)) == "int" {
				custom_print.Info("打印参数 %v", got)
			} else {
				t.Errorf("Time() ReturnType = %v, want %v", fmt.Sprint(reflect.TypeOf(got)), tt.wantType)
			}
		})

	}
}

func TestCheckDate(t *testing.T) {
	tests := []struct {
		name             string
		Year, Month, Day int
		want             bool
	}{
		{name: "Standard", Month: 1, Day: 31, Year: 2024, want: true},
		{name: "TestLeapYearDate", Month: 2, Day: 29, Year: 2024, want: true},
		{name: "TestErrorDate1", Month: 31, Day: 4, Year: 2024, want: false},
		{name: "TestErrorDate2", Month: 2, Day: 30, Year: 2024, want: false},
		{name: "TestErrorDate3", Month: 13, Day: 1, Year: 2024, want: false},
		{name: "TestErrorDate4", Month: 1, Day: 1, Year: -1, want: false},
		{name: "TestErrorDate5", Month: 1, Day: 100, Year: 2024, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := php.Checkdate(tt.Month, tt.Day, tt.Year); got != tt.want {
				t.Errorf("Time() ReturnType = %v, want %v", got, tt.want)
			}
		})

	}
}
