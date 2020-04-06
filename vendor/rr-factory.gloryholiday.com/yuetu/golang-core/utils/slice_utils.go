// slice_utils.go  * Created on  2019-06-03
// Copyright (c) 2019 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package utils

func ContainsString(list []string, text string) bool {
	for _, v := range list {
		if v == text {
			return true
		}
	}
	return false
}

func ContainsAnyString(list []string, vals []string) bool {
	for _, v := range vals {
		if ContainsString(list, v) {
			return true
		}
	}
	return false
}
