package main

import (
	"context"
	"gRPC/weblog"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Log(wr weblog.WebRequest) {

	log.Println(wr.String())

}

func WithLogging(h http.Handler, rpcserver weblog.WebLoggerClient) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(rw, r) // serve the original request

		delay := time.Since(start).Microseconds()

		wr := weblog.WebRequest{Ts: int64(start.Unix()), Method: r.Method, Url: r.RequestURI, Delay: delay}
		//Log(wr)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := rpcserver.Log(ctx, &wr)
		if err != nil {
			log.Fatalf("RPC invocation failed %v", err)
		}

	})
}

func main() {

	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't connect to RPC server %v", err)
	}
	defer conn.Close()
	rpcserver := weblog.NewWebLoggerClient(conn)

	log.Fatal(http.ListenAndServe(":8080",
		WithLogging(http.FileServer(http.Dir("htdocs")), rpcserver)))
}
