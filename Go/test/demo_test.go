package demo

import "testing"

/*
* @Author:hanyajun
* @Date:2019/10/19 8:56
* @Name:test
* @Function:
 */

func TestDemo1(t *testing.T) {
	t.Log("test1")
}

func TestDemo2(t *testing.T) {
	t.Log("test2")
}

func BenchmarkDemo3(b *testing.B) {
	b.Log("benchmark1")
}

func BenchmarkDemo4(b *testing.B) {
	b.Log("benchmark2")
}
