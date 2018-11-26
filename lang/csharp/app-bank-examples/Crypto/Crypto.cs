using System;
using System.IO;
using Org.BouncyCastle.Crypto;
using Org.BouncyCastle.OpenSsl;
namespace Crypto
{
    public static class Crypto
    {
        public static AsymmetricKeyParameter GetPublicKey(String pem){
            TextReader tR = new StringReader(pem);
            PemReader pR = new PemReader(tR);
            AsymmetricKeyParameter pubkey =(AsymmetricKeyParameter)r.ReadObject();
            return pubkey;
        }
        public static AsymmetricKeyParameter GetPrivateKey(String pem){
            TextReader tR = new StringReader(pem);
            PemReader pR = new PemReader(tR);
            AsymmetricCipherKeyPair keypair = (AsymmetricCipherKeyPair)r.ReadObject();
            return keypair.Private;
        }
    }
}
