package pkg

import (
	"encoding/binary"
	"encoding/hex"
	"net"

	"github.com/rs/zerolog/log"
)

func getMetadataRequest() []byte {
	buf := make([]byte, 1000) // 1000 is a random number large enough to hold the request
	offset := 4               // the first 4 bytes reserved for length

	// API key
	binary.BigEndian.PutUint16(buf[offset:], uint16(3))
	offset += 2

	// API version
	binary.BigEndian.PutUint16(buf[offset:], uint16(0))
	offset += 2

	// correlation id
	binary.BigEndian.PutUint32(buf[offset:], uint32(1))
	offset += 4

	// client id
	clientId := "test"
	binary.BigEndian.PutUint16(buf[offset:], uint16(len(clientId)))
	offset += 2
	copy(buf[offset:], []byte(clientId))
	offset += len(clientId)

	// metadata request
	binary.BigEndian.PutUint32(buf[offset:], uint32(0))
	offset += 4

	// write size
	size := offset - 4
	binary.BigEndian.PutUint32(buf[0:], uint32(size))

	buf = buf[:offset]

	log.Info().Msgf("request: %v", buf)
	log.Info().Msgf("request: %v", hex.EncodeToString(buf))
	return buf
}

type Broker struct {
	nodeId int32
	host   string
	port   int32
}

type MetaData struct {
	totalSize     int32
	correlationId int32
	brokers       []Broker
	// TODO: decode topic_metadata as well
}

func DecodeMetadataResponse(buf []byte) MetaData {
	offset := 0
	var data MetaData

	// get total size
	data.totalSize = int32(binary.BigEndian.Uint32(buf[offset:]))
	offset += 4

	// correlation id
	data.correlationId = int32(binary.BigEndian.Uint32(buf[offset:]))
	offset += 4

	// broker length
	brokerNum := int32(binary.BigEndian.Uint32(buf[offset:]))
	offset += 4
	for i := 0; i < int(brokerNum); i++ {
		var broker Broker

		broker.nodeId = int32(binary.BigEndian.Uint32(buf[offset:]))
		offset += 4

		// host
		strLen := int16(binary.BigEndian.Uint16(buf[offset:]))
		offset += 2
		broker.host = string(buf[offset : offset+int(strLen)])
		offset += int(strLen)

		// port
		broker.port = int32(binary.BigEndian.Uint32(buf[offset:]))
		offset += 4
		data.brokers = append(data.brokers, broker)
	}

	return data
}

func GetMetaData() {
	request := getMetadataRequest()
	serverAddr := AppConfig.BootstrapBrokers[0]
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Fatal().Msgf("Resolve server address %s failed", serverAddr)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal().Msgf("Dial failed: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write(request)
	if err != nil {
		log.Fatal().Msgf("Write to server failed: %v", err)
	}

	response := make([]byte, 10000) // big enough

	_, err = conn.Read(response)
	if err != nil {
		log.Fatal().Msgf("Write to server failed: %v", err)
	}

	metaData := DecodeMetadataResponse(response)

	log.Info().Msgf("kafka metadata: %v", metaData)
}
