package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/square/go-jose.v2"
)

var privateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func init() {
	// --- Memuat Private Key ---
	privateKeyBase64 := os.Getenv("POWERSYNC_PRIVATE_KEY") // Sesuai dengan nama di .env
	if privateKeyBase64 == "" {
		log.Fatal("privateKeyBase64 not found in .env")
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode private JWK key: %v", err)
	}

	// Parsing JWK private key
	var jwkPrivate jose.JSONWebKey
	err = json.Unmarshal(privateKeyBytes, &jwkPrivate)
	if err != nil {
		log.Fatalf("Failed unmarshal private JWK: %v", err)
	}

	// Konversi JWK ke *rsa.PrivateKey
	privateKey = jwkPrivate.Key.(*rsa.PrivateKey)
	if privateKey == nil {
		log.Fatal("Private key JWK tidak bisa dikonversi ke *rsa.PrivateKey")
	}
	log.Println("Private key JWK berhasil dimuat.")

	// --- Memuat Public Key ---
	publicKeyBase64 := os.Getenv("POWERSYNC_PUBLIC_KEY") // Sesuai dengan nama di .env
	if publicKeyBase64 == "" {
		log.Fatal("publicKeyBase64 not found in .env")
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode public JWK: %v", err)
	}

	// Parsing JWK public key
	var jwkPublic jose.JSONWebKey
	err = json.Unmarshal(publicKeyBytes, &jwkPublic)
	if err != nil {
		log.Fatalf("Failed unmarshal public JWK: %v", err)
	}

	// Konversi JWK ke *rsa.PublicKey
	PublicKey = jwkPublic.Key.(*rsa.PublicKey)
	if PublicKey == nil {
		log.Fatal("Public key JWK tidak bisa dikonversi ke *rsa.PublicKey")
	}
	log.Println("Public key JWK berhasil dimuat.")
}

func GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"sub":     userId,
		"aud":     "powersync-dev",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "powersync-c8dc2e6df0"
	signedToken, err := token.SignedString(privateKey)

	if err != nil {
		log.Println("Failed generate JWT Token " + err.Error())
		return "", err
	}

	return signedToken, nil
}
