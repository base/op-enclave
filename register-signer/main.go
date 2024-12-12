package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/base-org/op-enclave/bindings"
	"github.com/base-org/op-enclave/op-withdrawer/withdrawals"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hf/nitrite"
)

func main() {
	var attestationHex string
	var rpcUrl string
	var privateKeyHex string
	var configAddress string
	flag.StringVar(&attestationHex, "attestation", "", "attestation hex")
	flag.StringVar(&rpcUrl, "rpc", "https://sepolia.base.org", "rpc url")
	flag.StringVar(&privateKeyHex, "private-key", "", "private key")
	flag.StringVar(&configAddress, "address", "", "address of the SystemConfigGlobal proxy contract")
	flag.Parse()

	if attestationHex == "" || privateKeyHex == "" || configAddress == "" {
		flag.Usage()
		os.Exit(1)
	}

	attestation, err := hexutil.Decode(attestationHex)
	if err != nil {
		panic(err)
	}
	privateKey, err := hexutil.Decode(privateKeyHex)
	if err != nil {
		panic(err)
	}

	res, err := nitrite.Verify(attestation, nitrite.VerifyOptions{})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		panic(err)
	}

	key, err := crypto.ToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	signer := types.LatestSignerForChainID(chainId)
	auth := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(key.PublicKey),
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, signer, key)
		},
	}

	systemConfigGlobal, err := bindings.NewSystemConfigGlobal(common.HexToAddress(configAddress), client)
	if err != nil {
		panic(err)
	}

	pub, err := crypto.UnmarshalPubkey(res.Document.PublicKey)
	if err != nil {
		panic(err)
	}
	signerAddr := crypto.PubkeyToAddress(*pub)
	validSigner, err := systemConfigGlobal.ValidSigners(&bind.CallOpts{}, signerAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public key: %s\n", hexutil.Encode(res.Document.PublicKey))
	if validSigner {
		fmt.Printf("Signer already registered: %s\n", signerAddr.String())
		return
	} else {
		fmt.Printf("Registering signer: %s\n", signerAddr.String())
	}

	certManagerAddr, err := systemConfigGlobal.CertManager(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	certManager, err := bindings.NewCertManager(certManagerAddr, client)
	if err != nil {
		panic(err)
	}

	verifyCert := func(cert []byte, ca bool, parentCertHash common.Hash) common.Hash {
		certHash := crypto.Keccak256Hash(cert)
		verified, err := certManager.Verified(&bind.CallOpts{}, certHash)
		if err != nil {
			panic(err)
		}
		if len(verified) == 0 {
			tx, err := certManager.VerifyCert(auth, cert, ca, parentCertHash)
			if err != nil {
				panic(err)
			}
			receipt, err := withdrawals.WaitForReceipt(ctx, client, tx.Hash(), 2*time.Second)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Verified cert: %s, tx: %s\n", certHash.String(), receipt.TxHash.String())
		} else {
			fmt.Printf("Cert already verified: %s\n", certHash.String())
		}
		return certHash
	}

	parentCertHash := crypto.Keccak256Hash(res.Document.CABundle[0])
	for i := 0; i < len(res.Document.CABundle); i++ {
		cert := res.Document.CABundle[i]
		parentCertHash = verifyCert(cert, true, parentCertHash)
	}
	verifyCert(res.Document.Certificate, false, parentCertHash)

	pcr0Hash := crypto.Keccak256Hash(res.Document.PCRs[0])
	valid, err := systemConfigGlobal.ValidPCR0s(&bind.CallOpts{}, pcr0Hash)
	if err != nil {
		panic(err)
	}
	if !valid {
		tx, err := systemConfigGlobal.RegisterPCR0(auth, res.Document.PCRs[0])
		if err != nil {
			panic(err)
		}
		receipt, err := withdrawals.WaitForReceipt(ctx, client, tx.Hash(), 2*time.Second)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Registered PCR0, tx: %s\n", receipt.TxHash.String())
	} else {
		fmt.Printf("PCR0 already registered: %s\n", pcr0Hash.String())
	}

	tx, err := systemConfigGlobal.RegisterSigner(auth, res.COSESign1, res.Signature)
	if err != nil {
		panic(err)
	}
	receipt, err := withdrawals.WaitForReceipt(ctx, client, tx.Hash(), 2*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Registered signer, tx: %s\n", receipt.TxHash.String())
}
