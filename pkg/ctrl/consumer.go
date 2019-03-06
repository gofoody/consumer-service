package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/gofoody/consumer-service/pkg/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type ConsumerCtrl interface {
	BaseCtrl
	Show(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
}

type consumerCtrl struct {
	consumers map[int]*model.Consumer
	nextID    func() int
}

func NewConsumerCtrl() ConsumerCtrl {
	c := &consumerCtrl{}
	c.init()
	return c
}

func (c *consumerCtrl) Name() string {
	return "consumer controller"
}

func (c *consumerCtrl) Show(rw http.ResponseWriter, r *http.Request) {
	log.Debugf("consumer.Show(), url=%v", r.URL)

	id := mux.Vars(r)["consumerId"]
	conusmer := c.consumers[cast.ToInt(id)]
	if conusmer == nil {
		log.Debugf("consumer with id=%s not found", id)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	payload, err := json.Marshal(conusmer)
	if err != nil {
		log.Errorf("consumer to json failed, error:%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(payload))
}

func (c *consumerCtrl) Create(rw http.ResponseWriter, r *http.Request) {
	log.Debugf("consumer.Create()")

	decoder := json.NewDecoder(r.Body)
	var consumer model.Consumer
	err := decoder.Decode(&consumer)
	if err != nil {
		log.Errorf("json to consumer failed, error:%v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	consumer.ID = c.nextID()
	c.consumers[consumer.ID] = &consumer

	rw.Write([]byte(cast.ToString(consumer.ID)))
}

func (c *consumerCtrl) init() {
	c.nextID = IdGenerator()
	c.consumers = make(map[int]*model.Consumer)

	consumer := &model.Consumer{
		ID:   c.nextID(),
		Name: "Nurali",
	}
	c.consumers[consumer.ID] = consumer

	consumer = &model.Consumer{
		ID:   c.nextID(),
		Name: "John",
	}
	c.consumers[consumer.ID] = consumer
}

func IdGenerator() func() int {
	nextID := 0
	return func() int {
		nextID++
		return nextID
	}
}
