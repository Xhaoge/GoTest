package searchCase

import (
	"fmt"
)

type CaseStruct struct {
	CaseName	string
	Belong		string	
}

var CaseSearchList []interface{}

func init(){
	fmt.Println("this is SearchCase init")
	CaseSearchList = append(CaseSearchList,CaseSearchKeyValue0001)
}

// func (s CaseSearchKeyValue0001)TestInit(){
// 	fmt.Println("this is CaseSearchKeyValue0001 TestInit")
// }

// func (s CaseSearchKeyValue0001)TestProcess(){
// 	fmt.Println("this is CaseSearchKeyValue0001 TestProcess")
// }

// func (s CaseSearchKeyValue0001)TestResult(){
// 	fmt.Println("this is CaseSearchKeyValue0001 TestResult")
// }

