package main

import (
	"context"
	"fmt"
	"go_learn/grpc/stream/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"sync"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/4 16:38
 */

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := simple.NewSimpleServiceClient(conn)
	var globalWg sync.WaitGroup
	globalWg.Add(1)
	{
		stream, err := client.Hello(context.Background())
		if err != nil {
			panic(err)
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			i := 0
			for {
				i++
				err := stream.Send(&simple.HelloReq{
					Title: fmt.Sprintf("[Hello]Client Stream Send msg index:[%d]\n", i),
				})
				if err != nil {
					wg.Done()
					if err == io.EOF {
						fmt.Println("bye...")
						break
					}
					panic(err)
				}
				if i == 5 {
					fmt.Println("apply close.")
					_ = stream.CloseSend()
					wg.Done()
					break
				}
				time.Sleep(time.Second)
			}
		}()
		go func() {
			for {
				msg, err := stream.Recv()
				if err != nil {
					wg.Done()
					if err == io.EOF {
						fmt.Println("bye...")
						break
					}
					panic(err)
				}
				fmt.Printf("[Hello]receive server msg:%s\n", msg.Reply)
			}
		}()
		wg.Wait()
		globalWg.Done()
	}
	globalWg.Add(1)
	{
		stream, err := client.HelloForever(context.Background(), &simple.HelloReq{
			Title: "Hello Server",
		})
		if err != nil {
			panic(err)
		}
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for {
				msg, err := stream.Recv()
				if err != nil {
					wg.Done()
					if err == io.EOF {
						fmt.Println("bye...")
						break
					}
					panic(err)
				}
				fmt.Printf("[HelloForever]receive server msg:%s", msg.Reply)
			}
		}()
		wg.Wait()
		globalWg.Done()
	}
	globalWg.Wait()
}
