package visualization

import (
	"context"
	"net/http"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "visualization-server")

type Visualization struct {
	visualizer protocol.ChainVisualizer
	upgrader   websocket.Upgrader
}

func New(visualizer protocol.ChainVisualizer) *Visualization {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(j *http.Request) bool {
			return true
		},
	}
	return &Visualization{visualizer: visualizer, upgrader: upgrader}
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
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
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
