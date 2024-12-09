package shiji

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
