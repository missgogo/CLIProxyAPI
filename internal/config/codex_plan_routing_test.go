package config

import "testing"

func TestSanitizeCodexPlanRouting_NormalizesAndDropsInvalidRules(t *testing.T) {
	cfg := &Config{
		CodexPlanRouting: []CodexPlanRoute{
			{
				Models: []string{" Gpt-5-* ", "gpt-5-*", ""},
				Plan:   " Plus ",
			},
			{
				Models: []string{"   "},
				Plan:   "free",
			},
			{
				Models: []string{"gpt-4.1-mini"},
				Plan:   " ",
			},
		},
	}

	cfg.SanitizeCodexPlanRouting()

	if len(cfg.CodexPlanRouting) != 1 {
		t.Fatalf("len(CodexPlanRouting) = %d, want 1", len(cfg.CodexPlanRouting))
	}
	got := cfg.CodexPlanRouting[0]
	if got.Plan != "plus" {
		t.Fatalf("plan = %q, want %q", got.Plan, "plus")
	}
	if len(got.Models) != 1 || got.Models[0] != "gpt-5-*" {
		t.Fatalf("models = %#v, want []string{\"gpt-5-*\"}", got.Models)
	}
}
