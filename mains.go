package main
import( 
		"fmt"
		"github.com/streadway/amqp"
	)
func main(){
	body:="Data"
	cha:="Hello"
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare(cha,false,false,false,false,nil)
	err:= ch1.Publish("",q1.Name,false,false,amqp.Publishing{ContentType: "text/plain",Body: []byte(body)})
	if err!=nil{
		fmt.Println(err)
	}
	
}