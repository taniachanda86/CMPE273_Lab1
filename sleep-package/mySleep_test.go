package sleep
import (
    "time"
    "testing"
)

const (
    timeFormat = "2006-01-02 15:04 MST"
)
var duration1 int64
var duration2 int64

func buildInSleepTime(n int) {
	now1 := time.Now().Format(timeFormat)
	start1, _ := time.Parse(timeFormat, now1)
	time.Sleep(time.Duration(n)* time.Second)
	duration1 = int64(time.Since(start1)/time.Second)
	//return duration1
}

func mySleepTime(n int )  {
	now2 := time.Now().Format(timeFormat)
	start2, _ := time.Parse(timeFormat, now2)
	Sleep(n)
	duration2 = int64(time.Since(start2)/time.Second)
	//return duration2	
}

func TestSleep(t *testing.T) {

	for i:=1; i<3; i++{
		go buildInSleepTime(i)
		go mySleepTime(i)

		if duration2 != duration1{
			t.Error("Alert!!! This new sleep function is not working properly!!!", duration1, duration2)
		}
	}
	
}
// func TestSleep(t *testing.T){
//  	var i int64
// 	// //for i=1; i<3; i++{
// 		now := time.Now()
// 		nowFormatted := now.Format(timeFormat)
// 		start, _ := time.Parse(timeFormat, nowFormatted)
// 		Sleep(i)
// 		duration := int64(time.Since(start)/time.Second)
// 		if duration!= i{
// 			t.Error("Alert!!! This new sleep function is not working properly!!!", duration, i)
// 		}
// 	//}
// }