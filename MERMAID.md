# System Data Flow (Bid Placement)

This diagram visualizes the "System Data Flow" described in the `README.md`. It traces a user's bid through the full stack, highlighting the interaction between different nodes of the system.

```mermaid
sequenceDiagram
    participant User
    participant Console as "Console (Browser)"
    participant Caddy as "Caddy (Proxy)"
    participant API as "API (Go Backend)"
    participant DB as "Database"

    User->>+Console: 1. Enters bid in BidPanel (Node 6) on Auction Page (Node 4)

    Console->>Console: 2. useAuctionStream (Node 5) packages BidRequest (Node 7)

    Console->>+Caddy: 3. gRPC-Web stream with BidRequest
    Note over Caddy, API: Node 11 - Proxy rules
    Caddy->>+API: Forwards request to gRPC server

    API->>API: 4. gRPC Server (Node 3) receives request<br/>(auth middleware Node 18)

    API->>API: 5. Auction Service (Node 2) validates bid

    alt Bid is valid
        API->>+DB: 6. Persist bid via Querier (Node 17)<br/>using Schema (Node 16)
        DB-->>-API: Bid saved

        API->>API: 7. Calculate new state (price, anti-snipe) in Auction Service (Node 2)

        API-->>-Caddy: 8. Broadcast AuctionUpdate (Node 7) to all clients
        Caddy-->>-Console: Pushes update
    end

    Console->>-User: 9. UI re-renders with new price/history (Nodes 4, 5, 6)
```