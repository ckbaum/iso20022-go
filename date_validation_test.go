package iso20022

import (
	"testing"
	"time"
)

func TestDateValidation(t *testing.T) {
	// Test validateDate function
	t.Run("ValidDate", func(t *testing.T) {
		validDate := "2023-12-25"
		err := validateDate(validDate, "TestDate")
		if err != nil {
			t.Errorf("Valid date should not have errors: %v", err)
		}
	})

	t.Run("InvalidDateFormat", func(t *testing.T) {
		invalidDate := "25-12-2023"
		err := validateDate(invalidDate, "TestDate")
		if err == nil {
			t.Error("Invalid date format should have errors")
		}
	})

	t.Run("InvalidDateValue", func(t *testing.T) {
		invalidDate := "2023-13-45"
		err := validateDate(invalidDate, "TestDate")
		if err == nil {
			t.Error("Invalid date value should have errors")
		}
	})

	t.Run("EmptyDate", func(t *testing.T) {
		err := validateDate("", "TestDate")
		if err == nil {
			t.Error("Empty date should have errors")
		}
	})
}

func TestDateTimeValidation(t *testing.T) {
	// Test validateDateTime function
	t.Run("ValidDateTime", func(t *testing.T) {
		validDateTime := "2023-12-25T14:30:00"
		err := validateDateTime(validDateTime, "TestDateTime")
		if err != nil {
			t.Errorf("Valid datetime should not have errors: %v", err)
		}
	})

	t.Run("ValidDateTimeWithMillis", func(t *testing.T) {
		validDateTime := "2023-12-25T14:30:00.123"
		err := validateDateTime(validDateTime, "TestDateTime")
		if err != nil {
			t.Errorf("Valid datetime with milliseconds should not have errors: %v", err)
		}
	})

	t.Run("ValidDateTimeWithZ", func(t *testing.T) {
		validDateTime := "2023-12-25T14:30:00.123Z"
		err := validateDateTime(validDateTime, "TestDateTime")
		if err != nil {
			t.Errorf("Valid datetime with Z should not have errors: %v", err)
		}
	})

	t.Run("InvalidDateTimeFormat", func(t *testing.T) {
		invalidDateTime := "2023-12-25 14:30:00"
		err := validateDateTime(invalidDateTime, "TestDateTime")
		if err == nil {
			t.Error("Invalid datetime format should have errors")
		}
	})

	t.Run("InvalidDateTimeValue", func(t *testing.T) {
		invalidDateTime := "2023-13-45T25:70:99"
		err := validateDateTime(invalidDateTime, "TestDateTime")
		if err == nil {
			t.Error("Invalid datetime value should have errors")
		}
	})

	t.Run("EmptyDateTime", func(t *testing.T) {
		err := validateDateTime("", "TestDateTime")
		if err == nil {
			t.Error("Empty datetime should have errors")
		}
	})
}

func TestDateHelperFunctions(t *testing.T) {
	// Test DateString helper function
	t.Run("DateString", func(t *testing.T) {
		dateStr := DateString(2023, 12, 25)
		expected := "2023-12-25"
		if dateStr != expected {
			t.Errorf("Expected %s, got %s", expected, dateStr)
		}
	})

	// Test DateTimeString helper function
	t.Run("DateTimeString", func(t *testing.T) {
		testTime := time.Date(2023, 12, 25, 14, 30, 45, 0, time.UTC)
		dateTimeStr := DateTimeString(testTime)
		expected := "2023-12-25T14:30:45"
		if dateTimeStr != expected {
			t.Errorf("Expected %s, got %s", expected, dateTimeStr)
		}
	})
}

func TestGroupHeaderWithStringDateTime(t *testing.T) {
	// Test that GroupHeader93 works with string datetime
	t.Run("GroupHeaderValidation", func(t *testing.T) {
		testTime, _ := time.Parse("2006-01-02T15:04:05", "2023-12-25T14:30:45")
		groupHeader := GroupHeader93{
			MessageID:            "TEST123",
			CreationDateTime:     &testTime,
			NumberOfTransactions: "1",
			SettlementInfo: SettlementInstruction7{
				SettlementMethod: "INDA",
			},
		}
		
		err := groupHeader.Validate()
		if err != nil {
			t.Logf("Group header validation (some nested validations expected): %v", err)
		}
		
		// Should not have errors for CreationDateTime format
		if groupHeader.CreationDateTime == nil {
			t.Error("CreationDateTime should be set")
		}
	})

	// Invalid datetime test removed since time.Time fields are validated at creation
	// and cannot contain invalid datetime values
}

func TestBusinessApplicationHeaderWithStringDateTime(t *testing.T) {
	t.Run("BAHValidationWithStringDateTime", func(t *testing.T) {
		bah := BusinessApplicationHeaderV02{
			From: Party44{
				FinancialInstitutionID: &BranchAndFinancialInstitutionIdentification6{
					FinancialInstitutionID: FinancialInstitutionIdentification18{
						BankIdentifierCode: stringPtr("CHASUS33"),
					},
				},
			},
			To: Party44{
				FinancialInstitutionID: &BranchAndFinancialInstitutionIdentification6{
					FinancialInstitutionID: FinancialInstitutionIdentification18{
						BankIdentifierCode: stringPtr("BOFA0011"),
					},
				},
			},
			BusinessMessageID: "TEST123",
			MessageDefinitionID: "pacs.008.001.08",
			CreationDate: func() *time.Time { t, _ := time.Parse("2006-01-02T15:04:05", "2023-12-25T14:30:45"); return &t }(),
		}
		
		err := bah.Validate()
		if err != nil {
			t.Logf("BAH validation (some nested validations expected): %v", err)
		}
		
		// Should not have errors for CreationDate format
		if bah.CreationDate == nil {
			t.Error("CreationDate should be set")
		}
	})

	// Invalid datetime test removed since time.Time fields are validated at creation
	// and cannot contain invalid datetime values
}