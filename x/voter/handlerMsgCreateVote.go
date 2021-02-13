package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/drcyph3r/voter/x/voter/keeper"
	"github.com/drcyph3r/voter/x/voter/types"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateVote) (*sdk.Result, error) {
	k.CreateVote(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
