package main

type Aircon struct{
	OnOff bool
	Mode string
	Temp float32
	Speed int
}

func main(){
	var aircon = Aircon{}
	aircon.mode("Auto")
}

func (obj *Aircon)mode(s string) bool{
	switch s {
	case "Auto":
		obj.Mode = s
		return true
	case "Cool":
		obj.Mode = s
		return true
	case "Dry":
		obj.Mode = s
		return true
	case "Fan":
		obj.Mode = s
		return true
	default:
		return false
	}
}

func (obj *Aircon)temp(f float32) bool{
	if f < 17 || f > 38{
		return false
	}else{
		obj.Temp = f
		return true
	}
}

func (obj *Aircon)onoff(boo bool) bool{
	obj.OnOff = boo
	return obj.OnOff
}

func (obj *Aircon)speed(i int) bool{
	if i < 1 || i > 4 {
		return false
	}else{
		obj.Speed = i
		return true
	}
}