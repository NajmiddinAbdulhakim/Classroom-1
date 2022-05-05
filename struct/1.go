package main

import "fmt"

type Nothing struct {
	On bool
	Ammo int
	Power int
}

func main() {
	testStruct := Nothing{
		On: true,
		Ammo: 0,
		Power: 6,
	}
	testStruct.RideBike()
	testStruct.Shoot()
	fmt.Print(testStruct)
}

func (p *Nothing) Shoot() bool {
	if p.Ammo > 0 && p.On == true {
		p.Ammo--
		return true
	}else{ 
		return false 
	}
}

func (h *Nothing) RideBike() bool { 
	if h.Power > 0 && h.On == true {
		h.Power--
		return true
	}else{ 
		return false 
	}
}