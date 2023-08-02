package solimpl

func (a *AssertionChain) SetBackend(b ChainBackend) {
	a.backend = b
}

func (a *AssertionChain) SetClient(c BatchClient) {
	a.client = c
}
