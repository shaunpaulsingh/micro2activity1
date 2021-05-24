package main

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type Modes struct {
	arg1 string
	expected bool
}

type States struct {
	arg1 bool
	expected bool
}

type Speeds struct {
	arg1 int
	expected bool
}

type Temp struct {
	arg1 float32
	expected bool
}

var testModes = []Modes{
	{"Auto", true},
	{"Cool",true},
	{"Dry", true},
	{"Fan",true},
	{"Test",false},

}

var testStates = []States{
	{true, true},
	{false,false},
}

var testSpeeds = []Speeds{
	{0,false},
	{1,true},
	{2,true},
	{3,true},
	{4,true},
	{5,false},
}

var testTemp = []Temp{
	{16.5,false},
	{18.5, true},
	{20.0, true},
	{38.0, true},
	{39.2,false},
}

func TestModes(t *testing.T){
	for _, test := range testModes{
		var obj = Aircon{}
		if output := obj.mode(test.arg1); output != test.expected{
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

func TestOnOff(t *testing.T){
	for _, test := range testStates{
		var obj = Aircon{}
		if output := obj.onoff(test.arg1); output != test.expected{
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

func TestSpeeds(t *testing.T){
	for _, test := range testSpeeds{
		var obj = Aircon{}
		if output := obj.speed(test.arg1); output != test.expected{
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

func TestTemp(t *testing.T){
	for _, test := range testTemp{
		var obj = Aircon{}
		if output := obj.temp(test.arg1); output != test.expected{
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}