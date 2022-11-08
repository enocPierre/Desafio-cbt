package model

type Pagamento struct {
	ID            int
	TipoPagamento string    `json:"tipoPagamento"`
	Transacao     Transacao `json:"transaction"`
	Quantia       float64   `json:"quantia"`
	User          string    `json:"user"`
	Contas        int       `json:"contas"`
}
