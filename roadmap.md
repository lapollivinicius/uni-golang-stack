# Roadmap Completo para se Tornar Sênior em Go (Golang)

## 1. Fundamentos da Linguagem

### Dominar

- Variáveis
- Constantes
- Tipos básicos
- Arrays
- Slices
- Maps
- Structs
- Ponteiros
- Funções
- Métodos
- Interfaces
- Packages
- Modules
- Generics
- Closures
- Defer
- Panic
- Recover
- Reflection
- Embedding
- Type Assertions
- Type Switches
- Context

### Ser capaz de responder

- Quando usar ponteiros?
- Quando copiar structs?
- Quando usar interfaces?
- Quando evitar interfaces?
- Quando usar generics?

---

## 2. Gerenciamento de Dependências

### Conhecimentos

- go.mod
- go.sum
- Semantic Versioning
- Módulos privados
- Replace directives
- Workspaces

### Comandos

```bash
go mod init
go mod tidy
go mod vendor
go work init
```

---

## 3. Concorrência

### Dominar

- Goroutines
- Channels
- Buffered Channels
- Unbuffered Channels
- Select
- Worker Pools
- Pipelines
- Fan-In
- Fan-Out
- Context Cancellation

### Pacote sync

```go
sync.Mutex
sync.RWMutex
sync.WaitGroup
sync.Once
sync.Cond
sync.Pool
sync.Map
```

### Pacote atomic

```go
atomic.AddInt64
atomic.LoadInt64
atomic.StoreInt64
```

### Problemas clássicos

- Race Conditions
- Deadlocks
- Livelocks
- Starvation

### Ferramentas

```bash
go test -race
```

---

## 4. Runtime Interno do Go

### Entender

- Scheduler GMP
- Stack Growth
- Heap
- Garbage Collector
- Escape Analysis
- Memory Model
- Alocação de memória

### Ferramentas

```bash
go build -gcflags="-m"
```

---

## 5. Estruturas de Dados

### Implementar do zero

- Linked List
- Stack
- Queue
- Heap
- Priority Queue
- Hash Table
- Trie
- Trees
- BST
- Graphs

### Complexidade

- O(1)
- O(log n)
- O(n)
- O(n log n)
- O(n²)

---

## 6. Algoritmos

### Dominar

- Binary Search
- DFS
- BFS
- Dijkstra
- Dynamic Programming
- Topological Sort
- Sliding Window
- Two Pointers
- Backtracking

---

## 7. Testes

### Tipos

- Unit Tests
- Integration Tests
- Benchmarks
- Table Driven Tests
- Mocks
- Stubs
- Fakes

### Ferramentas

```bash
go test
go test ./...
go test -bench
```

### Bibliotecas

- Testify
- GoMock

---

## 8. Profiling e Performance

### Ferramentas

- pprof
- trace
- benchmark

### Conhecimentos

- CPU Profiling
- Heap Profiling
- Allocation Profiling
- Memory Leaks

---

## 9. HTTP e APIs

### Dominar

- net/http
- REST
- JSON
- Middleware
- Routing
- Rate Limiting
- Authentication

### Frameworks

- Gin
- Echo
- Fiber
- Chi

---

## 10. Bancos Relacionais

### SQL

- SELECT
- JOIN
- GROUP BY
- HAVING
- Window Functions
- Transactions
- Isolation Levels
- Locks
- Indexes

### Bancos

- PostgreSQL
- MySQL
- SQLite

### Drivers

- database/sql
- pgx

---

## 11. NoSQL

### Conhecer

- MongoDB
- Redis
- Cassandra

### Conceitos

- Documento
- Chave-Valor
- Colunar

---

## 12. Arquitetura

### Estudar

- Clean Architecture
- Hexagonal Architecture
- DDD
- CQRS
- Event-Driven Architecture

---

## 13. Microsserviços

### Dominar

- Service Discovery
- API Gateway
- Circuit Breaker
- Retry
- Backoff
- Saga Pattern
- Idempotência

---

## 14. Mensageria

### Ferramentas

- Kafka
- RabbitMQ
- NATS

### Conceitos

- Producer
- Consumer
- Offset
- Partition
- Ordering
- At Least Once
- Exactly Once

---

## 15. gRPC

### Conhecimentos

- Protocol Buffers
- Unary RPC
- Streaming
- Bidirectional Streaming
- Interceptors

---

## 16. Segurança

### Dominar

- JWT
- OAuth2
- OpenID Connect
- TLS
- HTTPS
- CORS
- CSRF
- XSS
- SQL Injection
- Secret Management

---

## 17. Docker

### Conhecer

- Dockerfile
- Multi-stage Builds
- Networks
- Volumes
- Docker Compose

---

## 18. Kubernetes

### Conhecer

- Pods
- Deployments
- Services
- ConfigMaps
- Secrets
- Ingress
- HPA
- StatefulSets

---

## 19. Cloud

### Escolher uma

- AWS
- GCP
- Azure

### Aprender

- IAM
- Storage
- Compute
- Containers
- Messaging
- Monitoring

---

## 20. Observabilidade

### Ferramentas

- OpenTelemetry
- Prometheus
- Grafana

### Conceitos

- Logs
- Metrics
- Tracing

---

## 21. CI/CD

### Conhecer

- Build
- Test
- Deploy
- Rollback

### Ferramentas

- GitHub Actions
- GitLab CI/CD
- Jenkins

---

## 22. Linux

### Comandos

```bash
grep
awk
sed
curl
ss
netstat
top
htop
lsof
strace
```

### Conceitos

- TCP/IP
- DNS
- Sockets
- Processos
- Threads
- Systemd

---

## 23. Git Avançado

```bash
git rebase
git cherry-pick
git bisect
git stash
git reset
git reflog
```

---

## 24. Engenharia de Software

### Dominar

- SOLID
- Design Patterns
- Refactoring
- Code Review
- Technical Debt
- System Design
- Trade-offs

---

## 25. Competências de Sênior

- Liderança técnica
- Mentoria
- Investigação de incidentes
- Comunicação com stakeholders
- Planejamento técnico
- Arquitetura de sistemas
- Análise de riscos
- Otimização de performance