# Gataca DID Method Specification
V0.3, Gataca

## Introduction
GATACA is a blockchain-based digital identity platform that provides Identity and Access Management (IAM) services using decentralized identifiers and verifiable credentials.

Gataca uses the specifications provided by the W3C in order to maximize interoperability between the different agents globally. [DID documents](https://www.w3.org/TR/did-core/) and [Verifiable Credentials documents](https://www.w3.org/TR/vc-data-model/). 

## Infrastructure
Gataca’s platform is based on a mobile identity portfolio, a set of APIs, and controllers for multiple blockchain networks. 

Gataca is agnostic to the blockchain network. We adapt our infrastructure to the third party’s preferred ledger. Gataca currently supports the public network Ethereum and private networks based on Hyperledger Fabric, Hyperledger Besu or Quorum. Other networks may be added as requested.

This document provides the DID method specs for how our DID schema is implemented on the Ethereum network.

The simple structure links an object to a DID with states and public keys. Users do not need privileges to read the information on the blockchain but do need them to write. Gataca is the unique user that can modify the smart contract.

## Specification
### Method DID Format
Gataca's DID scheme is defined as follows:
```
did                = "did:" method-name ":" method-specific-id
method-name        = "gatc"
method-specific-id = ^[a-km-zA-HJ-NP-Z1-9]{32}$
```
- **did**: DID scheme type
- **method-name**: Gataca uses *gatc* as its method name
- **identifier**: 32 characters using base58

Example of a Gataca DID - did:gatc:8fa938c89f82424a1r3b2fE5614fA2293

### Operations
Gataca provides CRUD operations to manage DIDs. A user can create a new DID with an associated key pair, read DIDs previously generated by themselves or created by other users, modify a DIDs associated status or revoke it permanently (logical deletion).

CRUD operations may be triggered:
1. Via the Gataca Identity app (Identity wallet). The wallet generates a key pair associated with that DID, and it sends the DID and public key to Gataca Backbone (middleware to communicate to DLTs). The private key is stored in a mobile device’s secure storage.
2. Via [Gataca Connect](https://docs.gatacaid.com/connect/) and Gataca Credential APIs. 

These services call Gataca Backbone to fetch the information from the appropriate DLT network. 

#### Create
A Create operation creates a DID and its associated key pair.
#### Read
A Read operation fetches the public key and status associated to a DID. Read operations are public.
#### Update
An Update operation modifies the data associated to a DID. A DID's associated data may be modified only by its owner.
#### Suspend
A Suspend operation temporarily blocks Read operations on a specific DID. DIDs may be suspended and reactivated only by their owner.
#### Revoke
A Revoke operation permanently blocks Read, Update or Suspend operations on a specific DID. DIDs may be revoked only by their owner.

## Security Considerations
### Crypto algorithms
Gataca's technology is capable of using different crypto algorithms. By default, it uses ED25519, but it can be adapted to alternative algorithms according to client needs. 
### AuthNFactors
Gataca uses security mechanisms to ensure the device and wallet owner are the same. AuthNFactors use the same structure as verifiable credentials, but its mechanisms ensure wallet and device ownership consistency by using challenges to verify authentication rather than storing information about the user.
### Consents - Anonymous relationship
Gataca uses Consent structures to provide evidence that the user and the verifier agree to share information between them. 

## Privacy Considerations
### DID-Relationships
With Gataca technology, users have a root DID that acts as a global identifier, but they have multiple DIDs related to each verifier relationship. This architecture provides anonymity to the user, protecting his root DID from inference processes. Additionally, this enables users to easily revoke access to their data.
### GDPR Compliance
Gataca is GDPR compliant. Gataca only stores public information in DLTs (DPKI). Gataca uses DLTs to provide integrity and resiliency and to act as a trusted third party, maintaining user privacy.
