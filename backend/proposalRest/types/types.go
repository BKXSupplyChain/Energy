package types

type Proposal struct {
    ID          string `json:"ID,omitemty"`
    Supplier    string `json:"Supplier,omitempty"`
    Consumer    string `json:"Consumer,omitempty"`
    Price       int32  `json:"Price,omitempty"`
    TotalCost   int32  `json:"TotalCost,omitempty"`
    DeltaPrice  int32  `json:"DeltaPrice,omitempty"`
    Salt        int32  `json:"Salt,omitempty"`
    //  Salt is secret code using between neighbours
    TTL         int32  `json:"TTL,omitempty"`
}

