package keys

import (
<<<<<<< HEAD
=======
	"bytes"
>>>>>>> main
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
<<<<<<< HEAD
=======
	"fmt"
	"io"
>>>>>>> main

	"github.com/DataDog/go-tuf/data"
)

func init() {
<<<<<<< HEAD
	VerifierMap.Store(data.KeyTypeRSASSA_PSS_SHA256, NewRsaVerifier)
	SignerMap.Store(data.KeyTypeRSASSA_PSS_SHA256, NewRsaSigner)
}

func NewRsaVerifier() Verifier {
	return &rsaVerifier{}
}

func NewRsaSigner() Signer {
=======
	VerifierMap.Store(data.KeyTypeRSASSA_PSS_SHA256, newRsaVerifier)
	SignerMap.Store(data.KeyTypeRSASSA_PSS_SHA256, newRsaSigner)
}

func newRsaVerifier() Verifier {
	return &rsaVerifier{}
}

func newRsaSigner() Signer {
>>>>>>> main
	return &rsaSigner{}
}

type rsaVerifier struct {
<<<<<<< HEAD
	PublicKey string `json:"public"`
=======
	PublicKey *PKIXPublicKey `json:"public"`
>>>>>>> main
	rsaKey    *rsa.PublicKey
	key       *data.PublicKey
}

func (p *rsaVerifier) Public() string {
<<<<<<< HEAD
	// Unique public key identifier, use a uniform encodng
	r, err := x509.MarshalPKIXPublicKey(p.rsaKey)
	if err != nil {
		// This shouldn't happen with a valid rsa key, but fallback on the
		// JSON public key string
		return string(p.PublicKey)
=======
	// This is already verified to succeed when unmarshalling a public key.
	r, err := x509.MarshalPKIXPublicKey(p.rsaKey)
	if err != nil {
		// TODO: Gracefully handle these errors.
		// See https://github.com/DataDog/go-tuf/issues/363
		panic(err)
>>>>>>> main
	}
	return string(r)
}

func (p *rsaVerifier) Verify(msg, sigBytes []byte) error {
	hash := sha256.Sum256(msg)

	return rsa.VerifyPSS(p.rsaKey, crypto.SHA256, hash[:], sigBytes, &rsa.PSSOptions{})
}

func (p *rsaVerifier) MarshalPublicKey() *data.PublicKey {
	return p.key
}

func (p *rsaVerifier) UnmarshalPublicKey(key *data.PublicKey) error {
<<<<<<< HEAD
	if err := json.Unmarshal(key.Value, p); err != nil {
		return err
	}
	var err error
	p.rsaKey, err = parseKey(p.PublicKey)
	if err != nil {
		return err
	}
=======
	// Prepare decoder limited to 512Kb
	dec := json.NewDecoder(io.LimitReader(bytes.NewReader(key.Value), MaxJSONKeySize))

	// Unmarshal key value
	if err := dec.Decode(p); err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return fmt.Errorf("tuf: the public key is truncated or too large: %w", err)
		}
		return err
	}

	rsaKey, ok := p.PublicKey.PublicKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("invalid public key")
	}

	if _, err := x509.MarshalPKIXPublicKey(rsaKey); err != nil {
		return fmt.Errorf("marshalling to PKIX key: invalid public key")
	}

	p.rsaKey = rsaKey
>>>>>>> main
	p.key = key
	return nil
}

<<<<<<< HEAD
// parseKey tries to parse a PEM []byte slice by attempting PKCS1 and PKIX in order.
func parseKey(data string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("tuf: pem decoding public key failed")
	}
	rsaPub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return rsaPub, nil
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		rsaPub, ok := key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("tuf: invalid rsa key")
		}
		return rsaPub, nil
	}
	return nil, errors.New("tuf: error unmarshalling rsa key")
}

=======
>>>>>>> main
type rsaSigner struct {
	*rsa.PrivateKey
}

<<<<<<< HEAD
type rsaPublic struct {
	// PEM encoded public key.
	PublicKey string `json:"public"`
}

func (s *rsaSigner) PublicData() *data.PublicKey {
	pub, _ := x509.MarshalPKIXPublicKey(s.Public().(*rsa.PublicKey))
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pub,
	})

	keyValBytes, _ := json.Marshal(rsaPublic{PublicKey: string(pubBytes)})
=======
type rsaPrivateKeyValue struct {
	Private string         `json:"private"`
	Public  *PKIXPublicKey `json:"public"`
}

func (s *rsaSigner) PublicData() *data.PublicKey {
	keyValBytes, _ := json.Marshal(rsaVerifier{PublicKey: &PKIXPublicKey{PublicKey: s.Public()}})
>>>>>>> main
	return &data.PublicKey{
		Type:       data.KeyTypeRSASSA_PSS_SHA256,
		Scheme:     data.KeySchemeRSASSA_PSS_SHA256,
		Algorithms: data.HashAlgorithms,
		Value:      keyValBytes,
	}
}

func (s *rsaSigner) SignMessage(message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	return rsa.SignPSS(rand.Reader, s.PrivateKey, crypto.SHA256, hash[:], &rsa.PSSOptions{})
}

func (s *rsaSigner) ContainsID(id string) bool {
	return s.PublicData().ContainsID(id)
}

func (s *rsaSigner) MarshalPrivateKey() (*data.PrivateKey, error) {
<<<<<<< HEAD
	return nil, errors.New("not implemented for test")
}

func (s *rsaSigner) UnmarshalPrivateKey(key *data.PrivateKey) error {
	return errors.New("not implemented for test")
=======
	priv := x509.MarshalPKCS1PrivateKey(s.PrivateKey)
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: priv})
	val, err := json.Marshal(rsaPrivateKeyValue{
		Private: string(pemKey),
		Public:  &PKIXPublicKey{PublicKey: s.Public()},
	})
	if err != nil {
		return nil, err
	}
	return &data.PrivateKey{
		Type:       data.KeyTypeRSASSA_PSS_SHA256,
		Scheme:     data.KeySchemeRSASSA_PSS_SHA256,
		Algorithms: data.HashAlgorithms,
		Value:      val,
	}, nil
}

func (s *rsaSigner) UnmarshalPrivateKey(key *data.PrivateKey) error {
	val := rsaPrivateKeyValue{}
	if err := json.Unmarshal(key.Value, &val); err != nil {
		return err
	}
	block, _ := pem.Decode([]byte(val.Private))
	if block == nil {
		return errors.New("invalid PEM value")
	}
	if block.Type != "RSA PRIVATE KEY" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}
	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	if _, err := json.Marshal(rsaVerifier{
		PublicKey: &PKIXPublicKey{PublicKey: k.Public()}}); err != nil {
		return fmt.Errorf("invalid public key: %s", err)
	}

	s.PrivateKey = k
	return nil
>>>>>>> main
}

func GenerateRsaKey() (*rsaSigner, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &rsaSigner{privkey}, nil
}
