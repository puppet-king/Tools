// Copyright 2024 The wangkai. ALL rights reserved.

/*
Package php
*/
package php

import (
	"time"
)

// Time 返回当前 Unix 时间戳
func Time() int64 {
	return time.Now().Unix()
}

// Strtotime 将任何英文文本日期时间描述解析为 Unix 时间戳, baseTimeStamp 没有提供则使用当前时间
// PHP 支持一系列的 日期与时间格式, 支持类似 BNF 的语法、 W3C RSS ISO8601 等等
// 而 BNF 语法要进行正则匹配, 有一些 W3C等这规范语法又需要穷举匹配, 非常消耗性能
// 而 +1 hour 这类的可以轻松使用 Add 来实现, 所以这里并不做支持
func Strtotime(layout, datetime string, baseTimeStamp *int64) int64 {
	// 根据 datetime 转成 layout 的格式, 用于解析
	var formatDateTime string
	if baseTimeStamp != nil {
		formatDateTime = time.Unix(*baseTimeStamp, 0).Format(layout)
	} else {
		// 判断 baseTimeStamp 参数
		formatDateTime = time.Now().Format(layout)
	}

	// 根据 datetime 转义成时间戳
	parseTime, err := time.ParseInLocation(layout, formatDateTime, time.Local)
	if err != nil {

		return -1
	} else {
		return parseTime.Unix()
	}
}

// Checkdate 验证一个格力高日期（公历）
func Checkdate(month, day, year int) bool {
	if month < 1 || month > 12 || year < 1 || year > 32767 || day < 1 || day > 31 {
		return false
	}

	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		// 闰年规则
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}

	return true
}
