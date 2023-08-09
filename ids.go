package main

import "fmt"

type Calculator struct {
   a int
   b int
}

func (c *Calculator)Add(){
   fmt.Println("Addition of two numbers: ", c.a + c.b)
}

func (c *Calculator)Mul(){
   fmt.Println("Multiplication of two numbers: ", c.a * c.b)
}

func (c *Calculator)Div(){
   fmt.Println("Division of two numbers: ", c.a / c.b)
}

func (c *Calculator)Sub(){
   fmt.Println("Subtraction of two numbers: ", c.a - c.b)
}

func test(){
   var a, b int
   fmt.Print("Enter the first number: ")
   fmt.Scanf("%d", &a)
   fmt.Print("Enter the second number: ")
   fmt.Scanf("%d", &b)
   cal := Calculator{
      a: a,
      b: b,
   }
   c:=1
   for c>=1 {
      fmt.Println("Enter 1 for Addition: ")
      fmt.Println("Enter 2 for Multiplication: ")
      fmt.Println("Enter 3 for Division: ")
      fmt.Println("Enter 4 for Subtraction: ")
      fmt.Print("Enter 5 for Exit: ")
      fmt.Scanf("%d", &c)
      switch c {
         case 1:
            cal.Add()
         case 2:
            cal.Mul()
         case 3:
            cal.Div()
         case 4:
            cal.Sub()
         case 5:
            c = 0
            break
         default:
         fmt.Println("Enter valid number.")
      }      
   }
}