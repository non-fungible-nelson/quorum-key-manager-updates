package stores

import (
	"context"
	"math/big"

	quorumtypes "github.com/consensys/quorum/core/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core"

	"github.com/consensys/quorum-key-manager/pkg/ethereum"
	"github.com/consensys/quorum-key-manager/src/stores/entities"
)

//go:generate mockgen -source=eth1.go -destination=mock/eth1.go -package=mock

type Eth1Store interface {
	// Info returns store information
	Info(context.Context) (*entities.StoreInfo, error)

	// Create creates an Ethereum account
	Create(ctx context.Context, id string, attr *entities.Attributes) (*entities.ETH1Account, error)

	// Import imports an externally created Ethereum account
	Import(ctx context.Context, id string, privKey []byte, attr *entities.Attributes) (*entities.ETH1Account, error)

	// Get gets an Ethereum account
	Get(ctx context.Context, addr common.Address) (*entities.ETH1Account, error)

	// List lists all Ethereum account addresses
	List(ctx context.Context) ([]common.Address, error)

	// Update updates Ethereum account attributes
	Update(ctx context.Context, addr common.Address, attr *entities.Attributes) (*entities.ETH1Account, error)

	// Delete deletes an account temporarily, by using Restore the account can be restored
	Delete(ctx context.Context, addr common.Address) error

	// GetDeleted Gets a deleted Ethereum accounts
	GetDeleted(ctx context.Context, addr common.Address) (*entities.ETH1Account, error)

	// ListDeleted lists all deleted Ethereum accounts
	ListDeleted(ctx context.Context) ([]common.Address, error)

	// Restore restores a previously deleted Ethereum account
	Restore(ctx context.Context, addr common.Address) error

	// Destroy destroys (purges) an Ethereum account permanently
	Destroy(ctx context.Context, addr common.Address) error

	// Sign signs a payload using the specified Ethereum account
	Sign(ctx context.Context, addr common.Address, data []byte) ([]byte, error)

	// SignHash signs a hash using the specified Ethereum account
	SignHash(ctx context.Context, addr common.Address, data []byte) ([]byte, error)

	// SignTypedData signs EIP-712 formatted data using the specified Ethereum account
	SignTypedData(ctx context.Context, addr common.Address, typedData *core.TypedData) ([]byte, error)

	// SignTransaction signs a public Ethereum transaction
	SignTransaction(ctx context.Context, addr common.Address, chainID *big.Int, tx *types.Transaction) ([]byte, error)

	// SignEEA signs an EEA transaction
	SignEEA(ctx context.Context, addr common.Address, chainID *big.Int, tx *types.Transaction, args *ethereum.PrivateArgs) ([]byte, error)

	// SignPrivate signs a Quorum private transaction
	SignPrivate(ctx context.Context, addr common.Address, tx *quorumtypes.Transaction) ([]byte, error)

	// ECRecover returns the Ethereum address from a signature and data
	ECRecover(ctx context.Context, data, sig []byte) (common.Address, error)

	// Verify verifies that a signature belongs to a given address
	Verify(ctx context.Context, addr common.Address, data, sig []byte) error

	// VerifyTypedData verifies that a typed data signature belongs to a given address
	VerifyTypedData(ctx context.Context, addr common.Address, typedData *core.TypedData, sig []byte) error

	// Encrypt encrypts any arbitrary data using a specified account
	Encrypt(ctx context.Context, addr common.Address, data []byte) ([]byte, error)

	// Decrypt decrypts a single block of encrypted data.
	Decrypt(ctx context.Context, addr common.Address, data []byte) ([]byte, error)
}