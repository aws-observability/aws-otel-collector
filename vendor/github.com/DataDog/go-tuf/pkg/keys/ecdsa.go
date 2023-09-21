package keys

import (
<<<<<<< HEAD
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/json"
	"errors"
	"math/big"
=======
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
>>>>>>> main

	"github.com/DataDog/go-tuf/data"
)

func init() {
<<<<<<< HEAD
	VerifierMap.Store(data.KeyTypeECDSA_SHA2_P256, NewEcdsaVerifier)
}

func NewEcdsaVerifier() Verifier {
	return &p256Verifier{}
}

type ecdsaSignature struct {
	R, S *big.Int
}

type p256Verifier struct {
	PublicKey data.HexBytes `json:"public"`
	key       *data.PublicKey
}

func (p *p256Verifier) Public() string {
	return p.PublicKey.String()
}

func (p *p256Verifier) Verify(msg, sigBytes []byte) error {
	x, y := elliptic.Unmarshal(elliptic.P256(), p.PublicKey)
	k := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	var sig ecdsaSignature
	if _, err := asn1.Unmarshal(sigBytes, &sig); err != nil {
		return err
	}

	hash := sha256.Sum256(msg)

	if !ecdsa.Verify(k, hash[:], sig.R, sig.S) {
=======
	// Note: we use LoadOrStore here to prevent accidentally overriding the
	// an explicit deprecated ECDSA verifier.
	// TODO: When deprecated ECDSA is removed, this can switch back to Store.
	VerifierMap.LoadOrStore(data.KeyTypeECDSA_SHA2_P256_OLD_FMT, NewEcdsaVerifier)
	VerifierMap.LoadOrStore(data.KeyTypeECDSA_SHA2_P256, NewEcdsaVerifier)
	SignerMap.Store(data.KeyTypeECDSA_SHA2_P256_OLD_FMT, newEcdsaSigner)
	SignerMap.Store(data.KeyTypeECDSA_SHA2_P256, newEcdsaSigner)
}

func NewEcdsaVerifier() Verifier {
	return &EcdsaVerifier{}
}

func newEcdsaSigner() Signer {
	return &ecdsaSigner{}
}

type EcdsaVerifier struct {
	PublicKey *PKIXPublicKey `json:"public"`
	ecdsaKey  *ecdsa.PublicKey
	key       *data.PublicKey
}

func (p *EcdsaVerifier) Public() string {
	// This is already verified to succeed when unmarshalling a public key.
	r, err := x509.MarshalPKIXPublicKey(p.ecdsaKey)
	if err != nil {
		// TODO: Gracefully handle these errors.
		// See https://github.com/DataDog/go-tuf/issues/363
		panic(err)
	}
	return string(r)
}

func (p *EcdsaVerifier) Verify(msg, sigBytes []byte) error {
	hash := sha256.Sum256(msg)

	if !ecdsa.VerifyASN1(p.ecdsaKey, hash[:], sigBytes) {
>>>>>>> main
		return errors.New("tuf: ecdsa signature verification failed")
	}
	return nil
}

<<<<<<< HEAD
func (p *p256Verifier) MarshalPublicKey() *data.PublicKey {
	return p.key
}

func (p *p256Verifier) UnmarshalPublicKey(key *data.PublicKey) error {
	if err := json.Unmarshal(key.Value, p); err != nil {
		return err
	}
	x, _ := elliptic.Unmarshal(elliptic.P256(), p.PublicKey)
	if x == nil {
		return errors.New("tuf: invalid ecdsa public key point")
	}
	p.key = key
	return nil
}
=======
func (p *EcdsaVerifier) MarshalPublicKey() *data.PublicKey {
	return p.key
}

func (p *EcdsaVerifier) UnmarshalPublicKey(key *data.PublicKey) error {
	// Prepare decoder limited to 512Kb
	dec := json.NewDecoder(io.LimitReader(bytes.NewReader(key.Value), MaxJSONKeySize))

	// Unmarshal key value
	if err := dec.Decode(p); err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return fmt.Errorf("tuf: the public key is truncated or too large: %w", err)
		}
		return err
	}

	ecdsaKey, ok := p.PublicKey.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("invalid public key")
	}

	if _, err := x509.MarshalPKIXPublicKey(ecdsaKey); err != nil {
		return fmt.Errorf("marshalling to PKIX key: invalid public key")
	}

	p.ecdsaKey = ecdsaKey
	p.key = key
	return nil
}

type ecdsaSigner struct {
	*ecdsa.PrivateKey
}

type ecdsaPrivateKeyValue struct {
	Private string         `json:"private"`
	Public  *PKIXPublicKey `json:"public"`
}

func (s *ecdsaSigner) PublicData() *data.PublicKey {
	// This uses a trusted public key JSON format with a trusted Public value.
	keyValBytes, _ := json.Marshal(EcdsaVerifier{PublicKey: &PKIXPublicKey{PublicKey: s.Public()}})
	return &data.PublicKey{
		Type:       data.KeyTypeECDSA_SHA2_P256,
		Scheme:     data.KeySchemeECDSA_SHA2_P256,
		Algorithms: data.HashAlgorithms,
		Value:      keyValBytes,
	}
}

func (s *ecdsaSigner) SignMessage(message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	return ecdsa.SignASN1(rand.Reader, s.PrivateKey, hash[:])
}

func (s *ecdsaSigner) MarshalPrivateKey() (*data.PrivateKey, error) {
	priv, err := x509.MarshalECPrivateKey(s.PrivateKey)
	if err != nil {
		return nil, err
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: priv})
	val, err := json.Marshal(ecdsaPrivateKeyValue{
		Private: string(pemKey),
		Public:  &PKIXPublicKey{PublicKey: s.Public()},
	})
	if err != nil {
		return nil, err
	}
	return &data.PrivateKey{
		Type:       data.KeyTypeECDSA_SHA2_P256,
		Scheme:     data.KeySchemeECDSA_SHA2_P256,
		Algorithms: data.HashAlgorithms,
		Value:      val,
	}, nil
}

func (s *ecdsaSigner) UnmarshalPrivateKey(key *data.PrivateKey) error {
	val := ecdsaPrivateKeyValue{}
	if err := json.Unmarshal(key.Value, &val); err != nil {
		return err
	}
	block, _ := pem.Decode([]byte(val.Private))
	if block == nil {
		return errors.New("invalid PEM value")
	}
	if block.Type != "EC PRIVATE KEY" {
		return fmt.Errorf("invalid block type: %s", block.Type)
	}
	k, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	if k.Curve != elliptic.P256() {
		return errors.New("unsupported ecdsa curve")
	}
	if _, err := json.Marshal(EcdsaVerifier{
		PublicKey: &PKIXPublicKey{PublicKey: k.Public()}}); err != nil {
		return fmt.Errorf("invalid public key: %s", err)
	}

	s.PrivateKey = k
	return nil
}

func GenerateEcdsaKey() (*ecdsaSigner, error) {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return &ecdsaSigner{privkey}, nil
}
>>>>>>> main
