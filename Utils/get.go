package get

import (
  "json"
	"https://fortnitecontent-website-prod07.ol.epicgames.com/content/api/pages/fortnite-game"
	"https://launcher-website-prod07.ol.epicgames.com/purchase%7Bparameters%7D"
	"https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/game/v2/creative/favorites/account_user/"
	"https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/storefront/v2/keychain"
	"https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/storefront/v2/catalog"
	"https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/calendar/v1/timeline"
)

func get() {
	generateAndWriteKeyPair()
}

func generateAndWriteKeyPair() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Cannot generate key, err: %v", err)
	}

	// Create pem directory if not exists
	if _, err := os.Stat("pem"); os.IsNotExist(err) {
		os.Mkdir("pem", os.ModeDir)
	}

	writePrivateKey(privateKey)
	writePublicKey(publicKey)
}

func writePrivateKey(privateKey ed25519.PrivateKey) {
	pemPrivateFile, err := os.Create("pem/private_key.pem")
	if err != nil {
		log.Fatalf("ERROR", err)
	}

	defer pemPrivateFile.Close()

	key, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("ERROR", err)
	}

	pemPrivateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: key,
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		log.Fatalf("ERROR", err)
	}
}

func writePublicKey(publicKey ed25519.PublicKey) {
	pemPublicFile, err := os.Create("pem/public_key.pem")
	if err != nil {
		log.Fatalf("ERROR", err)
	}

	defer pemPublicFile.Close()

	pubKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("ERROR", err)
	}

	pemPublicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKey,
	}

	err = pem.Encode(pemPublicFile, pemPublicBlock)
	if err != nil {
		log.Fatalf("ERROR, err)
	}
}
