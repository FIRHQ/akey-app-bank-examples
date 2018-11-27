this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)
require "grpc"
require "developer_services_pb"
require "user_services_pb"
require "tsv_services_pb"
require "openssl"
require "google/protobuf"
module AESCrypt
  # Decrypts a block of data (encrypted_data) given an encryption key
  # and an initialization vector (iv).  Keys, iv's, and the data
  # returned are all binary strings.  Cipher_type should be
  # "AES-256-CBC", "AES-256-ECB", or any of the cipher types
  # supported by OpenSSL.  Pass nil for the iv if the encryption type
  # doesn't use iv's (like ECB).
  #:return: => String
  #:arg: encrypted_data => String
  #:arg: key => String
  #:arg: iv => String
  #:arg: cipher_type => String
  def AESCrypt.decrypt(encrypted_data, key, iv, cipher_type)
    aes = OpenSSL::Cipher::Cipher.new(cipher_type)
    aes.decrypt
    aes.key = key
    aes.iv = iv if iv != nil
    aes.update(encrypted_data) + aes.final
  end

  # Encrypts a block of data given an encryption key and an
  # initialization vector (iv).  Keys, iv's, and the data returned
  # are all binary strings.  Cipher_type should be "AES-256-CBC",
  # "AES-256-ECB", or any of the cipher types supported by OpenSSL.
  # Pass nil for the iv if the encryption type doesn't use iv's (like
  # ECB).
  #:return: => String
  #:arg: data => String
  #:arg: key => String
  #:arg: iv => String
  #:arg: cipher_type => String
  def AESCrypt.encrypt(data, key, iv, cipher_type)
    aes = OpenSSL::Cipher::Cipher.new(cipher_type)
    aes.encrypt
    aes.key = key
    aes.iv = iv if iv != nil
    aes.update(data) + aes.final
  end
end
$grpc_developer_url = "open-api.akeywallet.com:80"
$grpc_tsv_url = "open-api-2sv.akeywallet.com:80"
$ak = "72b90c38e30cd418934753d873fb57344bc1b839"
$pk = OpenSSL::PKey::RSA.new File.read 'conf/public.pem'
$sk = OpenSSL::PKey::RSA.new File.read 'conf/private.pem'

def main
  stub_developer = Developer::DeveloperService::Stub.new($grpc_developer_url,:this_channel_is_insecure)
  # init session,get access key from your web console,CtyptoType select RSA or CURVE,curve25519 is a advanced cryto method
  message_session_request = Developer::SessionRequest.new(access_key: $ak,crypto_type: "RSA")
  message_session = stub_developer.init_session(message_session_request)
  p "session_id: #{message_session.session_id}"
  p "app_id: #{message_session.app_id}"
  p "crypto_type: #{message_session.crypto_type}"
  # save session secret key,first hex to byte string,then rsa decrypt it
  secret_key = message_session.secret_key.scan(/../).map { |x| x.hex }.pack('c*')
  origin_key = $sk.private_decrypt(secret_key)
  printf("origin_key hex is:#{origin_key.unpack('H*').first}")

  # get balance from your bank, need sign protobuf
  # marshal session
  message_session_buff = Google::Protobuf::encode(message_session)
  p "session buff:#{message_session_buff}"

  # sign session then binary to hex
  signed_session = $sk.sign("sha1",message_session_buff).unpack('H*').first

  p "signed session: #{signed_session}"
  balance = stub_developer.get_balance(message_session,{metadata:{sign:signed_session}})
  p "balance: #{balance.bank_currencies[0].coin_type}"
  p "balance: #{balance.bank_currencies[0].balance}"

  # then send currency to your address

  # sign  users to your app
  stub_user = App::User::AppUserService::Stub.new($grpc_developer_url,:this_channel_is_insecure)
  app_user1 = App::User::AppUser.new(session_id:message_session.session_id,app_user_id:"ABC5",crypto_type:"RSA")
  app_user2 = App::User::AppUser.new(session_id:message_session.session_id,app_user_id:"DEF5",crypto_type:"RSA")
  message_app_user1 = stub_user.sign_in_app_user(app_user1)
  message_app_user2 = stub_user.sign_in_app_user(app_user2)
  p "appuser1:#{message_app_user1.id}"
  p "appuser1's app_id:#{message_app_user1.app_id}"
  p "appuser2:#{message_app_user2.id}"
  # first hex to byte string ,then dcrypt aes
  user1_sk_buff = AESCrypt.decrypt(message_app_user1.secret_private_key.scan(/../).map { |x| x.hex }.pack('c*'),origin_key,origin_key,"AES-128-CBC")
  p "user1_sk_buff is: #{user1_sk_buff.force_encoding("UTF-8")}"
  user1_sk = OpenSSL::PKey::RSA.new(user1_sk_buff)
  p "user1 sk is private:#{user1_sk.private?}"
  # marshal user1
  user1_buff = Google::Protobuf::encode(app_user1)
  p "user1 buff is: #{user1_buff}"
  signed_user1 = user1_sk.sign("sha1",user1_buff).unpack('H*').first
  # get user1 balances
  user1_balance = stub_user.get_balance(app_user1,{metadata:{sign:signed_user1}})
  p "user1 balance: #{user1_balance.app_user_id}"
  p "user1 balance: #{user1_balance.appUserCurrencies[0].balance}"
  p "user1 balance: #{user1_balance.appUserCurrencies[0].coin_type}"

  # find a bank currency
  bwc = nil
  balance.bank_currencies.each do | c |
      if c.coin_type == "BWC"
          bwc = c
          break
      end
  end
  p "bwc is #{bwc.coin_type}:#{bwc.main_net}"
  # send bank tx
  bank_tx1 = Developer::BankTx.new(session_id:message_session.session_id,
                                   app_user_id:app_user1.app_user_id,
                                   currency_id:bwc.id,
                                   tx_type:"out",
                                   coin_type:bwc.coin_type,
                                   main_net:bwc.main_net,
                                   crypto_type:"RSA",
                                   amount:"100")
  # marshal tx
  tx1_buff = Google::Protobuf::encode(bank_tx1)
  # sign tx
  signed_tx1 = $sk.sign("sha1",tx1_buff).unpack('H*').first
  # create tx
  message_tx1 = stub_developer.send_tx_from_bank(bank_tx1,{metadata:{sign:signed_tx1}})
  p "tx1 received:#{message_tx1.id}"

  #send coin from one user to another user,before send,user need bind  2step verify to app
  # get user1 balance
  user1_balance = stub_user.get_balance(app_user1,{metadata:{sign:signed_user1}})
  p "user1 balance: #{user1_balance.app_user_id}"
  p "user1 balance: #{user1_balance.appUserCurrencies[0].balance}"
  p "user1 balance: #{user1_balance.appUserCurrencies[0].coin_type}"
  user1_bwc = nil
  user1_balance.appUserCurrencies.each do | c |
      if c.coin_type == "BWC"
          user1_bwc = c
          break
      end
  end
  p "user1 bwc currency is: #{user1_bwc.id}"
  app_user_tx1 = App::User::AppUserTx.new(app_id:message_app_user2.app_id,crypto_type:"RSA",amount:"10",from_user_currency_id:user1_bwc.id,to_user_id:message_app_user2.app_user_id)
  # marshal user tx1
  user_tx1_buff = Google::Protobuf::encode(app_user_tx1)
  # sign user tx1
  signed_user_tx1 = user1_sk.sign("sha1",user_tx1_buff).unpack('H*').first
  message_user_tx1 = stub_user.create_tx(app_user_tx1,{metadata:{sign:signed_user_tx1}})
  p "user tx1: #{message_user_tx1.id}"
  p "user tx1: #{message_user_tx1.amount}"
  p "user tx1: #{message_user_tx1.coin_type}"
  p "user tx1: #{message_user_tx1.main_net}"

  # bind 2step verify
  stub_tsv = Tsv::TSVService::Stub.new($grpc_tsv_url,:this_channel_is_insecure)
  # key request
  # key_request = Tsv::KeyRequest.new(session_id:message_session.session_id,label:app_user1.app_user_id)
  # message_key = stub_tsv.generate_key(key_request)
  # p "generate 2sv key:#{message_key.qr_url}"

  # bind key requrest
  #bind_key_request = Tsv::BindKeyRequest.new(key:"AVZJM4LWCCTXSACW",provider_name:"akey.io",verify_code:"479923",label:app_user1.app_user_id)
  #bind_status = stup_tsv.bind_key(bind_key_request)

  # bind 2step key at user device
  # then sendTx with 2step verify code
  # change code to real code
  send_tx1 = App::User::AppUserTx.new(id:message_user_tx1.id,session_id:message_session.session_id,code:"866642",crypto_type:"RSA")
  # marshal tx
  send_tx1_buff = Google::Protobuf::encode(send_tx1)
  # sign user tx1
  signed_send_tx1 = user1_sk.sign("sha1",send_tx1_buff).unpack('H*').first
  message_send_tx1 = stub_user.send_tx(send_tx1,{metadata:{sign:signed_send_tx1}})
  p "send tx1 amount: #{message_send_tx1.amount}"
end

main

