package main

import "fmt"

func main() {
	//gc := gcache.New(20).
	//	LRU().
	//	Build()
	//gc.Set("key", "ok")
	//value, err := gc.Get("key")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Get:", value)

	//fmt.Println((int(time.Unix(1598255760, 0).Weekday())+6)%7) //2020/8/24 15:56:0 周一 1 0
	//fmt.Println((int(time.Unix(1598687760, 0).Weekday())+6)%7) //2020/8/24 15:56:0 周六 6 5
	//fmt.Println((int(time.Unix(1598774160, 0).Weekday())+6)%7) //2020/8/24 15:56:0 周日 0 6

	//vecDssm := make([]float32, 10, 10)
	//numbers := []float32{0,1,2,3,4}
	//vecDssm = numbers
	//fmt.Println(vecDssm)
	//var TextArea []byte
	////TextArea=make([]byte,16)
	//copy(TextArea[:],[]byte("CN10"))
	//fmt.Println(""[:6])

	//nTime := time.Now()
	//limitTimeTwoDay := nTime.AddDate(0, 0, -2)
	//
	//fmt.Println(limitTimeTwoDay.Unix())
	//
	//timeLimit := time.Now().Add(-51 * time.Hour).Unix()
	//fmt.Println(timeLimit)
	//var buffer strings.Builder
	//buffer.WriteString(fmt.Sprintf("%d", 2000))//32
	//buffer.WriteString(fmt.Sprintf("|%s,%.4f,%d", "c0faaf8abcbd398bbd6e43eef92fd068",1.23, 1))
	//buffer.WriteString(fmt.Sprintf("%s", "c0faaf8abcbd398bbd6e43eef92fd068"))//32
	//buffer.WriteString(fmt.Sprintf("|%s", "c0faaf8abcbd398bbd6e43eef92fd068"))//33
	//buffer.WriteString(fmt.Sprintf("|%s,", "c0faaf8abcbd398bbd6e43eef92fd068"))//34
	//buffer.WriteString(fmt.Sprintf("|%s,%.4f", "c0faaf8abcbd398bbd6e43eef92fd068",1.234))//40
	//buffer.WriteString(fmt.Sprintf("|%s,%.4f,%.4f", "c0faaf8abcbd398bbd6e43eef92fd068",1.234, 1.23))//47
	//buffer.WriteString(fmt.Sprintf("|%s,%.4f,%.4f,%d", "c0faaf8abcbd398bbd6e43eef92fd068",10.234, 1.23,102))//51
	//buffer.WriteString(fmt.Sprintf("|%s,%.4f,%.4f,%d,%d", "c0faaf8abcbd398bbd6e43eef92fd068",10.2345, -10.2334,2000,3))
	//buffer.Grow(100)
	//32+20+
	//buffer.WriteString(fmt.Sprintf("%s,%s","c0faaf8abcbd398bbd6e43eef92fd068","comos:ivhvpwz1325968"))
	//buffer.WriteString(fmt.Sprintf("|%s,%s,%.4f,%.4f,%d,%d,%d","c0faaf8abcbd398bbd6e43eef92fd068","comos:ivhvpwz1325968",1000.0,12.3320,1203,2001,10000))
	//buffer.WriteString(fmt.Sprintf("|%s,%s,%.4f,%.4f,%.4f,%d,%d","c0faaf8abcbd398bbd6e43eef92fd068","comos:ivhvpwz1325968",1000.1234,2000.1234,3000.1234,2000,1000))
	//fmt.Println(buffer.String())
	//fmt.Println(buffer.Len())
	//fmt.Println(buffer.Cap())
	//var s1 []int     //len=0,cap=0,s1==nil
	//s2 := []int{}   //len=0,cap=0,s1!=nil
	//s3:=make([]int,0)//len=0,cap=0,s1!=nil
	//
	//fmt.Println(len(s1),cap(s1),s1==nil)
	//fmt.Println(len(s2),cap(s2),s2==nil)
	//fmt.Println(len(s3),cap(s3),s3==nil)

	//var sliceA =make([]int,1,2)
	//println(len(sliceA),cap(sliceA))
	//a := []float64{1, 2, 3}
	//b := []float64{4, 5, 6}
	//c:=floats.Dot(a,b)
	//fmt.Println(c)

	//MaxConcurrency:=20
	//len_item :=150000
	//start := 0
	//batchSize:=10000
	//for i := 0; i < MaxConcurrency; i++ {
	//	if start >=len_item{
	//		break
	//	}
	//	length := ((MaxConcurrency + i - 1) * batchSize) / (2 * (MaxConcurrency - 1))
	//	fmt.Println("------------length",length)
	//	if start+length > len_item {
	//		length = len_item - start
	//	}
	//	fmt.Println(start,start+length)
	//	start += length
	//}

	//format := "2006-01-02 15:04:05"
	//a, _ := time.Parse(format, "2021-03-10 11:00:00")
	//timeLimit := time.Now().Add(-24 * 15 * time.Hour)
	//if a.Before(timeLimit) {
	//	fmt.Println(a)
	//	fmt.Println(timeLimit)
	//} else {
	//	fmt.Println("---------")
	//}

	PcProduct := make([]uint8, 64)
	fmt.Println(len(PcProduct), cap(PcProduct))

	PcProduct = append(PcProduct, 1)
	fmt.Println(len(PcProduct), cap(PcProduct))

}
