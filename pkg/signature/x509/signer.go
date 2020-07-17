package x509

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/docker/libtrust"
	"github.com/notaryproject/nv2/pkg/signature"
)

type signer struct {
	key      libtrust.PrivateKey
	cert     *x509.Certificate
	rawCerts [][]byte
	hash     crypto.Hash
}

// NewSignerFromFiles creates a signer from files
func NewSignerFromFiles(keyPath, certPath string) (signature.Signer, error) {
	// Read key
	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	key, err := libtrust.UnmarshalPrivateKeyPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	// Read certificate
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	var (
		certs     []*x509.Certificate
		certBlock *pem.Block
	)
	for {
		certBlock, certBytes = pem.Decode(certBytes)
		if certBlock == nil {
			break
		}
		cert, err := x509.ParseCertificate(certBlock.Bytes)
		if err != nil {
			return nil, err
		}
		certs = append(certs, cert)
	}

	return NewSigner(key, certs)
}

// NewSigner creates a signer
func NewSigner(key libtrust.PrivateKey, certs []*x509.Certificate) (signature.Signer, error) {
	if len(certs) == 0 {
		return nil, errors.New("missing certificates")
	}

	cert := certs[0]
	publicKey, err := libtrust.FromCryptoPublicKey(crypto.PublicKey(cert.PublicKey))
	if err != nil {
		return nil, err
	}
	if key.KeyID() != publicKey.KeyID() {
		return nil, errors.New("key and certificate mismatch")
	}

	rawCerts := make([][]byte, 0, len(certs))
	for _, cert := range certs {
		rawCerts = append(rawCerts, cert.Raw)
	}

	return &signer{
		key:      key,
		cert:     cert,
		rawCerts: rawCerts,
		hash:     crypto.SHA256,
	}, nil
}

func (s *signer) Sign(raw []byte) (signature.Signature, error) {
	if err := verifyReferences(raw, s.cert); err != nil {
		return signature.Signature{}, err
	}
	sig, alg, err := s.key.Sign(bytes.NewReader(raw), s.hash)
	if err != nil {
		return signature.Signature{}, err
	}
	return signature.Signature{
		Type:      Type,
		Algorithm: alg,
		X5c:       s.rawCerts,
		Signature: sig,
	}, nil
}
