package groupcache

// peers.go defines how processes find and communicate with their peers.

import (
	pb "groupcache/pkg/groupcachepb"
)

// Context is an opaque value passed through calls to the
// ProtoGetter. It may be nil if your ProtoGetter implementation does
// not require a context.
type Context interface{}

// ProtoGetter is the interface that must be implemented by a peer
type ProtoGetter interface {
	Get(context Context, in *pb.GetRequest, out *pb.GetResponse) error
}

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer ProtoGetter, ok bool)
}

// NoPeers is an implementation of PeerPicker that nevers finds a peer
type NoPeers struct{}

func (NoPeers) PickPeer(key string) (peer ProtoGetter, ok bool) { return }

var (
	portPicker func(groupName string) PeerPicker
)

func RegisterPeerPicker(fn func() PeerPicker) {
	if portPicker != nil {
		panic("RegisterPeerPicker called more than once")
	}
	portPicker = func(_ string) PeerPicker { return fn() }
}

func RegisterPerGroupPeerPicker(fn func(groupName string) PeerPicker) {
	if portPicker != nil {
		panic("RegisterPeerPicker called more than once")
	}

	portPicker = fn
}

func getPeers(groupName string) PeerPicker {
	if portPicker == nil {
		return NoPeers{}
	}

	pk := portPicker(groupName)
	if pk == nil {
		pk = NoPeers{}
	}
	return pk
}
