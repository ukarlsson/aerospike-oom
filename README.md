# aerospike-oom

Run this command

```
go run main.go -host localhost -port 3000 -namespace test -elements 40000
```

Then wait a few seconds and the OOM killer kills the `asd` process.

```
Jul 03 2019 19:40:13 GMT: DEBUG (nsup): (nsup.c:907) {test} stop-writes limit not breached, memory sz:64 (64 + 0 + 0) limit:3865470566, disk avail-pct:100
Jul 03 2019 19:40:13 GMT: DEBUG (nsup): (nsup.c:907) {event} stop-writes limit not breached, memory sz:0 (0 + 0 + 0) limit:3865470566, disk avail-pct:100
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:162) NODE-ID bb90a986cf82800 CLUSTER-SIZE 1
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:243)    cluster-clock: skew-ms 0
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:270)    system-memory: free-kbytes 8919080 free-pct 56 heap-kbytes (15761969,15830444,22245376) heap-efficiency-pct 70.9
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:282)    in-progress: tsvc-q 0 info-q 0 rw-hash 1 proxy-hash 0 tree-gc-q 0
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:303)    fds: proto (1,7,6) heartbeat (0,0,0) fabric (0,0,0)
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:311)    heartbeat-received: self 198 foreign 0
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:341)    fabric-bytes-per-second: bulk (0,0) ctrl (0,0) meta (0,0) rw (0,0)
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:399) {test} objects: all 1 master 1 prole 0 non-replica 0
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:457) {test} migrations: complete
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:475) {test} memory-usage: total-bytes 64 index-bytes 64 sindex-bytes 0 data-bytes 0 used-pct 0.00
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:399) {event} objects: all 0 master 0 prole 0 non-replica 0
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:457) {event} migrations: complete
Jul 03 2019 19:40:13 GMT: INFO (info): (ticker.c:475) {event} memory-usage: total-bytes 0 index-bytes 0 sindex-bytes 0 data-bytes 0 used-pct 0.00
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:6817) dead nodes at quantum start: (empty)
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:6823) faulty nodes at quantum start: (empty)
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:6833) join requests at quantum start: (empty)
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:3510) clique based evicted nodes at quantum start: (empty)
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:2153) plugin data missing protocol identifier
Jul 03 2019 19:40:14 GMT: DEBUG (clustering): (clustering.c:6777) no cluster merge candidates found
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:6817) dead nodes at quantum start: (empty)
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:6823) faulty nodes at quantum start: (empty)
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:6833) join requests at quantum start: (empty)
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:3510) clique based evicted nodes at quantum start: (empty)
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:2153) plugin data missing protocol identifier
Jul 03 2019 19:40:16 GMT: DEBUG (clustering): (clustering.c:6777) no cluster merge candidates found
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:6817) dead nodes at quantum start: (empty)
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:6823) faulty nodes at quantum start: (empty)
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:6833) join requests at quantum start: (empty)
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:3510) clique based evicted nodes at quantum start: (empty)
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:2153) plugin data missing protocol identifier
Jul 03 2019 19:40:18 GMT: DEBUG (clustering): (clustering.c:6777) no cluster merge candidates found
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:6817) dead nodes at quantum start: (empty)
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:6823) faulty nodes at quantum start: (empty)
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:6833) join requests at quantum start: (empty)
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:3510) clique based evicted nodes at quantum start: (empty)
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:2153) plugin data missing protocol identifier
Jul 03 2019 19:40:20 GMT: DEBUG (clustering): (clustering.c:6777) no cluster merge candidates found
```

