package ac
import "fmt"

func main(){
	e := Employee{
		FirstName: "xhaoge",
        LastName: "sun",
        TotalLeaves: 44,
        LeavesTaken: 11,
	}
	e.LeavesRemaining()

}