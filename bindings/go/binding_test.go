package tree_sitter_sparql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_sparql "github.com/ioannisnezis/tree-sitter-sparql.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sparql.Language())
	if language == nil {
		t.Errorf("Error loading Sparql grammar")
	}
}
