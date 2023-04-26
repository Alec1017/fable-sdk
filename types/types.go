package types

type Contract struct {
	CodeId uint64
	Addr   string
}

type FactoryToken struct {
	CreatorAddr string
	Denom       string
	Subdenom    string
}
