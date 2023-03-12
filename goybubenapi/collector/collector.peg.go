package collector

// Code generated by peg collector/collector.peg DO NOT EDIT.

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	rulewords
	ruleword
	ruleanychar
)

var rul3s = [...]string{
	"Unknown",
	"words",
	"word",
	"anychar",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(w io.Writer, pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Fprintf(w, " ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Fprintf(w, "%v %v\n", rule, quote)
			} else {
				fmt.Fprintf(w, "\x1B[36m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(w io.Writer, buffer string) {
	node.print(w, false, buffer)
}

func (node *node32) PrettyPrint(w io.Writer, buffer string) {
	node.print(w, true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(os.Stdout, buffer)
}

func (t *tokens32) WriteSyntaxTree(w io.Writer, buffer string) {
	t.AST().Print(w, buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(os.Stdout, buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	tree, i := t.tree, int(index)
	if i >= len(tree) {
		t.tree = append(tree, token32{pegRule: rule, begin: begin, end: end})
		return
	}
	tree[i] = token32{pegRule: rule, begin: begin, end: end}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type Collector struct {
	Buffer string
	buffer []rune
	rules  [4]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *Collector) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *Collector) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *Collector
	max token32
}

func (e *parseError) Error() string {
	tokens, err := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		err += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return err
}

func (p *Collector) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *Collector) WriteSyntaxTree(w io.Writer) {
	p.tokens32.WriteSyntaxTree(w, p.Buffer)
}

func (p *Collector) SprintSyntaxTree() string {
	var bldr strings.Builder
	p.WriteSyntaxTree(&bldr)
	return bldr.String()
}

func Pretty(pretty bool) func(*Collector) error {
	return func(p *Collector) error {
		p.Pretty = pretty
		return nil
	}
}

func Size(size int) func(*Collector) error {
	return func(p *Collector) error {
		p.tokens32 = tokens32{tree: make([]token32, 0, size)}
		return nil
	}
}
func (p *Collector) Init(options ...func(*Collector) error) error {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	for _, option := range options {
		err := option(p)
		if err != nil {
			return err
		}
	}
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := p.tokens32
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 words <- <(word / anychar)*> */
		func() bool {
			{
				position1 := position
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position4, tokenIndex4 := position, tokenIndex
						if !_rules[ruleword]() {
							goto l5
						}
						goto l4
					l5:
						position, tokenIndex = position4, tokenIndex4
						if !_rules[ruleanychar]() {
							goto l3
						}
					}
				l4:
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				add(rulewords, position1)
			}
			return true
		},
		/* 1 word <- <('Ա' / 'Բ' / 'Գ' / 'Դ' / 'Ե' / 'Զ' / 'Է' / 'Ը' / 'Թ' / 'Ժ' / 'Ի' / 'Լ' / 'Խ' / 'Ծ' / 'Կ' / 'Հ' / 'Ձ' / 'Ղ' / 'Ճ' / 'Մ' / 'Յ' / 'Ն' / 'Շ' / 'Ո' / 'Չ' / 'Պ' / 'Ջ' / 'Ռ' / 'Ս' / 'Վ' / 'Տ' / 'Ր' / 'Ց' / 'Ւ' / 'Փ' / 'Ք' / 'Օ' / 'Ֆ' / 'Ո' / 'ա' / 'բ' / 'գ' / 'դ' / 'ե' / 'զ' / 'է' / 'ը' / 'թ' / 'ժ' / 'ի' / 'լ' / 'խ' / 'ծ' / 'կ' / 'հ' / 'ձ' / 'ղ' / 'ճ' / 'մ' / 'յ' / 'ն' / 'շ' / 'ո' / 'չ' / 'պ' / 'ջ' / 'ռ' / 'ս' / 'վ' / 'տ' / 'ր' / 'ց' / 'ւ' / 'փ' / 'ք' / 'օ' / 'ֆ' / 'ո' / 'և')+> */
		func() bool {
			position6, tokenIndex6 := position, tokenIndex
			{
				position7 := position
				{
					position10, tokenIndex10 := position, tokenIndex
					if buffer[position] != rune('Ա') {
						goto l11
					}
					position++
					goto l10
				l11:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Բ') {
						goto l12
					}
					position++
					goto l10
				l12:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Գ') {
						goto l13
					}
					position++
					goto l10
				l13:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Դ') {
						goto l14
					}
					position++
					goto l10
				l14:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ե') {
						goto l15
					}
					position++
					goto l10
				l15:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Զ') {
						goto l16
					}
					position++
					goto l10
				l16:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Է') {
						goto l17
					}
					position++
					goto l10
				l17:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ը') {
						goto l18
					}
					position++
					goto l10
				l18:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Թ') {
						goto l19
					}
					position++
					goto l10
				l19:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ժ') {
						goto l20
					}
					position++
					goto l10
				l20:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ի') {
						goto l21
					}
					position++
					goto l10
				l21:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Լ') {
						goto l22
					}
					position++
					goto l10
				l22:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Խ') {
						goto l23
					}
					position++
					goto l10
				l23:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ծ') {
						goto l24
					}
					position++
					goto l10
				l24:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Կ') {
						goto l25
					}
					position++
					goto l10
				l25:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Հ') {
						goto l26
					}
					position++
					goto l10
				l26:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ձ') {
						goto l27
					}
					position++
					goto l10
				l27:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ղ') {
						goto l28
					}
					position++
					goto l10
				l28:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ճ') {
						goto l29
					}
					position++
					goto l10
				l29:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Մ') {
						goto l30
					}
					position++
					goto l10
				l30:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Յ') {
						goto l31
					}
					position++
					goto l10
				l31:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ն') {
						goto l32
					}
					position++
					goto l10
				l32:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Շ') {
						goto l33
					}
					position++
					goto l10
				l33:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ո') {
						goto l34
					}
					position++
					goto l10
				l34:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Չ') {
						goto l35
					}
					position++
					goto l10
				l35:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Պ') {
						goto l36
					}
					position++
					goto l10
				l36:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ջ') {
						goto l37
					}
					position++
					goto l10
				l37:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ռ') {
						goto l38
					}
					position++
					goto l10
				l38:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ս') {
						goto l39
					}
					position++
					goto l10
				l39:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Վ') {
						goto l40
					}
					position++
					goto l10
				l40:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Տ') {
						goto l41
					}
					position++
					goto l10
				l41:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ր') {
						goto l42
					}
					position++
					goto l10
				l42:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ց') {
						goto l43
					}
					position++
					goto l10
				l43:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ւ') {
						goto l44
					}
					position++
					goto l10
				l44:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Փ') {
						goto l45
					}
					position++
					goto l10
				l45:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ք') {
						goto l46
					}
					position++
					goto l10
				l46:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Օ') {
						goto l47
					}
					position++
					goto l10
				l47:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ֆ') {
						goto l48
					}
					position++
					goto l10
				l48:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('Ո') {
						goto l49
					}
					position++
					goto l10
				l49:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ա') {
						goto l50
					}
					position++
					goto l10
				l50:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('բ') {
						goto l51
					}
					position++
					goto l10
				l51:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('գ') {
						goto l52
					}
					position++
					goto l10
				l52:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('դ') {
						goto l53
					}
					position++
					goto l10
				l53:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ե') {
						goto l54
					}
					position++
					goto l10
				l54:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('զ') {
						goto l55
					}
					position++
					goto l10
				l55:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('է') {
						goto l56
					}
					position++
					goto l10
				l56:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ը') {
						goto l57
					}
					position++
					goto l10
				l57:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('թ') {
						goto l58
					}
					position++
					goto l10
				l58:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ժ') {
						goto l59
					}
					position++
					goto l10
				l59:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ի') {
						goto l60
					}
					position++
					goto l10
				l60:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('լ') {
						goto l61
					}
					position++
					goto l10
				l61:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('խ') {
						goto l62
					}
					position++
					goto l10
				l62:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ծ') {
						goto l63
					}
					position++
					goto l10
				l63:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('կ') {
						goto l64
					}
					position++
					goto l10
				l64:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('հ') {
						goto l65
					}
					position++
					goto l10
				l65:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ձ') {
						goto l66
					}
					position++
					goto l10
				l66:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ղ') {
						goto l67
					}
					position++
					goto l10
				l67:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ճ') {
						goto l68
					}
					position++
					goto l10
				l68:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('մ') {
						goto l69
					}
					position++
					goto l10
				l69:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('յ') {
						goto l70
					}
					position++
					goto l10
				l70:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ն') {
						goto l71
					}
					position++
					goto l10
				l71:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('շ') {
						goto l72
					}
					position++
					goto l10
				l72:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ո') {
						goto l73
					}
					position++
					goto l10
				l73:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('չ') {
						goto l74
					}
					position++
					goto l10
				l74:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('պ') {
						goto l75
					}
					position++
					goto l10
				l75:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ջ') {
						goto l76
					}
					position++
					goto l10
				l76:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ռ') {
						goto l77
					}
					position++
					goto l10
				l77:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ս') {
						goto l78
					}
					position++
					goto l10
				l78:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('վ') {
						goto l79
					}
					position++
					goto l10
				l79:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('տ') {
						goto l80
					}
					position++
					goto l10
				l80:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ր') {
						goto l81
					}
					position++
					goto l10
				l81:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ց') {
						goto l82
					}
					position++
					goto l10
				l82:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ւ') {
						goto l83
					}
					position++
					goto l10
				l83:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('փ') {
						goto l84
					}
					position++
					goto l10
				l84:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ք') {
						goto l85
					}
					position++
					goto l10
				l85:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('օ') {
						goto l86
					}
					position++
					goto l10
				l86:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ֆ') {
						goto l87
					}
					position++
					goto l10
				l87:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('ո') {
						goto l88
					}
					position++
					goto l10
				l88:
					position, tokenIndex = position10, tokenIndex10
					if buffer[position] != rune('և') {
						goto l6
					}
					position++
				}
			l10:
			l8:
				{
					position9, tokenIndex9 := position, tokenIndex
					{
						position89, tokenIndex89 := position, tokenIndex
						if buffer[position] != rune('Ա') {
							goto l90
						}
						position++
						goto l89
					l90:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Բ') {
							goto l91
						}
						position++
						goto l89
					l91:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Գ') {
							goto l92
						}
						position++
						goto l89
					l92:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Դ') {
							goto l93
						}
						position++
						goto l89
					l93:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ե') {
							goto l94
						}
						position++
						goto l89
					l94:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Զ') {
							goto l95
						}
						position++
						goto l89
					l95:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Է') {
							goto l96
						}
						position++
						goto l89
					l96:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ը') {
							goto l97
						}
						position++
						goto l89
					l97:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Թ') {
							goto l98
						}
						position++
						goto l89
					l98:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ժ') {
							goto l99
						}
						position++
						goto l89
					l99:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ի') {
							goto l100
						}
						position++
						goto l89
					l100:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Լ') {
							goto l101
						}
						position++
						goto l89
					l101:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Խ') {
							goto l102
						}
						position++
						goto l89
					l102:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ծ') {
							goto l103
						}
						position++
						goto l89
					l103:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Կ') {
							goto l104
						}
						position++
						goto l89
					l104:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Հ') {
							goto l105
						}
						position++
						goto l89
					l105:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ձ') {
							goto l106
						}
						position++
						goto l89
					l106:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ղ') {
							goto l107
						}
						position++
						goto l89
					l107:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ճ') {
							goto l108
						}
						position++
						goto l89
					l108:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Մ') {
							goto l109
						}
						position++
						goto l89
					l109:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Յ') {
							goto l110
						}
						position++
						goto l89
					l110:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ն') {
							goto l111
						}
						position++
						goto l89
					l111:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Շ') {
							goto l112
						}
						position++
						goto l89
					l112:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ո') {
							goto l113
						}
						position++
						goto l89
					l113:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Չ') {
							goto l114
						}
						position++
						goto l89
					l114:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Պ') {
							goto l115
						}
						position++
						goto l89
					l115:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ջ') {
							goto l116
						}
						position++
						goto l89
					l116:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ռ') {
							goto l117
						}
						position++
						goto l89
					l117:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ս') {
							goto l118
						}
						position++
						goto l89
					l118:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Վ') {
							goto l119
						}
						position++
						goto l89
					l119:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Տ') {
							goto l120
						}
						position++
						goto l89
					l120:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ր') {
							goto l121
						}
						position++
						goto l89
					l121:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ց') {
							goto l122
						}
						position++
						goto l89
					l122:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ւ') {
							goto l123
						}
						position++
						goto l89
					l123:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Փ') {
							goto l124
						}
						position++
						goto l89
					l124:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ք') {
							goto l125
						}
						position++
						goto l89
					l125:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Օ') {
							goto l126
						}
						position++
						goto l89
					l126:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ֆ') {
							goto l127
						}
						position++
						goto l89
					l127:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('Ո') {
							goto l128
						}
						position++
						goto l89
					l128:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ա') {
							goto l129
						}
						position++
						goto l89
					l129:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('բ') {
							goto l130
						}
						position++
						goto l89
					l130:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('գ') {
							goto l131
						}
						position++
						goto l89
					l131:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('դ') {
							goto l132
						}
						position++
						goto l89
					l132:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ե') {
							goto l133
						}
						position++
						goto l89
					l133:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('զ') {
							goto l134
						}
						position++
						goto l89
					l134:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('է') {
							goto l135
						}
						position++
						goto l89
					l135:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ը') {
							goto l136
						}
						position++
						goto l89
					l136:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('թ') {
							goto l137
						}
						position++
						goto l89
					l137:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ժ') {
							goto l138
						}
						position++
						goto l89
					l138:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ի') {
							goto l139
						}
						position++
						goto l89
					l139:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('լ') {
							goto l140
						}
						position++
						goto l89
					l140:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('խ') {
							goto l141
						}
						position++
						goto l89
					l141:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ծ') {
							goto l142
						}
						position++
						goto l89
					l142:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('կ') {
							goto l143
						}
						position++
						goto l89
					l143:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('հ') {
							goto l144
						}
						position++
						goto l89
					l144:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ձ') {
							goto l145
						}
						position++
						goto l89
					l145:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ղ') {
							goto l146
						}
						position++
						goto l89
					l146:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ճ') {
							goto l147
						}
						position++
						goto l89
					l147:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('մ') {
							goto l148
						}
						position++
						goto l89
					l148:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('յ') {
							goto l149
						}
						position++
						goto l89
					l149:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ն') {
							goto l150
						}
						position++
						goto l89
					l150:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('շ') {
							goto l151
						}
						position++
						goto l89
					l151:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ո') {
							goto l152
						}
						position++
						goto l89
					l152:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('չ') {
							goto l153
						}
						position++
						goto l89
					l153:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('պ') {
							goto l154
						}
						position++
						goto l89
					l154:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ջ') {
							goto l155
						}
						position++
						goto l89
					l155:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ռ') {
							goto l156
						}
						position++
						goto l89
					l156:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ս') {
							goto l157
						}
						position++
						goto l89
					l157:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('վ') {
							goto l158
						}
						position++
						goto l89
					l158:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('տ') {
							goto l159
						}
						position++
						goto l89
					l159:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ր') {
							goto l160
						}
						position++
						goto l89
					l160:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ց') {
							goto l161
						}
						position++
						goto l89
					l161:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ւ') {
							goto l162
						}
						position++
						goto l89
					l162:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('փ') {
							goto l163
						}
						position++
						goto l89
					l163:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ք') {
							goto l164
						}
						position++
						goto l89
					l164:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('օ') {
							goto l165
						}
						position++
						goto l89
					l165:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ֆ') {
							goto l166
						}
						position++
						goto l89
					l166:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('ո') {
							goto l167
						}
						position++
						goto l89
					l167:
						position, tokenIndex = position89, tokenIndex89
						if buffer[position] != rune('և') {
							goto l9
						}
						position++
					}
				l89:
					goto l8
				l9:
					position, tokenIndex = position9, tokenIndex9
				}
				add(ruleword, position7)
			}
			return true
		l6:
			position, tokenIndex = position6, tokenIndex6
			return false
		},
		/* 2 anychar <- <.> */
		func() bool {
			position168, tokenIndex168 := position, tokenIndex
			{
				position169 := position
				if !matchDot() {
					goto l168
				}
				add(ruleanychar, position169)
			}
			return true
		l168:
			position, tokenIndex = position168, tokenIndex168
			return false
		},
	}
	p.rules = _rules
	return nil
}