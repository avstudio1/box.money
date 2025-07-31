# VTArchitect First-Order System Blueprint

**Codename:** Goldstream  
**Designation:** First-Order Truth  
**Purpose:** Define the full structure, stack, logic, and expansion rules for recursive system growth.

---

## System Mantra

> "He who builds faithfully at the node, builds the system."

---

## Project Layout

```
vtarchitect/
├── api/                    # GoFiber app
│   ├── main.go
│   ├── handlers/
│   ├── services/
│   ├── grpc/
│   └── events/            # NATS/Kafka integration
├── console/               # Vite + React frontend
│   ├── src/
│   ├── public/
│   └── vite.config.ts
├── shared/                # Shared Go/TS schema & utils
├── infra/                 # Docker, Caddy, deploy files
│   ├── docker-compose.yml
│   ├── caddy/Caddyfile
│   └── k8s/
├── .env
├── Makefile
└── README.md
```

---

## Stack Includes

**Backend:**  
GoFiber, grpc-go, nats.go, pgx, otel

**Frontend:**  
Vite, React, ShadCN, TanStack, Socket.io

**Infrastructure:**  
Docker, Caddy, Fly.io-ready

All components are modular and swappable. System integrity is governed by blueprint adherence.

---

## Nodes of Learning (Node Blueprint Table)

| Node # | Tree Path                             | Domain   | Prompt Summary                                          |
| ------ | ------------------------------------- | -------- | ------------------------------------------------------- |
| 1      | api/main.go                           | Backend  | Entrypoint: Fiber, gRPC, Redis, NATS, graceful shutdown |
| 2      | api/services/auction.go               | Backend  | Core auction logic: validation, timer, anti-snipe       |
| 3      | api/grpc/auction_server.go            | Backend  | gRPC stream: bid submit + push updates                  |
| 4      | console/src/pages/auction/[id].tsx    | Frontend | Dynamic auction page: Zustand, UI, stream               |
| 5      | console/src/hooks/useAuctionStream.ts | Frontend | Bid/Result streaming over gRPC-Web or WS                |
| 6      | console/src/components/BidPanel.tsx   | Frontend | UI input + bid history panel                            |
| 7      | shared/schema/auction.proto           | Shared   | Proto: BidRequest, BidUpdate, ResultUpdate              |
| 8      | shared/types/auction.d.ts             | Shared   | TS types from proto + helper interfaces                 |
| 9      | shared/utils/currency_format.go/ts    | Shared   | Cross-formatting utilities (Go + TS)                    |
| 10     | infra/docker-compose.yml              | Infra    | Full stack services: api, console, redis, nats, caddy   |
| 11     | infra/caddy/Caddyfile                 | Infra    | Proxy rules for /api and /rpc                           |
| 12     | infra/k8s/api-deploy.yaml             | Infra    | K8s deployment: backend + services                      |
| 13     | README.md                             | Root     | System description, flow, monetization engine           |
| 14     | Makefile                              | Root     | Dev scripts: proto, build, run, test                    |
| 15     | .env                                  | Root     | Unified secrets/config for all services                 |

---

## Invocation Format (per node)

```
## Node N — [file path]

**Prompt**:  
[Insert corresponding prompt text here]

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_
```

---

## Rules of Expansion

- Always maintain **3n** node count (e.g., 3, 6, 9, 12, 15…)  
- No node or file without tree alignment  
- All code must trace back to a documented prompt  
- Second-order documents serve node implementation, not initiation  
- First-order documents must remain immutable unless promoted through node update  
- The system is recursive: documentation → prompt → implementation → runtime → analytics → re-documentation

---

## Instructions to Agent(s)

1. Never violate node alignment or expand outside the blueprint without explicit directive.
2. Maintain calm, precise output.
3. Treat this document as canonical—every decision must point back to a node.
4. Preserve style, cadence, and intent.
5. The engine must flow with profit in mind—this is the $box$money core.

---

## Final Mantras

> *“He who builds faithfully at the node, builds the system.”*  
> *“Perfectly imperfect, imperfectly perfect.”*  
> *“The system remembers. The system responds.”*
