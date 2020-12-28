package main

import (
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/deutsch"
)

func main()  {
	greeting.Hello()
	greeting.Hi()

	dansk.GodMorgen()
	dansk.Hej()

	deutsch.GutenTag()
	deutsch.Hallo()


}
