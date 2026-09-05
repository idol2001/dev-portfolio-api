package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	privateKey     *rsa.PrivateKey
	publicKeyPEM   string
	keyOnce        sync.Once
	keyErr         error
)

// loadOrGenerateKeys 加载已有密钥或生成新密钥
func loadOrGenerateKeys() error {
	keyOnce.Do(func() {
		keyDir := getKeysDir()
		privPath := filepath.Join(keyDir, "rsa_private.pem")
		pubPath := filepath.Join(keyDir, "rsa_public.pem")

		// 尝试加载已有密钥
		if _, err := os.Stat(privPath); err == nil {
			privData, _ := os.ReadFile(privPath)
			pubData, _ := os.ReadFile(pubPath)
			if len(privData) > 0 && len(pubData) > 0 {
				block, _ := pem.Decode(privData)
				if block == nil {
					keyErr = fmt.Errorf("failed to decode PEM block")
					return
				}
				privateKey, keyErr = x509.ParsePKCS1PrivateKey(block.Bytes)
				if keyErr == nil {
					publicKeyPEM = string(pubData)
				}
				return
			}
		}

		// 生成新密钥对
		keyErr = generateKeys(privPath, pubPath)
	})
	return keyErr
}

// generateKeys 生成 RSA 密钥对并保存到文件
func generateKeys(privPath, pubPath string) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key: %w", err)
	}

	// 保存私钥
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	privBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}
	privPEM := pem.EncodeToMemory(privBlock)

	if err := os.MkdirAll(filepath.Dir(privPath), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(privPath, privPEM, 0600); err != nil {
		return fmt.Errorf("failed to write private key: %w", err)
	}

	// 保存公钥
	pubBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return err
	}
	pubBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}
	pubPEM := pem.EncodeToMemory(pubBlock)

	if err := os.WriteFile(pubPath, pubPEM, 0644); err != nil {
		return fmt.Errorf("failed to write public key: %w", err)
	}

	privateKey = key
	publicKeyPEM = string(pubPEM)
	return nil
}

// GetPublicKeyPEM 返回 PEM 格式的公钥（供前端获取）
func GetPublicKeyPEM() (string, error) {
	if err := loadOrGenerateKeys(); err != nil {
		return "", err
	}
	return publicKeyPEM, nil
}

// DecryptPassword 使用私钥解密前端传来的 Base64 编码密文
func DecryptPassword(encryptedBase64 string) (string, error) {
	if err := loadOrGenerateKeys(); err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %w", err)
	}

	return string(plaintext), nil
}

func getKeysDir() string {
	// 尝试从环境变量获取，否则使用默认路径
	dir := os.Getenv("RSA_KEYS_DIR")
	if dir == "" {
		dir = "./keys"
	}
	return dir
}
