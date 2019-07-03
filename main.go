package main

import (
	"flag"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"

	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	args := struct {
		host      *string
		port      *int
		namespace *string
		elements  *int
	}{
		flag.String("host", "localhost", "Aerospike host"),
		flag.Int("port", 3000, "Aerospike port"),
		flag.String("namespace", "event", "Aerospike namespace"),
		flag.Int("elements", 1000, "Map elements"),
	}

	flag.Parse()

	client, err := as.NewClient(*args.host, *args.port)
	if err != nil {
		logrus.WithError(err).Fatal("Error creating Aerospike client")
	}

	k, err := as.NewKey(*(args.namespace), "event", "key-1")
	if err != nil {
		logrus.WithError(err).Error("Error creating key")
		return
	}

	var operations []*as.Operation

	for index := 0; index < *args.elements; index++ {

		id := uuid.NewV1().Bytes()

		blob := make([]byte, 64)
		rand.Read(blob)

		operations = append(
			operations,
			as.MapPutOp(as.DefaultMapPolicy(), "bin-1", string(id), blob),
		)
	}

	policy := as.NewWritePolicy(0, 0)
	policy.SendKey = true

	_, err = client.Operate(
		policy,
		k,
		operations...)

	if err != nil {
		logrus.WithError(err).Error("Error writing map")
	}
}

