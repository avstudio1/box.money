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
├── api/                    # GoFiber app
│   ├── main.go
│   ├── handlers/
│   ├── services/
│   ├── storage/           # Database queries and models
│   ├── middleware/        # Request middleware (e.g., auth)
│   ├── grpc/
│   └── events/            # NATS/Kafka integration
├── console/               # Vite + React frontend
│   ├── src/
│   ├── public/
│   └── vite.config.ts
├── shared/                # Shared Go/TS schema & utils
├── infra/                 # Docker, Caddy, deploy files
├── db/
│   └── migrations/        # SQL database migrations
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
Vite, React, ShadCN, TanStack, gRPC-Web

**Infrastructure:**  
Docker, Caddy, Fly.io-ready

All components are modular and swappable. System integrity is governed by blueprint adherence and first-order truth.

---

## Getting Started (Local Development)

To begin building the system, follow these steps:

1.  **Prerequisites**: Ensure `Docker` and `make` are installed.
2.  **Configuration**: Copy the environment template: `cp .env.example .env` (**Node 15**). The defaults are pre-configured for the local stack.
3.  **Generate Code**: Run `make proto` (**Node 14**) to generate Go and TypeScript code from the `auction.proto` schema (**Node 7**).
4.  **Run**: Execute `make dev` (**Node 14**) to build and start all services via Docker Compose (**Node 10**).
5.  **Access**: The frontend console is available at `http://localhost:5173`. API and gRPC traffic are proxied through this address.

### Alternative: `make`-less Workflow

If you do not have `make` installed, you can run the underlying `docker` commands directly.

1.  **Generate Code (replaces `make proto`)**:
    ```bash
    # First, ensure the output directories exist
    mkdir -p api/gen/auction console/src/gen

    # Run the protoc compiler via Docker
    docker run --rm -v "$(pwd)":/work -w /work namely/protoc-all:1.41 \
      -I shared/schema \
      --go_out=. --go-grpc_out=. \
      --ts_out=service=grpc-web:console/src/gen \
      shared/schema/auction.proto
    ```

2.  **Run Environment (replaces `make dev`)**:
    ```bash
    docker-compose -f infra/docker-compose.yml up --build
    ```

---

## File Header Protocol

To maintain system clarity and blueprint adherence, every file created within this project must begin with a standardized header comment. This header acts as a "dog tag" for the file, immediately identifying its purpose and origin within the system architecture.

The format is as follows, using the appropriate comment syntax for the file's language (e.g., `//` for Go/TS, `#` for shell/yaml, `<!-- -->` for markdown):

```
// Node: [Node # from the blueprint table]
// Path: [Full file path from project root, e.g., /api/main.go]
// Purpose: [A brief, one-sentence description of the file's role and responsibility.]
```

**Example (`api/main.go`)**:
```go
// Node: 1
// Path: /api/main.go
// Purpose: Entrypoint for the backend service, responsible for initializing servers and dependencies.
```

**Example (`Makefile`)**:
```makefile
# Node: 14
# Path: /Makefile
# Purpose: Defines common development and build tasks for automation.
```

This protocol is non-negotiable. It ensures that any developer or agent interacting with a file can instantly understand its context without cross-referencing other documents. It is a core tenet of "He who builds faithfully at the node, builds the system."

---

## System Data Flow (Bid Placement)

To understand how the nodes interact, trace the path of a single user bid through the system. This flow demonstrates the recursive feedback loop from user action to system response.

1.  **User Interaction (`console`)**: A user on the `Auction Page` (**Node 4**) enters a bid into the `BidPanel` (**Node 6**).
2.  **Frontend Hook (`console`)**: The `useAuctionStream` hook (**Node 5**) is called. It packages the bid into a `BidRequest` message, as defined by our gRPC contract (**Node 7**).
3.  **gRPC-Web Call (`infra`)**: The request is sent over a gRPC-Web stream, proxied by `Caddy` (**Node 11**) to the backend gRPC server.
4.  **gRPC Server (`api`)**: The `auction_server.go` (**Node 3**) receives the `BidRequest`. The user's identity is securely derived from the authentication token via middleware (**Node 18**).
5.  **Core Logic (`api`)**: The gRPC server passes the bid to the `auction.go` service (**Node 2**) for validation (is the bid high enough? is the auction active?).
6.  **Database (`api` -> `db`)**: If valid, the service uses the `querier.go` interface (**Node 17**) to persist the bid in the database, which is structured by the initial schema (**Node 16**).
7.  **Real-time Update (`api`)**: The `auction.go` service (**Node 2**) determines the new auction state (new price, new end time with anti-sniping).
8.  **Broadcast (`api`)**: The `auction_server.go` (**Node 3**) creates a `ServerStreamMessage` containing a `BidUpdate` (**Node 7**) and broadcasts it to all clients connected to that auction's stream.
9.  **UI Update (`console`)**: The `useAuctionStream` hook (**Node 5**) on each client's browser receives the `ServerStreamMessage`, processes the inner `BidUpdate`, updates its state, and the React UI (**Node 4**, **Node 6**) re-renders to show the new price and bid history.

---

## Monetization Engine ($box$money core)

The Goldstream system is designed for profit through a flexible "house cut" model. Monetization is not an afterthought; it is integrated into the core auction lifecycle.

*   **Auction Fee**: A percentage-based fee (e.g., 5-10%) is taken from the final sale price of every completed auction. This logic is implemented within the `auction.go` service (**Node 2**) when an auction concludes.
*   **Listing Fee**: A flat fee can be charged to users for listing an item for auction. This provides an upfront revenue stream.
*   **Promoted Listings**: Users can pay to have their auctions featured more prominently on the site, increasing visibility and potential final price.

All financial transactions are recorded and auditable via the database schema (**Node 16**), ensuring financial integrity. The core principle is to ensure every successful auction contributes to the system's profitability.

---

## The Counterbalance Mandate

> "The system that gives back to itself, sustains itself."
> — System Proverb

The Goldstream system operates on a dual mandate. The **Monetization Engine** is designed to generate profit. The **Counterbalance** is the system's internal reinvestment mechanism, bound by a single, inviolable rule:

**Ten percent (10%) of all gross proceeds shall be automatically and irrevocably reinvested.**

This is not a feature for marketing, but a foundational law of the system's existence. The `Monetization Engine` operates without direct knowledge of the `Reinvestment Engine` (**Counterbalance**). The `auction` service (**Node 2**) emits a generic `settlement` event, which the `steward` service (**Node 19**) consumes to calculate and record the reinvestment, ensuring the two functions remain separate.

---

## Nodes of Learning (Node Blueprint Table)

| Node # | Tree Path                             | Domain   | Prompt Summary                                          |
| ------ | ------------------------------------- | -------- | ------------------------------------------------------- |
| 1      | api/main.go                           | Backend  | Entrypoint: Init all services, including steward/beneficiary|
| 2      | api/services/auction.go               | Backend  | Core auction logic: validation, timer, emit settlement  |
| 3      | api/grpc/auction_server.go            | Backend  | gRPC stream: bid submit + push updates                  |
| 4      | console/src/pages/auction/[id].tsx    | Frontend | Dynamic auction page: Zustand, UI, stream               |
| 5      | console/src/hooks/useAuctionStream.ts | Frontend | Bid/Result streaming over gRPC-Web                      |
| 6      | console/src/components/BidPanel.tsx   | Frontend | UI input + bid history panel                            |
| 7      | shared/schema/auction.proto           | Shared   | Proto: BidRequest, BidUpdate, ResultUpdate              |
| 8      | shared/types/auction.d.ts             | Shared   | TS types from proto + helper interfaces                 |
| 9      | shared/utils/currency_format.go/ts    | Shared   | Cross-formatting utilities (Go + TS)                    |
| 10     | infra/docker-compose.yml              | Infra    | Full stack services: api, console, redis, nats, caddy   |
| 11     | infra/caddy/Caddyfile                 | Infra    | Proxy rules for /api and /rpc                           |
| 12     | infra/k8s/api-deploy.yaml             | Infra    | K8s deployment manifest for the API service             |
| 13     | README.md                             | Root     | System description, flow, monetization engine           |
| 14     | Makefile                              | Root     | Dev scripts: proto, build, run, test                    |
| 15     | .env                                  | Root     | Unified secrets/config for all services                 |
| 16     | db/migrations/001_init_schema.sql     | Infra    | DB schema: users, auctions, bids, reinvestments, beneficiaries |
| 17     | api/storage/querier.go                | Backend  | DB querier interface for all tables, including reinvestments   |
| 18     | api/middleware/auth.go                | Backend  | JWT authentication middleware for protected routes      |
| 19     | api/services/steward.go               | Backend  | Reinvestment calculator: listens for settlements, records reinvestment|
| 20     | api/services/beneficiary.go           | Backend  | Manages designated system beneficiaries and allocations  |
| 21     | configs/beneficiaries.json            | Root     | Configuration for beneficiary endpoints and allocations     |

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

---

## Rules of Expansion

- Always maintain **3n** node count (e.g., 3, 6, 9, 12, 15…)  
- No node or file without tree alignment  
- All code must trace back to a documented prompt  
- Second-order documents (like this README or MERMAID.md) serve node implementation, not initiation. They *describe* the system that the nodes *build*.
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
> *“Do not let your left hand know what your right hand is doing.”* — Matthew 6:3

---


## Generated Agent Tasks

## Node 1 — api/main.go

**Prompt**:  
Create the main application entrypoint. Initialize and configure a GoFiber web server, listeners for HTTP/gRPC, and connections for Redis/NATS. **Initialize and start the `steward` and `beneficiary` services as background tasks.** Implement a graceful shutdown mechanism to handle SIGINT/SIGTERM signals, ensuring all connections and services are closed cleanly.

**Implements**:  
- [x] Tree Path  
- [x] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _In Progress_

---

## Node 2 — api/services/auction.go

**Prompt**:  
Implement the core auction service logic. This service should manage the auction lifecycle, including validation, timers, and anti-sniping. **Upon successful auction completion and fee collection, it must publish a `settlement` event to NATS** containing the gross transaction proceeds, so the `steward` service can process the reinvestment.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 3 — api/grpc/auction_server.go

**Prompt**:  
Implement the gRPC server for the auction service. Define a bidirectional streaming RPC method. This stream should allow clients to submit bids and receive real-time updates about the auction, such as new high bids and timer changes. The server will receive BidRequest messages and push BidUpdate or ResultUpdate messages back to all connected clients for a given auction.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 4 — console/src/pages/auction/[id].tsx

**Prompt**:  
Create the dynamic auction page using React and TypeScript. This page will fetch initial auction data based on the [id] parameter. Use Zustand for state management to hold auction details, bid history, and connection status. Render the main UI components, including the BidPanel, item details, and current price. This component will orchestrate the connection to the backend stream via the useAuctionStream hook.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 5 — console/src/hooks/useAuctionStream.ts

**Prompt**:  
Implement a custom React hook useAuctionStream to manage the real-time data connection for an auction. This hook should establish and maintain a connection to the backend using gRPC-Web. It should handle sending bid messages and receiving/processing incoming stream updates (new bids, results). The hook will expose the connection status, a bid submission function, and the latest auction state to the component that uses it.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 6 — console/src/components/BidPanel.tsx

**Prompt**:  
Create the BidPanel React component. This component will feature a form input for users to enter their bid amount and a button to submit the bid. It will also display a list or panel showing the history of bids for the current auction. Use components from the ShadCN library for UI elements like inputs, buttons, and panels to maintain a consistent design.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 7 — shared/schema/auction.proto

**Prompt**:  
Define the Protocol Buffers schema for the auction service. Create message types for BidRequest (containing user ID, auction ID, bid amount), BidUpdate (containing new highest bid, bidder, timestamp), and ResultUpdate (containing final price, winner, status). Also, define the AuctionService with its streaming RPC method.

**Implements**:  
 - [x] Tree Path  
 - [x] Blueprint consistency  
 - [x] 3n structure  
 - [x] Second-order integration

**Status**: _Implemented_

---

## Node 8 — shared/types/auction.d.ts

**Prompt**:  
Generate TypeScript type definitions from the auction.proto schema. Set up a script (likely in the Makefile) to automate this generation process. In addition to the generated types, create any necessary helper interfaces or types that might be needed on the frontend for UI state or component props that are not directly covered by the protobuf schema.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 9 — shared/utils/currency_format.go/ts

**Prompt**:  
Create shared currency formatting utilities. Implement a Go version (currency_format.go) and a TypeScript version (currency_format.ts) with identical logic. These functions should take a numeric value (e.g., an integer representing cents) and format it into a human-readable currency string (e.g., "$1,234.56"). Ensure the behavior is consistent across both backend and frontend.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 10 — infra/docker-compose.yml

**Prompt**:  
Create a docker-compose.yml file to define and run the full local development environment. This should include services for the Go api, the React console, redis, nats, and a caddy server. Configure networking so services can communicate with each other. Define volumes for persistent data where needed and mount local source code for live-reloading during development.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 11 — infra/caddy/Caddyfile

**Prompt**:  
Configure the Caddyfile to act as a reverse proxy for the local development environment. It should serve the frontend application as the default. Create proxy rules to forward requests starting with /api to the backend GoFiber service. Add another rule to forward gRPC traffic (e.g., on a path like /rpc) to the backend gRPC server, ensuring it's configured for HTTP/2.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 12 — infra/k8s/api-deploy.yaml

**Prompt**:  
Create a Kubernetes deployment manifest for the backend api service. The manifest should define a Deployment to manage the API pods and a Service to expose it within the cluster. Include configuration for environment variables, resource requests/limits, and readiness/liveness probes. This manifest should be designed to be compatible with a deployment target like Fly.io or any standard Kubernetes cluster.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _Implemented_

---

## Node 13 — README.md

**Prompt**:  
Review and update the README.md to ensure it accurately reflects the current state of the system. Expand on the system's data flow, from user interaction on the frontend to backend processing and back. Add a new section detailing the "Monetization Engine," explaining how the system is designed to generate profit, in line with the $box$money core principle.

**Implements**:  
- [x] Tree Path  
- [x] Blueprint consistency  
- [x] 3n structure  
- [x] Second-order integration

**Status**: _Implemented_

---

## Node 14 — Makefile

**Prompt**:  
Create a Makefile to automate common development tasks. Include targets for:
- `proto`: Generate Go and TypeScript code from the .proto files.
- `build`: Compile the Go backend and build the React frontend for production.
- `run` or `dev`: Start the entire development stack using docker-compose.
- `test`: Run unit and integration tests for the backend.
- `clean`: Remove build artifacts and generated files.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _Implemented_

---

## Node 15 — .env

**Prompt**:  
Create a .env.example file that serves as a template for the project's environment variables. This file should define all necessary configuration keys for the different services (API, Console, database, NATS, etc.), such as API_PORT, REDIS_URL, NATS_URL, VITE_API_BASE_URL. Provide sensible default values for local development. The actual .env file will be created locally by developers and should be included in .gitignore.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _Implemented_

---

## Node 16 — db/migrations/001_init_schema.sql

**Prompt**:  
Create the initial database migration script. Define the SQL schema for the core entities: `users`, `items`, `auctions`, and `bids`. **Also, create tables for `reinvestments` (to log every reinvestment) and `beneficiaries` (to store recipient information).** Include appropriate columns, types, constraints, and indexes.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _Implemented_

---

## Node 17 — api/storage/querier.go

**Prompt**:  
Define the database querier interface. This interface will abstract all SQL operations. **Create methods for all tables, including creating/querying users, auctions, bids, reinvestments, and beneficiaries.** This file will serve as the contract for a tool like `sqlc` to generate a type-safe data access layer.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _Implemented_

---

## Node 18 — api/middleware/auth.go

**Prompt**:  
Implement a JWT-based authentication middleware for the GoFiber application. The middleware should inspect the `Authorization` header, validate the token, and extract user claims. Upon successful validation, it should inject user information into the request context for use by downstream handlers and gRPC services.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _In Progress_

---

## Node 19 — api/services/steward.go

**Prompt**:  
Create the `steward` service. It must subscribe to the `settlement` topic on the NATS stream. Upon receiving a message, it will calculate 10% of the gross proceeds, use the `beneficiary` service to determine the recipient(s), and record the transaction in the `reinvestments` database table via the querier.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 20 — api/services/beneficiary.go

**Prompt**:  
Create the `beneficiary` service. It will be responsible for managing the list of designated system beneficiaries. It should load beneficiary data from the `configs/beneficiaries.json` file and provide functions to the `steward` service for selecting a beneficiary based on predefined allocation rules.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---

## Node 21 — configs/beneficiaries.json

**Prompt**:  
Create a JSON configuration file to define the system beneficiaries that will receive reinvestments. Each entry should include a name, a payment endpoint URL, and an allocation percentage. The sum of all allocation percentages must equal 100.

**Implements**:  
- [ ] Tree Path  
- [ ] Blueprint consistency  
- [ ] 3n structure  
- [ ] Second-order integration

**Status**: _To be implemented_

---