package signature

// Signer signs content
type Signer interface {
	Sign(content []byte) (Signature, error)
}

// Verifier verifies content
type Verifier interface {
	Verify(content []byte, signature Signature) error
}

// Scheme can sign and verify
type Scheme interface {
	Signer
	Verifier
}
