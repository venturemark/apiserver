package slate

func textLength(node Node) int {
	if node.Type == "" {
		return len(node.Text)
	}

	var totalLength int
	for _, n := range node.Children {
		totalLength += textLength(n)
	}

	return totalLength
}

func TextLength(nodes Nodes) int {
	var totalLength int
	for _, n := range nodes {
		totalLength += textLength(n)
	}

	return totalLength
}
