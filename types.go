package shiji

import "strings"

type CommaSeparatedQueryParam []string

func (t CommaSeparatedQueryParam) MarshalSchema() string {
	return strings.Join(t, ",")
}

type PaginatedResponse[T any] struct {
	Results []T `json:"results"`
	Paging  struct {
		PageNumber int `json:"pageNumber"`
		PageSize   int `json:"pageSize"`
		TotalCount int `json:"totalCount"`
	} `json:"paging"`
}

type ConfigurationSupportedType string

type AlternateName struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
}

type LocalizedString struct {
	LanguageCode string `json:"languageCode"`
	Content      string `json:"content"`
}

type AddressDetail struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	PostalCode      string `json:"postalCode"`
	StateCode       string `json:"stateCode"`
	CountryCode     string `json:"countryCode"`
	AddressTypeCode string `json:"addressTypeCode"`
}

type CommunicationDetail struct {
	Primary               bool   `json:"primary"`
	CommunicationTypeCode string `json:"communicationTypeCode"`
	Value                 string `json:"value"`
}

type CommunicationChannel struct {
	IsPrimary bool   `json:"isPrimary"`
	Details   string `json:"details"`
	ModeCode  string `json:"modeCode"`
}

type ContactDetail struct {
	Primary  bool   `json:"primary"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	JobTitle string `json:"jobTitle"`
}

type PropertyListItem struct {
	ID              string `json:"id"`
	Version         int64  `json:"version"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	LegalName       string `json:"legalName"`
	ParentID        string `json:"parentId"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	BusinessAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    string `json:"stateCode"`
		StateName    string `json::"stateName"`
		CountryCode  string `json:"countryCode"`
		CountryName  string `json:"countryName"`
	} `json:"businessAddress"`
	TaxID           string                 `json:"taxId"`
	StarRating      float64                `json:"starRating"`
	TypeID          string                 `json:"typeId"`
	StatusCode      string                 `json:"statusCode"`
	CustomFields    map[string]interface{} `json:"customFields"`
	DisplaySequence int32                  `json:"displaySequence"`
	RegionCode      string                 `json:"regionCode"`
	BrandID         string                 `json:"brandId"`
}

type Property struct {
	ID                     string                `json:"id"`
	Version                int64                 `json:"version"`
	Code                   string                `json:"code"`
	Name                   string                `json:"name"`
	ParentID               string                `json:"parentId"`
	DefaultLanguageCode    string                `json:"defaultLanguageCode"`
	TaxID                  string                `json:"taxId"`
	LegalName              string                `json:"legalName"`
	AlternateNames         []AlternateName       `json:"alternateNames"`
	SupportedLanguageCodes []string              `json:"supportedLanguageCodes"`
	AddressList            []AddressDetail       `json:"addressList"`
	CommunicationDetails   []CommunicationDetail `json:"communicationDetails"`
	Contacts               []ContactDetail       `json:"contacts"`
	LegalOwner             string                `json:"legalOwner"`
	Location               struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	CheckInTime              string                 `json:"checkInTime"`
	CheckOutTime             string                 `json:"checkOutTime"`
	FutureOperationalWindow  int32                  `json:"futureOperationalWindow"`
	AlternativePropertyCodes []string               `json:"alternativePropertyCodes"`
	AlternativePropertyIDs   []string               `json:"alternativePropertyIds"`
	CurrencyCode             string                 `json:"currencyCode"`
	CustomCurrencyFormat     string                 `json:"customCurrencyFormat"`
	ShortDataFormat          string                 `json:"shortDataFormat"`
	LongDateFormat           string                 `json:"longDateFormat"`
	TimeFormatCode           string                 `json:"timeFormatCode"`
	TimeZoneCode             string                 `json:"timeZoneCode"`
	StatusCode               string                 `json:"statusCode"`
	FiscalReportingPeriodID  string                 `json:"fiscalReportingPeriodId"`
	StarRating               float64                `json:"starRating"`
	TypeID                   string                 `json:"typeId"`
	BrandID                  string                 `json:"brandId"`
	OperatingModelCode       string                 `json:"operatingModelCode"`
	AttributeIDs             []string               `json:"attributeIds"`
	DistributionChannelIDs   []string               `json:"distributionChannelIds"`
	CustomFields             map[string]interface{} `json:"customFields"`
	ProfileStore             struct {
		PrimaryID string   `json:"primaryId"`
		OtherIDs  []string `json:"otherIds"`
	} `json:"profileStore"`

	EffectiveZoneIDs    []string `json:"effectiveZoneIds"`
	DisplaySequence     int32    `json:"displaySequence"`
	RegionCode          string   `json:"regionCode"`
	InventoryAllocation struct {
		StartTime string `json:"startTime"`
		EndTIme   string `json:"endTime"`
	} `json:"inventoryAllocation"`

	SupportedQuotationCurrencies        []string `json:"supportedQuotationCurrencies"`
	RatePlanQuotationCurrencies         []string `json:"ratePlanQuotationCurrencies"`
	PurchaseElementQuotationCurrencies  []string `json:"purchaseElementQuotationCurrencies"`
	MeetingAndEventsQuotationCurrencies []string `json:"meetingAndEventsQuotationCurrencies"`
	FirstDayOfWeek                      string   `json:"firstDayOfWeek"`
}

type TransactionCode struct {
	ID                     string            `json:"id"`
	Version                int64             `json:"version"`
	ZoneID                 string            `json:"zoneId"`
	Code                   string            `json:"code"`
	Description            []LocalizedString `json:"description"`
	IsActive               bool              `json:"isActive"`
	ForceToUse             bool              `json:"forceToUse"`
	DisplaySequence        int32             `json:"displaySequence"`
	SubCode                string            `json:"subCode"`
	CurrencyCode           string            `json:"currencyCode"`
	AcceptAsCashierPayment bool              `json:"acceptAsCashierPayment"`
	AllowToPostInAR        bool              `json:"allowToPostInAR"`
	OnlinePaymentMethod    string            `json:"onlinePaymentMethod"`
	AcceptAsDepositPayment bool              `json:"acceptAsDepositPayment"`
	AllowManualPosting     bool              `json:"allowManualPosting"`
	IsAdvancedDeposit      bool              `json:"isAdvancedDeposit"`
	IsCheckNumberMandatory bool              `json:"isCheckNumberMandatory"`
	PostingValueRange      struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"postingValueRange"`

	DefaultPostingValue      float64                `json:"defaultPostingValue"`
	AdjustmentID             string                 `json:"adjustmentId"`
	CustomFieldContainerJSON map[string]interface{} `json:"customFieldContainerJson"`
	IsLocked                 bool                   `json:"isLocked"`
	CreditCardSurcharge      struct {
		Percentage        float64 `json:"percentage"`
		TransactionCodeID string  `json:"transactionCodeId"`
	} `json:"creditCardSurcharge"`

	TransactionType         string `json:"transactionType"`
	IsNonRevenue            bool   `json:"isNonRevenue"`
	IsPostingAllowed        bool   `json:"isPostingAllowed"`
	TransactionSubGroupID   string `json:"transactionSubGroupId"`
	ParentTransactionCodeID string `json:"parentTransactionCodeId"`
	PaymentTypeCode         string `json:"paymentTypeCode"`
	CreditCardTypeCode      string `json:"creditCardTypeCode"`
	HasSubCodes             bool   `json:"hasSubCodes"`
	OfflinePaymentMethod    string `json:"offlinePaymentMethod"`
	IsGuestLedgerPrepayment bool   `json:"isGuestLedgerPrepayment"`
}

type TaxRule struct {
	ID                    string            `json:"id"`
	Version               int64             `json:"version"`
	ZoneID                string            `json:"zoneId"`
	Code                  string            `json:"code"`
	Description           []LocalizedString `json:"description"`
	IsActive              bool              `json:"isActive"`
	ForceToUse            bool              `json:"forceToUse"`
	DisplaySequence       int32             `json:"displaySequence"`
	IsCharge              bool              `json:"isCharge"`
	Category              string            `json:"category"`
	FlatAmount            float64           `json:"flatAmount"`
	PercentageAmount      float64           `json:"percentageAmount"`
	IsPercentage          bool              `json:"isPercentage"`
	AmountDefinition      string            `json:"amountDefinition"`
	AmountRangeDefinition struct {
		RateType         string `json:"rateType"`
		PriceType        string `json:"priceType"`
		BaseAmountType   string `json:"baseAmountType"`
		BaseAmountRanges []struct {
			From                    float64 `json:"from"`
			To                      float64 `json:"to"`
			ValueRangeType          string  `json:"valueRangeType"`
			Value                   float64 `json:"value"`
			IsIncrementationDefined bool    `json:"isIncrementationDefined"`
			IncrementationValue     float64 `json:"incrementationValue"`
			ThresholdValue          float64 `json:"thresholdValue"`
		} `json:"baseAmountRanges"`

		IsPercentage bool `json:"isPercentage"`
	} `json:"amountRangeDefinition"`

	BaseRuleID                      string `json:"baseRuleId"`
	IsBaseRule                      bool   `json:"isBaseRule"`
	IsTaxIncluded                   bool   `json:"isTaxIncluded"`
	PrintSeparately                 bool   `json:"printSeparately"`
	PostSeparately                  bool   `json:"postSeparately"`
	ChargeOnManualPosting           bool   `json:"chargeOnManualPosting"`
	ChargeMethodPerPersonDefinition struct {
		IsAdult    bool `json:"isAdult"`
		IsChildren bool `json:"isChildren"`
	} `json:"chargeMethodPerPersonDefinition"`

	ValidFromDate                DateTime `json:"validFromDate"`
	ValidToDate                  DateTime `json:"validToDate"`
	TransactionCodeID            string   `json:"transactionCodeId"`
	ApplicableTransactionCodeIDs []string `json:"applicableTransactionCodeIds"`
	AgeBucketIDs                 []string `json:"ageBucketIds"`
	LengthOfStay                 int32    `json:"lengthOfStay"`
	IsCommission                 bool     `json:"isCommission"`
	PurposeOfStayIDs             []string `json:"purposeOfStayIds"`
	ExcludeDayUse                bool     `json:"excludeDayUse"`
	ChargeMethod                 string   `json:"chargeMethod"`
}

type SubsidiaryListItem struct {
	ID              string `json:"id"`
	Version         int64  `json:"version"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	ParentID        string `json:"parentId"`
	ParentCode      string `json:"parentCode"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	BusinessAddress struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    string `json:"stateCode"`
		StateName    string `json:"stateName"`
		CountryCode  string `json:"countryCode"`
		CountryName  string `json:"countryName"`
	} `json:"businessAddress"`

	TaxID           string `json:"taxId"`
	StatusCode      string `json:"statusCode"`
	DisplaySequence int32  `json:"displaySequence"`
}

type AccountGroup struct {
	ID                             string            `json:"id"`
	Version                        int64             `json:"version"`
	ZoneID                         string            `json:"zoneId"`
	Code                           string            `json:"code"`
	Description                    []LocalizedString `json:"description"`
	IsActive                       bool              `json:"isActive"`
	DisplaySequence                int32             `json:"displaySequence"`
	TypeCode                       string            `json:"typeCode"`
	NumberCycleID                  string            `json:"numberCycleId"`
	ExcludeFromZeroCheckOutDefault bool              `json:"excludeFromZeroCheckOutDefault"`
}

type UserDetails struct {
	ID                        string   `json:"id"`
	AccountLoginProcedureType string   `json:"accountLoginProcedureType"`
	ExternalID                string   `json:"externalId"`
	Version                   int64    `json:"version"`
	UserName                  string   `json:"userName"`
	Email                     string   `json:"email"`
	IsSystemUser              bool     `json:"isSystemUser"`
	TenantID                  string   `json:"tenantId"`
	ActiveFrom                DateTime `json:"activeFrom"`
	ActiveTo                  DateTime `json:"activeTo"`
	DeactivationComment       string   `json:"deactivationComment"`
	BlockedTo                 DateTime `json:"blockedTo"`
	LockoutTo                 DateTime `json:"lockoutTo"`
	BlockComment              string   `json:"blockComment"`
	PasswordChangeRequired    bool     `json:"passwordChangeRequired"`
	PersonalData              struct {
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		GenderCode     string `json:"genderCode"`
		BirthDate      Date   `json:"birthDate"`
		EmployeeNumber string `json:"employeeNumber"`
		HireDate       Date   `json:"hireDate"`
	} `json:"personalData"`

	Address struct {
		CountryCode  string `json:"countryCode"`
		StateCode    string `json:"stateCode"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
	} `json:"address"`

	CommunicationChannels         []CommunicationDetail `json:"communicationChannels"`
	PasswordExpirationDisabled    bool                  `json:"passwordExpirationDisabled"`
	AutomaticDeactivationDisabled bool                  `json:"automaticDeactivationDisabled"`
	AuthenticatorRequired         bool                  `json:"authenticatorRequired"`
}

type UnitAssignmentListItem struct {
	UnitID          string `json:"unitId"`
	ParentID        string `json:"parentId"`
	Type            string `json:"type"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	IsPublished     bool   `json:"isPublished"`
	RegionCode      string `json:"regionCode"`
	DisplaySequence int32  `json:"displaySequence"`
}

type ContactProfileResponse struct {
	ContactProfileID                  string   `json:"contactProfileId"`
	ContactProfileRole                string   `json:"contactProfileRole"`
	IsPrimary                         bool     `json:"isPrimary"`
	Roles                             []string `json:"roles"`
	ContactProfileAvailabilityDetails struct {
		ProfileStoreID                   string   `json:"profileStoreId"`
		RegionCode                       string   `json:"regionCode"`
		IsAvailableOutsideOfSourceRegion bool     `json:"isAvailableOutsideOfSourceRegion"`
		AvailableInRegions               []string `json:"availableInRegions"`
	} `json:"contactProfileAvailabilityDetails"`
}

type CompanyProfileResponse struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	TypeCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"typeCode"`

	LinkedProfileCount int32 `json:"linkedProfileCount"`
	LinkedProfileIDs   []struct {
		ProfileRoleCode string `json:"profileRoleCode"`
		ProfileID       string `json:"profileId"`
	} `json:"linkedProfileIds"`

	Status struct {
		IsActive               bool `json:"isActive"`
		DeactivationReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"deactivationReasonCode"`

		IsRestricted          bool `json:"isRestricted"`
		RestrictionReasonCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"restrictionReasonCode"`

		RestrictionComment string `json:"restrictionComment"`
		IsLocked           bool   `json:"isLocked"`
		LockedReason       string `json:"lockedReason"`
		ForReview          bool   `json:"forReview"`
	} `json:"status"`

	Details struct {
		FullName      string `json:"fullName"`
		Abbreviation  string `json:"abbreviation"`
		CorporateID   string `json:"corporateId"`
		TaxID         string `json:"taxId"`
		IndustryCodes []struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"industryCodes"`

		BusinessSegmentCodes []struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"businessSegmentCodes"`

		PreferredLanguageProfileDetails struct {
			LanguageCode struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"languageCode"`

			FullName string `json:"fullName"`
		} `json:"preferredLanguageProfileDetails"`

		PreferredCommunicationLanguageCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"preferredCommunicationLanguageCode"`
	} `json:"details"`

	Addresses []struct {
		AddressTypeCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"addressTypeCode"`

		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		PostalCode   string `json:"postalCode"`
		StateCode    struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"stateCode"`

		CountryCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"countryCode"`

		LanguageCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"languageCode"`

		IsPrimary   bool `json:"isPrimary"`
		IsArPrimary bool `json:"isArPrimary"`
		DistrictID  struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"districtId"`

		ID string `json:"id"`
	} `json:"addresses"`

	CommunicationChannels []struct {
		ID       string `json:"id"`
		TypeCode struct {
			ID          string `json:"id"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"typeCode"`

		ModeCode struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"modeCode"`

		Details   string `json:"details"`
		IsPrimary bool   `json:"isPrimary"`
	} `json:"communicationChannels"`

	Metadata struct {
		CreatorID            string   `json:"creatorId"`
		CreationTime         DateTime `json:"creationTime"`
		LastUpdaterID        string   `json:"lastUpdaterId"`
		LastUpdateTime       DateTime `json:"lastUpdateTime"`
		CreationPropertyID   string   `json:"creationPropertyId"`
		LastUpdatePropertyID string   `json:"lastUpdatePropertyId"`
	} `json:"metadata"`

	NotDuplicateProfiles struct {
		Collection []struct {
			NotDuplicateProfileID string `json:"notDuplicateProfileId"`
			ProfileID             string `json:"profileId"`
			ProfileRoleCode       string `json:"profileRoleCode"`
		}

		Count int32 `json:"count"`
	} `json:"notDuplicateProfiles"`

	FreezeStatus struct {
		IsFrozen   bool   `json:"isFrozen"`
		Reason     string `json:"reason"`
		PropertyID string `json:"propertyId"`
		Source     string `json:"source"`
	} `json:"freezeStatus"`

	CustomFieldContainerJSON string `json:"customFieldContainerJson"`

	ExternalSystemIdentifiers []struct {
		ExternalSystemID string `json:"externalSystemId"`
		Identifier       string `json:"identifier"`
	} `json:"externalSystemIdentifiers"`

	NotesIdentifiers []string                 `json:"notesIdentifiers"`
	ContactProfiles  []ContactProfileResponse `json:"contactProfiles"`
	ContactFor       []ContactProfileResponse `json:"contactFor"`
	ProfileOwners    struct {
		Primary struct {
			ID string `json:"id"`
		} `json:"primary"`

		Others []struct {
			ID string `json:"id"`
		} `json:"others"`
	} `json:"profileOwners"`

	RelationshipDetails struct {
		IsRoot bool `json:"isRoot"`
	} `json:"relationshipDetails"`

	ExternalAccountReceivable struct {
		ExternalAccountReceivableNumber string `json:"externalAccountReceivableNumber"`
	} `json:"externalAccountReceivable"`

	HasImage bool `json:"hasImage"`
}

type Transaction struct {
	ID      string `json:"id"`
	Account struct {
		ID     string `json:"id"`
		Number string `json:"number"`
	} `json:"account"`

	OriginAccount struct {
		ID     string `json:"id"`
		Number string `json:"number"`
	} `json:"originAccount"`

	Version             int64    `json:"version"`
	TransactionDate     Date     `json:"transactionDate"`
	TransactionDateTime DateTime `json:"transactionDateTime"`
	BusinessDate        Date     `json:"businessDate"`
	RevenueDate         Date     `json:"revenueDate"`
	UpdatedAt           DateTime `json:"updatedAt"`
	User                struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`

	Cashier struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"cashier"`

	Comment      string `json:"comment"`
	PostingGroup struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"postingGroup"`

	PaymentMethod string `json:"paymentMethod"`
	FolioWindow   struct {
		ID        string `json:"id"`
		Number    int    `json:"number"`
		ProfileID string `json:"profileId"`
	} `json:"folioWindow"`

	InvoiceNumber   string `json:"invoiceNumber"`
	TransactionType string `json:"transactionType"`
	LedgerType      string `json:"ledgerType"`
	TransactionCode struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionCode"`

	TransactionGroup struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionGroup"`

	TransactionSubGroup struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"transactionSubGroup"`

	NetAmount      float64 `json:"netAmount"`
	GrossAmount    float64 `json:"grossAmount"`
	TaxAmount      float64 `json:"taxAmount"`
	PaymentAmount  float64 `json:"paymentAmount"`
	NetUnitPrice   float64 `json:"netUnitPrice"`
	GrossUnitPrice float64 `json:"grossUnitPrice"`
	Quantity       float64 `json:"quantity"`
	Currency       string  `json:"currency"`
	Metadata       struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"metaData"`
}

type Transactions []Transaction
