using System;
using Xunit;
using System.IO;
using System.Reflection;
using Org.BouncyCastle.Crypto;
using System.Text;
namespace Tests{
    public class CryptoTest{
        [Fact]
        public void TestGetPublicKey(){
            Stream pemSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testpk.pem");
            StreamReader sr = new StreamReader(pemSt);
            String pem = sr.ReadToEnd();
            Assert.False(Crypto.Crypto.GetPublicKey(pem).IsPrivate);
        }
        [Fact]
        public void TestGetPrivateKey()
        {
            Stream pemSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testsk.pem");
            StreamReader sr = new StreamReader(pemSt);
            String pem = sr.ReadToEnd();
            Assert.True(Crypto.Crypto.GetPrivateKey(pem).IsPrivate);

        }
        [Fact]
        public void TestRSASign()
        {
            Stream pemSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testsk.pem");
            StreamReader sr = new StreamReader(pemSt);
            String pem = sr.ReadToEnd();
            AsymmetricKeyParameter sk = Crypto.Crypto.GetPrivateKey(pem);
            Stream pemPkSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testpk.pem");
            StreamReader srPk = new StreamReader(pemPkSt);
            String pemPk = srPk.ReadToEnd();
            AsymmetricKeyParameter pk = Crypto.Crypto.GetPublicKey(pemPk);
            string src = "hello world!";
            string sig = Crypto.Crypto.SignWithRSA(Encoding.UTF8.GetBytes(src),sk);
            bool result = Crypto.Crypto.SignVerifyWithRSA(Encoding.UTF8.GetBytes(src), sig,pk);
            Assert.True(result);
        }
        [Fact]
        public void TestRSAEncryptDecrypt(){
            Stream pemSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testsk.pem");
            StreamReader sr = new StreamReader(pemSt);
            String pem = sr.ReadToEnd();
            AsymmetricKeyParameter sk = Crypto.Crypto.GetPrivateKey(pem);
            Stream pemPkSt = Assembly.GetExecutingAssembly().GetManifestResourceStream("app_bank_examples.Resources.testpk.pem");
            StreamReader srPk = new StreamReader(pemPkSt);
            String pemPk = srPk.ReadToEnd();
            AsymmetricKeyParameter pk = Crypto.Crypto.GetPublicKey(pemPk);
            string src = "hello world!";
            byte[] encrypted = Crypto.Crypto.EncryptPKCS1ByRSAPk(Encoding.UTF8.GetBytes(src),pk);
            byte[] decrypted = Crypto.Crypto.DecryptPKCS1ByRSASk(encrypted,sk);

            Assert.Equal(src, Encoding.UTF8.GetString(decrypted));
        }
        [Fact]
        public void TestAESEncryptDecrpt(){
            byte[] key = Crypto.Crypto.HexToByteArray("b42523e4cf19c60c367ce46df132d56f");
            string src = "hello world,中国!";
            byte[] encrypted = Crypto.Crypto.EncryptAES(Encoding.UTF8.GetBytes(src),key);
            byte[] decrypted = Crypto.Crypto.DecryptAES(encrypted, key);
            Assert.Equal(src, Encoding.UTF8.GetString(decrypted));
        }
    }
}
