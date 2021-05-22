package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/beego/beego/core/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka Producer failed, err:", err)
		return
	}
	logs.Debug("init kafka succ")
	return
}

func SendTokafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message faile, err:%v data:%v topic:%v", err, data, topic)
	}
	logs.Debug("send succ,pid:%v offset:%v topic %v", pid, offset, topic)
	return
}
