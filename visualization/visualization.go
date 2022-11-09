package visualization

import (
	"context"
	"net/http"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "visualization-server")

type Visualization struct {
	visualizer     protocol.ChainVisualizer
	eventsProvider protocol.EventsProvider
	eventsReceived chan protocol.AssertionChainEvent
	upgrader       websocket.Upgrader
}

func New(
	ctx context.Context,
	visualizer protocol.ChainVisualizer,
	eventsProvider protocol.EventsProvider,
) *Visualization {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(j *http.Request) bool {
			return true
		},
	}
	eventsReceived := make(chan protocol.AssertionChainEvent, 1)
	eventsProvider.SubscribeChainEvents(ctx, eventsReceived)
	return &Visualization{
		visualizer:     visualizer,
		upgrader:       upgrader,
		eventsProvider: eventsProvider,
		eventsReceived: eventsReceived,
	}
}

func (v *Visualization) Start(ctx context.Context) {
	http.HandleFunc("/assertion-chain", v.streamAssertionChainGraph(ctx))
	http.Handle("/", http.FileServer(http.Dir("visualization")))
	// TODO: Make configurable.
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func (v *Visualization) streamAssertionChainGraph(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := v.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.WithError(err).Error("Could not upgrade websocket connection")
			return
		}
		defer c.Close()
		for {
			select {
			case <-v.eventsReceived:
				g := v.visualizer.Visualize()
				if err = c.WriteMessage(1, []byte(g)); err != nil {
					log.WithError(err).Error("Could not write chain graph string over websocket")
					return
				}
			case <-ctx.Done():
			case <-r.Context().Done():
				return
			}
		}
	}
}
