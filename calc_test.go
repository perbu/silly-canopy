
package main

import (
	"testing"
	"strconv"
)

func TestAddition(t *testing.T) {

	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "2 + 3"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 5.0 {
	    t.Error("Expected 5.0, got ", resF)
  	}

}

func TestSub(t *testing.T) {
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "10 - 3"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 7 {
	    t.Error("Expected 7.0, got ", resF)
  	}

}

func TestMul(t *testing.T) {
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "8 * 7"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 56 {
	    t.Error("Expected 56, got ", resF)
  	}

}

func TestDiv(t *testing.T) {
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "64 / 8"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 8 {
	    t.Error("Expected 8, got ", resF)
  	}
}

func XTestDivZero(t *testing.T) {
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "64 / 0"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 0 {
	    t.Error("Expected 0, got ", resF)
  	}

}


func XTestInputWithoutSpace(t *testing.T) {
	inputC := make(chan string)
	outputC := make(chan string)
	go calc(inputC,outputC)
	inputC<- "2+2"
	res := <-outputC
	resF,_ := strconv.ParseFloat(res, 64);

 	if resF != 4 {
	    t.Error("Expected 4, got ", resF)
  	}

}



