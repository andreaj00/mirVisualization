// Code generated by Mir codegen. DO NOT EDIT.

package testerpbevents

import (
	types1 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types2 "github.com/filecoin-project/mir/pkg/pb/testerpb/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

func Tester(destModule types.ModuleID) *types1.Event {
	return &types1.Event{
		DestModule: destModule,
		Type: &types1.Event_Tester{
			Tester: &types2.Tester{},
		},
	}
}
