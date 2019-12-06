package main

/*
* @Author:hanyajun
* @Date:2019/10/10 16:51
* @Name:main
* @Function:
 */
type DemoInterface interface{}

type Demo4 struct {
	DemoInterface
	Age int
}

func main() {
	a := Demo4{
		Age: 10,
	}
	PrintType(a)
}

func PrintType(a interface{}) {

}
