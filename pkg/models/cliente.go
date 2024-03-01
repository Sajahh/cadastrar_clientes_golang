package models

type ClientePayload struct {
	Call      string        `json:"call"`
	AppKey    string        `json:"app_key"`
	AppSecret string        `json:"app_secret"`
	Param     []ClienteOmie `json:"param"`
}

type ClienteOmie struct {
	CodigoClienteIntegracao string `json:"codigo_cliente_integracao"`
	CodigoClienteOmie       string `json:"codigo_cliente_omie"`
	Email                   string `json:"email"`
	RazaoSocial             string `json:"razao_social"`
	CnpjCpf                 string `json:"cnpj_cpf"`
	Endereco                string `json:"endereco"`
	EnderecoNumero          string `json:"endereco_numero"`
	Complemento             string `json:"complemento"`
	Bairro                  string `json:"bairro"`
	Estado                  string `json:"estado"`
	Cidade                  string `json:"cidade"`
	Cep                     string `json:"cep"`
}
