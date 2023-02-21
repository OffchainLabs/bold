package goimpl

import (
	"context"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
)

type Inbox struct {
	messages [][]byte
	feed     *EventFeed[[]byte]
}

func NewInbox(ctx context.Context) *Inbox {
	return &Inbox{
		messages: [][]byte{},
		feed:     NewEventFeed[[]byte](ctx),
	}
}

func (inbox *Inbox) Subscribe(ctx context.Context, c chan<- []byte) {
	inbox.feed.Subscribe(ctx, c)
}

func (inbox *Inbox) SubscribeWithFilter(ctx context.Context, c chan<- []byte, filter func([]byte) bool) {
	inbox.feed.SubscribeWithFilter(ctx, c, filter)
}

func (inbox *Inbox) Append(tx protocol.ActiveTx, message []byte) {
	tx.VerifyReadWrite()
	inbox.messages = append(inbox.messages, message)
	inbox.feed.Append(message)
}

func (inbox *Inbox) NumMessages(tx protocol.ActiveTx) uint64 {
	tx.VerifyRead()
	return uint64(len(inbox.messages))
}

func (inbox *Inbox) GetMessage(tx protocol.ActiveTx, num uint64) ([]byte, error) {
	tx.VerifyRead()
	if num >= uint64(len(inbox.messages)) {
		return nil, ErrInvalidOp
	}
	return inbox.messages[num], nil
}
