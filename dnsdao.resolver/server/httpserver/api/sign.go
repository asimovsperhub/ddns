package api

//func NewWallet(auth string) (Wallet, error) {
//	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
//	if err != nil {
//		log.Println("Error generate account key: %v", err)
//		return nil, err
//	}
//
//	keyBytes := math.PaddedBigBytes(privateKeyECDSA.D, 32)
//	cryptoStruct, err := keystore.EncryptDataV3(keyBytes, []byte(auth), keystore.StandardScryptN, keystore.StandardScryptP)
//	if err != nil {
//		return nil, err
//	}
//
//	pub, pri, err := ed25519.GenerateKey(rand.Reader)
//	if err != nil {
//		logger.Errorf("Error generate network key: %v", err)
//		return nil, err
//	}
//	cipherTxt, err := encryptSubPriKey(pri, pub, auth)
//	if err != nil {
//		logger.Errorf("encrypt wallet err:%f", err)
//		return nil, err
//	}
//
//	obj := &PWallet{
//		Version:   WalletVersion,
//		MainAddr:  crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
//		SubAddr:   ConvertToID2(pub),
//		Crypto:    cryptoStruct,
//		SubCipher: cipherTxt,
//		key: &WalletKey{
//			SubPriKey:  pri,
//			MainPriKey: privateKeyECDSA,
//		},
//	}
//	return obj, nil
//}
//
//func (pw *PWallet) Sign(v []byte) ([]byte, error) {
//	return crypto.Sign(v, pw.key.MainPriKey)
//}
//
//func (pw *PWallet) VerifySig(message, signature []byte) bool {
//	hash := crypto.Keccak256Hash(message)
//	pk := crypto.FromECDSAPub(&pw.key.MainPriKey.PublicKey)
//	signature = signature[:len(signature)-1]
//	return crypto.VerifySignature(pk, hash[:], signature)
//}
