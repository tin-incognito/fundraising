package simapp

import (
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/tendermint/fundraising/app"
)

// defaultConsensusParams defines the default Tendermint consensus params used in testing.
var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// New creates application instance with in-memory database and disabled logging.
func New(dir string) *app.App {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	encodingCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	a := app.New(
		logger, db, nil, true, map[int64]bool{}, dir, 0, encodingCfg, simapp.EmptyAppOptions{},
	)

	// InitChain updates deliverState which is required when app.NewContext is called
	a.InitChain(
		abci.RequestInitChain{
			ConsensusParams: defaultConsensusParams,
			AppStateBytes:   []byte("{}"),
		},
	)

	fundraisingApp := a.(*app.App)

	return fundraisingApp
}
