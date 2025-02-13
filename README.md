<img width="1280" alt="Code Protobuf API" src="https://github.com/code-payments/code-protobuf-api/assets/5760385/e43106cd-4c50-4dfa-954f-a936290b8b86">

# Code Protobuf API

[![Release](https://img.shields.io/github/v/release/code-payments/code-protobuf-api.svg)](https://github.com/code-payments/code-protobuf-api/releases/latest)
[![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/code-payments/code-protobuf-api/blob/main/LICENSE.md)

The APIs and models for communication between Code clients and server.

## Services

- [Account](https://github.com/code-payments/code-protobuf-api/blob/main/proto/account/v1/account_service.proto)
- [Badge](https://github.com/code-payments/code-protobuf-api/blob/main/proto/badge/v1/badge_service.proto)
- [Chat](https://github.com/code-payments/code-protobuf-api/blob/main/proto/chat/v1/chat_service.proto)
- [Contact List](https://github.com/code-payments/code-protobuf-api/blob/main/proto/contact/v1/contact_list_service.proto)
- [Currency](https://github.com/code-payments/code-protobuf-api/blob/main/proto/currency/v1/currency_service.proto)
- [Device](https://github.com/code-payments/code-protobuf-api/blob/main/proto/device/v1/device_service.proto)
- [Invite](https://github.com/code-payments/code-protobuf-api/blob/main/proto/invite/v2/invite_service.proto)
- [Messaging](https://github.com/code-payments/code-protobuf-api/blob/main/proto/messaging/v1/messaging_service.proto)
- [Micro Payment](https://github.com/code-payments/code-protobuf-api/blob/main/proto/micropayment/v1/micro_payment_service.proto)
- [Phone Verification](https://github.com/code-payments/code-protobuf-api/blob/main/proto/phone/v1/phone_verification_service.proto)
- [Push](https://github.com/code-payments/code-protobuf-api/blob/main/proto/push/v1/push_service.proto)
- [Transaction](https://github.com/code-payments/code-protobuf-api/blob/main/proto/transaction/v2/transaction_service.proto)
- [User Identity](https://github.com/code-payments/code-protobuf-api/blob/main/proto/user/v1/identity_service.proto)

## Code Generation

Generated code can be found under the `generated/` directory. The following languages are directly supported:
- Go
- TypeScript/JavaScript (via [Protobuf-ES](https://github.com/bufbuild/protobuf-es))
- Php
- Python

To generate all code, run:

```bash
make
```

Or to generate a specific language, run:

```bash
make {go, protobuf-es, php, go}
```

## Getting Help

If you have any questions or need help integrating Code into your website or application, please reach out to us on [Discord](https://discord.gg/T8Tpj8DBFp) or [Twitter](https://twitter.com/getcode).

##  Contributing

For now the best way to contribute is to share feedback on [Discord](https://discord.gg/T8Tpj8DBFp). This will evolve as we continue to build out the platform and open up more ways to contribute.
