package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//专门往kafka写日志的模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

func Init(addrs []string, chan_max_size int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	//初始化loadDataChan
	logDataChan = make(chan *logData, chan_max_size)
	//开启后台的goroutine从通道中取数据发往kafka
	go sendToKafka()
	return
}

//给外部保留一个函数，改函数只把日志数据发送到一个内部channel 中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

//往kafka发送日志的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)

		default:
			time.Sleep(time.Millisecond * 50)
		}

	}

}
