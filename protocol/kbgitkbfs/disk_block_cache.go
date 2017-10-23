// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: kbgitkbfs-avdl/disk_block_cache.avdl

package kbgitkbfs1

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PrefetchStatus int

const (
	PrefetchStatus_NO_PREFETCH        PrefetchStatus = 0
	PrefetchStatus_TRIGGERED_PREFETCH PrefetchStatus = 1
	PrefetchStatus_FINISHED_PREFETCH  PrefetchStatus = 2
)

var PrefetchStatusMap = map[string]PrefetchStatus{
	"NO_PREFETCH":        0,
	"TRIGGERED_PREFETCH": 1,
	"FINISHED_PREFETCH":  2,
}

var PrefetchStatusRevMap = map[PrefetchStatus]string{
	0: "NO_PREFETCH",
	1: "TRIGGERED_PREFETCH",
	2: "FINISHED_PREFETCH",
}

// GetCachedBlockRes is the response from GetBlock.
type GetBlockRes struct {
	Buf            []byte         `codec:"buf" json:"buf"`
	ServerHalf     string         `codec:"serverHalf" json:"serverHalf"`
	PrefetchStatus PrefetchStatus `codec:"prefetchStatus" json:"prefetchStatus"`
}

// DeleteBlocksRes is the response from DeleteBlocks.
type DeleteBlocksRes struct {
	NumRemoved  int   `codec:"numRemoved" json:"numRemoved"`
	SizeRemoved int64 `codec:"sizeRemoved" json:"sizeRemoved"`
}

type GetBlockArg struct {
	TlfID   keybase1.TLFID `codec:"tlfID" json:"tlfID"`
	BlockID string         `codec:"blockID" json:"blockID"`
}

type PutBlockArg struct {
	TlfID      keybase1.TLFID `codec:"tlfID" json:"tlfID"`
	BlockID    string         `codec:"blockID" json:"blockID"`
	Buf        []byte         `codec:"buf" json:"buf"`
	ServerHalf string         `codec:"serverHalf" json:"serverHalf"`
}

type DeleteBlocksArg struct {
	BlockIDs []string `codec:"blockIDs" json:"blockIDs"`
}

type UpdateBlockMetadataArg struct {
	BlockID        string         `codec:"blockID" json:"blockID"`
	PrefetchStatus PrefetchStatus `codec:"prefetchStatus" json:"prefetchStatus"`
}

// DiskBlockCacheInterface specifies how to access a disk cache remotely.
type DiskBlockCacheInterface interface {
	// GetBlock gets a block from the disk cache.
	GetBlock(context.Context, GetBlockArg) (GetBlockRes, error)
	// PutBlock puts a block into the disk cache.
	PutBlock(context.Context, PutBlockArg) error
	// DeleteBlocks deletes a set of blocks from the disk cache.
	DeleteBlocks(context.Context, []string) (DeleteBlocksRes, error)
	// UpdateBlockMetadata updates the metadata for a block in the disk cache.
	UpdateBlockMetadata(context.Context, UpdateBlockMetadataArg) error
}

func DiskBlockCacheProtocol(i DiskBlockCacheInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "kbgitkbfs.1.DiskBlockCache",
		Methods: map[string]rpc.ServeHandlerDescription{
			"GetBlock": {
				MakeArg: func() interface{} {
					ret := make([]GetBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetBlockArg)(nil), args)
						return
					}
					ret, err = i.GetBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"PutBlock": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockArg)(nil), args)
						return
					}
					err = i.PutBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"DeleteBlocks": {
				MakeArg: func() interface{} {
					ret := make([]DeleteBlocksArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DeleteBlocksArg)
					if !ok {
						err = rpc.NewTypeError((*[]DeleteBlocksArg)(nil), args)
						return
					}
					ret, err = i.DeleteBlocks(ctx, (*typedArgs)[0].BlockIDs)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"UpdateBlockMetadata": {
				MakeArg: func() interface{} {
					ret := make([]UpdateBlockMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UpdateBlockMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]UpdateBlockMetadataArg)(nil), args)
						return
					}
					err = i.UpdateBlockMetadata(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type DiskBlockCacheClient struct {
	Cli rpc.GenericClient
}

// GetBlock gets a block from the disk cache.
func (c DiskBlockCacheClient) GetBlock(ctx context.Context, __arg GetBlockArg) (res GetBlockRes, err error) {
	err = c.Cli.Call(ctx, "kbgitkbfs.1.DiskBlockCache.GetBlock", []interface{}{__arg}, &res)
	return
}

// PutBlock puts a block into the disk cache.
func (c DiskBlockCacheClient) PutBlock(ctx context.Context, __arg PutBlockArg) (err error) {
	err = c.Cli.Call(ctx, "kbgitkbfs.1.DiskBlockCache.PutBlock", []interface{}{__arg}, nil)
	return
}

// DeleteBlocks deletes a set of blocks from the disk cache.
func (c DiskBlockCacheClient) DeleteBlocks(ctx context.Context, blockIDs []string) (res DeleteBlocksRes, err error) {
	__arg := DeleteBlocksArg{BlockIDs: blockIDs}
	err = c.Cli.Call(ctx, "kbgitkbfs.1.DiskBlockCache.DeleteBlocks", []interface{}{__arg}, &res)
	return
}

// UpdateBlockMetadata updates the metadata for a block in the disk cache.
func (c DiskBlockCacheClient) UpdateBlockMetadata(ctx context.Context, __arg UpdateBlockMetadataArg) (err error) {
	err = c.Cli.Call(ctx, "kbgitkbfs.1.DiskBlockCache.UpdateBlockMetadata", []interface{}{__arg}, nil)
	return
}
