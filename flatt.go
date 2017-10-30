package main

import "fmt"

type Category struct {
  Id int
  Parent_id int
  Category_name string
}

func newCategory(Id int, Parent_id int, Category_name string) *Category {
  c := new(Category)
  c.Id = Id
  c.Parent_id = Parent_id
  c.Category_name = Category_name
  return c
}

func main()  {
  var u *Category = newCategory(1,1,"ファッション")
  var v *Category = newCategory(2,2,"家具・インテリア")
  var w *Category = newCategory(29,1,"レディース")
  var x *Category = newCategory(30,1,"メンズ")
  var y *Category = newCategory(38,2,"収納家具")
  var z *Category = newCategory(200,29,"靴")
  fmt.Println(u)
  fmt.Println(v)
  fmt.Println(w)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(v.Category_name)
}
