// EvoChess project: Core
package main

import (
	lib "EvoChess/Lib"
	"fmt"
	"os"
	"time"
)

func main() {

	//определим базовые типы
	type typePosition [8][8]byte

	//0 пустое поле
	//1 пешка белая
	//2 ладья
	//3 конь
	//4 слон
	//5 ферзь
	//6 король

	//11 пешка
	//22 ладья
	//33 конь
	//44 слон
	//55 ферзь
	//66 король

	type списокСледующихПозиций struct {
		текущаяПозиция      typePosition
		вероятностьВыигрыша uint32
	}

	type игроваяСитуация struct {
		текущаяПозиция      typePosition
		былШахБелым         bool
		былШахЧерным        bool
		былаРокировкаБелых  bool
		былаРокировкаЧерных bool
	}
	startPosition := new(typePosition)
	fmt.Println("pos: ", startPosition)
	игра := new(игроваяСитуация)
	fmt.Println(игра.былШахБелым)

	//	fmt.Println(pos)
	//	type ChessBoard struct{}

	fmt.Println(os.Hostname())
	fmt.Println("Starting EvoChess.Core " + lib.Версия())
	fmt.Println("Starting EvoChess.Core " + lib.Версия1())

	i := 1
	//Пока не нажали CTRL+C крутим обучающий цикл
	ttime := time.Now()
	stime := ttime.Format("01.01.2006 23.59.59")
	for {
		//		break
		//		continue
		time.Sleep(1000 * time.Millisecond)
		i++
		ttime := time.Now()
		ftime := ttime.Format("02.01.2006 15.04.05")

		if stime != ftime {
			fmt.Println(i)
			fmt.Println(ftime)
			stime = ftime
			i = 0
		}

		//начинаем игру
		//		pos := new(position)
	}

	//	fmt.Println(TestPackage.Pi())

	//	board := "1000000000000000000000000000000000000000000000000000000000000000"

	//	fmt.Println(board[1])

}
