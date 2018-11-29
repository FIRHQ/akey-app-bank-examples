using System;
using System.IO;
using Org.BouncyCastle.Crypto;
using Org.BouncyCastle.OpenSsl;
using Org.BouncyCastle.Crypto.Parameters;
using Org.BouncyCastle.Security;
using Org.BouncyCastle.Crypto.Engines;
using Org.BouncyCastle.Crypto.Encodings;
using Org.BouncyCastle.Crypto.Modes;
using Org.BouncyCastle.Crypto.Paddings;

namespace Crypto
{
    public static class Crypto
    {
        public static byte[] HexToByteArray(String hex)
        {
            int NumberChars = hex.Length;
            byte[] bytes = new byte[NumberChars / 2];
            for (int i = 0; i < NumberChars; i += 2)
                bytes[i / 2] = Convert.ToByte(hex.Substring(i, 2), 16);
            return bytes;
        }
        public static AsymmetricKeyParameter GetPublicKey(String pem){
            TextReader tR = new StringReader(pem);
            PemReader pR = new PemReader(tR);
            AsymmetricKeyParameter pubkey =(AsymmetricKeyParameter)pR.ReadObject();
            return pubkey;
        }
        public static AsymmetricKeyParameter GetPrivateKey(String pem){
            TextReader tR = new StringReader(pem);
            PemReader pR = new PemReader(tR);
            AsymmetricCipherKeyPair keypair = (AsymmetricCipherKeyPair)pR.ReadObject();
            return keypair.Private;
        }
        public static string SignWithRSA(byte[] data, AsymmetricKeyParameter sk ){
            RsaKeyParameters rsaKeyParameters = (RsaKeyParameters)sk;
            ISigner sig = SignerUtilities.GetSigner("SHA1withRSA");
            sig.Init(true, rsaKeyParameters);
            sig.BlockUpdate(data, 0, data.Length);
            byte[] result = sig.GenerateSignature();
            var hex = BitConverter.ToString(result, 0).Replace("-", string.Empty).ToLower();
            return hex;
        }
        public static bool SignVerifyWithRSA(byte[] data,string hexSign, AsymmetricKeyParameter pk)
        {
            byte[] signedData = HexToByteArray(hexSign);
            RsaKeyParameters rsaKeyParameters = (RsaKeyParameters)pk;
            ISigner sig = SignerUtilities.GetSigner("SHA1withRSA");
            sig.Init(false, rsaKeyParameters);
            sig.BlockUpdate(data, 0, data.Length);
            bool result = sig.VerifySignature(signedData);
            return result;
        }
        public static byte[] DecryptPKCS1ByRSASk(byte[] data, AsymmetricKeyParameter sk){
            byte[] result = null;
            IAsymmetricBlockCipher cipher = new Pkcs1Encoding(new RsaEngine());
            cipher.Init(false, sk);
            try{
                result = cipher.ProcessBlock(data, 0, data.Length);
            }
            catch (InvalidCipherTextException e){
                System.Diagnostics.Debug.WriteLine(e.Message);
                result = null;
            }
            return result;
        }
        public static byte[] EncryptPKCS1ByRSAPk(byte[] data, AsymmetricKeyParameter pk)
        {
            byte[] result = null;
            IAsymmetricBlockCipher cipher = new Pkcs1Encoding(new RsaEngine());
            cipher.Init(true, pk);
            try
            {
                result = cipher.ProcessBlock(data, 0, data.Length);
            }
            catch (InvalidCipherTextException e)
            {
                System.Diagnostics.Debug.WriteLine(e.Message);
                result = null;
            }
            return result;
        }
        public static byte[] EncryptAES(byte[] data,byte[] key){
            AesEngine engine = new AesEngine();
            CbcBlockCipher blockCipher = new CbcBlockCipher(engine);
            PaddedBufferedBlockCipher cipher = new PaddedBufferedBlockCipher(blockCipher);
            KeyParameter keyParam = new KeyParameter(key);
            ParametersWithIV keyParamWithIV = new ParametersWithIV(keyParam, key, 0, key.Length);
            cipher.Init(true, keyParamWithIV);
            byte[] outputBytes = new byte[cipher.GetOutputSize(data.Length)];
            int length1 = cipher.ProcessBytes(data, outputBytes, 0);
            int length2 = cipher.DoFinal(outputBytes, length1); //Do the final block
            int actualLength = length1 + length2;
            byte[] result = new byte[actualLength];
            System.Buffer.BlockCopy(outputBytes,0,result,0,result.Length);
            return result;
        }
        public static byte[] DecryptAES(byte[] data,byte[] key){
            AesEngine engine = new AesEngine();
            CbcBlockCipher blockCipher = new CbcBlockCipher(engine);
            PaddedBufferedBlockCipher cipher = new PaddedBufferedBlockCipher(blockCipher);
            KeyParameter keyParam = new KeyParameter(key);
            ParametersWithIV keyParamWithIV = new ParametersWithIV(keyParam, key, 0, key.Length);
            cipher.Init(false, keyParamWithIV);
            byte[] comparisonBytes = new byte[cipher.GetOutputSize(data.Length)];
            int length1 = cipher.ProcessBytes(data, comparisonBytes, 0);
            int length2 = cipher.DoFinal(comparisonBytes, length1); //Do the final block
            int actualLength = length1 + length2;
            byte[] result = new byte[actualLength];
            System.Buffer.BlockCopy(comparisonBytes, 0, result, 0, result.Length);
            return result;
        }
    }
}
