// Copyright 2021 The psch Authors
// This file is part of the psch library.
//
// The psch library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The psch library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the psch library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/petercastron/PSCH/common"
	"github.com/petercastron/PSCH/rpc"
	"golang.org/x/crypto/sha3"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	MainnetGenesisHash2 = common.HexToHash("0xc01721172a4c69d1374e808d43d32dcb8b279f90fa252cd4ad8830558fb41adb")
	TestnetGenesisHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
}

// CheckpointOracles associates each known checkpoint oracles with the genesis hash of
// the chain it belongs to.
var CheckpointOracles = map[common.Hash]*CheckpointOracleConfig{
	MainnetGenesisHash: MainnetCheckpointOracle,
	TestnetGenesisHash: TestnetCheckpointOracle,
}

type BanlanceAlloc struct {
	Addr    string
	Balance string
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(188),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("ux0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    nil,
		BerlinBlock:         nil,
		LondonBlock:nil,
		EWASMBlock:nil,
		CatalystBlock:nil,
		Alien: &AlienConfig{
			Period:           10,
			Epoch:            210,
			MaxSignerCount:   21,
			TrantorBlock:     nil,
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(10), big.NewInt(1e+18)),
			GenesisTimestamp: 1641816014,
			SelfVoteSigners: []common.UnprefixedAddress{
				common.UnprefixedAddress(common.HexToAddress("82de6bd4b822c5af6110de34a133980c456708e0")),
				common.UnprefixedAddress(common.HexToAddress("feac212688fdc4d7f0f5af8caa02f981d55a7cf4")),
				common.UnprefixedAddress(common.HexToAddress("31f440fc8dd98bdbdb72ebd8a14a469439fc3433")),
			},
			SideChain:   false,
			MCRPCClient: nil,
			PBFTEnable:  false,
		},
	}

	MainnetBalanceAlloc = []BanlanceAlloc{
		BanlanceAlloc{Addr: "e6A6f49817Bfd9FDF4D0903D82f29A6Aa4D9E3bD", Balance: "5140000000000000000000000"},
		BanlanceAlloc{Addr: "15F4E608027052395d5c5c422F27937781563d7d", Balance: "7350000000000000000000000"},
		BanlanceAlloc{Addr: "bcfF5e8B6aa894459Ec3F59e0f29f032133962c9", Balance: "10500000000000000000000000"},
		BanlanceAlloc{Addr: "72eb423E8C12fAecBcAf53Cb982Dd9b323488bac", Balance: "8400000000000000000000000"},
		BanlanceAlloc{Addr: "A573d8C28a709ACbA1EB10e605694482a92c3593", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "7d51170F140c47e547664eAD4D1185ef864ba689", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "d9aac9B61571B9bE5717A275f41d772E9bfc745C", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "81ae1B55bb078102c965Bb8a4FaF48EcC4380F55", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "88eA42c6A2D9B23C52534b0e1eEcf3DEa0c6De76", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "f4955c8A120B1CDf3BFD7c6Dc43837Dd65360F01", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "D691EA3Fd19437BbD27A590BFcA3c435c9C07C38", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "208bc40a411786f9cE7b4A3D1f8424a4f59406E8", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "31f440fc8dd98bdbdb72ebd8a14a469439fc3433", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "82de6bd4b822c5af6110de34a133980c456708e0", Balance: "10000000000000000000000"},
		BanlanceAlloc{Addr: "feac212688fdc4d7f0f5af8caa02f981d55a7cf4", Balance: "10000000000000000000000"},
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 384,
		SectionHead:  common.HexToHash("0xb583a0ead70324849c4caf923476de3645c0d2f707c86221ec8e40078bdd6884"),
		CHTRoot:      common.HexToHash("0x6ecc993baad0c9f77fe9c4c13b89360112e5a0accae4d8502470b911211618b7"),
		BloomRoot:    common.HexToHash("0x66a30d8885c19921711704921de7b4bcbd1b49191b197ee79e34dafeed9a04d9"),
	}

	// MainnetCheckpointOracle contains a set of configs for the main network oracle.
	MainnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"),
		Signers: []common.Address{
			common.HexToAddress("0x1b2C260efc720BE89101890E4Db589b44E950527"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// RopstenChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(3),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("ux0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    nil,
		BerlinBlock:         nil,
		LondonBlock:nil,
		EWASMBlock:nil,
		CatalystBlock:nil,
		Alien: &AlienConfig{
			Period:           3,
			Epoch:            201600,
			MaxSignerCount:   21,
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(100), big.NewInt(1e+18)),
			TrantorBlock:     big.NewInt(0),
			GenesisTimestamp: 1554004800,
			SelfVoteSigners: []common.UnprefixedAddress{
				common.UnprefixedAddress(common.HexToAddress("t0be6865ffcbbe5f9746bef5c84b912f2ad9e52075")),
				common.UnprefixedAddress(common.HexToAddress("t04909b4e54395de9e313ad8a2254fe2dcda99e91c")),
				common.UnprefixedAddress(common.HexToAddress("t0a034350c8e80eb4d15ac62310657b29c711bb3d5")),
			},
		},
	}

	// RopstenTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 279,
		SectionHead:  common.HexToHash("0x4a4912848d4c06090097073357c10015d11c6f4544a0f93cbdd584701c3b7d58"),
		CHTRoot:      common.HexToHash("0x9053b7867ae921e80a4e2f5a4b15212e4af3d691ca712fb33dc150e9c6ea221c"),
		BloomRoot:    common.HexToHash("0x3dc04cb1be7ddc271f3f83469b47b76184a79d7209ef51d85b1539ea6d25a645"),
	}

	// RopstenCheckpointOracle contains a set of configs for the Ropsten test network oracle.
	TestnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xEF79475013f154E6A65b54cB2742867791bf0B84"),
		Signers: []common.Address{
			common.HexToAddress("0x32162F3581E88a5f62e8A61892B42C46E2c18f7b"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// SideChainConfig contains the chain parameters to run a node on the Ropsten test network.
	SideChainConfig = &ChainConfig{
		ChainID:             big.NewInt(8123),
		HomesteadBlock:      big.NewInt(1),
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: nil,
		Alien: &AlienConfig{
			Period:           5,
			Epoch:            201600,
			MaxSignerCount:   21,
			TrantorBlock:     big.NewInt(0),
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(100), big.NewInt(1e+18)),
			GenesisTimestamp: 1554004800,
			SelfVoteSigners:  []common.UnprefixedAddress{},
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the psch core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, new(EthashConfig), nil, nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the psch core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil}

	// AllAlienProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Alien consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllAlienProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, &AlienConfig{Period: 10, Epoch: 30000, MaxSignerCount: 21, MinVoterBalance: new(big.Int).Mul(big.NewInt(10000), big.NewInt(1000000000000000000)), GenesisTimestamp: 0, SelfVoteSigners: []common.UnprefixedAddress{}}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil, nil, nil, nil, &AlienConfig{Period: 5, Epoch: 30000, MaxSignerCount: 21, MinVoterBalance: new(big.Int).Mul(big.NewInt(10000), big.NewInt(1000000000000000000)), GenesisTimestamp: 0, SelfVoteSigners: []common.UnprefixedAddress{}}}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// HashEqual returns an indicator comparing the itself hash with given one.
func (c *TrustedCheckpoint) HashEqual(hash common.Hash) bool {
	if c.Empty() {
		return hash == common.Hash{}
	}
	return c.Hash() == hash
}

// Hash returns the hash of checkpoint's four key fields(index, sectionHead, chtRoot and bloomTrieRoot).
func (c *TrustedCheckpoint) Hash() common.Hash {
	var sectionIndex [8]byte
	binary.BigEndian.PutUint64(sectionIndex[:], c.SectionIndex)

	w := sha3.NewLegacyKeccak256()
	w.Write(sectionIndex[:])
	w.Write(c.SectionHead[:])
	w.Write(c.CHTRoot[:])
	w.Write(c.BloomRoot[:])

	var h common.Hash
	w.Sum(h[:0])
	return h
}

// Empty returns an indicator whether the checkpoint is regarded as empty.
func (c *TrustedCheckpoint) Empty() bool {
	return c.SectionHead == (common.Hash{}) || c.CHTRoot == (common.Hash{}) || c.BloomRoot == (common.Hash{})
}

// CheckpointOracleConfig represents a set of checkpoint contract(which acts as an oracle)
// config which used for light client checkpoint syncing.
type CheckpointOracleConfig struct {
	Address   common.Address   `json:"address"`
	Signers   []common.Address `json:"signers"`
	Threshold uint64           `json:"threshold"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	MuirGlacierBlock    *big.Int `json:"muirGlacierBlock,omitempty"`    // Eip-2384 (bomb delay) switch block (nil = no fork, 0 = already activated)
	BerlinBlock         *big.Int `json:"berlinBlock,omitempty"`         // Berlin switch block (nil = no fork, 0 = already on berlin)
	LondonBlock         *big.Int `json:"londonBlock,omitempty"`         // London switch block (nil = no fork, 0 = already on london)

	EWASMBlock    *big.Int `json:"ewasmBlock,omitempty"`    // EWASM switch block (nil = no fork, 0 = already activated)
	CatalystBlock *big.Int `json:"catalystBlock,omitempty"` // Catalyst switch block (nil = no fork, 0 = already on catalyst)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Alien  *AlienConfig  `json:"alien,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

type GenesisAccount struct {
	Balance string `json:"balance"`
}

// AlienLightConfig is the config for light node of alien
type AlienLightConfig struct {
	Alloc map[common.UnprefixedAddress]GenesisAccount `json:"alloc"`
}

// AlienConfig is the consensus engine configs for delegated-proof-of-stake based sealing.
type AlienConfig struct {
	Period           uint64                     `json:"period"`           // Number of seconds between blocks to enforce
	Epoch            uint64                     `json:"epoch"`            // Epoch length to reset votes and checkpoint
	MaxSignerCount   uint64                     `json:"maxSignersCount"`  // Max count of signers
	MinVoterBalance  *big.Int                   `json:"minVoterBalance"`  // Min voter balance to valid this vote
	GenesisTimestamp uint64                     `json:"genesisTimestamp"` // The LoopStartTime of first Block
	SelfVoteSigners  []common.UnprefixedAddress `json:"signers"`          // Signers vote by themselves to seal the block, make sure the signer accounts are pre-funded
	SideChain        bool                       `json:"sideChain"`        // If side chain or not
	MCRPCClient      *rpc.Client                // Main chain rpc client for side chain
	PBFTEnable       bool                       `json:"pbft"` //

	TrantorBlock  *big.Int          `json:"trantorBlock,omitempty"`  // Trantor switch block (nil = no fork)
	TerminusBlock *big.Int          `json:"terminusBlock,omitempty"` // Terminus switch block (nil = no fork)
	LightConfig   *AlienLightConfig `json:"lightConfig,omitempty"`
}

// String implements the stringer interface, returning the consensus engine details.
func (a *AlienConfig) String() string {
	return "alien"
}

// IsTrantor returns whether num is either equal to the Trantor block or greater.
func (a *AlienConfig) IsTrantor(num *big.Int) bool {
	return isForked(a.TrantorBlock, num)
}

// IsTerminus returns whether num is either equal to the Terminus block or greater.
func (a *AlienConfig) IsTerminus(num *big.Int) bool {
	return isForked(a.TerminusBlock, num)
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Alien != nil:
		engine = c.Alien
	case c.Clique != nil:
		engine = c.Clique
	case c.Ethash != nil:
		engine = c.Ethash
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v, Muir Glacier: %v, Berlin: %v, London: %v, Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.BerlinBlock,
		c.LondonBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsMuirGlacier returns whether num is either equal to the Muir Glacier (EIP-2384) fork block or greater.
func (c *ChainConfig) IsMuirGlacier(num *big.Int) bool {
	return isForked(c.MuirGlacierBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// IsBerlin returns whether num is either equal to the Berlin fork block or greater.
func (c *ChainConfig) IsBerlin(num *big.Int) bool {
	return isForked(c.BerlinBlock, num)
}

// IsLondon returns whether num is either equal to the London fork block or greater.
func (c *ChainConfig) IsLondon(num *big.Int) bool {
	return isForked(c.LondonBlock, num)
}

// IsCatalyst returns whether num is either equal to the Merge fork block or greater.
func (c *ChainConfig) IsCatalyst(num *big.Int) bool {
	return isForked(c.CatalystBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, psch isn't pluggable enough
// to guarantee that forks can be implemented in a different order than on official networks
func (c *ChainConfig) CheckConfigForkOrder() error {
	type fork struct {
		name     string
		block    *big.Int
		optional bool // if true, the fork may be nil and next fork is still allowed
	}
	var lastFork fork
	for _, cur := range []fork{
		{name: "homesteadBlock", block: c.HomesteadBlock},
		{name: "daoForkBlock", block: c.DAOForkBlock, optional: true},
		{name: "eip150Block", block: c.EIP150Block},
		{name: "eip155Block", block: c.EIP155Block},
		{name: "eip158Block", block: c.EIP158Block},
		{name: "byzantiumBlock", block: c.ByzantiumBlock},
		{name: "constantinopleBlock", block: c.ConstantinopleBlock},
		{name: "petersburgBlock", block: c.PetersburgBlock},
		{name: "istanbulBlock", block: c.IstanbulBlock},
		{name: "muirGlacierBlock", block: c.MuirGlacierBlock, optional: true},
		{name: "berlinBlock", block: c.BerlinBlock},
		{name: "londonBlock", block: c.LondonBlock},
	} {
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		// If it was optional and not set, then ignore it
		if !cur.optional || cur.block != nil {
			lastFork = cur
		}
	}
	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		// the only case where we allow Petersburg to be set in the past is if it is equal to Constantinople
		// mainly to satisfy fork ordering requirements which state that Petersburg fork be set if Constantinople fork is set
		if isForkIncompatible(c.ConstantinopleBlock, newcfg.PetersburgBlock, head) {
			return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
		}
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.MuirGlacierBlock, newcfg.MuirGlacierBlock, head) {
		return newCompatError("Muir Glacier fork block", c.MuirGlacierBlock, newcfg.MuirGlacierBlock)
	}
	if isForkIncompatible(c.BerlinBlock, newcfg.BerlinBlock, head) {
		return newCompatError("Berlin fork block", c.BerlinBlock, newcfg.BerlinBlock)
	}
	if isForkIncompatible(c.LondonBlock, newcfg.LondonBlock, head) {
		return newCompatError("London fork block", c.LondonBlock, newcfg.LondonBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsBerlin, IsLondon, IsCatalyst                          bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
		IsBerlin:         c.IsBerlin(num),
		IsLondon:         c.IsLondon(num),
		IsCatalyst:       c.IsCatalyst(num),
	}
}
