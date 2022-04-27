package hw

import (
	"math"
)

//type Geom struct {
//	X1, Y1, X2, Y2 float64
//}
//
//func (geom Geom) CalculateDistance() (distance float64) {
//
//	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
//		fmt.Println("Координаты не могут быть меньше нуля")
//		return -1
//	} else {
//		distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
//	}
//
//	// возврат расстояния между точками
//	return distance
//}

// ООП подход тут не нужен, убираем струтктуру Geom и принимаем на вход координаты.

func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	// Не понимаю, почему только положительные числа...
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 {
		// Если может возникнуть ошибка - возвращаем ошибку (p.s. имхо - лучше так, чем проверять -1, но оставил как было).
		return -1
	}

	// Если все успешно - возвращаем резултат, не сохраняем лишний раз в переменную.
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
