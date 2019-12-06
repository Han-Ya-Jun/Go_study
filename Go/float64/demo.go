package main

import (
	"fmt"
	"math"
	"net"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
* @Author:hanyajun
* @Date:2019/10/28 14:50
* @Name:demo1
* @Function:
 */

func main() {
	//a:=int64(72059002787201136)
	//log.SetFormatter(&log.JSONFormatter{})
	//fmt.Println(a)
	//b:=float64(a)
	//fmt.Println(b)
	//log.WithFields(log.Fields{
	//	"animal":int64(b),
	//}).Info("A walrus appears")
	a := int64(72059002787201136)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	fmt.Println(a)
	b := float64(a)
	log.Print(b)
	c := int64(b)
	log.Print(c)
	log.Debug().
		Str("Scale", "833 cents").IPAddr("localhost", net.ParseIP("127.0.0.1")).Msg("sdf")
}

//将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow(math.E, float64(retain)))
}

//将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

//精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}
