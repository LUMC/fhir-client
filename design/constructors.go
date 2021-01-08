package design

// This file contains manually generated convenience constructors for our
// models.

// NewOperationOutcome creates a pointer to an OperationOutcome and sets the
// severity, code and diagnostics for the first issue.
func NewOperationOutcome(severity, code, diagnostics string) *OperationOutcome {
	return &OperationOutcome{
		Issue: []OperationOutcomeIssueComponent{
			{
				Severity:    severity,
				Code:        code,
				Diagnostics: diagnostics,
			},
		},
	}
}
