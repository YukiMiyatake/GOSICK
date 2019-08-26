package main

import "testing"

type AInterface interface {
	Hoge()
}

type A struct {
	Get AInterface
}

type D struct {
	AInterface
}

type AImplements struct {
}

func NewAImplements() AInterface {
	return &AImplements{}
}

func (s AImplements) Hoge() {
}

type BImplements struct {
}

func (s *BImplements) Hoge() {
}

var a_ = &A{&AImplements{}}
var a2_ = &AImplements{}
var a3_ = NewAImplements()
var b_ = &BImplements{}
var c_ AInterface = &BImplements{}

var d_ = D{&AImplements{}}
var d2_ = D{&BImplements{}}

var e_ D

func init() {
	e_.AInterface = a2_

}

func Benchmark_A(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a_.Get.Hoge()
	}
}
func Benchmark_A2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a2_.Hoge()
	}
}
func Benchmark_A3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a3_.Hoge()
	}
}

func Benchmark_B(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b_.Hoge()
	}
}
func Benchmark_C(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c_.Hoge()
	}
}
func Benchmark_D(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d_.Hoge()
	}
}
func Benchmark_D2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d2_.Hoge()
	}
}

func Benchmark_E(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e_.Hoge()
	}
}
