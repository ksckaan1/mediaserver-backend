package password

import (
	"fmt"
	"strings"

	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
)

type Password struct {
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
}

func New() *Password {
	return &Password{
		time:      3,
		memory:    64 * 1024,
		threads:   4,
		keyLength: 32,
	}
}

func (p *Password) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (p *Password) HashPassword(pw string) (string, error) {
	salt, err := p.generateRandomBytes(16)
	if err != nil {
		return "", fmt.Errorf("generateRandomBytes: %w", err)
	}

	hash := argon2.IDKey([]byte(pw), salt, p.time, p.memory, p.threads, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// $argon2id$v=19$m=65536,t=3,p=4$salt$hash formatı
	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", p.memory, p.time, p.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

func (p *Password) VerifyPassword(password, encodedHash string) (bool, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return false, errors.New("invalid hash format")
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return false, err
	}
	if version != 19 {
		return false, fmt.Errorf("unsupported version: %d", version)
	}

	var memory, time uint32
	var threads uint8
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return false, err
	}

	keyLength := uint32(len(decodedHash))
	comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}
