package main

import "fmt"

type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func main() {
	pepe1 := NewPepeSchnele(10, 8, 7)
	pepe2 := NewPepeSchnele(5, 9, 6)

	rating1 := pepe1.GetRating()
	rating2 := pepe2.GetRating()

	fmt.Println(fmt.Sprintf("пепе шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d", pepe1.Speed, pepe1.Charisma, pepe1.Wisdom, rating1))
	fmt.Println(fmt.Sprintf("пепе шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d", pepe2.Speed, pepe2.Charisma, pepe2.Wisdom, rating2))

	if rating1 > rating2 {
		fmt.Println("первый пепе шнеле имеет более высокий фа рейтинг вотафа")
	} else if rating2 > rating1 {
		fmt.Println("второй пепе шнеле имеет более высокий фа рейтинг вотафа")
	} else {
		fmt.Println("рейтинги равны шнеле")
	}
}