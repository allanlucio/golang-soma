package main

import "testing"

func TestSoma(t *testing.T){
  resultado := Soma(5,5)
  if resultado != 10{
    t.Errorf("Soma esperada: %.2f, obtida: %.2f", 10.0, resultado)
  }
}