<?php
/**
 * Created by PhpStorm.
 * User: jeffma
 * Date: 2018-11-30
 * Time: 10:42
 */
require 'vendor/autoload.php';
function RSASign($data,$sk){
    openssl_sign($data, $sign, $sk, OPENSSL_ALGO_SHA1);
    $sign = bin2hex($sign);
    return $sign;
};
$pem_pk = file_get_contents("Conf/public.pem");
$pem_sk = file_get_contents("Conf/private.pem");
$pk = openssl_pkey_get_public($pem_pk);
$sk =  openssl_pkey_get_private($pem_sk);
$ak = "72b90c38e30cd418934753d873fb57344bc1b839";
$grpc_developer_url = "open-api.akeywallet.com:80";
$grpc_tsv_url = "open-api-2sv.akeywallet.com:80";
$developer_client = new \Developer\DeveloperServiceClient($grpc_developer_url, [
    'credentials' => Grpc\ChannelCredentials::createInsecure()]);
$user_client = new \App\User\AppUserServiceClient($grpc_developer_url, [
    'credentials' => Grpc\ChannelCredentials::createInsecure()]);
$tsv_client = new \Tsv\TSVServiceClient($grpc_tsv_url, [
    'credentials' => Grpc\ChannelCredentials::createInsecure()]);
$session_request = new \Developer\SessionRequest();
$session_request->setAccessKey($ak);
$session_request->setCryptoType("RSA");
list($session, $status) = $developer_client->initSession($session_request)->wait();
echo $session->getSecretKey();
echo "\n";
// save aes key
openssl_private_decrypt(hex2bin($session->getSecretKey()),$origin_key,$sk);

$signed_session = RSASign($session->serializeToString(),$sk);
list($bank_balance,$status) = $developer_client->getBalance($session, ['sign' => [$signed_session]])->wait();
// save bwc currency,bwc is a test currency
foreach($bank_balance->getBankCurrencies() as $currency){
    if($currency->getCoinType()=='BWC'){
        $bwc_currency = $currency;
        break;
    }
}
// sign in users
$user1 = new \App\User\AppUser();
$user1->setSessionId($session->getSessionId());
$user1->setCryptoType("RSA");
$user1->setAppUserId("ABC5");
$signed_user1 = RSASign($user1->serializeToString(),$sk);
list($app_user1,$status) =  $user_client->signInAppUser($user1,['sign' => [$signed_user1]])->wait();
echo "app user1 is:";
echo $app_user1->getId();
echo "\n";

$user2 = new \App\User\AppUser();
$user2->setSessionId($session->getSessionId());
$user2->setCryptoType("RSA");
$user2->setAppUserId("DEF5");
$signed_user2 = RSASign($user2->serializeToString(),$sk);
list($app_user2,$status) =  $user_client->signInAppUser($user2,['sign' => [$signed_user2]])->wait();
echo "app user2 is:";
echo $app_user2->getId();
echo "\n";

// decrypt user1's pem private key
$user1_sk_pem = openssl_decrypt(hex2bin($app_user1->getSecretPrivateKey()),'aes-128-cbc',$origin_key,OPENSSL_RAW_DATA,$origin_key);
$user1_sk = openssl_pkey_get_private($user1_sk_pem);
echo "user 1 sk is:";
echo $user1_sk_pem;
echo "\n";
// send bank tx
$bank_tx1 = new \Developer\BankTx();
$bank_tx1->setSessionId($session->getSessionId());
$bank_tx1->setAppUserId($user1->getAppUserId());
$bank_tx1->setTxType("out");
$bank_tx1->setCurrencyId($bwc_currency->getId());
$bank_tx1->setCoinType($bwc_currency->getCoinType());
$bank_tx1->setMainNet($bwc_currency->getMainNet());
$bank_tx1->setCryptoType("RSA");
$bank_tx1->setAmount("100");
$signed_bank_tx1 = RSASign($bank_tx1->serializeToString(),$sk);
list($message_bank_tx1,$status) = $developer_client->sendTxFromBank($bank_tx1, ['sign' => [$signed_bank_tx1]])->wait();
echo "bank tx1 id is:";
echo $message_bank_tx1->getId();
echo "\n";

// get user1 bwc balance
$signed_user1 = RSASign($user1->serializeToString(),$user1_sk);
list($user1_balance,$status) = $user_client->getBalance($user1,['sign' => [$signed_user1]])->wait();

// save user1 bwc currency,bwc is a test currency

foreach($user1_balance->getAppUserCurrencies() as $user_currency){
    if($user_currency->getCoinType()=='BWC'){
        $user1_bwc_currency = $user_currency;
        break;
    }
}

echo "user1 bwc balance is:";
echo $user1_bwc_currency->getBalance();
echo "\n";

//send coin from one user to another user,before send,user need bind  2step verify to app

$app_user_tx1 = new \App\User\AppUserTx();
$app_user_tx1->setAppId($app_user2->getAppId());
$app_user_tx1->setToUserId($app_user2->getAppUserId());
$app_user_tx1->setCryptoType("RSA");
$app_user_tx1->setAmount("10");
$app_user_tx1->setFromUserCurrencyId($user1_bwc_currency->getId());
$signed_user_tx1 = RSASign($app_user_tx1->serializeToString(),$user1_sk);
list($user_tx1,$status) = $user_client->createTx($app_user_tx1,['sign' => [$signed_user_tx1]])->wait();
echo "user tx1 id is:";
echo $user_tx1->getId();
echo "\n";

// bind 2step verify
// key request
//$key_request = new \Tsv\KeyRequest();
//$key_request->setSessionId($session->getSessionId());
//$key_request->setLabel($user1->getAppUserId());
//list($message_key,$status) = $tsv_client->generateKey($key_request)->wait();
//echo "qr url is:";
//echo $message_key->getQrUrl();
//echo "\n";


//bind key requrest
//$bind_key_request = new \Tsv\BindKeyRequest();
//$bind_key_request->setKey("AVZJM4LWCCTXSACW");
//$bind_key_request->setProviderName("akey.io");
//$bind_key_request->setVerifyCode("479923");
//$bind_key_request->setLabel($user1->getAppUserId());
//list($message_bind_key,$status) = $tsv_client->bindKey($bind_key_request)->wait();

// bind 2step key at user device
// then sendTx with 2step verify code
// change code to real code
//$send_tx1 = new \App\User\AppUserTx();
//$send_tx1->setId($user_tx1->getId());
//$send_tx1->setSessionId($session->getSessionId());
//$send_tx1->setCode("123456");
//$send_tx1->setCryptoType("RSA");
//$signed_send_tx1 = RSASign($send_tx1->serializeToString(),$user1_sk);
//list($message_send_tx1,$status) = $user_client->createTx($send_tx1,['sign' => [$signed_send_tx1]])->wait();


