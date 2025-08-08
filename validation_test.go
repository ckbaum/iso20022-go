package iso20022

import (
	"testing"
	"time"
)

func TestNewValidationFunctions(t *testing.T) {
	// Test InterestType1 validation
	t.Run("InterestType1", func(t *testing.T) {
		// Valid case - exactly one choice
		validChoice := InterestType1{
			Code: stringPtr("FIXED"),
		}
		if err := validChoice.Validate(); err != nil {
			t.Errorf("Valid InterestType1 should not have errors: %v", err)
		}

		// Invalid case - no choice
		noChoice := InterestType1{}
		err := noChoice.Validate()
		if err == nil {
			t.Error("InterestType1 with no choices should have validation error")
		}

		// Invalid case - both choices
		bothChoices := InterestType1{
			Code:        stringPtr("FIXED"),
			Proprietary: stringPtr("CUSTOM"),
		}
		err = bothChoices.Validate()
		if err == nil {
			t.Error("InterestType1 with both choices should have validation error")
		}
	})

	// Test ServiceLevel8 validation
	t.Run("ServiceLevel8", func(t *testing.T) {
		// Valid case
		validService := ServiceLevel8{
			Code: stringPtr("SEPA"),
		}
		if err := validService.Validate(); err != nil {
			t.Errorf("Valid ServiceLevel8 should not have errors: %v", err)
		}

		// Invalid case - no choice
		noChoice := ServiceLevel8{}
		err := noChoice.Validate()
		if err == nil {
			t.Error("ServiceLevel8 with no choices should have validation error")
		}
	})

	// Test GenericIdentification30 validation
	t.Run("GenericIdentification30", func(t *testing.T) {
		// Valid case
		validGeneric := GenericIdentification30{
			ID:         "TEST",
			Issuer:     "Test Issuer",
			SchemeName: stringPtr("TestScheme"),
		}
		if err := validGeneric.Validate(); err != nil {
			t.Errorf("Valid GenericIdentification30 should not have errors: %v", err)
		}

		// Invalid case - ID too short
		invalidGeneric := GenericIdentification30{
			ID:     "TEL", // Only 3 characters, should be 4
			Issuer: "Test Issuer",
		}
		err := invalidGeneric.Validate()
		if err == nil {
			t.Error("GenericIdentification30 with wrong ID length should have validation error")
		}

		// Invalid case - missing Issuer
		missingIssuer := GenericIdentification30{
			ID: "TEST",
			// Missing Issuer
		}
		err = missingIssuer.Validate()
		if err == nil {
			t.Error("GenericIdentification30 with missing Issuer should have validation error")
		}
	})

	// Test Rate4 validation
	t.Run("Rate4", func(t *testing.T) {
		// Valid case
		validRate := Rate4{
			Rate: floatPtr(2.5),
		}
		if err := validRate.Validate(); err != nil {
			t.Errorf("Valid Rate4 should not have errors: %v", err)
		}

		// Invalid case - negative rate
		negativeRate := Rate4{
			Rate: floatPtr(-1.0),
		}
		err := negativeRate.Validate()
		if err == nil {
			t.Error("Rate4 with negative rate should have validation error")
		}
	})

	// Test BalanceType10 validation
	t.Run("BalanceType10", func(t *testing.T) {
		// Valid case
		validBalance := BalanceType10{
			Code: stringPtr("CLBD"),
		}
		if err := validBalance.Validate(); err != nil {
			t.Errorf("Valid BalanceType10 should not have errors: %v", err)
		}

		// Invalid case - no choice
		noChoice := BalanceType10{}
		err := noChoice.Validate()
		if err == nil {
			t.Error("BalanceType10 with no choices should have validation error")
		}
	})

	t.Log("All new validation functions are working correctly!")
}

func TestAdditionalValidationFunctions(t *testing.T) {
	// Test PaymentIdentification7 validation
	t.Run("PaymentIdentification7", func(t *testing.T) {
		// Valid case
		validPayment := PaymentIdentification7{
			EndToEndID:    "END2END123",
			TransactionID: stringPtr("TXN456"),
		}
		if err := validPayment.Validate(); err != nil {
			t.Errorf("Valid PaymentIdentification7 should not have errors: %v", err)
		}

		// Invalid case - EndToEndID too long
		invalidPayment := PaymentIdentification7{
			EndToEndID: "ThisEndToEndIDIsWayTooLongAndShouldFailValidation",
		}
		err := invalidPayment.Validate()
		if err == nil {
			t.Error("PaymentIdentification7 with too long EndToEndID should have validation error")
		}

		// Invalid case - missing required EndToEndID
		missingEndToEnd := PaymentIdentification7{}
		err = missingEndToEnd.Validate()
		if err == nil {
			t.Error("PaymentIdentification7 with missing EndToEndID should have validation error")
		}
	})

	// Test PartyIdentification135 validation
	t.Run("PartyIdentification135", func(t *testing.T) {
		// Valid case
		validParty := PartyIdentification135{
			Name: stringPtr("Test Party Name"),
			PostalAddress: &PostalAddress24{
				TownName: stringPtr("Test City"),
				Country:  stringPtr("US"),
			},
		}
		if err := validParty.Validate(); err != nil {
			t.Errorf("Valid PartyIdentification135 should not have errors: %v", err)
		}

		// Invalid case - name too long
		invalidParty := PartyIdentification135{
			Name: stringPtr("This is an extremely long party name that exceeds the maximum allowed length of 140 characters and should therefore fail validation when tested"),
		}
		err := invalidParty.Validate()
		if err == nil {
			t.Error("PartyIdentification135 with too long name should have validation error")
		}
	})

	// Test PostalAddress24 validation
	t.Run("PostalAddress24", func(t *testing.T) {
		// Valid case
		validAddress := PostalAddress24{
			StreetName:     stringPtr("123 Main Street"),
			BuildingNumber: stringPtr("123"),
			PostCode:       stringPtr("12345"),
			TownName:       stringPtr("Anytown"),
			Country:        stringPtr("US"),
		}
		if err := validAddress.Validate(); err != nil {
			t.Errorf("Valid PostalAddress24 should not have errors: %v", err)
		}

		// Invalid case - invalid country code
		invalidAddress := PostalAddress24{
			TownName: stringPtr("Test City"),
			Country:  stringPtr("INVALID"),
		}
		err := invalidAddress.Validate()
		if err == nil {
			t.Error("PostalAddress24 with invalid country code should have validation error")
		}
	})

	// Test FinancialInstitutionIdentification18 validation
	t.Run("FinancialInstitutionIdentification18", func(t *testing.T) {
		// Valid case
		validFI := FinancialInstitutionIdentification18{
			BankIdentifierCode:    stringPtr("CHASUS33"),
			LegalEntityIdentifier: stringPtr("12345678901234567890"),
			Name:                  stringPtr("Test Bank"),
		}
		if err := validFI.Validate(); err != nil {
			t.Errorf("Valid FinancialInstitutionIdentification18 should not have errors: %v", err)
		}

		// Invalid case - invalid BIC
		invalidFI := FinancialInstitutionIdentification18{
			BankIdentifierCode: stringPtr("INVALID"),
			Name:               stringPtr("Test Bank"),
		}
		err := invalidFI.Validate()
		if err == nil {
			t.Error("FinancialInstitutionIdentification18 with invalid BIC should have validation error")
		}
	})

	// Test GenericAccountIdentification1 validation
	t.Run("GenericAccountIdentification1", func(t *testing.T) {
		// Valid case
		validGeneric := GenericAccountIdentification1{
			ID:     "1234567890",
			Issuer: stringPtr("Test Issuer"),
		}
		if err := validGeneric.Validate(); err != nil {
			t.Errorf("Valid GenericAccountIdentification1 should not have errors: %v", err)
		}

		// Invalid case - ID too long
		invalidGeneric := GenericAccountIdentification1{
			ID: "ThisAccountIdentificationIsWayTooLong",
		}
		err := invalidGeneric.Validate()
		if err == nil {
			t.Error("GenericAccountIdentification1 with too long ID should have validation error")
		}

		// Invalid case - missing required ID
		missingID := GenericAccountIdentification1{
			Issuer: stringPtr("Test Issuer"),
		}
		err = missingID.Validate()
		if err == nil {
			t.Error("GenericAccountIdentification1 with missing ID should have validation error")
		}
	})

	t.Log("All additional validation functions are working correctly!")
}

func TestMoreValidationFunctions(t *testing.T) {
	// Test CashAccount38 validation
	t.Run("CashAccount38", func(t *testing.T) {
		// Valid case with IBAN
		validAccount := CashAccount38{
			ID: AccountIdentification4{
				IBAN: stringPtr("DE89370400440532013000"),
			},
			Currency: stringPtr("EUR"),
			Name:     stringPtr("Test Account"),
		}
		if err := validAccount.Validate(); err != nil {
			t.Errorf("Valid CashAccount38 should not have errors: %v", err)
		}

		// Invalid case - currency too short
		invalidAccount := CashAccount38{
			ID: AccountIdentification4{
				IBAN: stringPtr("DE89370400440532013000"),
			},
			Currency: stringPtr("EU"), // Too short
		}
		err := invalidAccount.Validate()
		if err == nil {
			t.Error("CashAccount38 with invalid currency should have validation error")
		}

		// Invalid case - no account identification choice
		noChoiceAccount := CashAccount38{
			ID:       AccountIdentification4{},
			Currency: stringPtr("USD"),
		}
		err = noChoiceAccount.Validate()
		if err == nil {
			t.Error("CashAccount38 with no ID choice should have validation error")
		}
	})

	// Test AccountIdentification4 validation
	t.Run("AccountIdentification4", func(t *testing.T) {
		// Valid case - IBAN only
		validIBAN := AccountIdentification4{
			IBAN: stringPtr("DE89370400440532013000"),
		}
		if err := validIBAN.Validate(); err != nil {
			t.Errorf("Valid IBAN AccountIdentification4 should not have errors: %v", err)
		}

		// Valid case - Other only
		validOther := AccountIdentification4{
			Other: &GenericAccountIdentification1{
				ID:     "12345",
				Issuer: stringPtr("Test Issuer"),
			},
		}
		if err := validOther.Validate(); err != nil {
			t.Errorf("Valid Other AccountIdentification4 should not have errors: %v", err)
		}

		// Invalid case - no choice
		noChoice := AccountIdentification4{}
		err := noChoice.Validate()
		if err == nil {
			t.Error("AccountIdentification4 with no choice should have validation error")
		}

		// Invalid case - both choices
		bothChoices := AccountIdentification4{
			IBAN:  stringPtr("DE89370400440532013000"),
			Other: &GenericAccountIdentification1{ID: "12345"},
		}
		err = bothChoices.Validate()
		if err == nil {
			t.Error("AccountIdentification4 with both choices should have validation error")
		}

		// Invalid case - IBAN too short
		shortIBAN := AccountIdentification4{
			IBAN: stringPtr("DE123"),
		}
		err = shortIBAN.Validate()
		if err == nil {
			t.Error("AccountIdentification4 with too short IBAN should have validation error")
		}
	})

	// Test PaymentTypeInfo28 validation
	t.Run("PaymentTypeInfo28", func(t *testing.T) {
		// Valid case
		validPaymentType := PaymentTypeInfo28{
			InstructionPriority: stringPtr("HIGH"),
			SequenceType:        stringPtr("FRST"),
		}
		if err := validPaymentType.Validate(); err != nil {
			t.Errorf("Valid PaymentTypeInfo28 should not have errors: %v", err)
		}

		// Invalid case - instruction priority too long
		invalidPriority := PaymentTypeInfo28{
			InstructionPriority: stringPtr("VERYHIGHPRIORITY"), // Too long
		}
		err := invalidPriority.Validate()
		if err == nil {
			t.Error("PaymentTypeInfo28 with too long priority should have validation error")
		}

		// Invalid case - sequence type too long
		invalidSequence := PaymentTypeInfo28{
			SequenceType: stringPtr("VERYLONGSEQUENCETYPE"), // Too long
		}
		err = invalidSequence.Validate()
		if err == nil {
			t.Error("PaymentTypeInfo28 with too long sequence type should have validation error")
		}
	})

	// Test CreditTransferTransaction39 validation
	t.Run("CreditTransferTransaction39", func(t *testing.T) {
		// Valid case
		validTx := CreditTransferTransaction39{
			PaymentID: PaymentIdentification7{
				EndToEndID: "END2END123",
			},
			InterbankSettlementAmount: ActiveCurrencyAndAmount{
				Value:    1000.00,
				Currency: "USD",
			},
			ChargeBearer: "SLEV",
			Debtor: PartyIdentification135{
				Name: stringPtr("Test Debtor"),
			},
			DebtorAgent: BranchAndFinancialInstitutionIdentification6{
				FinancialInstitutionID: FinancialInstitutionIdentification18{
					BankIdentifierCode: stringPtr("CHASUS33"),
				},
			},
			Creditor: PartyIdentification135{
				Name: stringPtr("Test Creditor"),
			},
			CreditorAgent: BranchAndFinancialInstitutionIdentification6{
				FinancialInstitutionID: FinancialInstitutionIdentification18{
					BankIdentifierCode: stringPtr("BOFAUS3N"),
				},
			},
		}
		if err := validTx.Validate(); err != nil {
			t.Errorf("Valid CreditTransferTransaction39 should not have errors: %v", err)
		}

		// Invalid case - missing charge bearer
		invalidTx := CreditTransferTransaction39{
			PaymentID: PaymentIdentification7{
				EndToEndID: "END2END123",
			},
			InterbankSettlementAmount: ActiveCurrencyAndAmount{
				Value:    1000.00,
				Currency: "USD",
			},
			ChargeBearer: "", // Missing required field
		}
		err := invalidTx.Validate()
		if err == nil {
			t.Error("CreditTransferTransaction39 with missing ChargeBearer should have validation error")
		}

		// Invalid case - charge bearer too long
		longChargeBearer := CreditTransferTransaction39{
			PaymentID: PaymentIdentification7{
				EndToEndID: "END2END123",
			},
			InterbankSettlementAmount: ActiveCurrencyAndAmount{
				Value:    1000.00,
				Currency: "USD",
			},
			ChargeBearer: "VERYLONGCHARGEBEARER", // Too long
		}
		err = longChargeBearer.Validate()
		if err == nil {
			t.Error("CreditTransferTransaction39 with too long ChargeBearer should have validation error")
		}
	})

	// Test FIToFICustomerCreditTransferV08 validation
	t.Run("FIToFICustomerCreditTransferV08", func(t *testing.T) {
		// Valid case
		validTransfer := FIToFICustomerCreditTransferV08{
			GroupHeader: GroupHeader93{
				MessageID:            "MSG123",
				CreationDateTime:     func() *time.Time { t := time.Now(); return &t }(),
				NumberOfTransactions: "1",
				SettlementInfo: SettlementInstruction7{
					SettlementMethod: "INDA",
				},
			},
			CreditTransferTransactionInfo: []CreditTransferTransaction39{
				{
					PaymentID: PaymentIdentification7{
						EndToEndID: "END2END123",
					},
					InterbankSettlementAmount: ActiveCurrencyAndAmount{
						Value:    1000.00,
						Currency: "USD",
					},
					ChargeBearer: "SLEV",
					Debtor: PartyIdentification135{
						Name: stringPtr("Test Debtor"),
					},
					DebtorAgent: BranchAndFinancialInstitutionIdentification6{
						FinancialInstitutionID: FinancialInstitutionIdentification18{
							BankIdentifierCode: stringPtr("CHASUS33"),
						},
					},
					Creditor: PartyIdentification135{
						Name: stringPtr("Test Creditor"),
					},
					CreditorAgent: BranchAndFinancialInstitutionIdentification6{
						FinancialInstitutionID: FinancialInstitutionIdentification18{
							BankIdentifierCode: stringPtr("BOFAUS3N"),
						},
					},
				},
			},
		}
		if err := validTransfer.Validate(); err != nil {
			t.Errorf("Valid FIToFICustomerCreditTransferV08 should not have errors: %v", err)
		}

		// Invalid case - no transactions
		noTransactions := FIToFICustomerCreditTransferV08{
			GroupHeader: GroupHeader93{
				MessageID:            "MSG123",
				CreationDateTime:     func() *time.Time { t := time.Now(); return &t }(),
				NumberOfTransactions: "0",
				SettlementInfo: SettlementInstruction7{
					SettlementMethod: "INDA",
				},
			},
			CreditTransferTransactionInfo: []CreditTransferTransaction39{},
		}
		err := noTransactions.Validate()
		if err == nil {
			t.Error("FIToFICustomerCreditTransferV08 with no transactions should have validation error")
		}
	})

	t.Log("All new validation functions are working correctly!")
}

// Helper function to create float pointers
func floatPtr(f float64) *float64 {
	return &f
}