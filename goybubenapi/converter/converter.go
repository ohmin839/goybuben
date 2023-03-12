package converter

import "strings"

func (c *Converter) Evaluate() string {
	if ast := c.AST(); ast != nil {
		return c.Ruleletters(ast)
	}
	return ""
}

func (c *Converter) Ruleletters(node *node32) string {
	var b strings.Builder
	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleletter:
			b.WriteString(c.Ruleletter(node))
		}
		node = node.next
	}
	return b.String()
}

func (c *Converter) Ruleletter(node *node32) string {
	node = node.up
	switch node.pegRule {
	case rulealphabet:
		return c.Rulealphabet(node)
	case rulenonAlphabet:
		return c.RulenonAlphabet(node)
	default:
		return c.Ruleasis(node)
	}
}

func (c *Converter) Rulealphabet(node *node32) string {
	node = node.up
	switch node.pegRule {
	case rulelargeA:
		return c.RulelargeA(node)
	case rulelargeB:
		return c.RulelargeB(node)
	case rulelargeC:
		return c.RulelargeC(node)
	case rulelargeD:
		return c.RulelargeD(node)
	case rulelargeE:
		return c.RulelargeE(node)
	case rulelargeF:
		return c.RulelargeF(node)
	case rulelargeG:
		return c.RulelargeG(node)
	case rulelargeH:
		return c.RulelargeH(node)
	case rulelargeI:
		return c.RulelargeI(node)
	case rulelargeJ:
		return c.RulelargeJ(node)
	case rulelargeK:
		return c.RulelargeK(node)
	case rulelargeL:
		return c.RulelargeL(node)
	case rulelargeM:
		return c.RulelargeM(node)
	case rulelargeN:
		return c.RulelargeN(node)
	case rulelargeO:
		return c.RulelargeO(node)
	case rulelargeP:
		return c.RulelargeP(node)
	case rulelargeQ:
		return c.RulelargeQ(node)
	case rulelargeR:
		return c.RulelargeR(node)
	case rulelargeS:
		return c.RulelargeS(node)
	case rulelargeT:
		return c.RulelargeT(node)
	case rulelargeU:
		return c.RulelargeU(node)
	case rulelargeV:
		return c.RulelargeV(node)
	case rulelargeW:
		return c.RulelargeW(node)
	case rulelargeX:
		return c.RulelargeX(node)
	case rulelargeY:
		return c.RulelargeY(node)
	case rulelargeZ:
		return c.RulelargeZ(node)
	case rulesmallA:
		return c.RulesmallA(node)
	case rulesmallB:
		return c.RulesmallB(node)
	case rulesmallC:
		return c.RulesmallC(node)
	case rulesmallD:
		return c.RulesmallD(node)
	case rulesmallE:
		return c.RulesmallE(node)
	case rulesmallF:
		return c.RulesmallF(node)
	case rulesmallG:
		return c.RulesmallG(node)
	case rulesmallH:
		return c.RulesmallH(node)
	case rulesmallI:
		return c.RulesmallI(node)
	case rulesmallJ:
		return c.RulesmallJ(node)
	case rulesmallK:
		return c.RulesmallK(node)
	case rulesmallL:
		return c.RulesmallL(node)
	case rulesmallM:
		return c.RulesmallM(node)
	case rulesmallN:
		return c.RulesmallN(node)
	case rulesmallO:
		return c.RulesmallO(node)
	case rulesmallP:
		return c.RulesmallP(node)
	case rulesmallQ:
		return c.RulesmallQ(node)
	case rulesmallR:
		return c.RulesmallR(node)
	case rulesmallS:
		return c.RulesmallS(node)
	case rulesmallT:
		return c.RulesmallT(node)
	case rulesmallU:
		return c.RulesmallU(node)
	case rulesmallV:
		return c.RulesmallV(node)
	case rulesmallW:
		return c.RulesmallW(node)
	case rulesmallX:
		return c.RulesmallX(node)
	case rulesmallY:
		return c.RulesmallY(node)
	case rulesmallZ:
		return c.RulesmallZ(node)
	}
	return ""
}

func (c *Converter) RulelargeA(node *node32) string {
	return "\u0531"
}
func (c *Converter) RulelargeB(node *node32) string {
	return "\u0532"
}
func (c *Converter) RulelargeC(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0549"
	}
	if strings.Contains(text, "'") {
		return "\u053E"
	}
	return "\u0551"
}
func (c *Converter) RulelargeD(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "z") {
		return "\u0541"
	}
	return "\u0534"
}
func (c *Converter) RulelargeE(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "'") {
		return "\u0537"
	}
	return "\u0535"
}
func (c *Converter) RulelargeF(node *node32) string {
	return "\u0556"
}
func (c *Converter) RulelargeG(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0542"
	}
	return "\u0533"
}
func (c *Converter) RulelargeH(node *node32) string {
	return "\u0540"
}
func (c *Converter) RulelargeI(node *node32) string {
	return "\u053B"
}
func (c *Converter) RulelargeJ(node *node32) string {
	return "\u054B"
}
func (c *Converter) RulelargeK(node *node32) string {
	return "\u053F"
}
func (c *Converter) RulelargeL(node *node32) string {
	return "\u053C"
}
func (c *Converter) RulelargeM(node *node32) string {
	return "\u0544"
}
func (c *Converter) RulelargeN(node *node32) string {
	return "\u0546"
}
func (c *Converter) RulelargeO(node *node32) string {
	return "\u0555"
}
func (c *Converter) RulelargeP(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "'") {
		return "\u0553"
	}
	return "\u054A"
}
func (c *Converter) RulelargeQ(node *node32) string {
	return "\u0554"
}
func (c *Converter) RulelargeR(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "r") {
		return "\u054C"
	}
	return "\u0550"
}
func (c *Converter) RulelargeS(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0547"
	}
	return "\u054D"
}
func (c *Converter) RulelargeT(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "w") {
		return "\u0543"
	}
	if strings.Contains(text, "'") {
		return "\u0539"
	}
	return "\u054F"
}
func (c *Converter) RulelargeU(node *node32) string {
	return "\u0548\u0582"
}
func (c *Converter) RulelargeV(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "o") {
		return "\u0548"
	}
	return "\u054E"
}
func (c *Converter) RulelargeW(node *node32) string {
	return "\u0552"
}
func (c *Converter) RulelargeX(node *node32) string {
	return "\u053D"
}
func (c *Converter) RulelargeY(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "'") {
		return "\u0538"
	}
	return "\u0545"
}
func (c *Converter) RulelargeZ(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u053A"
	}
	return "\u0536"
}

func (c *Converter) RulesmallA(node *node32) string {
	return "\u0561"
}
func (c *Converter) RulesmallB(node *node32) string {
	return "\u0562"
}
func (c *Converter) RulesmallC(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0579"
	}
	if strings.Contains(text, "'") {
		return "\u056E"
	}
	return "\u0581"
}
func (c *Converter) RulesmallD(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "z") {
		return "\u0571"
	}
	return "\u0564"
}
func (c *Converter) RulesmallE(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "v") {
		return "\u0587"
	}
	if strings.Contains(text, "'") {
		return "\u0567"
	}
	return "\u0565"
}
func (c *Converter) RulesmallF(node *node32) string {
	return "\u0586"
}
func (c *Converter) RulesmallG(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0572"
	}
	return "\u0563"
}
func (c *Converter) RulesmallH(node *node32) string {
	return "\u0570"
}
func (c *Converter) RulesmallI(node *node32) string {
	return "\u056B"
}
func (c *Converter) RulesmallJ(node *node32) string {
	return "\u057B"
}
func (c *Converter) RulesmallK(node *node32) string {
	return "\u056F"
}
func (c *Converter) RulesmallL(node *node32) string {
	return "\u056C"
}
func (c *Converter) RulesmallM(node *node32) string {
	return "\u0574"
}
func (c *Converter) RulesmallN(node *node32) string {
	return "\u0576"
}
func (c *Converter) RulesmallO(node *node32) string {
	return "\u0585"
}
func (c *Converter) RulesmallP(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "'") {
		return "\u0583"
	}
	return "\u057A"
}
func (c *Converter) RulesmallQ(node *node32) string {
	return "\u0584"
}
func (c *Converter) RulesmallR(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if text == "rr" {
		return "\u057C"
	}
	return "\u0580"
}
func (c *Converter) RulesmallS(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u0577"
	}
	return "\u057D"
}
func (c *Converter) RulesmallT(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "w") {
		return "\u0573"
	}
	if strings.Contains(text, "'") {
		return "\u0569"
	}
	return "\u057F"
}
func (c *Converter) RulesmallU(node *node32) string {
	return "\u0578\u0582"
}
func (c *Converter) RulesmallV(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "o") {
		return "\u0578"
	}
	return "\u057E"
}
func (c *Converter) RulesmallW(node *node32) string {
	return "\u0582"
}
func (c *Converter) RulesmallX(node *node32) string {
	return "\u056D"
}
func (c *Converter) RulesmallY(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "'") {
		return "\u0568"
	}
	return "\u0575"
}
func (c *Converter) RulesmallZ(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "h") {
		return "\u056A"
	}
	return "\u0566"
}

func (c *Converter) RulenonAlphabet(node *node32) string {
	node = node.up
	switch node.pegRule {
	case ruledollar:
		return c.Ruledollar(node)
	case rulebackQuote:
		return c.RulebackQuote(node)
	case rulecolon:
		return c.Rulecolon(node)
	case ruleleftGuillemet:
		return c.RuleleftGuillemet(node)
	case rulerightGuillemet:
		return c.RulerightGuillemet(node)
	case rulequestion:
		return c.Rulequestion(node)
	case ruleexclamation:
		return c.Ruleexclamation(node)
	}
	return ""
}

func (c *Converter) Ruledollar(node *node32) string {
	return "\u058F"
}
func (c *Converter) RulebackQuote(node *node32) string {
	return "\u055D"
}
func (c *Converter) Rulecolon(node *node32) string {
	return "\u0589"
}
func (c *Converter) RuleleftGuillemet(node *node32) string {
	return "\u00AB"
}
func (c *Converter) RulerightGuillemet(node *node32) string {
	return "\u00BB"
}
func (c *Converter) Rulequestion(node *node32) string {
	return "\u055E"
}
func (c *Converter) Ruleexclamation(node *node32) string {
	text := string(c.buffer[node.begin:node.end])
	if strings.Contains(text, "~") {
		return "\u055C"
	}
	return "\u055B"
}

func (c *Converter) Ruleasis(node *node32) string {
	return string(c.buffer[node.begin:node.end])
}
