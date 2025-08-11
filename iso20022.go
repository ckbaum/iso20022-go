package iso20022

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// PACS.008.001.08 - FI to FI Customer Credit Transfer
// Pacs00800108Document represents the PACS.008.001.08 Financial Institution to Financial Institution Customer Credit Transfer message.
// This message is used by financial institutions to transfer funds on behalf of customers between different institutions,
// containing all necessary payment details including debtor/creditor information and settlement instructions.
type Pacs00800108Document struct {
	XMLName                  xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08 Document"`
	FICustomerCreditTransfer FIToFICustomerCreditTransferV08 `xml:"FIToFICstmrCdtTrf"`
}

// Pacs00900108Document represents the PACS.009.001.08 Financial Institution Credit Transfer message.
// This message is used for inter-bank credit transfers between financial institutions,
// typically for settlement purposes and institutional fund movements.
type Pacs00900108Document struct {
	XMLName          xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Document"`
	FICreditTransfer FinancialInstitutionCreditTransferV08 `xml:"FICdtTrf"`
}

// Pacs00200110Document represents the PACS.002.001.10 Financial Institution to Financial Institution Payment Status Report.
// This message provides status updates for payment instructions between financial institutions,
// reporting successful processing, rejections, or pending status with detailed reason codes.
type Pacs00200110Document struct {
	XMLName               xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Document"`
	FIPaymentStatusReport FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt"`
}

// Pacs00400110Document represents the PACS.004.001.10 Payment Return message.
// This message is used by financial institutions to return previously processed payments,
// typically due to insufficient funds, incorrect account details, or other processing issues.
type Pacs00400110Document struct {
	XMLName       xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Document"`
	PaymentReturn PaymentReturnV10 `xml:"PmtRtr"`
}

// Pacs02800103Document represents the PACS.028.001.03 Financial Institution to Financial Institution Payment Status Request.
// This message allows financial institutions to request status information about previously sent payments,
// enabling tracking and reconciliation of payment instructions in the clearing and settlement process.
type Pacs02800103Document struct {
	XMLName                xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 Document"`
	FIPaymentStatusRequest FIToFIPaymentStatusRequestV03 `xml:"FIToFIPmtStsReq"`
}

// Camt05200108Document represents the CAMT.052.001.08 Bank to Customer Account Report message.
// This message provides customers with account balance information and transaction summaries,
// enabling account monitoring and cash management for corporate and institutional clients.
type Camt05200108Document struct {
	XMLName           xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 Document"`
	BankAccountReport BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
}

// Camt05400108Document represents the CAMT.054.001.08 Bank to Customer Debit Credit Notification message.
// This message notifies customers of individual credit or debit entries posted to their accounts,
// providing detailed transaction information for reconciliation and cash management purposes.
type Camt05400108Document struct {
	XMLName                     xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.08 Document"`
	BankDebitCreditNotification BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
}

// Camt05500109Document represents the CAMT.055.001.09 Customer Payment Cancellation Request message.
// This message allows customers to request cancellation of previously submitted payment instructions,
// providing justification and reference details for the cancellation request.
type Camt05500109Document struct {
	XMLName                      xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.055.001.09 Document"`
	CustomerPaymentCancelRequest CustomerPaymentCancellationRequestV09 `xml:"CstmrPmtCxlReq"`
}

// Camt05600108Document represents the CAMT.056.001.08 Financial Institution to Financial Institution Payment Cancellation Request.
// This message enables financial institutions to request payment cancellations from other institutions,
// typically used for recall of pacs.008 messages with proper justification and reason codes.
type Camt05600108Document struct {
	XMLName                xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Document"`
	FIPaymentCancelRequest FIToFIPaymentCancellationRequestV08 `xml:"FIToFIPmtCxlReq"`
}

// Camt06000105Document represents the CAMT.060.001.05 Account Reporting Request message.
// This message allows customers to request account information and reports from their banks,
// specifying the type of report, date range, and level of detail required.
type Camt06000105Document struct {
	XMLName                 xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05 Document"`
	AccountReportingRequest AccountReportingRequestV05 `xml:"AcctRptgReq"`
}

// Camt02600107Document represents the CAMT.026.001.07 Unable To Apply message.
// This message is used when a financial institution cannot process or apply a received instruction,
// providing detailed information about the reason for non-processing and any corrective actions needed.
type Camt02600107Document struct {
	XMLName       xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Document"`
	UnableToApply UnableToApplyV07 `xml:"UblToApply"`
}

// Camt02800109Document represents the CAMT.028.001.09 Additional Payment Info message.
// This message provides supplementary information related to payments that could not be included
// in the original payment instruction, supporting enhanced payment processing and reconciliation.
type Camt02800109Document struct {
	XMLName               xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.09 Document"`
	AdditionalPaymentInfo AdditionalPaymentInfoV09 `xml:"AddtlPmtInf"`
}

// Camt02900109Document represents the CAMT.029.001.09 Resolution of Investigation message.
// This message communicates the final outcome and resolution of payment investigations
// between financial institutions, providing closure to exception handling processes.
type Camt02900109Document struct {
	XMLName                 xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Document"`
	InvestigationResolution ResolutionOfInvestigationV09 `xml:"RsltnOfInvstgtn"`
}

// Pain01300107Document represents the PAIN.013.001.07 Creditor Payment Activation Request message.
// This message allows creditors to request payment activation from debtors,
// commonly used for direct debit scenarios and electronic invoice presentment.
type Pain01300107Document struct {
	XMLName                          xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07 Document"`
	CreditorPaymentActivationRequest CreditorPaymentActivationRequestV07 `xml:"CdtrPmtActvtnReq"`
}

// Pain01400107Document represents the PAIN.014.001.07 Creditor Payment Activation Request Status Report message.
// This message provides status updates on creditor payment activation requests,
// indicating acceptance, rejection, or processing status of payment activation requests.
type Pain01400107Document struct {
	XMLName                               xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.07 Document"`
	CreditorPaymentActivationStatusReport CreditorPaymentActivationRequestStatusReportV07 `xml:"CdtrPmtActvtnReqStsRpt"`
}

// Admi00400102Document represents the ADMI.004.001.02 System Event Notification message.
// This administrative message notifies participants of system events such as
// maintenance windows, system availability changes, or operational status updates.
type Admi00400102Document struct {
	XMLName                 xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 Document"`
	SystemEventNotification SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
}

// Admi01100101Document represents the ADMI.011.001.01 System Event Acknowledgement message.
// This administrative message acknowledges receipt of system event notifications,
// confirming that participants have received and understood system status changes.
type Admi01100101Document struct {
	XMLName                    xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 Document"`
	SystemEventAcknowledgement SystemEventAcknowledgementV01 `xml:"SysEvtAck"`
}

// Admi00600101Document represents the ADMI.006.001.01 Resend Request message.
// This administrative message allows participants to request retransmission
// of previously sent messages when original messages were not received or processed correctly.
type Admi00600101Document struct {
	XMLName       xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.006.001.01 Document"`
	ResendRequest ResendRequestV01 `xml:"RsndReq"`
}

// Admi00700101Document represents the ADMI.007.001.01 Receipt Acknowledgement message.
// This administrative message acknowledges the successful receipt of messages,
// providing confirmation that transmitted messages have been properly received and processed.
type Admi00700101Document struct {
	XMLName                xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01 Document"`
	ReceiptAcknowledgement ReceiptAcknowledgementV01 `xml:"RctAck"`
}

// Admi99800102Document represents the ADMI.998.001.02 Administration Proprietary Message.
// This administrative message allows transmission of proprietary or custom administrative information
// between financial institutions that falls outside standard ISO 20022 message types.
type Admi99800102Document struct {
	XMLName               xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:admi.998.001.02 Document"`
	AdministrationMessage AdministrationProprietaryMessageV02 `xml:"AdmstnPrtryMsg"`
}

// FIToFICustomerCreditTransferV08 represents the core structure of a PACS.008.001.08 message.
// This structure contains the group header with message-level information and multiple
// credit transfer transaction details for inter-bank customer payment processing.
type FIToFICustomerCreditTransferV08 struct {
	GroupHeader                   GroupHeader93                 `xml:"GrpHdr"`
	CreditTransferTransactionInfo []CreditTransferTransaction39 `xml:"CdtTrfTxInf"`
	SupplementaryData             []SupplementaryData           `xml:"SplmtryData,omitempty"`
}

// GroupHeader93 contains message-level information that applies to all transactions within a PACS.008 message.
// It includes message identification, creation timestamp, settlement information, and agent details
// that are common across all credit transfer transactions in the message batch.
type GroupHeader93 struct {
	MessageID                      string                                        `xml:"MsgId"`
	CreationDateTime               *time.Time                                    `xml:"CreDtTm,omitempty"`
	BatchBooking                   *bool                                         `xml:"BtchBookg,omitempty"`
	NumberOfTransactions           string                                        `xml:"NbOfTxs"`
	ControlSum                     *float64                                      `xml:"CtrlSum,omitempty"`
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt,omitempty"`
	InterbankSettlementDate        *string                                       `xml:"IntrBkSttlmDt,omitempty"`
	SettlementInfo                 SettlementInstruction7                        `xml:"SttlmInf"`
	PaymentTypeInfo                *PaymentTypeInfo28                            `xml:"PmtTpInf,omitempty"`
	InstructingAgent               *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent                *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
}

// CreditTransferTransaction39 contains the detailed information for an individual credit transfer transaction.
// This includes payment identification, settlement amounts, charge information, and complete
// debtor/creditor party details required for processing inter-bank customer credit transfers.
type CreditTransferTransaction39 struct {
	PaymentID                        PaymentIdentification7                        `xml:"PmtId"`
	PaymentTypeInfo                  *PaymentTypeInfo28                            `xml:"PmtTpInf,omitempty"`
	InterbankSettlementAmount        ActiveCurrencyAndAmount                       `xml:"IntrBkSttlmAmt"`
	InterbankSettlementDate          *string                                       `xml:"IntrBkSttlmDt,omitempty"`
	SettlementPriority               *string                                       `xml:"SttlmPrty,omitempty"`
	SettlementTimeIndication         *SettlementDateTimeIndication                 `xml:"SttlmTmIndctn,omitempty"`
	SettlementTimeRequest            *SettlementTimeRequest                        `xml:"SttlmTmReq,omitempty"`
	AcceptanceDateTime               *time.Time                                    `xml:"AccptncDtTm,omitempty"`
	PoolingAdjustmentDate            *string                                       `xml:"PoolgAdjstmntDt,omitempty"`
	InstructedAmount                 *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty"`
	ExchangeRate                     *float64                                      `xml:"XchgRate,omitempty"`
	ChargeBearer                     string                                        `xml:"ChrgBr"`
	ChargesInfo                      []Charges7                                    `xml:"ChrgsInf,omitempty"`
	PreviousInstructingAgent1        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty"`
	PreviousInstructingAgent1Account *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty"`
	PreviousInstructingAgent2        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty"`
	PreviousInstructingAgent2Account *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty"`
	PreviousInstructingAgent3        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty"`
	PreviousInstructingAgent3Account *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty"`
	InstructingAgent                 *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent                  *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	IntermediaryAgent1               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntermediaryAgent1Account        *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntermediaryAgent2               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntermediaryAgent2Account        *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntermediaryAgent3               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	IntermediaryAgent3Account        *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty"`
	UltimateDebtor                   *PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	InitiatingParty                  *PartyIdentification135                       `xml:"InitgPty,omitempty"`
	Debtor                           PartyIdentification135                        `xml:"Dbtr"`
	DebtorAccount                    *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent                      BranchAndFinancialInstitutionIdentification6  `xml:"DbtrAgt"`
	DebtorAgentAccount               *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CreditorAgent                    BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	CreditorAgentAccount             *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Creditor                         PartyIdentification135                        `xml:"Cdtr"`
	CreditorAccount                  *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltimateCreditor                 *PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstructionsForCreditorAgent     []InstructionForCreditorAgent                 `xml:"InstrForCdtrAgt,omitempty"`
	InstructionsForNextAgent         []InstructionForNextAgent                     `xml:"InstrForNxtAgt,omitempty"`
	Purpose                          *Purpose                                      `xml:"Purp,omitempty"`
	RegulatoryReporting              []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty"` // max 10 elements
	Tax                              *TaxInfo                                      `xml:"Tax,omitempty"`
	RelatedRemittanceInfo            []RemittanceLocation                          `xml:"RltdRmtInf,omitempty"` // max 10 elements
	RemittanceInfo                   *RemittanceInfo                               `xml:"RmtInf,omitempty"`
	SupplementaryData                []SupplementaryData                           `xml:"SplmtryData,omitempty"`
}

// PaymentIdentification7 contains unique identifiers for payment transactions.
// This includes instruction ID, end-to-end ID for tracking, transaction ID,
// and UETR (Unique End-to-end Transaction Reference) for global payment traceability.
type PaymentIdentification7 struct {
	InstructionID           *string `xml:"InstrId,omitempty"`
	EndToEndID              string  `xml:"EndToEndId"`
	TransactionID           *string `xml:"TxId,omitempty"`
	UETR                    *string `xml:"UETR,omitempty"`
	ClearingSystemReference *string `xml:"ClrSysRef,omitempty"`
}

// PaymentIdentification provides legacy payment identification structure for backward compatibility.
// Contains basic payment identifiers including instruction ID, end-to-end ID, transaction ID and UETR.
type PaymentIdentification struct {
	InstructionID *string `xml:"InstrId,omitempty"`
	EndToEndID    string  `xml:"EndToEndId"`
	TransactionID string  `xml:"TxId"`
	UETR          *string `xml:"UETR,omitempty"`
}

// PaymentTypeInfo specifies the type and priority of payment processing instructions.
// Includes instruction priority, clearing channel, service level, local instruments and category purpose
// to guide how the payment should be processed by financial institutions.
type PaymentTypeInfo struct {
	InstructionPriority *string          `xml:"InstrPrty,omitempty"`
	ClearingChannel     *string          `xml:"ClrChanl,omitempty"`
	ServiceLevel        []ServiceLevel   `xml:"SvcLvl,omitempty"`
	LocalInstrument     *LocalInstrument `xml:"LclInstrm,omitempty"`
	SequenceType        *string          `xml:"SeqTp,omitempty"`
	CategoryPurpose     *CategoryPurpose `xml:"CtgyPurp,omitempty"`
}

// ActiveCurrencyAndAmount represents a monetary amount with an active currency code.
// Used throughout ISO 20022 messages to specify settlement amounts, fees, and other monetary values
// with their corresponding three-character ISO currency codes.
type ActiveCurrencyAndAmount struct {
	Value    float64 `xml:",chardata"`
	Currency string  `xml:"Ccy,attr"`
}

// ActiveOrHistoricCurrencyAndAmount represents a monetary amount with active or historic currency.
// Similar to ActiveCurrencyAndAmount but allows for historic currencies that are no longer in active use,
// supporting legacy transactions and reporting requirements.
type ActiveOrHistoricCurrencyAndAmount struct {
	Value    float64 `xml:",chardata"`
	Currency string  `xml:"Ccy,attr"`
}

// BranchAndFinancialInstitutionIdentification provides identification for financial institutions and their branches.
// Contains the core institution identification plus optional branch details for routing
// and processing payments through specific institutional locations.
type BranchAndFinancialInstitutionIdentification struct {
	FinancialInstitutionID FinancialInstitutionIdentification `xml:"FinInstnId"`
	BranchID               *BranchData                        `xml:"BrnchId,omitempty"`
}

// FinancialInstitutionIdentification contains comprehensive identification details for financial institutions.
// Includes BIC codes, clearing system member IDs, LEI, institution name, postal address
// and other identifiers used for routing and settlement purposes.
type FinancialInstitutionIdentification struct {
	BankIdentifierCode     *string                             `xml:"BICFI,omitempty"`
	ClearingSystemMemberID *ClearingSystemMemberIdentification `xml:"ClrSysMmbId,omitempty"`
	LegalEntityIdentifier  *string                             `xml:"LEI,omitempty"`
	Name                   *string                             `xml:"Nm,omitempty"`
	PostalAddress          *PostalAddress                      `xml:"PstlAdr,omitempty"`
	Other                  *GenericFinancialIdentification     `xml:"Othr,omitempty"`
}

// BranchData contains detailed information about a specific branch of a financial institution.
// Includes branch identifier, LEI, name and postal address for precise routing
// of payments to specific institutional locations.
type BranchData struct {
	ID                    *string        `xml:"Id,omitempty"`
	LegalEntityIdentifier *string        `xml:"LEI,omitempty"`
	Name                  *string        `xml:"Nm,omitempty"`
	PostalAddress         *PostalAddress `xml:"PstlAdr,omitempty"`
}

// BranchAndFinancialInstitutionIdentification6 provides PACS.008.001.08 specific institution identification.
// This version matches the exact XSD schema requirements for the pacs.008 message format,
// ensuring compliance with the specific version's data model requirements.
type BranchAndFinancialInstitutionIdentification6 struct {
	FinancialInstitutionID FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BranchID               *BranchData3                         `xml:"BrnchId,omitempty"`
}

// FinancialInstitutionIdentification18 contains pacs.008.001.08 specific institution identification details.
// Matches the exact XSD schema structure for financial institution identification
// within the PACS.008.001.08 message format specification.
type FinancialInstitutionIdentification18 struct {
	BankIdentifierCode     *string                             `xml:"BICFI,omitempty"`
	ClearingSystemMemberID *ClearingSystemMemberIdentification `xml:"ClrSysMmbId,omitempty"`
	LegalEntityIdentifier  *string                             `xml:"LEI,omitempty"`
	Name                   *string                             `xml:"Nm,omitempty"`
	PostalAddress          *PostalAddress                      `xml:"PstlAdr,omitempty"`
	Other                  *GenericFinancialIdentification     `xml:"Othr,omitempty"`
}

// BranchData3 contains PACS.008.001.08 specific branch information.
// This structure matches the exact XSD schema requirements for branch data
// within the pacs.008.001.08 message format.
type BranchData3 struct {
	ID                    *string        `xml:"Id,omitempty"`
	LegalEntityIdentifier *string        `xml:"LEI,omitempty"`
	Name                  *string        `xml:"Nm,omitempty"`
	PostalAddress         *PostalAddress `xml:"PstlAdr,omitempty"`
}

// PaymentTypeInfo28 provides PACS.008.001.08 specific payment type details.
// Contains instruction priority, service level, local instrument, sequence type and category purpose
// as defined by the pacs.008.001.08 XSD schema specification.
type PaymentTypeInfo28 struct {
	InstructionPriority *string          `xml:"InstrPrty,omitempty"`
	ServiceLevel        []ServiceLevel   `xml:"SvcLvl,omitempty"`
	LocalInstrument     *LocalInstrument `xml:"LclInstrm,omitempty"`
	SequenceType        *string          `xml:"SeqTp,omitempty"`
	CategoryPurpose     *CategoryPurpose `xml:"CtgyPurp,omitempty"`
}

// PartyIdentification135 contains PACS.008.001.08 specific party identification information.
// Provides comprehensive identification details for parties involved in credit transfer transactions,
// including name, postal address, identification details and contact information.
type PartyIdentification135 struct {
	Name               *string          `xml:"Nm,omitempty"`
	PostalAddress      *PostalAddress24 `xml:"PstlAdr,omitempty"`
	ID                 *Party38         `xml:"Id,omitempty"`
	CountryOfResidence *string          `xml:"CtryOfRes,omitempty"`
	ContactDetails     *Contact4        `xml:"CtctDtls,omitempty"`
}

// PostalAddress24 contains comprehensive postal address information for PACS.008.001.08 messages.
// Provides detailed address components including street, building, postal code, town, district,
// country and additional address lines for precise party location identification.
type PostalAddress24 struct {
	AddressType        *string  `xml:"AdrTp,omitempty"`
	Department         *string  `xml:"Dept,omitempty"`
	SubDepartment      *string  `xml:"SubDept,omitempty"`
	StreetName         *string  `xml:"StrtNm,omitempty"`
	BuildingNumber     *string  `xml:"BldgNb,omitempty"`
	BuildingName       *string  `xml:"BldgNm,omitempty"`
	Floor              *string  `xml:"Flr,omitempty"`
	PostBox            *string  `xml:"PstBx,omitempty"`
	Room               *string  `xml:"Room,omitempty"`
	PostCode           *string  `xml:"PstCd,omitempty"`
	TownName           *string  `xml:"TwnNm,omitempty"`
	TownLocationName   *string  `xml:"TwnLctnNm,omitempty"`
	DistrictName       *string  `xml:"DstrctNm,omitempty"`
	CountrySubDivision *string  `xml:"CtrySubDvsn,omitempty"`
	Country            *string  `xml:"Ctry,omitempty"`
	AddressLine        []string `xml:"AdrLine,omitempty"`
}

// Party38 contains party identification details for PACS.008.001.08 messages.
// Supports both organization identification (for corporate entities) and private identification
// (for individual persons) as defined by the pacs.008.001.08 XSD schema.
type Party38 struct {
	OrganizationID *OrganizationIdentification29 `xml:"OrgId,omitempty"`
	PrivateID      *PersonIdentification13       `xml:"PrvtId,omitempty"`
}

// Contact4 contains comprehensive contact information for parties in PACS.008.001.08 messages.
// Includes name, various communication methods (phone, mobile, fax, email), job details,
// department information and preferred communication methods for party contacts.
type Contact4 struct {
	NamePrefix      *string         `xml:"NmPrfx,omitempty"`
	Name            *string         `xml:"Nm,omitempty"`
	PhoneNumber     *string         `xml:"PhneNb,omitempty"`
	MobileNumber    *string         `xml:"MobNb,omitempty"`
	FaxNumber       *string         `xml:"FaxNb,omitempty"`
	EmailAddress    *string         `xml:"EmailAdr,omitempty"`
	EmailPurpose    *string         `xml:"EmailPurp,omitempty"`
	JobTitle        *string         `xml:"JobTitl,omitempty"`
	Responsibility  *string         `xml:"Rspnsblty,omitempty"`
	Department      *string         `xml:"Dept,omitempty"`
	Other           []OtherContact1 `xml:"Othr,omitempty"`
	PreferredMethod *string         `xml:"PrefrdMtd,omitempty"`
}

// CashAccount38 contains account identification and details for PACS.008.001.08 messages.
// Provides comprehensive account information including ID, type, currency, name and proxy details
// required for debtor and creditor account specification in credit transfer transactions.
type CashAccount38 struct {
	ID       AccountIdentification4       `xml:"Id"`
	Type     *CashAccountType2Choice      `xml:"Tp,omitempty"`
	Currency *string                      `xml:"Ccy,omitempty"`
	Name     *string                      `xml:"Nm,omitempty"`
	Proxy    *ProxyAccountIdentification1 `xml:"Prxy,omitempty"`
}

// AccountIdentification4 provides account identification using IBAN or other account identifiers.
// Supports both standardized IBAN format and proprietary account identification schemes
// for flexible account specification in payment messages.
type AccountIdentification4 struct {
	IBAN  *string                        `xml:"IBAN,omitempty"`
	Other *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// CashAccountType2 specifies the type of cash account using standard codes or proprietary values.
// Allows classification of accounts (current, savings, etc.) for proper payment processing
// and regulatory compliance requirements.
type CashAccountType2 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// ProxyAccountIdentification1 contains proxy account identification information.
// Enables account identification through proxy mechanisms such as mobile phone numbers,
// email addresses or other alternative identifiers for modern payment systems.
type ProxyAccountIdentification1 struct {
	Type *ProxyAccountType1 `xml:"Tp,omitempty"`
	ID   string             `xml:"Id"`
}

// ProxyAccountType1 specifies the type of proxy account identifier being used.
// Defines the format and nature of the proxy identifier (phone, email, etc.)
// to ensure correct interpretation by receiving systems.
type ProxyAccountType1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// GenericAccountIdentification1 provides generic account identification with custom schemes.
// Allows flexible account identification using proprietary or non-standard numbering schemes
// with optional scheme name and issuer information for context.
type GenericAccountIdentification1 struct {
	ID         string              `xml:"Id"`
	SchemeName *AccountSchemeName1 `xml:"SchmeNm,omitempty"`
	Issuer     *string             `xml:"Issr,omitempty"`
}

// AccountSchemeName1 specifies the naming scheme used for account identification.
// Provides both standardized codes and proprietary scheme names to describe
// the format and interpretation of account identifiers.
type AccountSchemeName1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// OrganizationIdentification29 contains identification details for organizational entities.
// Includes BIC codes, Legal Entity Identifiers (LEI) and other organizational identifiers
// required for compliance and routing purposes in financial transactions.
type OrganizationIdentification29 struct {
	AnyBankIdentifierCode *string                              `xml:"AnyBIC,omitempty"`
	LegalEntityIdentifier *string                              `xml:"LEI,omitempty"`
	Other                 []GenericOrganizationIdentification1 `xml:"Othr,omitempty"`
}

// GenericOrganizationIdentification1 provides flexible organizational identification.
// Supports custom identification schemes for organizations that may not have
// standard BIC or LEI identifiers, with scheme name and issuer context.
type GenericOrganizationIdentification1 struct {
	ID         string                                 `xml:"Id"`
	SchemeName *OrganizationIdentificationSchemeName1 `xml:"SchmeNm,omitempty"`
	Issuer     *string                                `xml:"Issr,omitempty"`
}

type OrganizationIdentificationSchemeName1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type PersonIdentification13 struct {
	DateAndPlaceOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty"`
	Other               []GenericPersonIdentification2 `xml:"Othr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDate       *string `xml:"BirthDt,omitempty"`
	ProvinceOfBirth *string `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth     string  `xml:"CityOfBirth"`
	CountryOfBirth  string  `xml:"CtryOfBirth"`
}

type GenericPersonIdentification2 struct {
	ID         string                           `xml:"Id"`
	SchemeName *PersonIdentificationSchemeName2 `xml:"SchmeNm,omitempty"`
	Issuer     *string                          `xml:"Issr,omitempty"`
}

type PersonIdentificationSchemeName2 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type OtherContact1 struct {
	ChannelType string  `xml:"ChanlTp"`
	ID          *string `xml:"Id,omitempty"`
}

// Charges7 for pacs.008.001.08 (exact XSD match)
type Charges7 struct {
	Amount ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agent  BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// SettlementInstruction7 for pacs.008.001.08 (exact XSD match)
type SettlementInstruction7 struct {
	SettlementMethod                     string                                        `xml:"SttlmMtd"`
	SettlementAccount                    *CashAccount                                  `xml:"SttlmAcct,omitempty"`
	ClearingSystem                       *ClearingSystemIdentificationSecondary        `xml:"ClrSys,omitempty"`
	InstructingReimbursementAgent        *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstructingReimbursementAgentAccount *CashAccount                                  `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstructedReimbursementAgent         *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstructedReimbursementAgentAccount  *CashAccount                                  `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThirdReimbursementAgent              *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThirdReimbursementAgentAccount       *CashAccount                                  `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// CashAccount represents a cash account
type CashAccount struct {
	ID       AccountIdentification       `xml:"Id"`
	Type     *CashAccountType            `xml:"Tp,omitempty"`
	Currency *string                     `xml:"Ccy,omitempty"`
	Name     *string                     `xml:"Nm,omitempty"`
	Proxy    *ProxyAccountIdentification `xml:"Prxy,omitempty"`
}

// AccountIdentification contains account identification choices
type AccountIdentification struct {
	IBAN  *string                       `xml:"IBAN,omitempty"`
	Other *GenericAccountIdentification `xml:"Othr,omitempty"`
}

// PartyIdentification contains party identification information
type PartyIdentification struct {
	Name               *string        `xml:"Nm,omitempty"`
	PostalAddress      *PostalAddress `xml:"PstlAdr,omitempty"`
	ID                 *Party         `xml:"Id,omitempty"`
	CountryOfResidence *string        `xml:"CtryOfRes,omitempty"`
	ContactDetails     *Contact       `xml:"CtctDtls,omitempty"`
}

// Party contains party identification choices
type Party struct {
	OrganizationID *OrganizationIdentification `xml:"OrgId,omitempty"`
	PrivateID      *PersonIdentification       `xml:"PrvtId,omitempty"`
}

// PostalAddress contains postal address information
type PostalAddress struct {
	AddressType        *string  `xml:"AdrTp,omitempty"`
	Department         *string  `xml:"Dept,omitempty"`
	SubDepartment      *string  `xml:"SubDept,omitempty"`
	StreetName         *string  `xml:"StrtNm,omitempty"`
	BuildingNumber     *string  `xml:"BldgNb,omitempty"`
	BuildingName       *string  `xml:"BldgNm,omitempty"`
	Floor              *string  `xml:"Flr,omitempty"`
	PostBox            *string  `xml:"PstBx,omitempty"`
	Room               *string  `xml:"Room,omitempty"`
	PostalCode         *string  `xml:"PstCd,omitempty"`
	TownName           *string  `xml:"TwnNm,omitempty"`
	TownLocationName   *string  `xml:"TwnLctnNm,omitempty"`
	DistrictName       *string  `xml:"DstrctNm,omitempty"`
	CountrySubDivision *string  `xml:"CtrySubDvsn,omitempty"`
	Country            *string  `xml:"Ctry,omitempty"`
	AddressLines       []string `xml:"AdrLine,omitempty"`
}

// Charges contains charge information
type Charges struct {
	Amount ActiveOrHistoricCurrencyAndAmount           `xml:"Amt"`
	Agent  BranchAndFinancialInstitutionIdentification `xml:"Agt"`
}

// TaxInfo contains tax information
type TaxInfo struct {
	Creditor               *TaxPartyCreditor                  `xml:"Cdtr,omitempty"`
	Debtor                 *TaxPartyDebtor                    `xml:"Dbtr,omitempty"`
	UltimateDebtor         *TaxPartyDebtor                    `xml:"UltmtDbtr,omitempty"`
	AdministrationZone     *string                            `xml:"AdmstnZone,omitempty"`
	ReferenceNumber        *string                            `xml:"RefNb,omitempty"`
	Method                 *string                            `xml:"Mtd,omitempty"`
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TotalTaxAmount         *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Date                   *string                            `xml:"Dt,omitempty"`
	SequenceNumber         *float64                           `xml:"SeqNb,omitempty"`
	Record                 []TaxRecord                        `xml:"Rcrd,omitempty"`
}

// TaxPartyCreditor contains creditor tax party information
type TaxPartyCreditor struct {
	TaxID          *string `xml:"TaxId,omitempty"`
	RegistrationID *string `xml:"RegnId,omitempty"`
	TaxType        *string `xml:"TaxTp,omitempty"`
}

// TaxPartyDebtor contains debtor tax party information
type TaxPartyDebtor struct {
	TaxID          *string           `xml:"TaxId,omitempty"`
	RegistrationID *string           `xml:"RegnId,omitempty"`
	TaxType        *string           `xml:"TaxTp,omitempty"`
	Authorization  *TaxAuthorization `xml:"Authstn,omitempty"`
}

// RemittanceLocation contains remittance location information
type RemittanceLocation struct {
	RemittanceID                        *string                  `xml:"RmtId,omitempty"`
	RemittanceLocationDetails           []RemittanceLocationData `xml:"RmtLctnDtls,omitempty"`
	RemittanceLocationElectronicAddress *string                  `xml:"RmtLctnElctrncAdr,omitempty"`
	RemittanceLocationPostalAddress     *NameAndAddress          `xml:"RmtLctnPstlAdr,omitempty"`
}

// RemittanceInfo contains remittance information
type RemittanceInfo struct {
	Unstructured []string                   `xml:"Ustrd,omitempty"`
	Structured   []StructuredRemittanceInfo `xml:"Strd,omitempty"`
}

// SupplementaryData contains supplementary data
type SupplementaryData struct {
	PlaceAndName *string                   `xml:"PlcAndNm,omitempty"`
	Envelope     SupplementaryDataEnvelope `xml:"Envlp"`
}

// SupplementaryDataEnvelope contains the envelope for supplementary data
type SupplementaryDataEnvelope struct {
	// Any XML content can go here - this is intentionally flexible for extensibility
	Content string `xml:",innerxml"`
}

// SupplementaryData1 - XSD-specific version
type SupplementaryData1 struct {
	PlaceAndName *string                    `xml:"PlcAndNm,omitempty"`
	Envelope     SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Content string `xml:",innerxml"`
}

// Additional supporting types for pacs.009.001.08

type SettlementDateTimeIndication1 struct {
	DebitDateTime  *time.Time `xml:"DbtDtTm,omitempty"`
	CreditDateTime *time.Time `xml:"CdtDtTm,omitempty"`
}

type SettlementTimeRequest2 struct {
	ContinuousLinkedSettlementTime *time.Time `xml:"CLSTm,omitempty"`
	TillTime                       *time.Time `xml:"TillTm,omitempty"`
	FromTime                       *time.Time `xml:"FrTm,omitempty"`
	RejectTime                     *time.Time `xml:"RjctTm,omitempty"`
}

type InstructionForCreditorAgent2 struct {
	Code            *string `xml:"Cd,omitempty"`
	InstructionInfo *string `xml:"InstrInf,omitempty"`
}

type InstructionForNextAgent1 struct {
	Code            *string `xml:"Cd,omitempty"`
	InstructionInfo *string `xml:"InstrInf,omitempty"`
}

type Purpose2 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type RemittanceInfo2 struct {
	Unstructured []string `xml:"Ustrd,omitempty"`
}

type CreditTransferTransaction37 struct {
	UltimateDebtor                   *PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	InitiatingParty                  *PartyIdentification135                       `xml:"InitgPty,omitempty"`
	Debtor                           PartyIdentification135                        `xml:"Dbtr"`
	DebtorAccount                    *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent                      BranchAndFinancialInstitutionIdentification6  `xml:"DbtrAgt"`
	DebtorAgentAccount               *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	PreviousInstructingAgent1        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty"`
	PreviousInstructingAgent1Account *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty"`
	PreviousInstructingAgent2        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty"`
	PreviousInstructingAgent2Account *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty"`
	PreviousInstructingAgent3        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty"`
	PreviousInstructingAgent3Account *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty"`
	IntermediaryAgent1               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntermediaryAgent1Account        *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntermediaryAgent2               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntermediaryAgent2Account        *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntermediaryAgent3               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	IntermediaryAgent3Account        *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty"`
	CreditorAgent                    BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	CreditorAgentAccount             *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Creditor                         PartyIdentification135                        `xml:"Cdtr"`
	CreditorAccount                  *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltimateCreditor                 *PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstructionForCreditorAgent      []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty"`
	InstructionForNextAgent          []InstructionForNextAgent1                    `xml:"InstrForNxtAgt,omitempty"`
	Tax                              *TaxInfo8                                     `xml:"Tax,omitempty"`
	RemittanceInfo                   *RemittanceInfo16                             `xml:"RmtInf,omitempty"`
	InstructedAmount                 *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty"`
}

type InstructionForCreditorAgent1 struct {
	Code            *string `xml:"Cd,omitempty"`
	InstructionInfo *string `xml:"InstrInf,omitempty"`
}

type TaxInfo8 struct {
	Creditor               *TaxParty1                         `xml:"Cdtr,omitempty"`
	Debtor                 *TaxParty2                         `xml:"Dbtr,omitempty"`
	AdministrationZone     *string                            `xml:"AdmstnZone,omitempty"`
	ReferenceNumber        *string                            `xml:"RefNb,omitempty"`
	Method                 *string                            `xml:"Mtd,omitempty"`
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TotalTaxAmount         *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Date                   *string                            `xml:"Dt,omitempty"`
	SequenceNumber         *float64                           `xml:"SeqNb,omitempty"`
	Record                 []TaxRecord2                       `xml:"Rcrd,omitempty"`
}

type RemittanceInfo16 struct {
	Unstructured []string                     `xml:"Ustrd,omitempty"`
	Structured   []StructuredRemittanceInfo16 `xml:"Strd,omitempty"`
}

type TaxParty1 struct {
	TaxID          *string `xml:"TaxId,omitempty"`
	RegistrationID *string `xml:"RegnId,omitempty"`
	TaxType        *string `xml:"TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxID          *string            `xml:"TaxId,omitempty"`
	RegistrationID *string            `xml:"RegnId,omitempty"`
	TaxType        *string            `xml:"TaxTp,omitempty"`
	Authorization  *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

type TaxAuthorisation1 struct {
	Title *string `xml:"Titl,omitempty"`
	Name  *string `xml:"Nm,omitempty"`
}

type TaxRecord2 struct {
	Type            *string     `xml:"Tp,omitempty"`
	Category        *string     `xml:"Ctgy,omitempty"`
	CategoryDetails *string     `xml:"CtgyDtls,omitempty"`
	DebtorStatus    *string     `xml:"DbtrSts,omitempty"`
	CertificateID   *string     `xml:"CertId,omitempty"`
	FormsCode       *string     `xml:"FrmsCd,omitempty"`
	Period          *TaxPeriod2 `xml:"Prd,omitempty"`
	TaxAmount       *TaxAmount2 `xml:"TaxAmt,omitempty"`
	AdditionalInfo  *string     `xml:"AddtlInf,omitempty"`
}

type TaxPeriod2 struct {
	Year       *string      `xml:"Yr,omitempty"`
	Type       *string      `xml:"Tp,omitempty"`
	FromToDate *DatePeriod2 `xml:"FrToDt,omitempty"`
}

type TaxAmount2 struct {
	Rate              *float64                           `xml:"Rate,omitempty"`
	TaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TotalAmount       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Details           []TaxRecordDetails2                `xml:"Dtls,omitempty"`
}

type TaxRecordDetails2 struct {
	Period *TaxPeriod2                       `xml:"Prd,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DatePeriod2 struct {
	FromDate *string `xml:"FrDt,omitempty"`
	ToDate   *string `xml:"ToDt,omitempty"`
}

type StructuredRemittanceInfo16 struct {
	ReferredDocumentInfo     []ReferredDocumentInfo7 `xml:"RfrdDocInf,omitempty"`
	ReferredDocumentAmount   *RemittanceAmount2      `xml:"RfrdDocAmt,omitempty"`
	CreditorReferenceInfo    *CreditorReferenceInfo2 `xml:"CdtrRefInf,omitempty"`
	Invoicer                 *PartyIdentification135 `xml:"Invcr,omitempty"`
	Invoicee                 *PartyIdentification135 `xml:"Invcee,omitempty"`
	TaxRemittance            *TaxInfo7               `xml:"TaxRmt,omitempty"`
	GarnishmentRemittance    *Garnishment3           `xml:"GrnshmtRmt,omitempty"`
	AdditionalRemittanceInfo []string                `xml:"AddtlRmtInf,omitempty"` // max 3 elements
}

type ReferredDocumentInfo7 struct {
	Type        *ReferredDocumentType4 `xml:"Tp,omitempty"`
	Number      *string                `xml:"Nb,omitempty"`
	RelatedDate *string                `xml:"RltdDt,omitempty"`
	LineDetails []DocumentLineInfo1    `xml:"LineDtls,omitempty"`
}

type ReferredDocumentType4 struct {
	CodeOrProprietary ReferredDocumentType3 `xml:"CdOrPrtry"`
	Issuer            *string               `xml:"Issr,omitempty"`
}

type ReferredDocumentType3 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type DocumentLineInfo1 struct {
	Identification []DocumentLineIdentification1 `xml:"Id"`
	Description    *string                       `xml:"Desc,omitempty"`
	Amount         *RemittanceAmount3            `xml:"Amt,omitempty"`
}

type DocumentLineIdentification1 struct {
	Type        *DocumentLineType1 `xml:"Tp,omitempty"`
	Number      *string            `xml:"Nb,omitempty"`
	RelatedDate *time.Time         `xml:"RltdDt,omitempty"`
}

type DocumentLineType1 struct {
	CodeOrProprietary DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issuer            *string                 `xml:"Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type RemittanceAmount2 struct {
	DuePayableAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DiscountAppliedAmount     []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CreditNoteAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmount                 []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjustmentAmountAndReason []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RemittedAmount            *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePayableAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DiscountAppliedAmount     []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty"`
	CreditNoteAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmount                 []TaxAmountAndType1                `xml:"TaxAmt,omitempty"`
	AdjustmentAmountAndReason []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty"`
	RemittedAmount            *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

type CreditorReferenceInfo2 struct {
	Type      *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Reference *string                 `xml:"Ref,omitempty"`
}

type CreditorReferenceType2 struct {
	CodeOrProprietary CreditorReferenceType1 `xml:"CdOrPrtry"`
	Issuer            *string                `xml:"Issr,omitempty"`
}

type CreditorReferenceType1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type TaxInfo7 struct {
	Creditor               *TaxParty1                         `xml:"Cdtr,omitempty"`
	Debtor                 *TaxParty2                         `xml:"Dbtr,omitempty"`
	UltimateDebtor         *TaxParty2                         `xml:"UltmtDbtr,omitempty"`
	AdministrationZone     *string                            `xml:"AdmstnZone,omitempty"`
	ReferenceNumber        *string                            `xml:"RefNb,omitempty"`
	Method                 *string                            `xml:"Mtd,omitempty"`
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TotalTaxAmount         *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Date                   *string                            `xml:"Dt,omitempty"`
	SequenceNumber         *float64                           `xml:"SeqNb,omitempty"`
	Record                 []TaxRecord2                       `xml:"Rcrd,omitempty"`
}

type Garnishment3 struct {
	Type                            GarnishmentType1                   `xml:"Tp"`
	Garnishee                       *PartyIdentification135            `xml:"Grnshee,omitempty"`
	GarnishmentAdministrator        *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty"`
	ReferenceNumber                 *string                            `xml:"RefNb,omitempty"`
	Date                            *time.Time                         `xml:"Dt,omitempty"`
	RemittedAmount                  *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
	FamilyMedicalInsuranceIndicator *bool                              `xml:"FmlyMdclInsrncInd,omitempty"`
	EmployeeTerminationIndicator    *bool                              `xml:"MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CodeOrProprietary GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issuer            *string                `xml:"Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type DiscountAmountAndType1 struct {
	Type   *DiscountAmountType1              `xml:"Tp,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// DiscountAmountType1 matches XSD type
type DiscountAmountType1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// TaxAmountAndType1 matches XSD type
type TaxAmountAndType1 struct {
	Type   *TaxAmountType1                   `xml:"Tp,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

// TaxAmountType1 matches XSD type
type TaxAmountType1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// DocumentAdjustment1 matches XSD type
type DocumentAdjustment1 struct {
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator *string                           `xml:"CdtDbtInd,omitempty"`
	Reason               *string                           `xml:"Rsn,omitempty"`
	AdditionalInfo       *string                           `xml:"AddtlInf,omitempty"`
}

// DateAndDateTime2 matches XSD type
type DateAndDateTime2 struct {
	Date     *string    `xml:"Dt,omitempty"`
	DateTime *time.Time `xml:"DtTm,omitempty"`
}

// Supporting types for completeness
type ServiceLevel struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type LocalInstrument struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type CategoryPurpose struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type ClearingSystemMemberIdentification struct {
	ClearingSystemID *ClearingSystemIdentification `xml:"ClrSysId,omitempty"`
	MemberID         string                        `xml:"MmbId"`
}

type ClearingSystemIdentification struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type GenericFinancialIdentification struct {
	ID         string                             `xml:"Id"`
	SchemeName *FinancialIdentificationSchemeName `xml:"SchmeNm,omitempty"`
	Issuer     *string                            `xml:"Issr,omitempty"`
}

type FinancialIdentificationSchemeName struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type CashAccountType struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type ProxyAccountIdentification struct {
	Type *ProxyAccountType `xml:"Tp,omitempty"`
	ID   string            `xml:"Id"`
}

type ProxyAccountType struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type GenericAccountIdentification struct {
	ID         string             `xml:"Id"`
	SchemeName *AccountSchemeName `xml:"SchmeNm,omitempty"`
	Issuer     *string            `xml:"Issr,omitempty"`
}

type AccountSchemeName struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type OrganizationIdentification struct {
	AnyBankIdentifierCode *string                             `xml:"AnyBIC,omitempty"`
	LegalEntityIdentifier *string                             `xml:"LEI,omitempty"`
	Other                 []GenericOrganizationIdentification `xml:"Othr,omitempty"`
}

type PersonIdentification struct {
	DateAndPlaceOfBirth *DateAndPlaceOfBirth          `xml:"DtAndPlcOfBirth,omitempty"`
	Other               []GenericPersonIdentification `xml:"Othr,omitempty"`
}

type GenericOrganizationIdentification struct {
	ID         string                                `xml:"Id"`
	SchemeName *OrganizationIdentificationSchemeName `xml:"SchmeNm,omitempty"`
	Issuer     *string                               `xml:"Issr,omitempty"`
}

type OrganizationIdentificationSchemeName struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDate       *string `xml:"BirthDt,omitempty"`
	ProvinceOfBirth *string `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth     string  `xml:"CityOfBirth"`
	CountryOfBirth  string  `xml:"CtryOfBirth"`
}

type GenericPersonIdentification struct {
	ID         string                          `xml:"Id"`
	SchemeName *PersonIdentificationSchemeName `xml:"SchmeNm,omitempty"`
	Issuer     *string                         `xml:"Issr,omitempty"`
}

type PersonIdentificationSchemeName struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type Contact struct {
	NamePrefix      *string        `xml:"NmPrfx,omitempty"`
	Name            *string        `xml:"Nm,omitempty"`
	PhoneNumber     *string        `xml:"PhneNb,omitempty"`
	MobileNumber    *string        `xml:"MobNb,omitempty"`
	FaxNumber       *string        `xml:"FaxNb,omitempty"`
	EmailAddress    *string        `xml:"EmailAdr,omitempty"`
	EmailPurpose    *string        `xml:"EmailPurp,omitempty"`
	JobTitle        *string        `xml:"JobTitl,omitempty"`
	Responsibility  *string        `xml:"Rspnsblty,omitempty"`
	Department      *string        `xml:"Dept,omitempty"`
	Other           []OtherContact `xml:"Othr,omitempty"`
	PreferredMethod *string        `xml:"PrefrdMtd,omitempty"`
}

type OtherContact struct {
	ChannelType string  `xml:"ChanlTp"`
	ID          *string `xml:"Id,omitempty"`
}

type TaxAuthorization struct {
	Title *string `xml:"Titl,omitempty"`
	Name  *string `xml:"Nm,omitempty"`
}

type TaxRecord struct {
	Type            *string    `xml:"Tp,omitempty"`
	Category        *string    `xml:"Ctgy,omitempty"`
	CategoryDetails *string    `xml:"CtgyDtls,omitempty"`
	DebtorStatus    *string    `xml:"DbtrSts,omitempty"`
	CertificateID   *string    `xml:"CertId,omitempty"`
	FormsCode       *string    `xml:"FrmsCd,omitempty"`
	Period          *TaxPeriod `xml:"Prd,omitempty"`
	TaxAmount       *TaxAmount `xml:"TaxAmt,omitempty"`
	AdditionalInfo  *string    `xml:"AddtlInf,omitempty"`
}

type TaxPeriod struct {
	Year       *time.Time  `xml:"Yr,omitempty"`
	Type       *string     `xml:"Tp,omitempty"`
	FromToDate *DatePeriod `xml:"FrToDt,omitempty"`
}

type DatePeriod struct {
	FromDate *string `xml:"FrDt,omitempty"`
	ToDate   *string `xml:"ToDt,omitempty"`
}

type TaxAmount struct {
	Rate              *float64                           `xml:"Rate,omitempty"`
	TaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TotalAmount       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Details           []TaxRecordDetails                 `xml:"Dtls,omitempty"`
}

type TaxRecordDetails struct {
	Period *TaxPeriod                        `xml:"Prd,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type RemittanceLocationData struct {
	Method            string          `xml:"Mtd"`
	ElectronicAddress *string         `xml:"ElctrncAdr,omitempty"`
	PostalAddress     *NameAndAddress `xml:"PstlAdr,omitempty"`
}

type NameAndAddress struct {
	Name    string        `xml:"Nm"`
	Address PostalAddress `xml:"Adr"`
}

type StructuredRemittanceInfo struct {
	ReferredDocumentInfo     []ReferredDocumentInfo   `xml:"RfrdDocInf,omitempty"`
	ReferredDocumentAmount   *RemittanceAmountPrimary `xml:"RfrdDocAmt,omitempty"`
	CreditorReferenceInfo    *CreditorReferenceInfo   `xml:"CdtrRefInf,omitempty"`
	Invoicer                 *PartyIdentification     `xml:"Invcr,omitempty"`
	Invoicee                 *PartyIdentification     `xml:"Invcee,omitempty"`
	TaxRemittance            *TaxInfoSecondary        `xml:"TaxRmt,omitempty"`
	GarnishmentRemittance    *Garnishment             `xml:"GrnshmtRmt,omitempty"`
	AdditionalRemittanceInfo *string                  `xml:"AddtlRmtInf,omitempty"`
}

type ReferredDocumentInfo struct {
	Type        *ReferredDocumentType `xml:"Tp,omitempty"`
	Number      *string               `xml:"Nb,omitempty"`
	RelatedDate *time.Time            `xml:"RltdDt,omitempty"`
	LineDetails []DocumentLineInfo    `xml:"LineDtls,omitempty"`
}

type ReferredDocumentType struct {
	CodeOrProprietary ReferredDocumentTypeOption `xml:"CdOrPrtry"`
	Issuer            *string                    `xml:"Issr,omitempty"`
}

type ReferredDocumentTypeOption struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type DocumentLineInfo struct {
	ID          []DocumentLineIdentification `xml:"Id"`
	Description *string                      `xml:"Desc,omitempty"`
	Amount      *RemittanceAmountSecondary   `xml:"Amt,omitempty"`
}

type DocumentLineIdentification struct {
	Type        *DocumentLineType `xml:"Tp,omitempty"`
	Number      *string           `xml:"Nb,omitempty"`
	RelatedDate *time.Time        `xml:"RltdDt,omitempty"`
}

type DocumentLineType struct {
	CodeOrProprietary DocumentLineTypeOption `xml:"CdOrPrtry"`
	Issuer            *string                `xml:"Issr,omitempty"`
}

type DocumentLineTypeOption struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type RemittanceAmountPrimary struct {
	DuePayableAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DiscountAppliedAmount     []DiscountAmountAndType            `xml:"DscntApldAmt,omitempty"`
	CreditNoteAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmount                 []TaxAmountAndType                 `xml:"TaxAmt,omitempty"`
	AdjustmentAmountAndReason []DocumentAdjustment               `xml:"AdjstmntAmtAndRsn,omitempty"`
	RemittedAmount            *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

type RemittanceAmountSecondary struct {
	DuePayableAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DiscountAppliedAmount     []DiscountAmountAndType            `xml:"DscntApldAmt,omitempty"`
	CreditNoteAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmount                 []TaxAmountAndType                 `xml:"TaxAmt,omitempty"`
	AdjustmentAmountAndReason []DocumentAdjustment               `xml:"AdjstmntAmtAndRsn,omitempty"`
	RemittedAmount            *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

type DiscountAmountAndType struct {
	Type   *DiscountAmountType               `xml:"Tp,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DiscountAmountType struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type TaxAmountAndType struct {
	Type   *TaxAmountType                    `xml:"Tp,omitempty"`
	Amount ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type TaxAmountType struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type DocumentAdjustment struct {
	Amount                ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator  *string                           `xml:"CdtDbtInd,omitempty"`
	Reason                *string                           `xml:"Rsn,omitempty"`
	AdditionalInformation *string                           `xml:"AddtlInf,omitempty"`
}

type CreditorReferenceInfo struct {
	Type      *CreditorReferenceType `xml:"Tp,omitempty"`
	Reference *string                `xml:"Ref,omitempty"`
}

type CreditorReferenceType struct {
	CodeOrProprietary CreditorReferenceTypeOption `xml:"CdOrPrtry"`
	Issuer            *string                     `xml:"Issr,omitempty"`
}

type CreditorReferenceTypeOption struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type TaxInfoSecondary struct {
	Creditor               *TaxPartyCreditor                  `xml:"Cdtr,omitempty"`
	Debtor                 *TaxPartyDebtor                    `xml:"Dbtr,omitempty"`
	AdministrationZone     *string                            `xml:"AdmstnZone,omitempty"`
	ReferenceNumber        *string                            `xml:"RefNb,omitempty"`
	Method                 *string                            `xml:"Mtd,omitempty"`
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TotalTaxAmount         *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Date                   *string                            `xml:"Dt,omitempty"`
	SequenceNumber         *float64                           `xml:"SeqNb,omitempty"`
	Record                 []TaxRecord                        `xml:"Rcrd,omitempty"`
}

type Garnishment struct {
	Type                            GarnishmentType                    `xml:"Tp"`
	Garnishee                       *PartyIdentification               `xml:"Grnshee,omitempty"`
	GarnishmentAdministrator        *PartyIdentification               `xml:"GrnshmtAdmstr,omitempty"`
	ReferenceNumber                 *string                            `xml:"RefNb,omitempty"`
	Date                            *time.Time                         `xml:"Dt,omitempty"`
	RemittedAmount                  *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
	FamilyMedicalInsuranceIndicator *bool                              `xml:"FmlyMdclInsrncInd,omitempty"`
	EmployeeTerminationIndicator    *bool                              `xml:"MplyeeTermntnInd,omitempty"`
}

type GarnishmentType struct {
	CodeOrProprietary GarnishmentTypeOption `xml:"CdOrPrtry"`
	Issuer            *string               `xml:"Issr,omitempty"`
}

type GarnishmentTypeOption struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// Additional supporting types
type SettlementInstruction struct {
	SettlementMethod                     string                                       `xml:"SttlmMtd"`
	SettlementAccount                    *CashAccount                                 `xml:"SttlmAcct,omitempty"`
	ClearingSystem                       *ClearingSystemIdentificationSecondary       `xml:"ClrSys,omitempty"`
	InstructingReimbursementAgent        *BranchAndFinancialInstitutionIdentification `xml:"InstgRmbrsmntAgt,omitempty"`
	InstructingReimbursementAgentAccount *CashAccount                                 `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstructedReimbursementAgent         *BranchAndFinancialInstitutionIdentification `xml:"InstdRmbrsmntAgt,omitempty"`
	InstructedReimbursementAgentAccount  *CashAccount                                 `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThirdReimbursementAgent              *BranchAndFinancialInstitutionIdentification `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThirdReimbursementAgentAccount       *CashAccount                                 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

type ClearingSystemIdentificationSecondary struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type SettlementDateTimeIndication struct {
	DebitDateTime  *time.Time `xml:"DbtDtTm,omitempty"`
	CreditDateTime *time.Time `xml:"CdtDtTm,omitempty"`
}

type SettlementTimeRequest struct {
	ClearingSystemTime *time.Time `xml:"CLSTm,omitempty"`
	TillTime           *time.Time `xml:"TillTm,omitempty"`
	FromTime           *time.Time `xml:"FrTm,omitempty"`
	RejectTime         *time.Time `xml:"RjctTm,omitempty"`
}

type InstructionForCreditorAgent struct {
	Code            *string `xml:"Cd,omitempty"`
	InstructionInfo *string `xml:"InstrInf,omitempty"`
}

type InstructionForNextAgent struct {
	Code            *string `xml:"Cd,omitempty"`
	InstructionInfo *string `xml:"InstrInf,omitempty"`
}

type Purpose struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type RegulatoryReporting3 struct {
	DebitCreditReportingIndicator *string                          `xml:"DbtCdtRptgInd,omitempty"`
	Authority                     *RegulatoryAuthority2            `xml:"Authrty,omitempty"`
	Dtls                          []StructuredRegulatoryReporting3 `xml:"Dtls,omitempty"`
}

type RegulatoryAuthority2 struct {
	Name    *string `xml:"Nm,omitempty"`
	Country *string `xml:"Ctry,omitempty"`
}

type StructuredRegulatoryReporting3 struct {
	Type        *string                            `xml:"Tp,omitempty"`
	Date        *time.Time                         `xml:"Dt,omitempty"`
	Country     *string                            `xml:"Ctry,omitempty"`
	Code        *string                            `xml:"Cd,omitempty"`
	Amount      *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	Information []string                           `xml:"Inf,omitempty"`
}

// FinancialInstitutionCreditTransferV08 - pacs.009.001.08
type FinancialInstitutionCreditTransferV08 struct {
	GroupHeader                   GroupHeader93                 `xml:"GrpHdr"`
	CreditTransferTransactionInfo []CreditTransferTransaction36 `xml:"CdtTrfTxInf"`
	SupplementaryData             []SupplementaryData1          `xml:"SplmtryData,omitempty"`
}

// FIToFIPaymentStatusReportV10 - pacs.002.001.10
type FIToFIPaymentStatusReportV10 struct {
	GroupHeader                       GroupHeader91           `xml:"GrpHdr"`
	OriginalGroupInformationAndStatus []OriginalGroupHeader17 `xml:"OrgnlGrpInfAndSts,omitempty"`
	TransactionInfoAndStatus          []PaymentTransaction110 `xml:"TxInfAndSts,omitempty"`
	SupplementaryData                 []SupplementaryData1    `xml:"SplmtryData,omitempty"`
}

// PaymentReturnV10 - pacs.004.001.10
type PaymentReturnV10 struct {
	GroupHeader       GroupHeader90              `xml:"GrpHdr"`
	OriginalGroupInfo OriginalGroupInfo29        `xml:"OrgnlGrpInf"`
	TransactionInfo   []PaymentTransactionInfo51 `xml:"TxInf,omitempty"`
	SupplementaryData []SupplementaryData        `xml:"SplmtryData,omitempty"`
}

// FIToFIPaymentStatusRequestV03 - pacs.028.001.03
type FIToFIPaymentStatusRequestV03 struct {
	GroupHeader       GroupHeader91              `xml:"GrpHdr"`
	OriginalGroupInfo *OriginalGroupHeader17     `xml:"OrgnlGrpInf,omitempty"`
	TransactionInfo   []PaymentTransactionInfo50 `xml:"TxInf,omitempty"`
	SupplementaryData []SupplementaryData        `xml:"SplmtryData,omitempty"`
}

// BankToCustomerAccountReportV08 - camt.052.001.08
type BankToCustomerAccountReportV08 struct {
	GroupHeader       GroupHeader81        `xml:"GrpHdr"`
	Report            []AccountReport25    `xml:"Rpt"`
	SupplementaryData []SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

// BankToCustomerDebitCreditNotificationV08 - camt.054.001.08
type BankToCustomerDebitCreditNotificationV08 struct {
	GroupHeader       GroupHeader81           `xml:"GrpHdr"`
	Notification      []AccountNotification17 `xml:"Ntfctn"`
	SupplementaryData []SupplementaryData1    `xml:"SplmtryData,omitempty"`
}

// CustomerPaymentCancellationRequestV09 - camt.055.001.09
type CustomerPaymentCancellationRequestV09 struct {
	Assignment        CaseAssignment5         `xml:"Assgnmt"`
	Case              *Case5                  `xml:"Case,omitempty"`
	ControlData       *ControlData1           `xml:"CtrlData,omitempty"`
	Underlying        UnderlyingTransaction22 `xml:"Undrlyg"`
	SupplementaryData []SupplementaryData1    `xml:"SplmtryData,omitempty"`
}

// FIToFIPaymentCancellationRequestV08 - camt.056.001.08
type FIToFIPaymentCancellationRequestV08 struct {
	Assignment        CaseAssignment5         `xml:"Assgnmt"`
	Case              *Case5                  `xml:"Case,omitempty"`
	ControlData       *ControlData1           `xml:"CtrlData,omitempty"`
	Underlying        UnderlyingTransaction23 `xml:"Undrlyg"`
	SupplementaryData []SupplementaryData1    `xml:"SplmtryData,omitempty"`
}

// AccountReportingRequestV05 - camt.060.001.05
type AccountReportingRequestV05 struct {
	GroupHeader       GroupHeader81        `xml:"GrpHdr"`
	ReportingRequest  []ReportingRequest5  `xml:"RptgReq"`
	SupplementaryData []SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

// UnableToApplyV07 - camt.026.001.07
type UnableToApplyV07 struct {
	Assignment        CaseAssignment5             `xml:"Assgnmt"`
	Case              *Case5                      `xml:"Case,omitempty"`
	Underlying        UnderlyingTransaction21     `xml:"Undrlyg"`
	Justification     UnableToApplyJustification3 `xml:"Justfn"`
	SupplementaryData []SupplementaryData         `xml:"SplmtryData,omitempty"`
}

// AdditionalPaymentInfoV09 - camt.028.001.09
type AdditionalPaymentInfoV09 struct {
	Assignment        CaseAssignment5           `xml:"Assgnmt"`
	Case              *Case5                    `xml:"Case,omitempty"`
	Underlying        UnderlyingTransaction21   `xml:"Undrlyg"`
	Info              PaymentComplementaryInfo9 `xml:"Inf"`
	SupplementaryData []SupplementaryData       `xml:"SplmtryData,omitempty"`
}

// ResolutionOfInvestigationV09 - camt.029.001.09
type ResolutionOfInvestigationV09 struct {
	Assignment            CaseAssignment5            `xml:"Assgnmt"`
	ResolvedCase          *Case5                     `xml:"RslvdCase,omitempty"`
	Status                InvestigationStatus5       `xml:"Sts"`
	CancellationDetails   []UnderlyingTransaction22  `xml:"CxlDtls,omitempty"`
	ModificationDetails   *PaymentTransaction91      `xml:"ModDtls,omitempty"`
	ClaimNonReceipt       *ClaimNonReceipt2          `xml:"ClmNonRct,omitempty"`
	StatementDetails      *StatementResolutionEntry4 `xml:"StmtDtls,omitempty"`
	CorrectionTransaction *CorrectiveTransaction4    `xml:"CrrctnTx,omitempty"`
	ResolutionRelatedInfo *ResolutionData2           `xml:"RsltnRltdInf,omitempty"`
	SupplementaryData     []SupplementaryData        `xml:"SplmtryData,omitempty"`
}

// CreditorPaymentActivationRequestV07 - pain.013.001.07
type CreditorPaymentActivationRequestV07 struct {
	GroupHeader       GroupHeader78          `xml:"GrpHdr"`
	PaymentInfo       []PaymentInstruction31 `xml:"PmtInf"`
	SupplementaryData []SupplementaryData    `xml:"SplmtryData,omitempty"`
}

// CreditorPaymentActivationRequestStatusReportV07 - pain.014.001.07
type CreditorPaymentActivationRequestStatusReportV07 struct {
	GroupHeader         GroupHeader86                  `xml:"GrpHdr"`
	OriginalGroupInfo   OriginalGroupHeader17          `xml:"OrgnlGrpInf"`
	OriginalPaymentInfo []OriginalPaymentInstruction32 `xml:"OrgnlPmtInf,omitempty"`
	SupplementaryData   []SupplementaryData            `xml:"SplmtryData,omitempty"`
}

// SystemEventNotificationV02 - admi.004.001.02
type SystemEventNotificationV02 struct {
	EventInfo         Event2              `xml:"EvtInf"`
	SupplementaryData []SupplementaryData `xml:"SplmtryData,omitempty"`
}

// SystemEventAcknowledgementV01 - admi.011.001.01
type SystemEventAcknowledgementV01 struct {
	AcknowledgementDetails Acknowledgement1    `xml:"AckDtls"`
	SupplementaryData      []SupplementaryData `xml:"SplmtryData,omitempty"`
}

// ResendRequestV01 - admi.006.001.01
type ResendRequestV01 struct {
	RequestDetails    ResendSearchCriteria1 `xml:"ReqDtls"`
	SupplementaryData []SupplementaryData   `xml:"SplmtryData,omitempty"`
}

// MessageHeader10 represents message identification and optional creation date/time for admi.007.001.01
type MessageHeader10 struct {
	XMLName          xml.Name   `xml:",omitempty"`
	MessageID        string     `xml:"MsgId"`
	CreationDateTime *time.Time `xml:"CreDtTm,omitempty"`
	QueryName        *string    `xml:"QryNm,omitempty"`
}

// MessageReference1 contains a reference to the original message and optional issuer
type MessageReference1 struct {
	XMLName         xml.Name                `xml:",omitempty"`
	Reference       string                  `xml:"Ref"`
	MessageName     *string                 `xml:"MsgNm,omitempty"`
	ReferenceIssuer *PartyIdentification136 `xml:"RefIssr,omitempty"`
}

// RequestHandling2 contains status information for the receipt acknowledgement
type RequestHandling2 struct {
	XMLName        xml.Name   `xml:",omitempty"`
	StatusCode     string     `xml:"StsCd"`
	StatusDateTime *time.Time `xml:"StsDtTm,omitempty"`
	Description    *string    `xml:"Desc,omitempty"`
}

// ReceiptAcknowledgementReport2 contains the related reference and request handling information
type ReceiptAcknowledgementReport2 struct {
	XMLName          xml.Name          `xml:",omitempty"`
	RelatedReference MessageReference1 `xml:"RltdRef"`
	RequestHandling  RequestHandling2  `xml:"ReqHdlg"`
}

// PartyIdentification120Choice represents different ways to identify a party
type PartyIdentification120Choice struct {
	XMLName        xml.Name                 `xml:",omitempty"`
	AnyBIC         *string                  `xml:"AnyBIC,omitempty"`
	ProprietaryID  *GenericIdentification36 `xml:"PrtryId,omitempty"`
	NameAndAddress *NameAndAddress5         `xml:"NmAndAdr,omitempty"`
}

// PartyIdentification136 contains party identification with optional LEI
type PartyIdentification136 struct {
	XMLName xml.Name                     `xml:",omitempty"`
	ID      PartyIdentification120Choice `xml:"Id"`
	LEI     *string                      `xml:"LEI,omitempty"`
}

// GenericIdentification36 represents a generic identification scheme
type GenericIdentification36 struct {
	XMLName    xml.Name `xml:",omitempty"`
	ID         string   `xml:"Id"`
	Issuer     string   `xml:"Issr"`
	SchemeName *string  `xml:"SchmeNm,omitempty"`
}

// NameAndAddress5 contains party name and optional postal address
type NameAndAddress5 struct {
	XMLName xml.Name        `xml:",omitempty"`
	Name    string          `xml:"Nm"`
	Address *PostalAddress1 `xml:"Adr,omitempty"`
}

// PostalAddress1 contains postal address information for admi.007.001.01
type PostalAddress1 struct {
	XMLName            xml.Name `xml:",omitempty"`
	AddressType        *string  `xml:"AdrTp,omitempty"`
	AddressLine        []string `xml:"AdrLine,omitempty"`
	StreetName         *string  `xml:"StrtNm,omitempty"`
	BuildingNumber     *string  `xml:"BldgNb,omitempty"`
	PostCode           *string  `xml:"PstCd,omitempty"`
	TownName           *string  `xml:"TwnNm,omitempty"`
	CountrySubDivision *string  `xml:"CtrySubDvsn,omitempty"`
	Country            string   `xml:"Ctry"`
}

// ReceiptAcknowledgementV01 - admi.007.001.01
type ReceiptAcknowledgementV01 struct {
	XMLName           xml.Name                        `xml:",omitempty"`
	MessageID         MessageHeader10                 `xml:"MsgId"`
	Report            []ReceiptAcknowledgementReport2 `xml:"Rpt"`
	SupplementaryData []SupplementaryData1            `xml:"SplmtryData,omitempty"`
}

// Admi00200101Document represents the ADMI.002.001.01 Message Rejection message.
// This administrative message is used to reject a previously received message when it cannot be processed,
// providing detailed information about the rejection reason, error location, and additional diagnostic data.
type Admi00200101Document struct {
	XMLName          xml.Name            `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Document"`
	MessageRejection MessageRejectionV01 `xml:"admi.002.001.01"`
}

// MessageRejectionV01 represents the core structure of an ADMI.002.001.01 message.
// Contains the related message reference and the detailed rejection reason information
// explaining why the original message could not be processed.
type MessageRejectionV01 struct {
	RelatedReference MessageReference `xml:"RltdRef"`
	Reason           RejectionReason2 `xml:"Rsn"`
}

// MessageReference contains a reference to the original message being rejected.
// Provides the unique identifier reference to link this rejection back to the original message.
type MessageReference struct {
	Reference string `xml:"Ref"`
}

// RejectionReason2 contains detailed information about why the message was rejected.
// Includes the rejecting party's reason code, optional rejection timestamp, error location,
// descriptive reason, and additional diagnostic data for troubleshooting.
type RejectionReason2 struct {
	RejectingPartyReason string     `xml:"RjctgPtyRsn"`
	RejectionDateTime    *time.Time `xml:"RjctnDtTm,omitempty"`
	ErrorLocation        *string    `xml:"ErrLctn,omitempty"`
	ReasonDescription    *string    `xml:"RsnDesc,omitempty"`
	AdditionalData       *string    `xml:"AddtlData,omitempty"`
}

// AdministrationProprietaryMessageV02 - admi.998.001.02
type AdministrationProprietaryMessageV02 struct {
	MessageID         *MessageReference   `xml:"MsgId,omitempty"`
	Related           *MessageReference   `xml:"Rltd,omitempty"`
	ProprietaryData   ProprietaryData6    `xml:"PrtryData"`
	SupplementaryData []SupplementaryData `xml:"SplmtryData,omitempty"`
}

// Transaction and Group Header types
type GroupHeader90 struct {
	MessageID        string                                       `xml:"MsgId"`
	CreatedDateTime  *time.Time                                   `xml:"CreDtTm,omitempty"`
	InstructingAgent *BranchAndFinancialInstitutionIdentification `xml:"InstgAgt,omitempty"`
	InstructedAgent  *BranchAndFinancialInstitutionIdentification `xml:"InstdAgt,omitempty"`
}

type GroupHeader91 struct {
	MessageID        string                                       `xml:"MsgId"`
	CreatedDateTime  *time.Time                                   `xml:"CreDtTm,omitempty"`
	InstructingAgent *BranchAndFinancialInstitutionIdentification `xml:"InstgAgt,omitempty"`
	InstructedAgent  *BranchAndFinancialInstitutionIdentification `xml:"InstdAgt,omitempty"`
}

type GroupHeader81 struct {
	MsgID                 string               `xml:"MsgId"`
	CreatedDateTime       *time.Time           `xml:"CreDtTm,omitempty"`
	MessageRecipient      *PartyIdentification `xml:"MsgRcpt,omitempty"`
	MessagePagination     *Pagination1         `xml:"MsgPgntn,omitempty"`
	OriginalBusinessQuery *string              `xml:"OrgnlBizQry,omitempty"`
	AdditionalInformation *string              `xml:"AddtlInf,omitempty"`
}

type GroupHeader78 struct {
	MessageID            string                                       `xml:"MsgId"`
	CreatedDateTime      *time.Time                                   `xml:"CreDtTm,omitempty"`
	Authorization        []Authorization1                             `xml:"Authstn,omitempty"`
	NumberOfTransactions string                                       `xml:"NbOfTxs"`
	ControlSum           *float64                                     `xml:"CtrlSum,omitempty"`
	InitiatingParty      PartyIdentification                          `xml:"InitgPty"`
	ForwardingAgent      *BranchAndFinancialInstitutionIdentification `xml:"FwdgAgt,omitempty"`
}

type GroupHeader86 struct {
	MessageID       string                                       `xml:"MsgId"`
	CreatedDateTime *time.Time                                   `xml:"CreDtTm,omitempty"`
	InitiatingParty PartyIdentification                          `xml:"InitgPty"`
	ForwardingAgent *BranchAndFinancialInstitutionIdentification `xml:"FwdgAgt,omitempty"`
}

// Transaction information types

// CreditTransferTransaction36 - for pacs.009.001.08 (exact XSD match)
type CreditTransferTransaction36 struct {
	PaymentID                        PaymentIdentification7                        `xml:"PmtId"`
	PaymentTypeInfo                  *PaymentTypeInfo28                            `xml:"PmtTpInf,omitempty"`
	InterbankSettlementAmount        ActiveCurrencyAndAmount                       `xml:"IntrBkSttlmAmt"`
	InterbankSettlementDate          *string                                       `xml:"IntrBkSttlmDt,omitempty"`
	SettlementPriority               *string                                       `xml:"SttlmPrty,omitempty"`
	SettlementTimeIndication         *SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty"`
	SettlementTimeRequest            *SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty"`
	PreviousInstructingAgent1        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty"`
	PreviousInstructingAgent1Account *CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty"`
	PreviousInstructingAgent2        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty"`
	PreviousInstructingAgent2Account *CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty"`
	PreviousInstructingAgent3        *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty"`
	PreviousInstructingAgent3Account *CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty"`
	InstructingAgent                 *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent                  *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	IntermediaryAgent1               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntermediaryAgent1Account        *CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntermediaryAgent2               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntermediaryAgent2Account        *CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntermediaryAgent3               *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	IntermediaryAgent3Account        *CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty"`
	UltimateDebtor                   *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtDbtr,omitempty"`
	Debtor                           BranchAndFinancialInstitutionIdentification6  `xml:"Dbtr"`
	DebtorAccount                    *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent                      *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	DebtorAgentAccount               *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CreditorAgent                    *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	CreditorAgentAccount             *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Creditor                         BranchAndFinancialInstitutionIdentification6  `xml:"Cdtr"`
	CreditorAccount                  *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltimateCreditor                 *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtCdtr,omitempty"`
	InstructionForCreditorAgent      []InstructionForCreditorAgent2                `xml:"InstrForCdtrAgt,omitempty"`
	InstructionForNextAgent          []InstructionForNextAgent1                    `xml:"InstrForNxtAgt,omitempty"`
	Purpose                          *Purpose2                                     `xml:"Purp,omitempty"`
	RemittanceInfo                   *RemittanceInfo2                              `xml:"RmtInf,omitempty"`
	UnderlyingCustomerCreditTransfer *CreditTransferTransaction37                  `xml:"UndrlygCstmrCdtTrf,omitempty"`
	SupplementaryData                []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

type CreditTransferTransactionInfo35 struct {
	PaymentID                    PaymentIdentification                        `xml:"PmtId"`
	PaymentTypeInfo              *PaymentTypeInfo                             `xml:"PmtTpInf,omitempty"`
	InterbankSettlementAmount    ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt"`
	InterbankSettlementDate      *time.Time                                   `xml:"IntrBkSttlmDt,omitempty"`
	SettlementPriority           *string                                      `xml:"SttlmPrty,omitempty"`
	SettlementTimeIndication     *SettlementDateTimeIndication                `xml:"SttlmTmIndctn,omitempty"`
	SettlementTimeRequest        *SettlementTimeRequest                       `xml:"SttlmTmReq,omitempty"`
	InstructingAgent             *BranchAndFinancialInstitutionIdentification `xml:"InstgAgt,omitempty"`
	InstructedAgent              *BranchAndFinancialInstitutionIdentification `xml:"InstdAgt,omitempty"`
	IntermediaryAgent1           *BranchAndFinancialInstitutionIdentification `xml:"IntrmyAgt1,omitempty"`
	IntermediaryAgent1Account    *CashAccount                                 `xml:"IntrmyAgt1Acct,omitempty"`
	IntermediaryAgent2           *BranchAndFinancialInstitutionIdentification `xml:"IntrmyAgt2,omitempty"`
	IntermediaryAgent2Account    *CashAccount                                 `xml:"IntrmyAgt2Acct,omitempty"`
	IntermediaryAgent3           *BranchAndFinancialInstitutionIdentification `xml:"IntrmyAgt3,omitempty"`
	IntermediaryAgent3Account    *CashAccount                                 `xml:"IntrmyAgt3Acct,omitempty"`
	CreditorAgent                BranchAndFinancialInstitutionIdentification  `xml:"CdtrAgt"`
	CreditorAgentAccount         *CashAccount                                 `xml:"CdtrAgtAcct,omitempty"`
	Creditor                     *PartyIdentification                         `xml:"Cdtr,omitempty"`
	CreditorAccount              *CashAccount                                 `xml:"CdtrAcct,omitempty"`
	InstructionsForCreditorAgent []InstructionForCreditorAgent                `xml:"InstrForCdtrAgt,omitempty"`
	InstructionsForNextAgent     []InstructionForNextAgent                    `xml:"InstrForNxtAgt,omitempty"`
	Purpose                      *Purpose                                     `xml:"Purp,omitempty"`
	RemittanceInfo               *RemittanceInfo                              `xml:"RmtInf,omitempty"`
	SupplementaryData            []SupplementaryData                          `xml:"SplmtryData,omitempty"`
}

// Additional supporting types
type PaymentTransactionInfo52 struct {
	StatusID                         *string                                       `xml:"StsId,omitempty"`
	OriginalInstructionID            *string                                       `xml:"OrgnlInstrId,omitempty"`
	OriginalEndToEndID               *string                                       `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID            *string                                       `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                     *string                                       `xml:"OrgnlUETR,omitempty"`
	TransactionStatus                *string                                       `xml:"TxSts,omitempty"`
	StatusReasonInfo                 []StatusReasonInfo12                          `xml:"StsRsnInf,omitempty"`
	ChargesInfo                      []Charges                                     `xml:"ChrgsInf,omitempty"`
	AcceptanceDateTime               *time.Time                                    `xml:"AccptncDtTm,omitempty"`
	EffectiveInterbankSettlementDate *DateAndDateTime2                             `xml:"FctvIntrBkSttlmDt,omitempty"`
	AccountServicerReference         *string                                       `xml:"AcctSvcrRef,omitempty"`
	ClearingSystemReference          *string                                       `xml:"ClrSysRef,omitempty"`
	InstructingAgent                 *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent                  *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	OriginalTransactionReference     *OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty"`
	SupplementaryData                []SupplementaryData                           `xml:"SplmtryData,omitempty"`
}

type PaymentTransactionInfo51 struct {
	ReturnID                          *string                            `xml:"RtrId,omitempty"`
	OriginalGroupInfo                 *OriginalGroupInfo29               `xml:"OrgnlGrpInf,omitempty"`
	OriginalInstructionID             *string                            `xml:"OrgnlInstrId,omitempty"`
	OriginalEndToEndID                *string                            `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID             *string                            `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                      *string                            `xml:"OrgnlUETR,omitempty"`
	OriginalClearingSystemReference   *string                            `xml:"OrgnlClrSysRef,omitempty"`
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`
	ReturnedInterbankSettlementAmount ActiveCurrencyAndAmount            `xml:"RtrdIntrBkSttlmAmt"`
	InterbankSettlementDate           *time.Time                         `xml:"IntrBkSttlmDt,omitempty"`
	ReturnedInstructedAmount          *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`
	ExchangeRate                      *float64                           `xml:"XchgRate,omitempty"`
	CompensationAmount                *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`
	ReturnReason                      ReturnReason5                      `xml:"RtrRsn"`
	OriginalTransactionReference      *OriginalTransactionReference28    `xml:"OrgnlTxRef,omitempty"`
	SupplementaryData                 []SupplementaryData                `xml:"SplmtryData,omitempty"`
}

type PaymentTransactionInfo50 struct {
	StatusRequestID                 *string             `xml:"StsReqId,omitempty"`
	OriginalInstructionID           *string             `xml:"OrgnlInstrId,omitempty"`
	OriginalEndToEndID              *string             `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID           *string             `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                    *string             `xml:"OrgnlUETR,omitempty"`
	OriginalClearingSystemReference *string             `xml:"OrgnlClrSysRef,omitempty"`
	StatusReason                    *StatusReason6      `xml:"StsRsn,omitempty"`
	SupplementaryData               []SupplementaryData `xml:"SplmtryData,omitempty"`
}

// Supporting choice and reference types
type OriginalGroupHeader17 struct {
	OriginalMessageID             string                           `xml:"OrgnlMsgId"`
	OriginalMessageNameID         string                           `xml:"OrgnlMsgNmId"`
	OriginalCreatedDateTime       *time.Time                       `xml:"OrgnlCreDtTm,omitempty"`
	OriginalNumberOfTransactions  *string                          `xml:"OrgnlNbOfTxs,omitempty"`
	OriginalControlSum            *float64                         `xml:"OrgnlCtrlSum,omitempty"`
	GroupStatus                   *string                          `xml:"GrpSts,omitempty"`
	StatusReasonInfo              []StatusReasonInfo12             `xml:"StsRsnInf,omitempty"`
	NumberOfTransactionsPerStatus []NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`
}

type OriginalGroupInfo29 struct {
	OriginalMessageID            string                `xml:"OrgnlMsgId"`
	OriginalMessageNameID        string                `xml:"OrgnlMsgNmId"`
	OriginalCreatedDateTime      *time.Time            `xml:"OrgnlCreDtTm,omitempty"`
	OriginalNumberOfTransactions *string               `xml:"OrgnlNbOfTxs,omitempty"`
	OriginalControlSum           *float64              `xml:"OrgnlCtrlSum,omitempty"`
	ReturnReason                 *PaymentReturnReason5 `xml:"RtrRsn,omitempty"`
}

// PaymentTransaction110 - for pacs.002.001.10 (exact XSD match)
type PaymentTransaction110 struct {
	StatusID                         *string                                       `xml:"StsId,omitempty"`
	OriginalGroupInfo                *OriginalGroupInfo29                          `xml:"OrgnlGrpInf,omitempty"`
	OriginalInstructionID            *string                                       `xml:"OrgnlInstrId,omitempty"`
	OriginalEndToEndID               *string                                       `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID            *string                                       `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                     *string                                       `xml:"OrgnlUETR,omitempty"`
	TransactionStatus                *string                                       `xml:"TxSts,omitempty"`
	StatusReasonInfo                 []StatusReasonInfo12                          `xml:"StsRsnInf,omitempty"`
	ChargesInfo                      []Charges7                                    `xml:"ChrgsInf,omitempty"`
	AcceptanceDateTime               *time.Time                                    `xml:"AccptncDtTm,omitempty"`
	EffectiveInterbankSettlementDate *DateAndDateTime2                             `xml:"FctvIntrBkSttlmDt,omitempty"`
	AccountServicerReference         *string                                       `xml:"AcctSvcrRef,omitempty"`
	ClearingSystemReference          *string                                       `xml:"ClrSysRef,omitempty"`
	InstructingAgent                 *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent                  *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	OriginalTransactionReference     *OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty"`
	SupplementaryData                []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// Additional base types
type Event2 struct {
	EventCode        string     `xml:"EvtCd"`
	EventParameter   []string   `xml:"EvtParam,omitempty"`
	EventDescription *string    `xml:"EvtDesc,omitempty"`
	EventTime        *time.Time `xml:"EvtTm,omitempty"`
}

type Acknowledgement1 struct {
	AcknowledgedMessageID string                  `xml:"AckdMsgId"`
	ReportOrError         AcknowledgementOrError2 `xml:"RptOrErr"`
}

type ProprietaryData6 struct {
	Type string           `xml:"Tp"`
	Data ProprietaryData5 `xml:"Data"`
}

type ProprietaryData5 struct {
	Envelope string `xml:"Envlp"`
}

// Additional choice types

type ReturnReason5 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type PaymentReturnReason5 struct {
	Reason                ReturnReason5 `xml:"Rsn"`
	AdditionalInformation []string      `xml:"AddtlInf,omitempty"`
}

type StatusReasonInfo12 struct {
	Originator            *PartyIdentification `xml:"Orgtr,omitempty"`
	Reason                *StatusReason6       `xml:"Rsn,omitempty"`
	AdditionalInformation []string             `xml:"AddtlInf,omitempty"`
}

type NumberOfTransactionsPerStatus5 struct {
	DetailedNumberOfTransactions string   `xml:"DtldNbOfTxs"`
	DetailedStatus               string   `xml:"DtldSts"`
	DetailedControlSum           *float64 `xml:"DtldCtrlSum,omitempty"`
}

type Pagination1 struct {
	PageNumber    string `xml:"PgNb"`
	LastPageIndex bool   `xml:"LastPgInd"`
}

type Authorization1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

// AccountReport25 - Bank to Customer Account Report according to CAMT.052.001.08 XSD
type AccountReport25 struct {
	ID                       string              `xml:"Id"`                     // Max35Text - required
	ReportPagination         *Pagination1        `xml:"RptPgntn,omitempty"`     // Optional
	ElectronicSequenceNumber *float64            `xml:"ElctrncSeqNb,omitempty"` // Number - optional
	ReportingSequence        *SequenceRange1     `xml:"RptgSeq,omitempty"`      // Optional
	LegalSequenceNumber      *float64            `xml:"LglSeqNb,omitempty"`     // Number - optional
	CreationDateTime         *time.Time          `xml:"CreDtTm,omitempty"`      // ISODateTime - optional
	FromToDate               *DateTimePeriod1    `xml:"FrToDt,omitempty"`       // Optional
	CopyDuplicateIndicator   *string             `xml:"CpyDplctInd,omitempty"`  // CopyDuplicate1Code - optional
	ReportingSource          *ReportingSource1   `xml:"RptgSrc,omitempty"`      // Optional
	Account                  CashAccount39       `xml:"Acct"`                   // Required
	RelatedAccount           *CashAccount38      `xml:"RltdAcct,omitempty"`     // Optional
	Interest                 []AccountInterest4  `xml:"Intrst,omitempty"`       // 0..unbounded
	Balance                  []CashBalance8      `xml:"Bal,omitempty"`          // 0..unbounded
	TransactionsSummary      *TotalTransactions6 `xml:"TxsSummry,omitempty"`    // Optional
	Entry                    []ReportEntry10     `xml:"Ntry,omitempty"`         // 0..unbounded
	AdditionalReportInfo     *string             `xml:"AddtlRptInf,omitempty"`  // Max500Text - optional
}

// AccountNotification19 - Alias for AccountNotification17 (version mismatch correction)
type AccountNotification19 = AccountNotification17

// AccountNotification17 - Bank to Customer Debit Credit Notification according to CAMT.054.001.08 XSD
type AccountNotification17 struct {
	ID                         string              `xml:"Id"`                       // Max35Text - required
	NotificationPagination     *Pagination1        `xml:"NtfctnPgntn,omitempty"`    // Optional
	ElectronicSequenceNumber   *float64            `xml:"ElctrncSeqNb,omitempty"`   // Number - optional
	ReportingSequence          *SequenceRange1     `xml:"RptgSeq,omitempty"`        // Optional
	LegalSequenceNumber        *float64            `xml:"LglSeqNb,omitempty"`       // Number - optional
	CreationDateTime           *time.Time          `xml:"CreDtTm,omitempty"`        // ISODateTime - optional
	FromToDate                 *DateTimePeriod1    `xml:"FrToDt,omitempty"`         // Optional
	CopyDuplicateIndicator     *string             `xml:"CpyDplctInd,omitempty"`    // CopyDuplicate1Code - optional
	ReportingSource            *ReportingSource1   `xml:"RptgSrc,omitempty"`        // Optional
	Account                    CashAccount39       `xml:"Acct"`                     // Required
	RelatedAccount             *CashAccount38      `xml:"RltdAcct,omitempty"`       // Optional
	Interest                   []AccountInterest4  `xml:"Intrst,omitempty"`         // 0..unbounded
	TransactionsSummary        *TotalTransactions6 `xml:"TxsSummry,omitempty"`      // Optional
	Entry                      []ReportEntry10     `xml:"Ntry,omitempty"`           // 0..unbounded
	AdditionalNotificationInfo *string             `xml:"AddtlNtfctnInf,omitempty"` // Max500Text - optional
}

// CaseAssignment5 - Case assignment for investigation messages
type CaseAssignment5 struct {
	ID               string    `xml:"Id"`      // Max35Text - required
	Assigner         Party40   `xml:"Assgnr"`  // Required
	Assignee         Party40   `xml:"Assgne"`  // Required
	CreationDateTime time.Time `xml:"CreDtTm"` // ISODateTime - required
}

// Case5 - Case information for investigation messages
type Case5 struct {
	ID                   string  `xml:"Id"`                       // Max35Text - required
	Creator              Party40 `xml:"Cretr"`                    // Required
	ReopenCaseIndication *bool   `xml:"ReopCaseIndctn,omitempty"` // YesNoIndicator - optional
}

// ControlData1 - Control data for investigations and cancellations
type ControlData1 struct {
	NumberOfTransactions string   `xml:"NbOfTxs"`           // Max15NumericText - required
	ControlSum           *float64 `xml:"CtrlSum,omitempty"` // DecimalNumber - optional
}

// UnderlyingTransaction21 - Underlying transaction information
type UnderlyingTransaction21 struct {
	// Simplified structure for now - can be expanded with specific transaction details
	OriginalGroupInfo          *OriginalGroupInfo3            `xml:"OrgnlGrpInfAndSts,omitempty"`
	OriginalPaymentInformation []OriginalPaymentInstruction32 `xml:"OrgnlPmtInfAndSts,omitempty"`
	TransactionInfo            []PaymentTransaction91         `xml:"TxInfAndSts,omitempty"`
}

// UnderlyingTransaction22 - Similar to UnderlyingTransaction21 but for different message types
type UnderlyingTransaction22 struct {
	// Simplified structure for now - can be expanded with specific transaction details
	OriginalGroupInfo          *OriginalGroupInfo3            `xml:"OrgnlGrpInfAndSts,omitempty"`
	OriginalPaymentInformation []OriginalPaymentInstruction32 `xml:"OrgnlPmtInfAndSts,omitempty"`
	TransactionInfo            []PaymentTransaction91         `xml:"TxInfAndSts,omitempty"`
}

// UnderlyingTransaction23 - Similar to UnderlyingTransaction21/22 but for different message types
type UnderlyingTransaction23 struct {
	// Simplified structure for now - can be expanded with specific transaction details
	OriginalGroupInfo          *OriginalGroupInfo3            `xml:"OrgnlGrpInfAndSts,omitempty"`
	OriginalPaymentInformation []OriginalPaymentInstruction32 `xml:"OrgnlPmtInfAndSts,omitempty"`
	TransactionInfo            []PaymentTransaction91         `xml:"TxInfAndSts,omitempty"`
}

// ReportingRequest5 - Reporting request information
type ReportingRequest5 struct {
	ReportingPeriod   *Period2                                      `xml:"RptgPrd,omitempty"`
	ReportingSequence *SequenceRange1                               `xml:"RptgSeq,omitempty"`
	RequestedMessage  []MessageIdentification2                      `xml:"ReqdMsgNmId,omitempty"`
	Account           []AccountIdentification4                      `xml:"Acct,omitempty"`
	Owner             *Party40                                      `xml:"AcctOwnr,omitempty"`
	Servicer          *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty"`
}

// UnableToApplyJustification3 - Justification for inability to apply
type UnableToApplyJustification3 struct {
	OriginalMessageIdentification string                        `xml:"OrgnlMsgId"`
	OriginalMessageNameID         string                        `xml:"OrgnlMsgNmId"`
	OriginalCreationDateTime      *time.Time                    `xml:"OrgnlCreDtTm,omitempty"`
	OriginalNumberOfTransactions  *string                       `xml:"OrgnlNbOfTxs,omitempty"`
	OriginalControlSum            *float64                      `xml:"OrgnlCtrlSum,omitempty"`
	UnableToApplyDetails          []UnableToApplyIncorrectInfo4 `xml:"UblToApplyDtls"`
}

// PaymentComplementaryInfo9 - Additional payment information
type PaymentComplementaryInfo9 struct {
	InstructionID           *string                                       `xml:"InstrId,omitempty"`
	EndToEndID              *string                                       `xml:"EndToEndId,omitempty"`
	TransactionID           *string                                       `xml:"TxId,omitempty"`
	UETR                    *string                                       `xml:"UETR,omitempty"`
	Amount                  *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt,omitempty"`
	RequestedExecutionDate  *string                                       `xml:"ReqdExctnDt,omitempty"`
	RequestedCollectionDate *string                                       `xml:"ReqdColltnDt,omitempty"`
	Debtor                  *PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DebtorAccount           *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent             *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	Creditor                *PartyIdentification135                       `xml:"Cdtr,omitempty"`
	CreditorAccount         *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	CreditorAgent           *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
}

// InvestigationStatus5 - Status of an investigation
type InvestigationStatus5 struct {
	ConfirmationStatus *string                      `xml:"ConfSts,omitempty"`   // ExternalInvestigationExecutionConfirmation1Code
	ModificationStatus *string                      `xml:"ModSts,omitempty"`    // ModificationStatus
	CancellationStatus *string                      `xml:"CxlSts,omitempty"`    // CancellationStatus
	AcceptedStatus     *string                      `xml:"AccptdSts,omitempty"` // AcceptedStatus
	PendingStatus      *PendingStatus4              `xml:"PndgSts,omitempty"`
	RejectedStatus     *RejectionReason31           `xml:"RjctdSts,omitempty"`
	DuplicateStatus    *DuplicateStatus             `xml:"DplctSts,omitempty"`
	ProprietaryStatus  *ProprietaryStatusAndReason6 `xml:"PrtrySts,omitempty"`
}

// PaymentTransaction91 - Payment transaction information
type PaymentTransaction91 struct {
	StatusID                     *string                                       `xml:"StsId,omitempty"`
	OriginalInstruction          *OriginalTransactionReference31               `xml:"OrgnlInstrRef,omitempty"`
	OriginalEndToEndID           *string                                       `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID        *string                                       `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                 *string                                       `xml:"OrgnlUETR,omitempty"`
	TransactionStatus            *string                                       `xml:"TxSts,omitempty"` // ExternalPaymentTransactionStatus1Code
	StatusReason                 []StatusReason6                               `xml:"StsRsnInf,omitempty"`
	ChargesInfo                  []Charges2                                    `xml:"ChrgsInf,omitempty"`
	AcceptanceDateTime           *time.Time                                    `xml:"AccptncDtTm,omitempty"`
	AccountServicerReference     *string                                       `xml:"AcctSvcrRef,omitempty"`
	ClearingSystemReference      *string                                       `xml:"ClrSysRef,omitempty"`
	InstructingAgent             *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent              *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	OriginalTransactionReference *OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty"`
	SupplementaryData            []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// ClaimNonReceipt2 - Claim for non-receipt of funds
type ClaimNonReceipt2 struct {
	DateAndPlaceOfBirth    *DateAndPlaceOfBirth1    `xml:"DtAndPlcOfBirth,omitempty"`
	ContactDetails         *Contact4                `xml:"CtctDtls,omitempty"`
	MissingOrIncorrectInfo *MissingOrIncorrectInfo3 `xml:"MssngOrIncrrctInf,omitempty"`
}

// StatementResolutionEntry4 - Statement resolution entry information
type StatementResolutionEntry4 struct {
	OriginalGroupInfo                *OriginalGroupInfo3         `xml:"OrgnlGrpInf,omitempty"`
	OriginalStatementID              *string                     `xml:"OrgnlStmtId,omitempty"`
	OriginalAccountServicerReference *string                     `xml:"OrgnlAcctSvcrRef,omitempty"`
	Account                          *CashAccount38              `xml:"Acct,omitempty"`
	RelatedAccount                   *CashAccount38              `xml:"RltdAcct,omitempty"`
	Statement                        []StatementResolutionEntry4 `xml:"Stmt,omitempty"`
}

// CorrectiveTransaction4 - Corrective transaction information
type CorrectiveTransaction4 struct {
	InstructionID             *string                                       `xml:"InstrId,omitempty"`
	EndToEndID                *string                                       `xml:"EndToEndId,omitempty"`
	TransactionID             *string                                       `xml:"TxId,omitempty"`
	UETR                      *string                                       `xml:"UETR,omitempty"`
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	InterbankSettlementDate   *time.Time                                    `xml:"IntrBkSttlmDt,omitempty"`
	RequestedCollectionDate   *time.Time                                    `xml:"ReqdColltnDt,omitempty"`
	RequestedExecutionDate    *time.Time                                    `xml:"ReqdExctnDt,omitempty"`
	Debtor                    *PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DebtorAccount             *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent               *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	Creditor                  *PartyIdentification135                       `xml:"Cdtr,omitempty"`
	CreditorAccount           *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	CreditorAgent             *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
}

// ResolutionData2 - Resolution data for investigation
type ResolutionData2 struct {
	EndToEndID                *string                            `xml:"EndToEndId,omitempty"`
	TransactionID             *string                            `xml:"TxId,omitempty"`
	UETR                      *string                            `xml:"UETR,omitempty"`
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`
	InterbankSettlementDate   *time.Time                         `xml:"IntrBkSttlmDt,omitempty"`
	ClearingChannel           *string                            `xml:"ClrChanl,omitempty"`
	DebtorName                *string                            `xml:"DbtrNm,omitempty"`
	CreditorName              *string                            `xml:"CdtrNm,omitempty"`
	CreditorReference         *CreditorReferenceInfo2            `xml:"CdtrRefInf,omitempty"`
}

// PaymentInstruction31 - Payment instruction information
type PaymentInstruction31 struct {
	PaymentInfoID             string                                        `xml:"PmtInfId"` // Required
	PaymentMethod             string                                        `xml:"PmtMtd"`   // Required - PaymentMethod3Code
	BatchBooking              *bool                                         `xml:"BtchBookg,omitempty"`
	NumberOfTransactions      *string                                       `xml:"NbOfTxs,omitempty"` // Max15NumericText
	ControlSum                *float64                                      `xml:"CtrlSum,omitempty"`
	PaymentTypeInfo           *PaymentTypeInfo19                            `xml:"PmtTpInf,omitempty"`
	RequestedExecutionDate    *time.Time                                    `xml:"ReqdExctnDt,omitempty"`
	Pooler                    *PartyIdentification135                       `xml:"Poolgr,omitempty"`
	Debtor                    *PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DebtorAccount             *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent               *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	CreditTransferTransaction []CreditTransferTransaction39                 `xml:"CdtTrfTxInf,omitempty"`
	SupplementaryData         []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// OriginalPaymentInstruction32 - Original payment instruction reference
type OriginalPaymentInstruction32 struct {
	OriginalPaymentInfoID             *string                            `xml:"OrgnlPmtInfId,omitempty"`
	OriginalInstructionID             *string                            `xml:"OrgnlInstrId,omitempty"`
	OriginalEndToEndID                *string                            `xml:"OrgnlEndToEndId,omitempty"`
	OriginalTransactionID             *string                            `xml:"OrgnlTxId,omitempty"`
	OriginalUETR                      *string                            `xml:"OrgnlUETR,omitempty"`
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`
	OriginalInterbankSettlementDate   *string                            `xml:"OrgnlIntrBkSttlmDt,omitempty"`
	ReversalReasonInformation         []PaymentReversalReason7           `xml:"RvslRsnInf,omitempty"`
	OriginalTransactionReference      *OriginalTransactionReference31    `xml:"OrgnlTxRef,omitempty"`
	SupplementaryData                 []SupplementaryData1               `xml:"SplmtryData,omitempty"`
}

// ResendSearchCriteria1 - Search criteria for resend request
type ResendSearchCriteria1 struct {
	MessageNameID     []string         `xml:"MsgNmId,omitempty"`   // Max35Text
	BusinessMessageID []string         `xml:"BizMsgIdr,omitempty"` // Max35Text
	CreationDateTime  *DateTimePeriod1 `xml:"CreDtTm,omitempty"`
	FileReference     []string         `xml:"FileRef,omitempty"` // Max35Text
	Originator        []Party44        `xml:"Orgtr,omitempty"`
	Destination       []Party44        `xml:"Dest,omitempty"`
}

// RequestHandling1 - Request handling information
type RequestHandling1 struct {
	Identification  string     `xml:"Id"`                // Max35Text - required
	RequestType     string     `xml:"ReqTp"`             // RequestType4 - required
	RequestDateTime *time.Time `xml:"ReqDtTm,omitempty"` // ISODateTime - required
	Description     *string    `xml:"Desc,omitempty"`    // Max350Text
	Reference       []string   `xml:"Ref,omitempty"`     // Max35Text
}

// RequestReportOrError1 - Request report or error information
type RequestReportOrError1 struct {
	BusinessReport   *string          `xml:"BizRpt,omitempty"` // Document content or reference
	OperationalError []ErrorHandling5 `xml:"OprlErr,omitempty"`
}

// AcknowledgementOrError2 - Acknowledgement or error information
type AcknowledgementOrError2 struct {
	AcknowledgementDetails *Acknowledgement1 `xml:"AckDtls,omitempty"`
	OperationalError       []ErrorHandling5  `xml:"OprlErr,omitempty"`
}

// OriginalTransactionReference28 - Original transaction reference information
type OriginalTransactionReference28 struct {
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty"`
	Amount                    *AmountType4                                  `xml:"Amt,omitempty"`
	InterbankSettlementDate   *time.Time                                    `xml:"IntrBkSttlmDt,omitempty"`
	RequestedCollectionDate   *time.Time                                    `xml:"ReqdColltnDt,omitempty"`
	RequestedExecutionDate    *time.Time                                    `xml:"ReqdExctnDt,omitempty"`
	CreditorSchemeID          *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty"`
	SettlementInfo            *SettlementInstruction7                       `xml:"SttlmInf,omitempty"`
	PaymentTypeInfo           *PaymentTypeInfo19                            `xml:"PmtTpInf,omitempty"`
	PaymentMethod             *string                                       `xml:"PmtMtd,omitempty"`
	MandateRelatedInfo        *MandateRelatedInfo14                         `xml:"MndtRltdInf,omitempty"`
	RemittanceInfo            *RemittanceInfo16                             `xml:"RmtInf,omitempty"`
	UltimateDebtor            *PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	Debtor                    *PartyIdentification135                       `xml:"Dbtr,omitempty"`
	DebtorAccount             *CashAccount38                                `xml:"DbtrAcct,omitempty"`
	DebtorAgent               *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	DebtorAgentAccount        *CashAccount38                                `xml:"DbtrAgtAcct,omitempty"`
	CreditorAgent             *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	CreditorAgentAccount      *CashAccount38                                `xml:"CdtrAgtAcct,omitempty"`
	Creditor                  *PartyIdentification135                       `xml:"Cdtr,omitempty"`
	CreditorAccount           *CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltimateCreditor          *PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	SupplementaryData         []SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// OriginalTransactionReference31 - Similar to OriginalTransactionReference28 but for different message types
type OriginalTransactionReference31 = OriginalTransactionReference28

type SequenceRange1 struct {
	FromSequence     *string  `xml:"FrSeq,omitempty"`
	ToSequence       *string  `xml:"ToSeq,omitempty"`
	EqualSequence    *string  `xml:"EQSeq,omitempty"`
	NotEqualSequence []string `xml:"NEQSeq,omitempty"`
}

type DateTimePeriod1 struct {
	FromDateTime *time.Time `xml:"FrDtTm,omitempty"`
	ToDateTime   *time.Time `xml:"ToDtTm,omitempty"`
}

type ReportingSource1 struct {
	Code        *string `xml:"Cd,omitempty"`
	Proprietary *string `xml:"Prtry,omitempty"`
}

type CashAccount39 struct {
	ID       AccountIdentification4       `xml:"Id"`
	Type     *CashAccountType2Choice      `xml:"Tp,omitempty"`
	Currency *string                      `xml:"Ccy,omitempty"`
	Name     *string                      `xml:"Nm,omitempty"`
	Proxy    *ProxyAccountIdentification1 `xml:"Prxy,omitempty"`
}

type AccountInterest4 struct {
	Type       *InterestType1   `xml:"Tp,omitempty"`
	Rate       []Rate4          `xml:"Rate,omitempty"`
	FromToDate *DateTimePeriod1 `xml:"FrToDt,omitempty"`
	Reason     *string          `xml:"Rsn,omitempty"`
	Tax        *TaxCharges2     `xml:"Tax,omitempty"`
}

type CashBalance8 struct {
	Type                 BalanceType10                     `xml:"Tp"`
	Line                 []BalanceSubType1                 `xml:"Line,omitempty"`
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator string                            `xml:"CdtDbtInd"`
	Date                 DateAndDateTime2                  `xml:"Dt"`
	Availability         []CashAvailability1               `xml:"Avlbty,omitempty"`
}

type TotalTransactions6 struct {
	TotalEntries       *NumberAndSumOfTransactions4 `xml:"TtlNtries,omitempty"`
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`
	TotalDebitEntries  *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`
}

type ReportEntry10 struct {
	EntryReference       *string                           `xml:"NtryRef,omitempty"`
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator string                            `xml:"CdtDbtInd"`
	Status               string                            `xml:"Sts"`
	BookingDate          *DateAndDateTime2                 `xml:"BookgDt,omitempty"`
	ValueDate            *DateAndDateTime2                 `xml:"ValDt,omitempty"`
	TransactionDetails   []EntryTransaction10              `xml:"NtryDtls,omitempty"`
	AdditionalEntryInfo  *string                           `xml:"AddtlNtryInf,omitempty"`
}

type AmountType4 struct {
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EquivalentAmount *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

type PaymentTypeInfo19 struct {
	InstructionPriority *string           `xml:"InstrPrty,omitempty"`
	ClearingChannel     *string           `xml:"ClrChanl,omitempty"`
	ServiceLevel        []ServiceLevel8   `xml:"SvcLvl,omitempty"`
	LocalInstrument     *LocalInstrument2 `xml:"LclInstrm,omitempty"`
	SequenceType        *string           `xml:"SeqTp,omitempty"`
	CategoryPurpose     *CategoryPurpose1 `xml:"CtgyPurp,omitempty"`
}

type MandateRelatedInfo14 struct {
	MandateID            *string                 `xml:"MndtId,omitempty"`
	DateOfSignature      *string                 `xml:"DtOfSgntr,omitempty"`
	AmentmentIndicator   *bool                   `xml:"AmdmntInd,omitempty"`
	AmendmentInfoDetails *AmendmentInfoDetails13 `xml:"AmdmntInfDtls,omitempty"`
	ElectronicSignature  *string                 `xml:"ElctrncSgntr,omitempty"`
	FirstCollectionDate  *string                 `xml:"FrstColltnDt,omitempty"`
	FinalCollectionDate  *string                 `xml:"FnlColltnDt,omitempty"`
	Frequency            *string                 `xml:"Frqcy,omitempty"`
	Reason               *MandateSetupReason1    `xml:"Rsn,omitempty"`
	TrackingDays         *string                 `xml:"TrckgDays,omitempty"`
}

// Simplified placeholders for complex types - these can be expanded later
// InterestType1 - Interest type selection
type InterestType1 struct {
	Code        *string `xml:"Cd,omitempty"`    // InterestType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// Rate4 - Interest rate information
type Rate4 struct {
	Type          *RateType4                               `xml:"Tp,omitempty"`
	ValidityRange *ActiveOrHistoricCurrencyAndAmountRange2 `xml:"VldtyRg,omitempty"`
	Rate          *float64                                 `xml:"Rate,omitempty"` // PercentageRate
}

// TaxCharges2 - Tax charges information
type TaxCharges2 struct {
	ID     *string                            `xml:"Id,omitempty"`
	Rate   *float64                           `xml:"Rate,omitempty"` // PercentageRate
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// BalanceType10 - Balance type selection
type BalanceType10 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalBalanceType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// BalanceSubType1 - Balance sub-type selection
type BalanceSubType1 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalBalanceSubType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// CashAvailability1 - Cash availability information
type CashAvailability1 struct {
	Date                 DateAndDateTime2                  `xml:"Dt"`
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator string                            `xml:"CdtDbtInd"`
}

// NumberAndSumOfTransactions4 - Number and sum of transactions
type NumberAndSumOfTransactions4 struct {
	NumberOfEntries *string                `xml:"NbOfNtries,omitempty"` // Max15NumericText
	Sum             *float64               `xml:"Sum,omitempty"`        // DecimalNumber
	TotalNetEntry   *TotalNetEntryDetails1 `xml:"TtlNetNtry,omitempty"`
}

// NumberAndSumOfTransactions1 - Simplified number and sum
type NumberAndSumOfTransactions1 struct {
	NumberOfEntries *string  `xml:"NbOfNtries,omitempty"` // Max15NumericText
	Sum             *float64 `xml:"Sum,omitempty"`        // DecimalNumber
}

// EntryTransaction10 - Entry transaction details
type EntryTransaction10 struct {
	References                        *TransactionReferences6            `xml:"Refs,omitempty"`
	Amount                            *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CreditDebitIndicator              *string                            `xml:"CdtDbtInd,omitempty"`
	AmountDetails                     *AmountAndCurrencyExchange3        `xml:"AmtDtls,omitempty"`
	Availability                      []CashAvailability1                `xml:"Avlbty,omitempty"`
	BankTransactionCode               *BankTransactionCodeStructure4     `xml:"BkTxCd,omitempty"`
	Charges                           *Charges6                          `xml:"Chrgs,omitempty"`
	TechnicalInputChannel             *string                            `xml:"TechInptChanl,omitempty"`
	Interest                          *TransactionInterest4              `xml:"Intrst,omitempty"`
	RelatedParties                    *TransactionParties6               `xml:"RltdPties,omitempty"`
	RelatedAgents                     *TransactionAgents5                `xml:"RltdAgts,omitempty"`
	LegalSequenceNumber               *float64                           `xml:"LglSeqNb,omitempty"`
	Purpose                           *Purpose2                          `xml:"Purp,omitempty"`
	RelatedRemittanceInfo             *RemittanceLocation7               `xml:"RltdRmtInf,omitempty"`
	RemittanceInfo                    *RemittanceInfo16                  `xml:"RmtInf,omitempty"`
	RelatedDates                      *TransactionDates3                 `xml:"RltdDts,omitempty"`
	RelatedPrice                      *TransactionPrice4                 `xml:"RltdPric,omitempty"`
	RelatedQuantities                 *TransactionQuantities3            `xml:"RltdQties,omitempty"`
	FinancialInstrumentIdentification *SecurityIdentification19          `xml:"FinInstrmId,omitempty"`
	Tax                               *TaxInfo8                          `xml:"Tax,omitempty"`
	ReturnInfo                        *PaymentReturnReason5              `xml:"RtrInf,omitempty"`
	CorporateAction                   *CorporateActionInfo2              `xml:"CorpActn,omitempty"`
	SafekeepingPlace                  *SafekeepingPlaceFormat28          `xml:"SfkpgPlc,omitempty"`
	AdditionalTransactionInfo         *string                            `xml:"AddtlTxInf,omitempty"`
	SupplementaryData                 []SupplementaryData1               `xml:"SplmtryData,omitempty"`
}

// EquivalentAmount2 - Equivalent amount in different currency
type EquivalentAmount2 struct {
	Amount             ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CurrencyOfTransfer string                            `xml:"CcyOfTrf"`
}

// ServiceLevel8 - Service level for payment instructions
type ServiceLevel8 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalServiceLevel1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// LocalInstrument2 - Local clearing system instrument
type LocalInstrument2 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalLocalInstrument1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// CategoryPurpose1 - Category purpose for payments
type CategoryPurpose1 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalCategoryPurpose1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// AmendmentInfoDetails13 - Amendment information for mandates
type AmendmentInfoDetails13 struct {
	OriginalMandateID            *string                                       `xml:"OrgnlMndtId,omitempty"`
	OriginalCreditorSchemeID     *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty"`
	OriginalCreditorAgent        *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty"`
	OriginalCreditorAgentAccount *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OriginalDebtor               *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty"`
	OriginalDebtorAccount        *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty"`
	OriginalDebtorAgent          *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty"`
	OriginalDebtorAgentAccount   *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OriginalFinalCollectionDate  *string                                       `xml:"OrgnlFnlColltnDt,omitempty"`
	OriginalFrequency            *Frequency36                                  `xml:"OrgnlFrqcy,omitempty"`
	OriginalReason               *MandateSetupReason1                          `xml:"OrgnlRsn,omitempty"`
	OriginalTrackingDays         *string                                       `xml:"OrgnlTrckgDays,omitempty"`
}

// MandateSetupReason1 - Reason for mandate setup
type MandateSetupReason1 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalMandateSetupReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max70Text
}

// CashAccountType2Choice - Choice between account type code or proprietary
type CashAccountType2Choice struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalCashAccountType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// Additional missing types - proper XSD-based implementations

// Party40 - Choice between party identification or agent
type Party40 struct {
	Party *PartyIdentification135                       `xml:"Pty,omitempty"`
	Agent *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
}

// OriginalGroupInfo3 - Original group information for investigations
type OriginalGroupInfo3 struct {
	OriginalMessageID            string     `xml:"OrgnlMsgId"`
	OriginalMessageNameID        string     `xml:"OrgnlMsgNmId"`
	OriginalCreationDateTime     *time.Time `xml:"OrgnlCreDtTm,omitempty"`
	OriginalNumberOfTransactions *string    `xml:"OrgnlNbOfTxs,omitempty"`
	OriginalControlSum           *float64   `xml:"OrgnlCtrlSum,omitempty"`
	GroupCancellationID          *string    `xml:"GrpCxlId,omitempty"`
}

// Period2 - Period specification choice
type Period2 struct {
	FromDateTime *time.Time `xml:"FrDtTm,omitempty"`
	ToDateTime   *time.Time `xml:"ToDtTm,omitempty"`
	FromDate     *string    `xml:"FrDt,omitempty"`
	ToDate       *string    `xml:"ToDt,omitempty"`
}

// MessageIdentification2 - Message identification
type MessageIdentification2 struct {
	MessageNameID string `xml:"MsgNmId"` // Max35Text
	MessageID     string `xml:"MsgId"`   // Max35Text
}

// UnableToApplyIncorrectInfo4 - Information about unable to apply
type UnableToApplyIncorrectInfo4 struct {
	MissingOrIncorrectInfo       *MissingOrIncorrectInfo3 `xml:"MssngOrIncrrctInf,omitempty"`
	PossibleDuplicateInstruction *bool                    `xml:"PssblDplctInstr,omitempty"`
}

// PendingStatus4 - Status information for pending items
type PendingStatus4 struct {
	Reason                *PendingReason16 `xml:"Rsn,omitempty"`
	AdditionalInformation *string          `xml:"AddtlInf,omitempty"` // Max105Text
}

// RejectionReason31 - Rejection reason choice
type RejectionReason31 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalStatusReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// DuplicateStatus - Duplicate status information
type DuplicateStatus struct {
	DuplicateOf *string `xml:"DplctOf,omitempty"` // Max35Text
}

// ProprietaryStatusAndReason6 - Proprietary status and reason
type ProprietaryStatusAndReason6 struct {
	ProprietaryStatus string   `xml:"PrtrySts"`      // Max35Text
	Reason            []string `xml:"Rsn,omitempty"` // Max35Text
}

// StatusReason6 - Status reason choice
type StatusReason6 struct {
	RejectionReason       *RejectionReason31    `xml:"RjctnRsn,omitempty"`
	CancellationReason    *CancellationReason33 `xml:"CxlRsn,omitempty"`
	ModificationReason    *ModificationReason2  `xml:"ModRsn,omitempty"`
	AdditionalInformation []string              `xml:"AddtlInf,omitempty"` // Max105Text
}

// Charges2 - Charges information
type Charges2 struct {
	Amount ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agent  BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// PaymentReversalReason7 - Payment reversal reason
type PaymentReversalReason7 struct {
	Originator            *PartyIdentification135 `xml:"Orgtr,omitempty"`
	Reason                *ReversalReason4        `xml:"Rsn,omitempty"`
	AdditionalInformation []string                `xml:"AddtlInf,omitempty"` // Max105Text
}

// MissingOrIncorrectInfo3 - Missing or incorrect information details
type MissingOrIncorrectInfo3 struct {
	AMLRequest    *bool                     `xml:"AMLReq,omitempty"`
	MissingInfo   []UnableToApplyMissing1   `xml:"MssngInf,omitempty"`
	IncorrectInfo []UnableToApplyIncorrect1 `xml:"IncrrctInf,omitempty"`
}

// ErrorHandling5 - Error handling information
type ErrorHandling5 struct {
	ErrorCode   string  `xml:"Err"`            // ExternalSystemErrorHandling1Code
	Description *string `xml:"Desc,omitempty"` // Max140Text
}

// Party44 - Similar to Party40 but for different contexts (ADMI messages)
type Party44 struct {
	OrganisationIdentification *PartyIdentification135                       `xml:"OrgId,omitempty"`
	FinancialInstitutionID     *BranchAndFinancialInstitutionIdentification6 `xml:"FIId,omitempty"`
}

// ValidationError represents a validation error with field context
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Field '%s': %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (errs ValidationErrors) Error() string {
	if len(errs) == 0 {
		return ""
	}
	var messages []string
	for _, err := range errs {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

func (errs ValidationErrors) HasErrors() bool {
	return len(errs) > 0
}

// validateRequired checks if a field has a non-zero value
func validateRequired(value interface{}, fieldName string) error {
	v := reflect.ValueOf(value)

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ValidationError{Field: fieldName, Message: "is required but is nil"}
		}
		v = v.Elem()
	}

	// Check if value is zero
	if v.IsZero() {
		return ValidationError{Field: fieldName, Message: "is required but is empty"}
	}

	return nil
}

// validateStringLength validates string length constraints from XSD
func validateStringLength(value string, minLen, maxLen int, fieldName string) error {
	length := len(value)
	if length < minLen {
		return ValidationError{Field: fieldName, Message: fmt.Sprintf("length %d is below minimum %d", length, minLen)}
	}
	if length > maxLen {
		return ValidationError{Field: fieldName, Message: fmt.Sprintf("length %d exceeds maximum %d", length, maxLen)}
	}
	return nil
}

// validatePattern validates string against regex pattern from XSD
func validatePattern(value string, pattern string, fieldName string) error {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return ValidationError{Field: fieldName, Message: fmt.Sprintf("pattern validation failed: %s", err.Error())}
	}
	if !matched {
		return ValidationError{Field: fieldName, Message: fmt.Sprintf("does not match required pattern '%s'", pattern)}
	}
	return nil
}

// validateEnumeration validates that value is in allowed enumeration values
func validateEnumeration(value string, validValues []string, fieldName string) error {
	for _, valid := range validValues {
		if value == valid {
			return nil
		}
	}
	return ValidationError{Field: fieldName, Message: fmt.Sprintf("'%s' is not a valid enumeration value", value)}
}

// validateCurrency validates currency code format (ISO 4217)
func validateCurrency(code string, fieldName string) error {
	if err := validateStringLength(code, 3, 3, fieldName); err != nil {
		return err
	}
	// Pattern for ISO 4217 currency codes (3 uppercase letters)
	return validatePattern(code, `^[A-Z]{3}$`, fieldName)
}

// validateCountryCode validates country code format (ISO 3166-1 alpha-2)
func validateCountryCode(code string, fieldName string) error {
	if err := validateStringLength(code, 2, 2, fieldName); err != nil {
		return err
	}
	// Pattern for ISO 3166-1 alpha-2 country codes (2 uppercase letters)
	return validatePattern(code, `^[A-Z]{2}$`, fieldName)
}

// validateBIC validates BIC (Bank Identifier Code) format
func validateBIC(bic string, fieldName string) error {
	if err := validateStringLength(bic, 8, 11, fieldName); err != nil {
		return err
	}
	// BIC pattern: 4 letters (institution), 2 letters (country), 2 alphanumeric (location), optional 3 alphanumeric (branch)
	return validatePattern(bic, `^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`, fieldName)
}

// validateIBAN validates IBAN format
func validateIBAN(iban string, fieldName string) error {
	if err := validateStringLength(iban, 15, 34, fieldName); err != nil {
		return err
	}
	// Basic IBAN pattern (country code + 2 digits + up to 30 alphanumeric)
	return validatePattern(iban, `^[A-Z]{2}[0-9]{2}[A-Z0-9]{1,30}$`, fieldName)
}

// validateLEI validates Legal Entity Identifier format
func validateLEI(lei string, fieldName string) error {
	if err := validateStringLength(lei, 20, 20, fieldName); err != nil {
		return err
	}
	// LEI pattern: 18 alphanumeric characters + 2 check digits
	return validatePattern(lei, `^[A-Z0-9]{18}[0-9]{2}$`, fieldName)
}

// validateUUID validates UUID v4 format
func validateUUID(uuid string, fieldName string) error {
	if err := validateStringLength(uuid, 36, 36, fieldName); err != nil {
		return err
	}
	// UUID v4 pattern
	return validatePattern(uuid, `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`, fieldName)
}

// validateDate validates date string in YYYY-MM-DD format
func validateDate(date string, fieldName string) error {
	if date == "" {
		return ValidationError{Field: fieldName, Message: "date is required"}
	}

	// Check format using regex
	if err := validatePattern(date, `^\d{4}-\d{2}-\d{2}$`, fieldName); err != nil {
		return ValidationError{Field: fieldName, Message: "date must be in YYYY-MM-DD format"}
	}

	// Parse and validate actual date values
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ValidationError{Field: fieldName, Message: "invalid date value"}
	}

	return nil
}

// validateDateTime validates datetime string in YYYY-MM-DDTHH:MM:SS or YYYY-MM-DDTHH:MM:SS.000Z format
func validateDateTime(dateTime string, fieldName string) error {
	if dateTime == "" {
		return ValidationError{Field: fieldName, Message: "datetime is required"}
	}

	// Check format using regex (supports both with and without milliseconds/timezone)
	if err := validatePattern(dateTime, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3}Z?)?$`, fieldName); err != nil {
		return ValidationError{Field: fieldName, Message: "datetime must be in YYYY-MM-DDTHH:MM:SS format"}
	}

	// Parse and validate actual datetime values
	layouts := []string{
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05.000Z",
	}

	var parseErr error
	for _, layout := range layouts {
		_, parseErr = time.Parse(layout, dateTime)
		if parseErr == nil {
			break
		}
	}

	if parseErr != nil {
		return ValidationError{Field: fieldName, Message: "invalid datetime value"}
	}

	return nil
}

// DateString creates a date string in YYYY-MM-DD format
func DateString(year int, month, day int) string {
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// DateTimeString creates a datetime string in YYYY-MM-DDTHH:MM:SS format
func DateTimeString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}

// Validate ActiveCurrencyAndAmount according to XSD constraints
func (a *ActiveCurrencyAndAmount) Validate() error {
	var errs ValidationErrors

	// Value validation - must be positive
	if a.Value <= 0 {
		errs = append(errs, ValidationError{Field: "Value", Message: "must be positive"})
	}

	// Currency validation (ISO 4217 - 3 letter code)
	if a.Currency == "" {
		errs = append(errs, ValidationError{Field: "Ccy", Message: "is required"})
	} else if err := validateCurrency(a.Currency, "Ccy"); err != nil {
		errs = append(errs, err.(ValidationError))
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate GroupHeader93 according to XSD constraints
func (g *GroupHeader93) Validate() error {
	var errs ValidationErrors

	// MessageID: Max35Text (required)
	if err := validateRequired(g.MessageID, "MsgId"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else if err := validateStringLength(g.MessageID, 1, 35, "MsgId"); err != nil {
		errs = append(errs, err.(ValidationError))
	}

	// NumberOfTransactions: Max15NumericText (required)
	if err := validateRequired(g.NumberOfTransactions, "NbOfTxs"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := validatePattern(g.NumberOfTransactions, `^[0-9]{1,15}$`, "NbOfTxs"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	// CreationDateTime is required
	if g.CreationDateTime == nil {
		errs = append(errs, ValidationError{Field: "CreDtTm", Message: "is required"})
	} else {
		// time.Time is already validated by Go's type system
	}

	// SettlementInfo validation
	if err := validateRequired(g.SettlementInfo, "SttlmInf"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := validateRequired(g.SettlementInfo.SettlementMethod, "SttlmMtd"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate Pacs00800108Document according to pacs.008.001.08 XSD
func (d *Pacs00800108Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.FICustomerCreditTransfer, "FIToFICstmrCdtTrf"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate GroupHeader (required)
		if err := d.FICustomerCreditTransfer.GroupHeader.Validate(); err != nil {
			if valErrs, ok := err.(ValidationErrors); ok {
				errs = append(errs, valErrs...)
			} else {
				errs = append(errs, ValidationError{Field: "GrpHdr", Message: err.Error()})
			}
		}

		// Validate CreditTransferTransactionInfo array (required, min 1)
		if len(d.FICustomerCreditTransfer.CreditTransferTransactionInfo) == 0 {
			errs = append(errs, ValidationError{Field: "CdtTrfTxInf", Message: "at least one transaction is required"})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to camt.052.001.08 XSD
func (d *Camt05200108Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.BankAccountReport, "BkToCstmrAcctRpt"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.BankAccountReport
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to camt.054.001.08 XSD
func (d *Camt05400108Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.BankDebitCreditNotification, "BkToCstmrDbtCdtNtfctn"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.BankDebitCreditNotification
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to camt.060.001.05 XSD
func (d *Camt06000105Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.AccountReportingRequest, "AcctRptgReq"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.AccountReportingRequest
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to pain.013.001.07 XSD
func (d *Pain01300107Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.CreditorPaymentActivationRequest, "CdtrPmtActvtnReq"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.CreditorPaymentActivationRequest
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to pain.014.001.07 XSD
func (d *Pain01400107Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.CreditorPaymentActivationStatusReport, "CdtrPmtActvtnReqStsRpt"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.CreditorPaymentActivationStatusReport
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to admi.004.001.02 XSD
func (d *Admi00400102Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.SystemEventNotification, "SysEvtNtfctn"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.SystemEventNotification
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to admi.011.001.01 XSD
func (d *Admi01100101Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.SystemEventAcknowledgement, "SysEvtAck"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.SystemEventAcknowledgement
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to admi.006.001.01 XSD
func (d *Admi00600101Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.ResendRequest, "RsndReq"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.ResendRequest
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to admi.007.001.01 XSD
func (d *Admi00700101Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.ReceiptAcknowledgement, "RctAck"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.ReceiptAcknowledgement
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs comprehensive validation according to admi.998.001.02 XSD
func (d *Admi99800102Document) Validate() error {
	var errs ValidationErrors

	// Validate required fields
	if err := validateRequired(d.AdministrationMessage, "AdmstnPrtryMsg"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Validate nested structure - placeholder for now
		// TODO: Implement full validation based on XSD constraints
		_ = d.AdministrationMessage
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Additional missing types for compilation - XSD-based implementations

// RateType4 - Rate type selection
type RateType4 struct {
	Percentage *float64 `xml:"Pctg,omitempty"` // PercentageRate
	Other      *string  `xml:"Othr,omitempty"` // Max35Text
}

// ActiveOrHistoricCurrencyAndAmountRange2 - Currency amount range
type ActiveOrHistoricCurrencyAndAmountRange2 struct {
	Amount   AmountRangeBoundary1 `xml:"Amt"`
	Currency string               `xml:"Ccy"` // ActiveOrHistoricCurrencyCode
}

// AmountRangeBoundary1 - Amount range boundary
type AmountRangeBoundary1 struct {
	BoundaryAmount float64 `xml:"BdryAmt"` // DecimalNumber
	Included       bool    `xml:"Incl"`    // YesNoIndicator
}

// TotalNetEntryDetails1 - Total net entry details
type TotalNetEntryDetails1 struct {
	NumberOfEntries *string               `xml:"NbOfNtries,omitempty"` // Max15NumericText
	Sum             *float64              `xml:"Sum,omitempty"`        // DecimalNumber
	TotalNetEntry   *AmountAndDirection35 `xml:"TtlNetNtry,omitempty"`
}

// AmountAndDirection35 - Amount with debit/credit direction
type AmountAndDirection35 struct {
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator string                            `xml:"CdtDbtInd"` // CreditDebitCode
}

// TransactionReferences6 - Transaction reference information
type TransactionReferences6 struct {
	MessageID                         *string `xml:"MsgId,omitempty"`             // Max35Text
	AccountServicerRef                *string `xml:"AcctSvcrRef,omitempty"`       // Max35Text
	PaymentInfoID                     *string `xml:"PmtInfId,omitempty"`          // Max35Text
	InstructionID                     *string `xml:"InstrId,omitempty"`           // Max35Text
	EndToEndID                        *string `xml:"EndToEndId,omitempty"`        // Max35Text
	TransactionID                     *string `xml:"TxId,omitempty"`              // Max35Text
	MandateID                         *string `xml:"MndtId,omitempty"`            // Max35Text
	CheckNumber                       *string `xml:"ChqNb,omitempty"`             // Max35Text
	ClearingSystemRef                 *string `xml:"ClrSysRef,omitempty"`         // Max35Text
	AccountOwnerTransactionID         *string `xml:"AcctOwnrTxId,omitempty"`      // Max35Text
	AccountServicerTransactionID      *string `xml:"AcctSvcrTxId,omitempty"`      // Max35Text
	MarketInfrastructureTransactionID *string `xml:"MktInfrstrctrTxId,omitempty"` // Max35Text
	ProcessingID                      *string `xml:"PrcgId,omitempty"`            // Max35Text
}

// AmountAndCurrencyExchange3 - Amount with currency exchange
type AmountAndCurrencyExchange3 struct {
	InstructedAmount       *AmountAndCurrencyExchangeDetails4  `xml:"InstdAmt,omitempty"`
	TransactionAmount      *AmountAndCurrencyExchangeDetails4  `xml:"TxAmt,omitempty"`
	CounterValueAmount     *AmountAndCurrencyExchangeDetails4  `xml:"CntrValAmt,omitempty"`
	AnnouncedPostingAmount *AmountAndCurrencyExchangeDetails4  `xml:"AnncdPstngAmt,omitempty"`
	ProprietaryAmount      []AmountAndCurrencyExchangeDetails5 `xml:"PrtryAmt,omitempty"`
}

// AmountAndCurrencyExchangeDetails4 - Amount and exchange details
type AmountAndCurrencyExchangeDetails4 struct {
	Amount           ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CurrencyExchange *CurrencyExchange5                `xml:"CcyXchg,omitempty"`
}

// AmountAndCurrencyExchangeDetails5 - Proprietary amount details
type AmountAndCurrencyExchangeDetails5 struct {
	Amount           ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CurrencyExchange *CurrencyExchange5                `xml:"CcyXchg,omitempty"`
	Type             string                            `xml:"Tp"` // Max35Text
}

// CurrencyExchange5 - Currency exchange information
type CurrencyExchange5 struct {
	SourceCurrency string   `xml:"SrcCcy"`             // ActiveOrHistoricCurrencyCode
	TargetCurrency *string  `xml:"TrgtCcy,omitempty"`  // ActiveOrHistoricCurrencyCode
	UnitCurrency   *string  `xml:"UnitCcy,omitempty"`  // ActiveOrHistoricCurrencyCode
	ExchangeRate   *float64 `xml:"XchgRate,omitempty"` // BaseOneRate
	ContractID     *string  `xml:"CtrctId,omitempty"`  // Max35Text
	QuotationDate  *string  `xml:"QtnDt,omitempty"`    // ISODate
}

// BankTransactionCodeStructure4 - Bank transaction code structure
type BankTransactionCodeStructure4 struct {
	Domain    BankTransactionCodeStructure5  `xml:"Domn"`
	Family    BankTransactionCodeStructure6  `xml:"Fmly"`
	SubFamily *BankTransactionCodeStructure7 `xml:"SubFmlyCd,omitempty"`
}

// BankTransactionCodeStructure5 - Bank transaction domain
type BankTransactionCodeStructure5 struct {
	Code   string `xml:"Cd"`   // ExternalBankTransactionDomain1Code
	Family string `xml:"Fmly"` // ExternalBankTransactionFamily1Code
}

// BankTransactionCodeStructure6 - Bank transaction family
type BankTransactionCodeStructure6 struct {
	Code          string `xml:"Cd"`        // ExternalBankTransactionFamily1Code
	SubFamilyCode string `xml:"SubFmlyCd"` // ExternalBankTransactionSubFamily1Code
}

// BankTransactionCodeStructure7 - Bank transaction sub-family
type BankTransactionCodeStructure7 struct {
	Code string `xml:"Cd"` // ExternalBankTransactionSubFamily1Code
}

// Charges6 - Charges information
type Charges6 struct {
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`
	Record                   []ChargesRecord3                   `xml:"Rcrd,omitempty"`
}

// ChargesRecord3 - Individual charge record
type ChargesRecord3 struct {
	Amount                   ActiveOrHistoricCurrencyAndAmount             `xml:"Amt"`
	CreditDebitIndicator     *string                                       `xml:"CdtDbtInd,omitempty"` // CreditDebitCode
	ChargesIncludedIndicator *bool                                         `xml:"ChrgInclInd,omitempty"`
	Type                     *ChargeType3                                  `xml:"Tp,omitempty"`
	Rate                     *float64                                      `xml:"Rate,omitempty"` // PercentageRate
	Bearer                   *ChargeBearerType1Code                        `xml:"Br,omitempty"`
	Agent                    *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty"`
	Tax                      *TaxCharges2                                  `xml:"Tax,omitempty"`
}

// ChargeType3 - Charge type selection
type ChargeType3 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalChargeType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// ChargeBearerType1Code - Charge bearer type
type ChargeBearerType1Code string

// TransactionInterest4 - Transaction interest information
type TransactionInterest4 struct {
	TotalInterestAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlIntrsTAndTaxAmt,omitempty"`
	Record                    []InterestRecord2                  `xml:"Rcrd,omitempty"`
}

// InterestRecord2 - Individual interest record
type InterestRecord2 struct {
	Amount               ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CreditDebitIndicator string                            `xml:"CdtDbtInd"` // CreditDebitCode
	Type                 *InterestType1                    `xml:"Tp,omitempty"`
	Rate                 *Rate4                            `xml:"Rate,omitempty"`
	FromToDate           *DateTimePeriod1                  `xml:"FrToDt,omitempty"`
	Reason               *string                           `xml:"Rsn,omitempty"` // Max35Text
	Tax                  *TaxCharges2                      `xml:"Tax,omitempty"`
}

// TransactionParties6 - Transaction parties
type TransactionParties6 struct {
	InitiatingParty  *PartyIdentification135 `xml:"InitgPty,omitempty"`
	Debtor           *PartyIdentification135 `xml:"Dbtr,omitempty"`
	DebtorAccount    *CashAccount38          `xml:"DbtrAcct,omitempty"`
	UltimateDebtor   *PartyIdentification135 `xml:"UltmtDbtr,omitempty"`
	Creditor         *PartyIdentification135 `xml:"Cdtr,omitempty"`
	CreditorAccount  *CashAccount38          `xml:"CdtrAcct,omitempty"`
	UltimateCreditor *PartyIdentification135 `xml:"UltmtCdtr,omitempty"`
	TradingParty     *PartyIdentification135 `xml:"TradgPty,omitempty"`
	Proprietary      []ProprietaryParty5     `xml:"Prtry,omitempty"`
}

// ProprietaryParty5 - Proprietary party information
type ProprietaryParty5 struct {
	Type  string                 `xml:"Tp"` // Max35Text
	Party PartyIdentification135 `xml:"Pty"`
}

// Frequency36 - Frequency choice
type Frequency36 struct {
	Type        *Frequency6Code      `xml:"Tp,omitempty"`
	Period      *FrequencyPeriod1    `xml:"Prd,omitempty"`
	PointInTime *FrequencyAndMoment1 `xml:"PtInTm,omitempty"`
}

// Frequency6Code - Frequency code
type Frequency6Code string

// FrequencyPeriod1 - Frequency period
type FrequencyPeriod1 struct {
	Type           Frequency6Code `xml:"Tp"`
	CountPerPeriod int            `xml:"CntPerPrd"`
}

// FrequencyAndMoment1 - Frequency and moment
type FrequencyAndMoment1 struct {
	Type        Frequency6Code    `xml:"Tp"`
	PointInTime Exact2NumericText `xml:"PtInTm"`
}

// Exact2NumericText - Exactly 2 numeric characters
type Exact2NumericText string

// Additional missing transaction and security types

// TransactionAgents5 - Transaction agents
type TransactionAgents5 struct {
	InstructingAgent   *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty"`
	InstructedAgent    *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty"`
	DebtorAgent        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty"`
	CreditorAgent      *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty"`
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	ReceivingAgent     *BranchAndFinancialInstitutionIdentification6 `xml:"RcvgAgt,omitempty"`
	DeliveringAgent    *BranchAndFinancialInstitutionIdentification6 `xml:"DlvrgAgt,omitempty"`
	IssuingAgent       *BranchAndFinancialInstitutionIdentification6 `xml:"IssgAgt,omitempty"`
	SettlementAgent    *BranchAndFinancialInstitutionIdentification6 `xml:"SttlmAgt,omitempty"`
	Proprietary        []ProprietaryAgent4                           `xml:"Prtry,omitempty"`
}

// ProprietaryAgent4 - Proprietary agent information
type ProprietaryAgent4 struct {
	Type  string                                       `xml:"Tp"` // Max35Text
	Agent BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// RemittanceLocation7 - Remittance location details
type RemittanceLocation7 struct {
	RemittanceID              *string                   `xml:"RmtId,omitempty"` // Max35Text
	RemittanceLocationDetails []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty"`
}

// RemittanceLocationData1 - Remittance location data
type RemittanceLocationData1 struct {
	Method            string           `xml:"Mtd"`                  // RemittanceLocationMethod2Code
	ElectronicAddress *string          `xml:"ElctrncAdr,omitempty"` // Max2048Text
	PostalAddress     *PostalAddress24 `xml:"PstlAdr,omitempty"`
}

// TransactionDates3 - Transaction dates
type TransactionDates3 struct {
	AcceptanceDateTime                  *time.Time         `xml:"AccptncDtTm,omitempty"`            // ISODateTime
	TradeActivityContractSettlementDate *string            `xml:"TradActvtyCtrctSttlmDt,omitempty"` // ISODate
	TradeDate                           *string            `xml:"TradDt,omitempty"`                 // ISODate
	InterbankSettlementDate             *time.Time         `xml:"IntrBkSttlmDt,omitempty"`          // ISODate
	StartDate                           *string            `xml:"StartDt,omitempty"`                // ISODate
	EndDate                             *string            `xml:"EndDt,omitempty"`                  // ISODate
	TransactionDateTime                 *time.Time         `xml:"TxDtTm,omitempty"`                 // ISODateTime
	Proprietary                         []ProprietaryDate3 `xml:"Prtry,omitempty"`
}

// ProprietaryDate3 - Proprietary date information
type ProprietaryDate3 struct {
	Type     string     `xml:"Tp"`             // Max35Text
	Date     *string    `xml:"Dt,omitempty"`   // ISODate
	DateTime *time.Time `xml:"DtTm,omitempty"` // ISODateTime
}

// TransactionPrice4 - Transaction price choice
type TransactionPrice4 struct {
	Deal        *string             `xml:"Deal,omitempty"` // Max35Text
	Proprietary []ProprietaryPrice2 `xml:"Prtry,omitempty"`
}

// ProprietaryPrice2 - Proprietary price information
type ProprietaryPrice2 struct {
	Type  string `xml:"Tp"`   // Max35Text
	Price string `xml:"Pric"` // Max35Text
}

// TransactionQuantities3 - Transaction quantities choice
type TransactionQuantities3 struct {
	Proprietary []ProprietaryQuantity1 `xml:"Prtry"`
}

// ProprietaryQuantity1 - Proprietary quantity information
type ProprietaryQuantity1 struct {
	Type     string `xml:"Tp"`  // Max35Text
	Quantity string `xml:"Qty"` // Max35Text
}

// SecurityIdentification19 - Security identification
type SecurityIdentification19 struct {
	ISIN                *string                `xml:"ISIN,omitempty"` // ISINOct2015Identifier
	OtherIdentification []OtherIdentification1 `xml:"OthrId,omitempty"`
	Description         *string                `xml:"Desc,omitempty"` // Max140Text
}

// OtherIdentification1 - Other identification
type OtherIdentification1 struct {
	ID     string                `xml:"Id"`            // Max35Text
	Suffix *string               `xml:"Sfx,omitempty"` // Max16Text
	Type   IdentificationSource3 `xml:"Tp"`
}

// IdentificationSource3 - Identification source choice
type IdentificationSource3 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalFinancialInstrumentIdentificationType1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// CorporateActionInfo2 - Corporate action information
type CorporateActionInfo2 struct {
	CodeOrProprietary CorporateActionCodeAndProprietary `xml:"CdOrPrtry"`
	Description       *string                           `xml:"Desc,omitempty"` // Max140Text
}

// CorporateActionCodeAndProprietary - Corporate action code and proprietary
type CorporateActionCodeAndProprietary struct {
	Code        *string `xml:"Cd,omitempty"`    // CorporateActionEventType33Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// SafekeepingPlaceFormat28 - Safekeeping place format choice
type SafekeepingPlaceFormat28 struct {
	Identification        *SafekeepingPlaceTypeAndText6             `xml:"Id,omitempty"`
	Country               *string                                   `xml:"Ctry,omitempty"` // CountryCode
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId,omitempty"`
	Proprietary           *GenericIdentification30                  `xml:"Prtry,omitempty"`
}

// SafekeepingPlaceTypeAndText6 - Safekeeping place type and text
type SafekeepingPlaceTypeAndText6 struct {
	Type           string  `xml:"Tp"`           // SafekeepingPlace1Code
	Identification *string `xml:"Id,omitempty"` // Max35Text
}

// SafekeepingPlaceTypeAndAnyBICIdentifier1 - Safekeeping place type and BIC
type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	Type           string `xml:"Tp"` // SafekeepingPlace1Code
	Identification string `xml:"Id"` // AnyBICDec2014Identifier
}

// ReversalReason4 - Reversal reason choice
type ReversalReason4 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalReversalReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// Additional missing types to resolve compilation errors

// PendingReason16 - Pending reason choice
type PendingReason16 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalPendingProcessingReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// CancellationReason33 - Cancellation reason choice
type CancellationReason33 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalCancellationReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// ModificationReason2 - Modification reason choice
type ModificationReason2 struct {
	Code        *string `xml:"Cd,omitempty"`    // ExternalModificationReason1Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// UnableToApplyMissing1 - Unable to apply missing information
type UnableToApplyMissing1 struct {
	Code        *string `xml:"Cd,omitempty"`    // UnableToApplyMissingInformation3Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// UnableToApplyIncorrect1 - Unable to apply incorrect information
type UnableToApplyIncorrect1 struct {
	Code        *string `xml:"Cd,omitempty"`    // UnableToApplyIncorrectInfo4Code
	Proprietary *string `xml:"Prtry,omitempty"` // Max35Text
}

// GenericIdentification30 - Generic identification with exact 4 alphanumeric text
type GenericIdentification30 struct {
	ID         string  `xml:"Id"`                // Exact4AlphaNumericText
	Issuer     string  `xml:"Issr"`              // Max35Text
	SchemeName *string `xml:"SchmeNm,omitempty"` // Max35Text
}

// Core validation functions for key struct types that don't have them

// Validate performs validation for InterestType1
func (i *InterestType1) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if i.Code != nil {
		choiceCount++
		if err := validateStringLength(*i.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if i.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*i.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for Rate4
func (r *Rate4) Validate() error {
	var errs ValidationErrors

	if r.Rate != nil {
		if *r.Rate < 0 {
			errs = append(errs, ValidationError{Field: "Rate", Message: "percentage rate cannot be negative"})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for TaxCharges2
func (t *TaxCharges2) Validate() error {
	var errs ValidationErrors

	if t.ID != nil {
		if err := validateStringLength(*t.ID, 1, 35, "ID"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if t.Rate != nil {
		if *t.Rate < 0 {
			errs = append(errs, ValidationError{Field: "Rate", Message: "percentage rate cannot be negative"})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for BalanceType10
func (b *BalanceType10) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if b.Code != nil {
		choiceCount++
		if err := validateStringLength(*b.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if b.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*b.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for BalanceSubType1
func (b *BalanceSubType1) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if b.Code != nil {
		choiceCount++
		if err := validateStringLength(*b.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if b.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*b.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for ServiceLevel8
func (s *ServiceLevel8) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if s.Code != nil {
		choiceCount++
		if err := validateStringLength(*s.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if s.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*s.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for LocalInstrument2
func (l *LocalInstrument2) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if l.Code != nil {
		choiceCount++
		if err := validateStringLength(*l.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if l.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*l.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for CategoryPurpose1
func (c *CategoryPurpose1) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if c.Code != nil {
		choiceCount++
		if err := validateStringLength(*c.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if c.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*c.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for MandateSetupReason1
func (m *MandateSetupReason1) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if m.Code != nil {
		choiceCount++
		if err := validateStringLength(*m.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if m.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*m.Proprietary, 1, 70, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for CashAccountType2
func (c *CashAccountType2Choice) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if c.Code != nil {
		choiceCount++
		if err := validateStringLength(*c.Code, 1, 35, "Code"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}
	if c.Proprietary != nil {
		choiceCount++
		if err := validateStringLength(*c.Proprietary, 1, 35, "Proprietary"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if choiceCount != 1 {
		errs = append(errs, ValidationError{Field: "Choice", Message: "exactly one choice must be present"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for GenericIdentification30
func (g *GenericIdentification30) Validate() error {
	var errs ValidationErrors

	if err := validateRequired(g.ID, "ID"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if len(g.ID) != 4 {
			errs = append(errs, ValidationError{Field: "ID", Message: "must be exactly 4 characters"})
		}
	}

	if err := validateRequired(g.Issuer, "Issuer"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := validateStringLength(g.Issuer, 1, 35, "Issuer"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if g.SchemeName != nil {
		if err := validateStringLength(*g.SchemeName, 1, 35, "SchemeName"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Additional focused validation functions for core structs

// Validate performs validation for PaymentIdentification7
func (p *PaymentIdentification7) Validate() error {
	var errs ValidationErrors

	if err := validateRequired(p.EndToEndID, "EndToEndID"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := validateStringLength(p.EndToEndID, 1, 35, "EndToEndID"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.TransactionID != nil {
		if err := validateStringLength(*p.TransactionID, 1, 35, "TransactionID"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for PartyIdentification135
func (p *PartyIdentification135) Validate() error {
	var errs ValidationErrors

	if p.Name != nil {
		if err := validateStringLength(*p.Name, 1, 140, "Name"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.PostalAddress != nil {
		if err := p.PostalAddress.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "PostalAddress", Message: err.Error()})
		}
	}

	if p.ContactDetails != nil {
		if err := p.ContactDetails.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "ContactDetails", Message: err.Error()})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for PostalAddress24
func (p *PostalAddress24) Validate() error {
	var errs ValidationErrors

	if p.Department != nil {
		if err := validateStringLength(*p.Department, 1, 70, "Department"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.StreetName != nil {
		if err := validateStringLength(*p.StreetName, 1, 70, "StreetName"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.BuildingNumber != nil {
		if err := validateStringLength(*p.BuildingNumber, 1, 16, "BuildingNumber"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.PostCode != nil {
		if err := validateStringLength(*p.PostCode, 1, 16, "PostCode"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.TownName != nil {
		if err := validateStringLength(*p.TownName, 1, 35, "TownName"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.Country != nil {
		if err := validateCountryCode(*p.Country, "Country"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for Contact4
func (c *Contact4) Validate() error {
	var errs ValidationErrors

	if c.Name != nil {
		if err := validateStringLength(*c.Name, 1, 140, "Name"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if c.EmailAddress != nil {
		if err := validateStringLength(*c.EmailAddress, 1, 2048, "EmailAddress"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if c.PhoneNumber != nil {
		if err := validateStringLength(*c.PhoneNumber, 1, 35, "PhoneNumber"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for BranchAndFinancialInstitutionIdentification6
func (b *BranchAndFinancialInstitutionIdentification6) Validate() error {
	var errs ValidationErrors

	if err := validateRequired(b.FinancialInstitutionID, "FinancialInstitutionID"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := b.FinancialInstitutionID.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "FinancialInstitutionID", Message: err.Error()})
		}
	}

	if b.BranchID != nil {
		if err := b.BranchID.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "BranchID", Message: err.Error()})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for FinancialInstitutionIdentification18
func (f *FinancialInstitutionIdentification18) Validate() error {
	var errs ValidationErrors

	if f.BankIdentifierCode != nil {
		if err := validateBIC(*f.BankIdentifierCode, "BankIdentifierCode"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if f.LegalEntityIdentifier != nil {
		if err := validateLEI(*f.LegalEntityIdentifier, "LegalEntityIdentifier"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if f.Name != nil {
		if err := validateStringLength(*f.Name, 1, 140, "Name"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for BranchData3
func (b *BranchData3) Validate() error {
	var errs ValidationErrors

	if b.ID != nil {
		if err := validateStringLength(*b.ID, 1, 35, "ID"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if b.LegalEntityIdentifier != nil {
		if err := validateLEI(*b.LegalEntityIdentifier, "LegalEntityIdentifier"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if b.Name != nil {
		if err := validateStringLength(*b.Name, 1, 140, "Name"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	// Note: PostalAddress validation skipped as it uses different type

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for GenericAccountIdentification1
func (g *GenericAccountIdentification1) Validate() error {
	var errs ValidationErrors

	if err := validateRequired(g.ID, "ID"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		if err := validateStringLength(g.ID, 1, 34, "ID"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	// Note: SchemeName validation skipped as it uses custom type

	if g.Issuer != nil {
		if err := validateStringLength(*g.Issuer, 1, 35, "Issuer"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for CashAccount38
func (c *CashAccount38) Validate() error {
	var errs ValidationErrors

	// ID is required - delegating to AccountIdentification4 validation
	if err := c.ID.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "ID", Message: err.Error()})
	}

	if c.Currency != nil {
		if err := validateStringLength(*c.Currency, 3, 3, "Currency"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if c.Name != nil {
		if err := validateStringLength(*c.Name, 1, 70, "Name"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for AccountIdentification4
func (a *AccountIdentification4) Validate() error {
	var errs ValidationErrors

	// Must have exactly one choice - either IBAN or Other
	hasIBAN := a.IBAN != nil
	hasOther := a.Other != nil

	if !hasIBAN && !hasOther {
		errs = append(errs, ValidationError{Field: "Choice", Message: "must have either IBAN or Other"})
	} else if hasIBAN && hasOther {
		errs = append(errs, ValidationError{Field: "Choice", Message: "cannot have both IBAN and Other"})
	}

	if hasIBAN {
		// IBAN should be validated with specific IBAN format
		if err := validateStringLength(*a.IBAN, 15, 34, "IBAN"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
		// Additional IBAN format validation could be added here
	}

	if hasOther {
		if err := a.Other.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "Other", Message: err.Error()})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for PaymentTypeInfo28
func (p *PaymentTypeInfo28) Validate() error {
	var errs ValidationErrors

	if p.InstructionPriority != nil {
		// InstructionPriority should be a valid priority code
		if err := validateStringLength(*p.InstructionPriority, 1, 4, "InstructionPriority"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	if p.SequenceType != nil {
		// SequenceType validation - common values: FRST, RCUR, FNAL, OOFF
		if err := validateStringLength(*p.SequenceType, 1, 4, "SequenceType"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	// Note: ServiceLevel, LocalInstrument, CategoryPurpose validation skipped due to custom types

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for CreditTransferTransaction39
func (c *CreditTransferTransaction39) Validate() error {
	var errs ValidationErrors

	// PaymentID is required
	if err := c.PaymentID.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "PaymentID", Message: err.Error()})
	}

	// InterbankSettlementAmount is required
	if err := c.InterbankSettlementAmount.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "InterbankSettlementAmount", Message: err.Error()})
	}

	// ChargeBearer is required and should be valid charge bearer code
	if err := validateRequired(c.ChargeBearer, "ChargeBearer"); err != nil {
		errs = append(errs, err.(ValidationError))
	} else {
		// Common values: DEBT, CRED, SHAR, SLEV
		if err := validateStringLength(c.ChargeBearer, 1, 4, "ChargeBearer"); err != nil {
			errs = append(errs, err.(ValidationError))
		}
	}

	// Debtor is required
	if err := c.Debtor.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "Debtor", Message: err.Error()})
	}

	// DebtorAgent is required
	if err := c.DebtorAgent.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "DebtorAgent", Message: err.Error()})
	}

	// Creditor is required
	if err := c.Creditor.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "Creditor", Message: err.Error()})
	}

	// CreditorAgent is required
	if err := c.CreditorAgent.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "CreditorAgent", Message: err.Error()})
	}

	// Optional fields
	if c.PaymentTypeInfo != nil {
		if err := c.PaymentTypeInfo.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "PaymentTypeInfo", Message: err.Error()})
		}
	}

	if c.DebtorAccount != nil {
		if err := c.DebtorAccount.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "DebtorAccount", Message: err.Error()})
		}
	}

	if c.CreditorAccount != nil {
		if err := c.CreditorAccount.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "CreditorAccount", Message: err.Error()})
		}
	}

	if c.UltimateDebtor != nil {
		if err := c.UltimateDebtor.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "UltimateDebtor", Message: err.Error()})
		}
	}

	if c.UltimateCreditor != nil {
		if err := c.UltimateCreditor.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "UltimateCreditor", Message: err.Error()})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate performs validation for FIToFICustomerCreditTransferV08
func (f *FIToFICustomerCreditTransferV08) Validate() error {
	var errs ValidationErrors

	// GroupHeader is required
	if err := f.GroupHeader.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "GroupHeader", Message: err.Error()})
	}

	// CreditTransferTransactionInfo is required and must have at least one item
	if len(f.CreditTransferTransactionInfo) == 0 {
		errs = append(errs, ValidationError{Field: "CreditTransferTransactionInfo", Message: "at least one credit transfer transaction is required"})
	} else {
		for i, tx := range f.CreditTransferTransactionInfo {
			if err := tx.Validate(); err != nil {
				errs = append(errs, ValidationError{Field: fmt.Sprintf("CreditTransferTransactionInfo[%d]", i), Message: err.Error()})
			}
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Business Application Header V02 (head.001.001.02)
// BusinessApplicationHeaderV02 represents the Business Application Header used to wrap ISO 20022 messages.
// It provides message identification, routing information, and processing metadata for business application messages.
type BusinessApplicationHeaderV02 struct {
	CharacterSet           *string                       `xml:"CharSet,omitempty"`    // Character set used in the message (UnicodeChartsCode)
	From                   Party44                       `xml:"Fr"`                   // Message originator
	To                     Party44                       `xml:"To"`                   // Message recipient
	BusinessMessageID      string                        `xml:"BizMsgIdr"`            // Unique business message identifier (Max35Text)
	MessageDefinitionID    string                        `xml:"MsgDefIdr"`            // Message definition identifier (Max35Text)
	BusinessService        *string                       `xml:"BizSvc,omitempty"`     // Business service identifier (Max35Text)
	MarketPractice         *ImplementationSpecification1 `xml:"MktPrctc,omitempty"`   // Market practice specification
	CreationDate           *time.Time                    `xml:"CreDt,omitempty"`      // Creation date and time (ISODateTime)
	BusinessProcessingDate *time.Time                    `xml:"BizPrcgDt,omitempty"`  // Business processing date (ISODateTime)
	CopyDuplicate          *CopyDuplicate1Code           `xml:"CpyDplct,omitempty"`   // Copy/duplicate indicator
	PossibleDuplicate      *bool                         `xml:"PssblDplct,omitempty"` // Possible duplicate flag (YesNoIndicator)
	Priority               *BusinessMessagePriorityCode  `xml:"Prty,omitempty"`       // Message priority
	Signature              *SignatureEnvelope            `xml:"Sgntr,omitempty"`      // Digital signature
	Related                []BusinessApplicationHeader5  `xml:"Rltd,omitempty"`       // Related headers
}

// CopyDuplicate1Code represents the copy/duplicate indicator
type CopyDuplicate1Code string

const (
	CopyDuplicateCodeCoDu CopyDuplicate1Code = "CODU"
	CopyDuplicateCodeCopy CopyDuplicate1Code = "COPY"
	CopyDuplicateCodeDupl CopyDuplicate1Code = "DUPL"
)

// BusinessMessagePriorityCode represents message priority levels
type BusinessMessagePriorityCode string

const (
	BusinessMessagePriorityHigh   BusinessMessagePriorityCode = "HIGH"
	BusinessMessagePriorityNormal BusinessMessagePriorityCode = "NORM"
	BusinessMessagePriorityUrgent BusinessMessagePriorityCode = "URGT"
)

// ImplementationSpecification1 represents market practice specification
type ImplementationSpecification1 struct {
	Registry *string `xml:"Regy"` // Max350Text - Registry information
	ID       *string `xml:"Id"`   // Max2048Text - Implementation identifier
}

// BusinessApplicationHeader5 represents related header information (used in V02)
type BusinessApplicationHeader5 struct {
	CharacterSet        *string                      `xml:"CharSet,omitempty"`    // Character set used in the message
	From                Party44                      `xml:"Fr"`                   // Message originator
	To                  Party44                      `xml:"To"`                   // Message recipient
	BusinessMessageID   string                       `xml:"BizMsgIdr"`            // Unique business message identifier
	MessageDefinitionID string                       `xml:"MsgDefIdr"`            // Message definition identifier
	BusinessService     *string                      `xml:"BizSvc,omitempty"`     // Business service identifier
	CreationDate        *time.Time                   `xml:"CreDt,omitempty"`      // Creation date and time
	CopyDuplicate       *CopyDuplicate1Code          `xml:"CpyDplct,omitempty"`   // Copy/duplicate indicator
	PossibleDuplicate   *bool                        `xml:"PssblDplct,omitempty"` // Possible duplicate flag
	Priority            *BusinessMessagePriorityCode `xml:"Prty,omitempty"`       // Message priority
	Signature           *SignatureEnvelope           `xml:"Sgntr,omitempty"`      // Digital signature
}

// SignatureEnvelope represents a digital signature wrapper
type SignatureEnvelope struct {
	// Implementation depends on specific signature requirements
	// This is a placeholder for XML signature structure
	Value string `xml:",chardata"`
}

// Validate validates the BusinessApplicationHeaderV02 structure
func (b *BusinessApplicationHeaderV02) Validate() error {
	var errs ValidationErrors

	// BusinessMessageID is required and has format restrictions
	if b.BusinessMessageID == "" {
		errs = append(errs, ValidationError{Field: "BusinessMessageID", Message: "business message identifier is required"})
	} else if len(b.BusinessMessageID) > 35 {
		errs = append(errs, ValidationError{Field: "BusinessMessageID", Message: "business message identifier must not exceed 35 characters"})
	}

	// MessageDefinitionID is required
	if b.MessageDefinitionID == "" {
		errs = append(errs, ValidationError{Field: "MessageDefinitionID", Message: "message definition identifier is required"})
	} else {
		// Validate message definition identifier format (e.g., pacs.008.001.08)
		msgDefPattern := regexp.MustCompile(`^[a-z]{4}\.\d{3}\.\d{3}\.\d{2}$`)
		if !msgDefPattern.MatchString(b.MessageDefinitionID) {
			errs = append(errs, ValidationError{Field: "MessageDefinitionID", Message: "message definition identifier must follow format like 'pacs.008.001.08'"})
		}
	}

	// From is required
	if err := b.From.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "From", Message: err.Error()})
	}

	// To is required
	if err := b.To.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "To", Message: err.Error()})
	}

	// CreationDate is required (string pointer value check)
	if b.CreationDate == nil {
		errs = append(errs, ValidationError{Field: "CreationDate", Message: "creation date is required"})
	} else {
		// time.Time is already validated by Go's type system
	}

	// Validate optional fields if present
	if b.CharacterSet != nil && *b.CharacterSet != "" {
		if len(*b.CharacterSet) > 35 {
			errs = append(errs, ValidationError{Field: "CharacterSet", Message: "character set must not exceed 35 characters"})
		}
	}

	if b.BusinessService != nil && *b.BusinessService != "" {
		if len(*b.BusinessService) > 35 {
			errs = append(errs, ValidationError{Field: "BusinessService", Message: "business service must not exceed 35 characters"})
		}
	}

	// V02 specific validations
	if b.MarketPractice != nil {
		if err := b.MarketPractice.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "MarketPractice", Message: err.Error()})
		}
	}

	// Validate Related headers if present
	for i, related := range b.Related {
		if err := related.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: fmt.Sprintf("Related[%d]", i), Message: err.Error()})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate validates the Party44 structure
func (p *Party44) Validate() error {
	var errs ValidationErrors

	// Exactly one choice must be present
	choiceCount := 0
	if p.FinancialInstitutionID != nil {
		choiceCount++
		if err := p.FinancialInstitutionID.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "FinancialInstitutionID", Message: err.Error()})
		}
	}
	if p.OrganisationIdentification != nil {
		choiceCount++
		if err := p.OrganisationIdentification.Validate(); err != nil {
			errs = append(errs, ValidationError{Field: "OrganisationIdentification", Message: err.Error()})
		}
	}

	if choiceCount == 0 {
		errs = append(errs, ValidationError{Field: "Party44", Message: "either FinancialInstitutionID or OrganisationIdentification must be provided"})
	} else if choiceCount > 1 {
		errs = append(errs, ValidationError{Field: "Party44", Message: "only one of FinancialInstitutionID or OrganisationIdentification can be provided"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate validates the ImplementationSpecification1 structure
func (i *ImplementationSpecification1) Validate() error {
	var errs ValidationErrors

	// Registry is required
	if i.Registry == nil || *i.Registry == "" {
		errs = append(errs, ValidationError{Field: "Registry", Message: "registry is required"})
	} else if len(*i.Registry) > 350 {
		errs = append(errs, ValidationError{Field: "Registry", Message: "registry must not exceed 350 characters"})
	}

	// ID is required
	if i.ID == nil || *i.ID == "" {
		errs = append(errs, ValidationError{Field: "ID", Message: "ID is required"})
	} else if len(*i.ID) > 2048 {
		errs = append(errs, ValidationError{Field: "ID", Message: "ID must not exceed 2048 characters"})
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// Validate validates the BusinessApplicationHeader5 structure
func (b *BusinessApplicationHeader5) Validate() error {
	var errs ValidationErrors

	// BusinessMessageID is required
	if b.BusinessMessageID == "" {
		errs = append(errs, ValidationError{Field: "BusinessMessageID", Message: "business message identifier is required"})
	} else if len(b.BusinessMessageID) > 35 {
		errs = append(errs, ValidationError{Field: "BusinessMessageID", Message: "business message identifier must not exceed 35 characters"})
	}

	// MessageDefinitionID is required
	if b.MessageDefinitionID == "" {
		errs = append(errs, ValidationError{Field: "MessageDefinitionID", Message: "message definition identifier is required"})
	} else if len(b.MessageDefinitionID) > 35 {
		errs = append(errs, ValidationError{Field: "MessageDefinitionID", Message: "message definition identifier must not exceed 35 characters"})
	}

	// From is required
	if err := b.From.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "From", Message: err.Error()})
	}

	// To is required
	if err := b.To.Validate(); err != nil {
		errs = append(errs, ValidationError{Field: "To", Message: err.Error()})
	}

	// CreationDate is required (string pointer value check)
	if b.CreationDate == nil {
		errs = append(errs, ValidationError{Field: "CreationDate", Message: "creation date is required"})
	} else {
		// time.Time is already validated by Go's type system
	}

	// Validate optional fields if present
	if b.CharacterSet != nil && *b.CharacterSet != "" {
		if len(*b.CharacterSet) > 35 {
			errs = append(errs, ValidationError{Field: "CharacterSet", Message: "character set must not exceed 35 characters"})
		}
	}

	if b.BusinessService != nil && *b.BusinessService != "" {
		if len(*b.BusinessService) > 35 {
			errs = append(errs, ValidationError{Field: "BusinessService", Message: "business service must not exceed 35 characters"})
		}
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// BusinessApplicationHeaderDocument represents a complete BAH message envelope
type BusinessApplicationHeaderDocument struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Document"`
	AppHdr  BusinessApplicationHeaderV02 `xml:"AppHdr"`
}

// Validate validates the BusinessApplicationHeaderDocument
func (b *BusinessApplicationHeaderDocument) Validate() error {
	return b.AppHdr.Validate()
}
