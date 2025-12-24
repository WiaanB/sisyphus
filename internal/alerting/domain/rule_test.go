package domain

import "testing"

func TestRule_IsViolated(t *testing.T) {
	tests := []struct {
		name  string
		rule  Rule
		value float64
		want  bool
	}{
		{
			name: "name is violated",
			rule: Rule{
				ID:        "1",
				Metric:    "p95_perf",
				Threshold: 10,
				Severity:  AlarmSeverityCritical,
			},
			value: 11,
			want:  true,
		},
		{
			name: "name is not violated",
			rule: Rule{
				ID:        "1",
				Metric:    "p95_perf",
				Threshold: 10,
				Severity:  AlarmSeverityCritical,
			},
			value: 10,
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.rule.IsViolated(test.value)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
