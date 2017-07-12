// EvoChess project: Core
package main

import (
	общийМодуль "EvoChess/Библиотека"
	"fmt"
	"os"
	"time"
)

func main() {

	//определим базовые типы
	type позиция [8][8]byte

	type списокСледующихПозиций struct {
		текущаяПозиция      позиция
		вероятностьВыигрыша uint32
	}

	type игроваяСитуация struct {
		текущаяПозиция      позиция
		былШахБелым         bool
		былШахЧерным        bool
		былаРокировкаБелых  bool
		былаРокировкаЧерных bool
	}
	игра := new(игроваяСитуация)
	fmt.Println(игра.былШахБелым)

	//	fmt.Println(pos)
	//	type ChessBoard struct{}

	fmt.Println(os.Hostname())
	fmt.Println("Starting EvoChess.Core " + общийМодуль.Version())

	//Пока не нажали CTRL+C крутим обучающий цикл
	ttime := time.Now()
	stime := ttime.Format("01.01.2006 00.00.00")
	for {
		break
		continue
		ttime := time.Now()
		ftime := ttime.Format("01.01.2006 00.00.00")

		if stime != ftime {
			fmt.Println(ftime)
			stime = ftime
		}

		//начинаем игру
		//		pos := new(position)
	}

	//	fmt.Println(TestPackage.Pi())

	//	board := "1000000000000000000000000000000000000000000000000000000000000000"

	//	fmt.Println(board[1])

}
