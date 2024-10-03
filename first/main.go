package main

import (
	"fmt"
	"math"
)

// Объявляем интерфейс с одним методом Move
type Shape interface {
	Area() float64
}

// Создадим структуру для круга с полем радиус типа float64
type Circle struct {
	Radius float64
}

// Создадим структуру для прямоугольника с полями высота/ширина типа float64
type Rectangle struct {
	Width  float64
	Height float64
}

// Сделаем метод для структуры круга и посчитаем площадь
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Сделаем метод для структуры прямоугольника и посчитаем площадь
func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Важный момент! (на будущее)
// Обрати внимание на звездочку в методе, так мы показываем то, что передаем
// именно указатель на структуру, иначе (если ее не будет)
// в функцию полетит копия структуры, и если мы будем менять там поля,
// то это коснется лишь копии структуры
// Подробнее здесь: https://metanit.com/go/tutorial/4.5.php

func main() {
	// Объявляем интерфейсы, так как интерфейсом является объект, который
	// содержит все его методы
	var shapeFirst = Circle{Radius: 1}
	var shapeSecond = Rectangle{Width: 5, Height: 5}
	// Получаем знаечения
	fmt.Println(shapeFirst.Area())
	fmt.Println(shapeSecond.Area())
}
