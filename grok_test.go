package grok

import "testing"

func TestDenormalizeGlobalPatterns(t *testing.T) {
	if denormalized, err := DenormalizePatternsFromMap(defalutPatterns); err != nil {
		t.Error(err)
	} else {
		if len(defalutPatterns) != len(denormalized) {
			t.Error("len(GlobalPatterns) != len(denormalized)")
		}
		for k := range denormalized {
			if _, ok := defalutPatterns[k]; !ok {
				t.Errorf("%s not exists", k)
			}
		}
	}
}

func TestParse(t *testing.T) {
	patternINT, err := DenormalizePattern(defalutPatterns["INT"])
	if err != nil {
		t.Error(err)
	}

	patterns := map[string]string{
		"INT": patternINT,
	}

	denormalized, err := DenormalizePatternsFromMap(defalutPatterns, patterns)
	if err != nil {
		t.Error(err)
	}
	g, err := CompilePattern("%{DAY:day}", denormalized)
	if err != nil {
		t.Error(err)
	}
	ret, err := g.Run("Tue qds")
	if err != nil {
		t.Error(err)
	}
	if ret["day"] != "Tue" {
		t.Fatalf("day should be 'Tue' have '%s'", ret["day"])
	}
}

func TestParseFromPathPattern(t *testing.T) {
	pathPatterns, err := DenormalizePatternsFromPath("./patterns")
	if err != nil {
		t.Error(err)
	}
	de, err := DenormalizePatternsFromMap(pathPatterns)
	if err != nil {
		t.Error(err)
	}
	g, err := CompilePattern("%{DAY:day}", de)
	if err != nil {
		t.Error(err)
	}
	ret, err := g.Run("Tue qds")
	if err != nil {
		t.Error(err)
	}
	if ret["day"] != "Tue" {
		t.Fatalf("day should be 'Tue' have '%s'", ret["day"])
	}
}

func TestDenormalizePatternsFromPathErr(t *testing.T) {
	_, err := DenormalizePatternsFromPath("./Lorem ipsum Minim qui in.")
	if err == nil {
		t.Fatalf("AddPatternsFromPath should returns an error when path is invalid")
	}
}
