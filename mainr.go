package main
import (
		"fmt"
	    "github.com/streadway/amqp"
)
func main(){
	cha:="Hello"
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare(cha,false,false,false,false,nil)
	msgs, _ := ch1.Consume(q1.Name,"",true,false,false,false,nil)
	go func() {
		for d := range msgs {
			time.Sleep(500*time.Millisecond)
			log.Printf("Received: %s", d.Body)
		}
	}()
}

}