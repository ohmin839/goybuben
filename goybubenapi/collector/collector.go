package collector

func (c *Collector) Evaluate() []string {
	if ast := c.AST(); ast != nil {
		return c.Rulewords(c.AST())
	}
	return []string{}
}

func (c *Collector) Rulewords(node *node32) []string {
	var result []string
	node = node.up
	for node != nil {
		switch node.pegRule {
		case ruleword:
			result = append(result, c.Ruleword(node))
		}
		node = node.next
	}
	return result
}

func (c *Collector) Ruleword(node *node32) string {
	return string(c.buffer[node.begin:node.end])
}
