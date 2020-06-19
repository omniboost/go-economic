package economic

type Journal struct {
	Name      string `json:"name"`
	Vouchers  string `json:"vouchers"`
	Entries   string `json:"entries"`
	Templates struct {
		FinanceVoucher        string `json:"financeVoucher"`
		ManualCustomerInvoice string `json:"manualCustomerInvoice"`
		Self                  string `json:"self"`
	} `json:"templates"`
	JournalNumber int    `json:"journalNumber"`
	Self          string `json:"self"`
	Settings      struct {
		VoucherNumbers struct {
			MinimumVoucherNumber int `json:"minimumVoucherNumber"`
		} `json:"voucherNumbers"`
		EntryTypeRestrictedTo string `json:"entryTypeRestrictedTo"`
		ContraAccounts        struct {
			CustomerPayments struct {
				AccountNumber int    `json:"accountNumber"`
				Self          string `json:"self"`
			} `json:"customerPayments"`
		} `json:"contraAccounts"`
	} `json:"settings,omitempty"`
}
