using System;
using Grpc.Core;
using Developer;
using Tsv;
using App.User;
using System.Text;
using Google.Protobuf;
namespace app_bank_examples
{
    class Program
    {
        static void Main(string[] args)
        {
            String addressAppBank = "open-api.akeywallet.com:80";
            String addressTsv = "open-api-2sv.akeywallet.com:80";
            Channel channelAppBank = new Channel(addressAppBank, ChannelCredentials.Insecure);
            Channel channelTsv = new Channel(addressTsv, ChannelCredentials.Insecure);
            var developerClient = new DeveloperService.DeveloperServiceClient(channelAppBank);
            var appUserClient = new AppUserService.AppUserServiceClient(channelAppBank);
            var tsvClient = new TSVService.TSVServiceClient(channelTsv);
            String pkString = System.IO.File.ReadAllText(@"conf/public.pem");
            String skString = System.IO.File.ReadAllText(@"conf/private.pem");
            // replace to your own accessKey
            String accessKey = "72b90c38e30cd418934753d873fb57344bc1b839";
            // read rsa string from pem
            var pk = Crypto.Crypto.GetPublicKey(pkString);
            var sk = Crypto.Crypto.GetPrivateKey(skString);
            // init session,get access key from your web console,CtyptoType select RSA or CURVE,curve25519 is a advanced cryto method
            Session session = developerClient.initSession(new SessionRequest { AccessKey = accessKey, CryptoType = "RSA" });
            // save session secret key
            var secretKeyBuff = Crypto.Crypto.HexToByteArray(session.SecretKey);
            var secretKey = Crypto.Crypto.DecryptPKCS1ByRSASk(secretKeyBuff, sk);
            // get balance from your bank, need sign protobuf
            // marshal session
            byte[] sessionBuff = new byte[session.CalculateSize()];
            CodedOutputStream sessionStream = new CodedOutputStream(sessionBuff);
            session.WriteTo(sessionStream);
            // sign session
            string signedSession = Crypto.Crypto.SignWithRSA(sessionBuff, sk);
            var metaSession = new Metadata();
            metaSession.Add("sign", signedSession);
            var bankBalance = developerClient.getBalance(session, metaSession);
            // save a currency named BWC
            AppBankCurrency bwcCurrency = null;
            for (int i = 0; i < bankBalance.BankCurrencies.Count; i++)
            {
                if (bankBalance.BankCurrencies[i].CoinType == "BWC")
                {
                    bwcCurrency = bankBalance.BankCurrencies[i];
                }
            }
            Console.WriteLine(bwcCurrency.CoinType);
            Console.WriteLine(bwcCurrency.Balance);
            // sign in user 1 user 2
            var user1 = appUserClient.signInAppUser(new AppUser { AppUserId = "ABC5", 
                SessionId = session.SessionId,CryptoType = "RSA"});
            var user2 = appUserClient.signInAppUser(new AppUser { AppUserId = "DEF5", 
                SessionId = session.SessionId,CryptoType = "RSA"});
            Console.WriteLine("user 1 is:{0}",user1.Id);
            Console.WriteLine("user 2 is:{0}",user2.Id);
            // decrypt user's pk sk;
            var user1SkSecretBuff = Crypto.Crypto.HexToByteArray(user1.SecretPrivateKey);
            var user1PemSk = Crypto.Crypto.DecryptAES(user1SkSecretBuff, secretKey);
            Console.WriteLine("user1 sk is:{0}", Encoding.UTF8.GetString(user1PemSk));
            var user1Sk = Crypto.Crypto.GetPrivateKey(Encoding.UTF8.GetString(user1PemSk));
            // send some coin from bank to user1,ensure your account has enough money
            var bankTx1 = new BankTx 
            { 
                SessionId = session.SessionId,
                TxType = "out",
                AppUserId = "ABC5",
                CoinType = bwcCurrency.CoinType,
                MainNet = bwcCurrency.MainNet,
                CurrencyId = bwcCurrency.Id,
                CryptoType = "RSA",
                Amount = "100"
            };
            // marshal banktx
            byte[] bankTx1Buff = new byte[bankTx1.CalculateSize()];
            CodedOutputStream bankTx1Stream = new CodedOutputStream(bankTx1Buff);
            bankTx1.WriteTo(bankTx1Stream);
            // sign user1
            string signedBank1 = Crypto.Crypto.SignWithRSA(bankTx1Buff, sk);
            var metaBankTx1 = new Metadata();
            metaBankTx1.Add("sign", signedBank1);
            var tx1 = developerClient.sendTxFromBank(bankTx1, metaBankTx1);
            Console.WriteLine("tx1 id is:{0}",tx1.Id);

            // get user1 balance
            var user1Pb = new AppUser
            {
                AppUserId = "ABC5",
                SessionId = session.SessionId,
                CryptoType = "RSA"
            };
            // marshal user1
            byte[] user1Buff = new byte[user1Pb.CalculateSize()];
            CodedOutputStream user1Stream = new CodedOutputStream(user1Buff);
            user1Pb.WriteTo(user1Stream);
            // sign user1
            string signedUser1 = Crypto.Crypto.SignWithRSA(user1Buff, user1Sk);
            var metaSignedUser1 = new Metadata();
            metaSignedUser1.Add("sign", signedUser1);
            var user1Balance = appUserClient.getBalance(user1Pb,metaSignedUser1);
            Console.WriteLine("user1 bwc :{0}:{1}",user1Balance.AppUserCurrencies[0].CoinType,user1Balance.AppUserCurrencies[0].Balance);
            // save user1 bwc currency
            AppUserCurrency user1BwcCurrency = null;
            for (int i = 0; i < user1Balance.AppUserCurrencies.Count; i++)
            {
                if (user1Balance.AppUserCurrencies[i].CoinType == "BWC")
                {
                    user1BwcCurrency = user1Balance.AppUserCurrencies[i];
                }
            }
            // send coin from one user to another user,before send,user need bind 2step verify
            var userTx1 = new AppUserTx
            {
                AppId = session.AppId,
                CryptoType = "RSA",
                FromUserCurrencyId = user1BwcCurrency.Id,
                ToUserId = user2.AppUserId,
                Amount = "10"
            };
            // marshal userTx1
            byte[] userTx1Buff = new byte[userTx1.CalculateSize()];
            CodedOutputStream userTx1Stream = new CodedOutputStream(userTx1Buff);
            userTx1.WriteTo(userTx1Stream);
            // sign user1
            string signedUserTx1 = Crypto.Crypto.SignWithRSA(userTx1Buff, user1Sk);
            var metaSignedUserTx1 = new Metadata();
            metaSignedUserTx1.Add("sign", signedUserTx1);
            var txUser1 = appUserClient.createTx(userTx1,metaSignedUserTx1);
            Console.WriteLine("txuser1 id:{0}",txUser1.Id);
            /*
            // bind two step verify example
            var keyRequest = new KeyRequest
            {
                SessionId = session.SessionId,
                Label = user1.AppUserId
            };
            var keyTsv = tsvClient.generateKey(keyRequest);
            Console.WriteLine("tsv qr url is:{0}",keyTsv.QrUrl);
            var bindKeyRequest = new BindKeyRequest
            {
                Key = "EH7EBYJFRNVPJLYW",
                ProviderName = "akey.io",
                VerifyCode = "191651",
                Label = user1.AppUserId
            };
            var bindKey = tsvClient.bindKey(bindKeyRequest);
            */
            var sendUserTx1 = new AppUserTx
            {
                Id = txUser1.Id,
                CryptoType = "RSA",
                Code = "866642",
                SessionId = session.SessionId
            };
            // marshal sendUserTx1
            byte[] sendUserTx1Buff = new byte[sendUserTx1.CalculateSize()];
            CodedOutputStream sendUserTx1Stream = new CodedOutputStream(sendUserTx1Buff);
            sendUserTx1.WriteTo(sendUserTx1Stream);
            // sign user1
            string signedSendUserTx1 = Crypto.Crypto.SignWithRSA(sendUserTx1Buff, user1Sk);
            var metaSignedSendUserTx1 = new Metadata();
            metaSignedSendUserTx1.Add("sign", signedSendUserTx1);
            var sendTxUser1 = appUserClient.sendTx(sendUserTx1, metaSignedSendUserTx1);

        }
    }
}
