// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package share_entity

// EduFdCurrency is the golang structure for table edu_fd_currency.
type EduFdCurrency struct {
	Code          string  `json:"code"          description:"国家编码"`
	EnName        string  `json:"enName"        description:"国家英文名称"`
	CnName        string  `json:"cnName"        description:"国家中文名称"`
	CurrencyCode  string  `json:"currencyCode"  description:"货币编码"`
	CurrencyCn    string  `json:"currencyCn"    description:"货币中文名称"`
	CurrencyEn    string  `json:"currencyEn"    description:"货币英文名称"`
	Symbol        string  `json:"symbol"        description:"货币符号"`
	SymbolNative  string  `json:"symbolNative"  description:"货币原生符号"`
	IsLegalTender int     `json:"isLegalTender" description:"是否法定货币：1是，0否"`
	CurrencyRate  float64 `json:"currencyRate"  description:"货币汇率，本币为人民币"`
}
