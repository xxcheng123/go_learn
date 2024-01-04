package main

import (
	"fmt"
	"go_learn/grpc/stream/pb/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
	"sync"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 11:46
 */

var _ simple.SimpleServiceServer = new(SimpleServiceServer)

type SimpleServiceServer struct {
	simple.UnsafeSimpleServiceServer
}

func (s *SimpleServiceServer) Hello(stream simple.SimpleService_HelloServer) error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		i := 0
		for {
			i++
			err := stream.Send(&simple.HelloResp{
				Reply: fmt.Sprintf("[Hello]Server Stream Send msg index:[%d]", i),
			})
			if err != nil {
				wg.Done()
				panic(err)
			}
			if i == 5 {
				fmt.Println("stop send.")
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
					fmt.Println("client already closed.bye...")
					break
				}
				panic(err)
			}
			fmt.Printf("[Hello]receive client msg:%s", msg.Title)
		}
	}()
	wg.Wait()
	return nil
}
func (s *SimpleServiceServer) HelloForever(req *simple.HelloReq, stream simple.SimpleService_HelloForeverServer) error {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Printf("[HelloForever]receive client msg:%s\n", req.Title)
	go func() {
		i := 0
		for {
			i++
			err := stream.Send(&simple.HelloResp{
				Reply: fmt.Sprintf("[HelloForever]Server Stream Send msg index:[%d]\n", i),
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
				wg.Done()
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	service := new(SimpleServiceServer)
	simple.RegisterSimpleServiceServer(s, service)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("server working...")
	err = s.Serve(listener)
	defer s.Stop()
	if err != nil {
		panic(err)
	}
}
