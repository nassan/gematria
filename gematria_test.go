package gematria

import (
	"testing"
)

//Testgematria tests that gematria is giving the correct result.
func TestGematria(t *testing.T) {
	g := Load()

	if r := g.Get(0); r != "" {
		t.Error("Expected an empty string, got ", r)
	}
	if r := g.Get(1); r != "א" {
		t.Error("Expected א, got ", r)
	}

	if r := g.Get(10); r != "י" {
		t.Error("Expected י, got ", r)
	}

	if r := g.Get(13); r != "יג" {
		t.Error("Expected יג, got ", r)
	}

	if r := g.Get(15); r != "טו" {
		t.Error("Expected טו, got ", r)
	}

	if r := g.Get(16); r != "טז" {
		t.Error("Expected טז, got ", r)
	}

	if r := g.Get(17); r != "יז" {
		t.Error("Expected יז, got ", r)
	}

	if r := g.Get(20); r != "כ" {
		t.Error("Expected כ, got ", r)
	}

	if r := g.Get(25); r != "כה" {
		t.Error("Expected כה, got ", r)
	}

	if r := g.Get(59); r != "נט" {
		t.Error("Expected נט, got ", r)
	}

	if r := g.Get(100); r != "ק" {
		t.Error("Expected ק, got ", r)
	}

	if r := g.Get(105); r != "קה" {
		t.Error("Expected קה, got ", r)
	}

	if r := g.Get(115); r != "קטו" {
		t.Error("Expected קטו, got ", r)
	}

	if r := g.Get(1115); r != "תתשטו" {
		t.Error("Expected קטו, got ", r)
	}
}
