package main

import (
  "fmt"
)

type Person struct {
  FirstName string
  Age       int
}

func main() {
  // 初期値を与えて作成
  a := Person{
    FirstName: "John",
    Age:       20,
  }
  fmt.Printf("%#v\n",a)

  aa := &a
  aa.Age = 25

  fmt.Printf("%#v\n",a.Age)
  fmt.Printf("%#v\n",aa.Age)

  // new関数で初期化
  var b = new(Person)
  b.FirstName = "Paul"
  b.Age = 30
  // fmt.Println(b)
  fmt.Printf("%#v\n",b);

  bb:= b
  bb.Age = 35

  fmt.Printf("%#v\n",b.Age)
  fmt.Printf("%#v\n",bb.Age)

  // make関数で初期化
  // 配列のスライス、map、チャネル専用
  c := make(map[string]int)
  c["John"] = 20
  c["Paul"] = 30
  // fmt.Println(c)
  fmt.Printf("%#v\n",c);

  var d struct {
    FirstName string
    Age       int
  }
  d.FirstName = "John"
  d.Age = 20
  fmt.Printf("%#v\n",d)
  // fmt.Println(d)
}
