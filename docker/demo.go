package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
) /*
* @Author: HanYaJun
* @Date: 2019/12/12 3:43 下午
* @Name: docker
* @Function:
 */

func main() {
	cli, err := client.NewClient("tcp://192.168.1.230:2375", "", nil, nil)
	if err != nil {
		panic(err)
	}
	//resp, err := cli.ContainerCreate(context.Background(), &container.Config{
	//
	//	Image: "registry.cn-hangzhou.aliyuncs.com/hanyajun/resource-manage-svc:dev",
	//	//WorkingDir: "/testfilespoutaa",
	//	/*Volumes: map[string]struct{}{
	//		"/home/deepglint/data_0": {},
	//	},*/
	//	//Cmd: []string{"./resource-manage-svc", "-addr", addr, "-file", fileName},
	//	//		Shell:        []string{"pwd"},
	//	ExposedPorts: map[nat.Port]struct{}{
	//		"8000": {},
	//	},
	//	AttachStdin:  true,
	//	AttachStdout: true,
	//	AttachStderr: true,
	//	Tty:          true,
	//}, &container.HostConfig{
	//
	//	PortBindings: map[nat.Port][]nat.PortBinding{
	//		"8000": []nat.PortBinding{
	//			nat.PortBinding{
	//				HostPort: "8093",
	//			},
	//		},
	//	},
	//	//Binds: []string{"/home/deepglint/data_0/111.264:/testfilespoutaa/data/111.264", "/home/deepglint/data_0/test.h264:/testfilespoutaa/data/test.h264"},
	//}, nil, "test")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.ID)
	//startContainer(resp.ID, cli)
	top, err := cli.ContainerTop(context.Background(), "5341f3a9f82bae1a7c215d60aa75997989c8245d15701e54b90e13bb89594884", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(top)

}

func startContainer(containerID string, cli *client.Client) {
	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	log.Println(err)
	if err == nil {
		fmt.Println("容器", containerID, "启动成功")
	}
}
