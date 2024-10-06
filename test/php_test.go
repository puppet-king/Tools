// Copyright 2024 The wangkai. ALL rights reserved.

/*
Package test
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"tools/internal/custom_print"
	"tools/php"
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

func TestDate(t *testing.T) {
	timestamp1 := int64(1728157288)

	tests := []struct {
		name      string
		format    string
		timestamp *int64
		want      string
	}{
		{name: "Standard", format: "2006/01/02", timestamp: &timestamp1, want: "2024/10/06"},
		{name: "TestError", format: "aaa", timestamp: &timestamp1, want: "aaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// php2go.Date();
			if got := php.Date(tt.format, tt.timestamp); got != tt.want {
				t.Errorf("Time() ReturnType = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestSleep(t *testing.T) {
	tests := []struct {
		name  string
		t, t2 int64
		want  int64
	}{
		{name: "Standard", t: int64(1), t2: int64(1000000), want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			php.Sleep(tt.t)
			php.Usleep(tt.t2)

			diff := time.Since(now)
			// 浮点数是不精准的, 所以无法类比
			if seconds := diff.Seconds(); int64(seconds) != tt.want {
				t.Errorf("Time() ReturnType = %v, want %v", seconds, tt.want)
			}
		})

	}
}

func TestStrpos(t *testing.T) {
	tests := []struct {
		name             string
		haystack, needle string
		offset           int
		want             int
	}{
		{name: "Standard", haystack: "hello world", needle: "w", offset: 0, want: 6},
		{name: "TestNegative", haystack: "hello world", needle: "w", offset: -5, want: 6},
		{name: "TestUTF8", haystack: "一二三四五六七八九十", needle: "六", offset: 0, want: 15}, // UTF8 一般中文 3个字符
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if pos, ok := php.Strpos(tt.haystack, tt.needle, tt.offset); !ok || pos != tt.want {
				t.Errorf("Time() ReturnType = %v, want %v", pos, tt.want)
			}

		})
	}
}
