package goybuben_test

import (
	"reflect"
	"testing"

	goybuben "github.com/ohmin839/goybuben/goybubenapi"
)

func TestToAybuben(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		// upper cases
		{"A", "Ա"},
		{"B", "Բ"},
		{"G", "Գ"},
		{"D", "Դ"},
		{"E", "Ե"},
		{"Z", "Զ"},
		{"E'", "Է"},
		{"Y'", "Ը"},
		{"T'", "Թ"},
		{"Zh", "Ժ"},
		{"I", "Ի"},
		{"L", "Լ"},
		{"X", "Խ"},
		{"C'", "Ծ"},
		{"K", "Կ"},
		{"H", "Հ"},
		{"Dz", "Ձ"},
		{"Gh", "Ղ"},
		{"Tw", "Ճ"},
		{"M", "Մ"},
		{"Y", "Յ"},
		{"N", "Ն"},
		{"Sh", "Շ"},
		{"Vo", "Ո"},
		{"Ch", "Չ"},
		{"P", "Պ"},
		{"J", "Ջ"},
		{"Rr", "Ռ"},
		{"S", "Ս"},
		{"V", "Վ"},
		{"T", "Տ"},
		{"R", "Ր"},
		{"C", "Ց"},
		{"W", "Ւ"},
		{"P'", "Փ"},
		{"Q", "Ք"},
		{"O", "Օ"},
		{"F", "Ֆ"},
		{"U", "Ու"},
		// lower cases
		{"a", "ա"},
		{"b", "բ"},
		{"g", "գ"},
		{"d", "դ"},
		{"e", "ե"},
		{"z", "զ"},
		{"e'", "է"},
		{"y'", "ը"},
		{"t'", "թ"},
		{"zh", "ժ"},
		{"i", "ի"},
		{"l", "լ"},
		{"x", "խ"},
		{"c'", "ծ"},
		{"k", "կ"},
		{"h", "հ"},
		{"dz", "ձ"},
		{"gh", "ղ"},
		{"tw", "ճ"},
		{"m", "մ"},
		{"y", "յ"},
		{"n", "ն"},
		{"sh", "շ"},
		{"vo", "ո"},
		{"ch", "չ"},
		{"p", "պ"},
		{"j", "ջ"},
		{"rr", "ռ"},
		{"s", "ս"},
		{"v", "վ"},
		{"t", "տ"},
		{"r", "ր"},
		{"c", "ց"},
		{"w", "ւ"},
		{"p'", "փ"},
		{"q", "ք"},
		{"o", "օ"},
		{"f", "ֆ"},
		{"u", "ու"},
		{"ev", "և"},
		// others
		{"$", "֏"},
		{"1234567890", "1234567890"},
		{",", ","},
		{".", "."},
		{"`", "՝"},
		{":", "։"},
		{"-", "-"},
		{"(", "("},
		{")", ")"},
		{"<<", "«"},
		{">>", "»"},
		{"?", "՞"},
		{"!", "՛"},
		{"!~", "՜"},
		{" ", " "},
		{"\t", "\t"},
		{"\n", "\n"},
		{"\r\n", "\r\n"},
		{"あ", "あ"},
		{"Barev Dzez:", "Բարև Ձեզ։"},
	}
	for _, c := range cases {
		got := goybuben.ToAybuben(c.in)
		if got != c.want {
			t.Errorf("Expected: %s, Actual: %s", c.want, got)
		}
	}
}

func TestToHayerenWords(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"", []string{}},
		{"Բարև Ձեզ։", []string{"Բարև", "Ձեզ"}},
	}
	for _, c := range cases {
		got := goybuben.ToHayerenWords(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Expected: %v, Actual: %v", c.want, got)
		}
	}
}
