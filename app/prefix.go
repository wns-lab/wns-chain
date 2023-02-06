package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcfg "github.com/evmos/ethermint/cmd/config"
	cmdcfg "github.com/wns-lab/wns-chain/cmd/wnsd/config"
)

func SetConfig() {
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	ethcfg.SetBip44CoinType(config)
	// Make sure address is compatible with ethereum
	config.SetAddressVerifier(VerifyAddressFormat)
	config.Seal()
}
