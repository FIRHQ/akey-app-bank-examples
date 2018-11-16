package main

import (
	pb_developer "akey-app-bank-examples/grpc/app-bank/developer"
	pb_user "akey-app-bank-examples/grpc/app-bank/user"
	pb_tsv "akey-app-bank-examples/grpc/tsv/tsv"
	"bytes"
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
)

const (
	address_app_bank     = "open-api.akeywallet.com:80"
	address_tsv = "open-api-2sv.akeywallet.com:80"
)

func padding(src []byte,blocksize int) []byte {
	padnum:=blocksize-len(src)%blocksize
	pad:=bytes.Repeat([]byte{byte(padnum)},padnum)
	return append(src,pad...)
}

func unpadding(src []byte) []byte {
	n:=len(src)
	unpadnum:=int(src[n-1])
	return src[:n-unpadnum]
}

func encryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	src=padding(src,block.BlockSize())
	blockmode:=cipher.NewCBCEncrypter(block,key)
	blockmode.CryptBlocks(src,src)
	return src
}

func decryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	blockmode:=cipher.NewCBCDecrypter(block,key)
	blockmode.CryptBlocks(src,src)
	src=unpadding(src)
	return src
}

func parsePemPrivate(src []byte) (*rsa.PrivateKey, error){
	block, _ := pem.Decode(src)
	if block == nil {
		return nil, errors.New("pem read error")
	}
	// parse sk
	sk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}else{
		return sk, nil
	}
}
func parsePemPublic(src []byte) (*rsa.PublicKey, error){
	block, _ := pem.Decode(src)
	if block == nil {
		return nil, errors.New("pem read error")
	}
	// parse pk
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}else{
		return pk.(*rsa.PublicKey), nil
	}
}

func main() {
	// replace to your own accessKey
	accessKey := "72b90c38e30cd418934753d873fb57344bc1b839"
	// read rsa string from pem file
	pemPrivate,err := ioutil.ReadFile("./conf/private.pem")
	if err != nil {
		log.Fatalf("rsa read error:%v",err)
	}
	pemPublic,err := ioutil.ReadFile("./conf/public.pem")
	if err != nil {
		log.Fatalf("rsa read error:%v",err)
	}
	sk, err := parsePemPrivate([]byte(pemPrivate))
	if err != nil {
		log.Fatal("private key read error1:%v" ,err)
	}else{
		log.Println(sk)
	}
	pk, err := parsePemPublic([]byte(pemPublic))
	if err != nil {
		log.Fatal("public key read error:%v" ,err)
	}else{
		log.Println(pk)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address_app_bank, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	connTsv, err := grpc.Dial(address_tsv, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect tsv: %v", err)
	}
	defer connTsv.Close()

	clientDeveloper := pb_developer.NewDeveloperServiceClient(conn)
	// init session,get access key from your web console,CtyptoType select RSA or CURVE,curve25519 is a advanced cryto method
    session, err := clientDeveloper.InitSession(context.Background(),&pb_developer.SessionRequest{AccessKey: accessKey, CryptoType: "RSA"})
    if err != nil{
    	log.Fatal("init session error:%v",err)
	}else{
		log.Println("session init finished")
    	log.Println(session)
	}
    // save session secret key
    secretKeyBuff,err := hex.DecodeString(session.SecretKey)
    if err!= nil {
    	log.Fatal("save secret key err:%v",err)
	}
	secretKey,err := rsa.DecryptPKCS1v15(rand.Reader, sk, secretKeyBuff)
	if err!= nil {
		log.Fatal("save secret key err:%v",err)
	}else{
		log.Println("aes key:",secretKey)
	}

    // get balance from your bank, need sign protobuf
    // marshal session
	sessionBuff, err := proto.Marshal(session)
	if err != nil{
		log.Fatal("get session buff error:%v",err)
	}else{
		log.Println(string(sessionBuff))
	}
	// sign session
	hash := sha1.New()
	hash.Write(sessionBuff)
	signData, err := rsa.SignPKCS1v15(rand.Reader, sk, crypto.SHA1, hash.Sum(nil))
	if err != nil{
		log.Fatal("get session buff error:%v",err)
	}else{
		log.Println(signData)
	}
    // signData to hex
	hexSign := hex.EncodeToString(signData)
	log.Println(hexSign)
	md := metadata.Pairs("sign", hexSign)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
    balance, err := clientDeveloper.GetBalance(ctx,session)
    if err != nil{
		log.Fatal("get balance error:%v",err)
	}else{
		log.Println(balance.GetBankCurrencies())
	}
    // then send currency to your address

    // sign a user to your app
    clientUser := pb_user.NewAppUserServiceClient(conn)

	// CtyptoType select RSA or CURVE
	appUser1, err := clientUser.SignInAppUser(context.Background(),&pb_user.AppUser{SessionId:session.SessionId,AppUserId:"ABC5",CryptoType:"RSA"})
	if err != nil{
		log.Fatal("sign in user error:%v",err)
	}else{
		log.Println(appUser1)
	}
	appUser2, err := clientUser.SignInAppUser(context.Background(),&pb_user.AppUser{SessionId:session.SessionId,AppUserId:"DEF5",CryptoType:"RSA"})
	if err != nil{
		log.Fatal("sign in user error:%v",err)
	}else{
		log.Println(appUser2)
	}
    // docrypt user's pk sk;
/*
    user1PkSecretBuff,err := hex.DecodeString(appUser1.SecretPublicKey)
    if err != nil{
	   log.Fatal("get user1 pk buff error:%v",err)
    }
*/
	user1SkSecretBuff,err := hex.DecodeString(appUser1.SecretPrivateKey)
	if err != nil{
		log.Fatal("get user1 sk buff error:%v",err)
	}

	user1SkBuff := decryptAES(user1SkSecretBuff,secretKey)
	// get user1 balances
	user1 := &pb_user.AppUser{AppId:session.AppId,AppUserId:appUser1.AppUserId,CryptoType:"RSA"}

	// marshal session
	user1Buff, err := proto.Marshal(user1)
	if err != nil{
		log.Fatal("get user1 buff error:%v",err)
	}else{
		log.Println(string(sessionBuff))
	}
	// sign session
	hash_user1 := sha1.New()
	hash_user1.Write(user1Buff)
	user1Sk, err := parsePemPrivate([]byte(user1SkBuff))
	if err != nil {
		log.Fatal("private key read error2:%v" ,err)
	}else{
		log.Println(sk)
	}

	signUser1Data, err := rsa.SignPKCS1v15(rand.Reader, user1Sk, crypto.SHA1, hash_user1.Sum(nil))
	if err != nil{
		log.Fatal("get user1 buff error:%v",err)
	}else{
		log.Println(signUser1Data)
	}

	// signData to hex
	hexUser1Sign := hex.EncodeToString(signUser1Data)
	log.Println(hexUser1Sign)
	mdUser1 := metadata.Pairs("sign", hexUser1Sign)
	ctxUser1 := metadata.NewOutgoingContext(context.Background(), mdUser1)
	balanceUser1, err := clientUser.GetBalance(ctxUser1,user1)
	if err != nil{
		log.Fatal("get user balance error:%v",err)
	}else{
		log.Println(balanceUser1)
	}
	// send some coin from bank to user1,ensure your account has enough money
    var balanceBWC *pb_developer.AppBankCurrency
    bankCurrencies := balance.GetBankCurrencies()
	for i := 0; i < len(bankCurrencies); i++ {
		if bankCurrencies[i].CoinType=="BWC" {
			balanceBWC = bankCurrencies[i]
			break
		}
	}
	// send bank tx
	bankTx1 := &pb_developer.BankTx{SessionId:session.SessionId,TxType:"out",AppId:session.AppId,AppUserId:appUser1.AppUserId,CoinType:balanceBWC.CoinType,MainNet:balanceBWC.MainNet,CurrencyId:balanceBWC.Id,CryptoType:"RSA",Amount:"100"}
	// marshal tx
	tx1Buff, err := proto.Marshal(bankTx1)
	if err != nil{
		log.Fatal("get tx1 buff error:%v",err)
	}else{
		log.Println(string(sessionBuff))
	}
	// sign tx
	hashTx1 := sha1.New()
	hashTx1.Write(tx1Buff)
	signDataTx1, err := rsa.SignPKCS1v15(rand.Reader, sk, crypto.SHA1, hashTx1.Sum(nil))
	if err != nil{
		log.Fatal("sign  error:%v",err)
	}else{
		log.Println(signDataTx1)
	}
	// signData to hex
	hexSignTx1 := hex.EncodeToString(signDataTx1)
	log.Println(hexSignTx1)
	mdSignTx1 := metadata.Pairs("sign", hexSignTx1)
	ctxSignTx1 := metadata.NewOutgoingContext(context.Background(), mdSignTx1)
	// create tx
	tx1, err := clientDeveloper.SendTxFromBank(ctxSignTx1,bankTx1)
	if err != nil{
		log.Fatal("create tx error:%v",err)
	}else{
		log.Println(tx1)
	}

	// send coin from one user to another user,before send,user need bind 2step verify
	// get user currency
	var user1BwcCurrency *pb_user.AppUserCurrency
	user1Currencies := balanceUser1.GetAppUserCurrencies()
	for i := 0; i < len(user1Currencies); i++ {
		if bankCurrencies[i].CoinType=="BWC" {
			user1BwcCurrency = user1Currencies[i]
			break
		}
	}
	userTx1 := &pb_user.AppUserTx{SessionId:session.SessionId,AppId:session.AppId,CryptoType:"RSA",Amount:"10",FromUserCurrencyId:user1BwcCurrency.Id,ToUserId:appUser2.AppUserId,FromUserId:appUser1.AppUserId}
	// marshal tx
	userTx1Buff, err := proto.Marshal(userTx1)
	if err != nil{
		log.Fatal("get tx1 buff error:%v",err)
	}else{
		log.Println(string(userTx1Buff))
	}
	// sign tx
	hashUserTx1 := sha1.New()
	hashUserTx1.Write(userTx1Buff)
	signDataUserTx1, err := rsa.SignPKCS1v15(rand.Reader, user1Sk, crypto.SHA1, hashUserTx1.Sum(nil))
	if err != nil{
		log.Fatal("sign user tx  error:%v",err)
	}else{
		log.Println(signDataUserTx1)
	}
	// signData to hex
	hexSignUserTx1 := hex.EncodeToString(signDataUserTx1)
	log.Println(hexSignUserTx1)
	mdSignUserTx1 := metadata.Pairs("sign", hexSignUserTx1)
	ctxUserSignTx1 := metadata.NewOutgoingContext(context.Background(), mdSignUserTx1)
	// create tx
	txUser1, err := clientUser.CreateTx(ctxUserSignTx1,userTx1)
	if err != nil{
		log.Fatal("create tx error:",err)
	}else{
		log.Println(txUser1)
	}

	// bind 2step verify
	clientTsv := pb_tsv.NewTSVServiceClient(connTsv)
	keyTsv, err := clientTsv.GenerateKey(context.Background(),&pb_tsv.KeyRequest{SessionId:session.SessionId,Label:appUser1.AppUserId})
	if err != nil{
		log.Fatal("generate key error:%v",err)
	}else{
		log.Println("==generate key finished")
		log.Println(keyTsv)
	}
	// bind 2step key at user device
    // then sendTx with 2step verify code
    // change code to real code
    userSendTx1 := &pb_user.AppUserTx{Code:"abc",SessionId:session.SessionId}
	proto.Merge(userSendTx1,txUser1)
	log.Println("merged:",userSendTx1)


	// marshal tx
	userSendTx1Buff, err := proto.Marshal(userSendTx1)
	if err != nil{
		log.Fatal("get tx1 buff error:%v",err)
	}else{
		log.Println(string(userSendTx1Buff))
	}
	// sign tx
	hashUserSendTx1 := sha1.New()
	hashUserSendTx1.Write(userSendTx1Buff)
	signDataUserSendTx1, err := rsa.SignPKCS1v15(rand.Reader, user1Sk, crypto.SHA1, hashUserSendTx1.Sum(nil))
	if err != nil{
		log.Fatal("sign user tx  error:%v",err)
	}else{
		log.Println(signDataUserSendTx1)
	}
	// signData to hex
	hexSignUserSendTx1 := hex.EncodeToString(signDataUserSendTx1)
	log.Println(hexSignUserTx1)
	mdSignUserSendTx1 := metadata.Pairs("sign", hexSignUserSendTx1)
	ctxUserSendSignTx1 := metadata.NewOutgoingContext(context.Background(), mdSignUserSendTx1)
	// create tx
	txUserSend1, err := clientUser.SendTx(ctxUserSendSignTx1,userSendTx1)
	if err != nil{
		log.Fatal("create tx error:",err)
	}else{
		log.Println(txUserSend1)
	}
}
