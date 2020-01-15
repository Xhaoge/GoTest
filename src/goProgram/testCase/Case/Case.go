package Case

import (
	"fmt"
	"Pro_golang/Yuetu/Case/search"
)

var AllCaseList []interface{}


// type CaseList struct{
// 	SearchList 		[]CaseStruct
// 	VerifyList		[]CaseStruct
// 	OrderLise		[]CaseStruct
// 	OthersList		[]CaseStruct
// }


func init(){
	fmt.Println("this is case init")
	fmt.Println("CaseSearchList:",searchCase.CaseSearchList)
	AllCaseList = append(AllCaseList,searchCase.CaseSearchList)
	fmt.Println("AllCaseList:",AllCaseList)
}

func CaseTest(){
	fmt.Println("this is testcase package")
}