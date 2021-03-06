# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: user.proto for package 'app.user'

require 'grpc'
require 'user_pb'

module App
  module User
    module AppUserService
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'app.user.AppUserService'

        # sign in user
        rpc :signInAppUser, AppUser, AppUser
        # create tx
        rpc :createTx, AppUserTx, AppUserTx
        # send coin
        rpc :sendTx, AppUserTx, AppUserTx
        # get balances
        rpc :getBalance, AppUser, AppUserBalance
        # withdraw
        rpc :withdraw, AppUserWithdraw, AppUserWithdraw
        # create user currency address
        rpc :createUserAddress, UserAddress, UserAddress
      end

      Stub = Service.rpc_stub_class
    end
  end
end
