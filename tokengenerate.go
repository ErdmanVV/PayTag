package tokengenerate

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

// tokenGenerator - структура для генерации токенов
type tokenGenerator struct {
	mu     sync.Mutex
	tokens map[string]bool
}

// newTokenGenerator - конструктор для tokenGenerator
func newTokenGenerator() *tokenGenerator {
	return &tokenGenerator{
		tokens: make(map[string]bool),
	}
}

// GenerateToken генерирует случайный токен
func (tg *tokenGenerator) GenerateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(b)
	tg.mu.Lock()
	defer tg.mu.Unlock()
	if tg.tokens[token] {
		return tg.GenerateToken()
	}
	tg.tokens[token] = true
	return token, nil
}
