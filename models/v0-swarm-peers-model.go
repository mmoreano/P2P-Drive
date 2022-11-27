package models

type SwarmPeersResponse struct {
	Peers []peer `json:"Peers"`
	Test  string
}

type peer struct {
	Addr      string   `json:"Addr"`
	Direction int      `json:"Direction"`
	Latency   string   `json:"Latency"`
	Muxer     string   `json:"Muxer"`
	Peer      string   `json:"Peer"`
	Streams   []stream `json:"Streams"`
}

type stream struct {
	Protocol string `json:"Protocol"`
}