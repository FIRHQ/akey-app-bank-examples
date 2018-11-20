# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "app.user.AppUser" do
    optional :id, :string, 1
    optional :app_id, :string, 2
    optional :session_id, :string, 3
    optional :secret_private_key, :string, 4
    optional :secret_public_key, :string, 5
    optional :nonce, :string, 6
    optional :app_user_id, :string, 7
    optional :crypto_type, :string, 8
  end
  add_message "app.user.AppUserCurrency" do
    optional :id, :string, 1
    optional :sys_user_id, :string, 2
    optional :bank_id, :string, 3
    optional :coin_type, :string, 4
    optional :main_net, :string, 5
    optional :balance, :string, 6
    optional :address_balance, :string, 7
    optional :address, :string, 8
    optional :decimals, :uint32, 9
  end
  add_message "app.user.AppUserTx" do
    optional :id, :string, 1
    optional :app_id, :string, 2
    optional :from_user_id, :string, 3
    optional :to_user_id, :string, 4
    optional :from_user_currency_id, :string, 5
    optional :amount, :string, 6
    optional :bank_id, :string, 7
    optional :fee, :uint64, 8
    optional :coin_type, :string, 9
    optional :address, :string, 10
    optional :main_net, :string, 11
    optional :trade_at, :uint64, 12
    optional :session_id, :string, 13
    optional :decimals, :uint32, 14
    optional :tx_status, :uint32, 15
    optional :code, :string, 16
    optional :crypto_type, :string, 17
  end
  add_message "app.user.AppUserBalance" do
    optional :app_user_id, :string, 1
    optional :app_id, :string, 2
    repeated :app_user_currencies, :message, 3, "app.user.AppUserCurrency"
  end
  add_message "app.user.AppUserWithdraw" do
    optional :id, :string, 1
    optional :addr, :string, 2
    optional :currency_id, :string, 3
    optional :amount, :string, 4
    optional :tx_id, :string, 5
    optional :app_user_id, :string, 6
    optional :crypto_type, :string, 7
  end
  add_message "app.user.UserAddress" do
    optional :id, :string, 1
    optional :address, :string, 2
    optional :owner, :string, 3
    optional :main_net, :string, 4
    optional :crypto_type, :string, 5
  end
end

module App
  module User
    AppUser = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.AppUser").msgclass
    AppUserCurrency = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.AppUserCurrency").msgclass
    AppUserTx = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.AppUserTx").msgclass
    AppUserBalance = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.AppUserBalance").msgclass
    AppUserWithdraw = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.AppUserWithdraw").msgclass
    UserAddress = Google::Protobuf::DescriptorPool.generated_pool.lookup("app.user.UserAddress").msgclass
  end
end
