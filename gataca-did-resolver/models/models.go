package models

type DIDDocument struct {
	Context         	string                `json:"@context,omitempty"`
	Id					string             `json:"id,omitempty"`
	PublicKey			[]*PublicKeyEd25519 `json:"publicKey,omitempty"`
	Authentication		[]string           `json:"authentication,omitempty"`
	Service      		*Service              `json:"service,omitempty"`
	Ledger				string             `json:"ledger,omitempty"`
}

type PublicKeyEd25519 struct {
	Context				[]string			`json:"@context,omitempty"`
	Id					string    			`json:"id,omitempty"`
	Type				string    			`json:"type,omitempty"`
	Controller			string    			`json:"controller,omitempty"`
	Key					string    			`json:"publicKeyBase58,omitempty"`
}

type Service struct {
	Id					string    			`json:"id,omitempty"`
	Type				string    			`json:"type,omitempty"`
	ServiceEndpoint		string    			`json:"serviceEndpoint,omitempty"`
}
