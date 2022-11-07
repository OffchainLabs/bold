package visualization

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/gorilla/websocket"
)

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
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func (v *Visualization) streamAssertionChainGraph(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := v.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
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
					log.Println("Error writing:", err)
					break
				}
			case <-ctx.Done():
			case <-r.Context().Done():
				log.Println("Context done:", err)
				break
			}
		}
	}
}
