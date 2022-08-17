package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	//创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		fmt.Println("fail to start consumer", err)
	}
	//根据topic获取所有的分区列表
	partitionList, err := consumer.Partitions("weatherStation")
	if err != nil {
		fmt.Println("fail to get list of partition,err:", err)
	}
	fmt.Println(partitionList)
	//遍历所有的分区
	for p := range partitionList {
		//针对每一个分区创建一个对应分区的消费者
		pc, err := consumer.ConsumePartition("weatherStation", int32(p), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d Offse:%d Key:%v Value:%s \n",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
