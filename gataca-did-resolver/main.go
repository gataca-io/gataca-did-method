package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gatacaid/gataca-backbone/apps/did"
	"github.com/gatacaid/gataca-did-method/gataca-did-resolver/contract"
	"github.com/gatacaid/gataca-did-method/gataca-did-resolver/models"
	"github.com/labstack/gommon/log"
	"math/big"
)

const (
	//Blockchain node URL (Infura: https://rinkeby.infura.io, Ganache: http://127.0.0.1:8545
	rawURL = "https://rinkeby.infura.io/v3/89cdc30a5a8c440abfa323d028e0c70f"
	contractAddress = "0x3843656e2Ba58f1Bd26f7a43d977F01516f7524d"
	hexKey = "hexKey"
	contextW3C = "https://www.w3.org/ns/did/v1"
)
var contextW3CSecurity = []string{"https://www.w3.org/ns/did/v1", "https://w3id.org/security/v1"}

type resolver struct {
	HexKey string
	Blockchain *ethclient.Client
	SCAddress *common.Address
	Contract *contract.GatacaID
	Auth *bind.TransactOpts
}

func deploy(key *ecdsa.PrivateKey, blockchain *ethclient.Client) (*contract.GatacaID, string, error) {

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("error casting public key to ECDSA")
		return nil, "", nil
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error(err)
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     			// in wei
	auth.GasLimit = uint64(3000000) 			// in units
	auth.GasPrice = gasPrice

	address, _, c, err := contract.DeployGatacaID(
		auth,
		blockchain,
	)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	return c, address.String(), nil
}

func loadSmartContract(addressS string, blockchain bind.ContractBackend) (*contract.GatacaID, error) {
	address := common.HexToAddress(addressS)

	c, err := contract.NewGatacaID(address, blockchain)
	if err != nil {
		log.Errorf("Unable to bind to deployed instance of contract: %s", address)
		return nil, err
	}

	return c, nil
}

func (r *resolver) GetDID(didS string) (*models.DIDDocument, error) {

	entity, err := r.Contract.GetEntity(nil, didS)
	if err != nil {
		log.Debugf("smart contract error. it can be did does not exist, it's suspended or it's revoked: %s", didS)
		return nil, err
	}

	publicKeyO := &models.PublicKeyEd25519{
		Context: contextW3CSecurity,
		Id: didS + "#keys-1",
		Type: did.Ed25519KeyType,
		Controller: didS,
		Key: entity.PubKey,
	}

	var publicKeyList []*models.PublicKeyEd25519
	publicKeyList = append(publicKeyList, publicKeyO)

	var authorizationList []string
	authorizationList = append(authorizationList, didS + "#keys-1")

	didO := &models.DIDDocument{
		Context: contextW3C,
		Id: didS,
		PublicKey: publicKeyList,
		Authentication: authorizationList,
	}

	return didO, err
}

func main() {
	log.Info("Executing Gataca Resolver")

	didS := "did:gatc:50e05cff9d34cc409f2095d4006698167821f439c263ff4f9693274df71ce7e2"

	blockchain, err := ethclient.Dial(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	defer blockchain.Close()

	c, err := loadSmartContract(contractAddress, blockchain)
	if err != nil {
		log.Fatalf("error loading smart contract. %v", err)
	}

	r := &resolver{
		HexKey: hexKey,
		Blockchain: blockchain,
		Contract: c,
		Auth: nil,
	}

	// Build W3C document
	didD, err := r.GetDID(didS)
	didB, err := json.MarshalIndent(didD, "", "    ")

	fmt.Printf("DID document:\n %s", didB)
}
