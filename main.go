package main


import (
	"fmt"
  "math/rand"
  "time"
)
func randPoint()(x,x1,x2,y,y1,y2 int){
  x=rand.Int()
  x1=rand.Int()
  x2=rand.Int()
  y=rand.Int()
  y1=rand.Int()
  y2=rand.Int()
  return
}
func coordRecalc(x,x1,x2,y,y1,y2 int)(int,int, int, int, int, int){
  x=2*x+10
  x1=2*x1+10
  x2=2*x2+10
  y=-3*y-5
  y1=-3*y1-5
  y2=-3*y2-5
  return x,x1,x2,y,y1,y2
}
func main() {
  rand.Seed(time.Now().UnixNano())
	fmt.Println("Функция, возвращающая несколько значений\n")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru \n\n")
  x,x1,x2,y,y1,y2:=randPoint()
  x,x1,x2,y,y1,y2=coordRecalc(x,x1,x2,y,y1,y2)
  //fmt.Println(x,x1,x2,y,y1,y2) в условиях задачи не было ничего про вывод
}
/*Задание 2. Функция, возвращающая несколько значений
Что нужно сделать
Напишите программу, которая с помощью функции генерирует три случайные точки в двумерном пространстве (две координаты), а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.*/
