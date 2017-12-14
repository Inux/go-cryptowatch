package cryptowatchmodels

type Index struct {
	Allowance Index_sub1 `json:"allowance"`
	Result    Index_sub2 `json:"result"`
}

type Index_sub1 struct {
	Cost      float64 `json:"cost"`
	Remaining float64 `json:"remaining"`
}

type Index_sub2 struct {
	Documentation string   `json:"documentation"`
	Indexes       []string `json:"indexes"`
	Revision      string   `json:"revision"`
	Uptime        string   `json:"uptime"`
}
