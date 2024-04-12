package main

import (
	"bytes"
	"fmt"
	"log"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
	if err != nil {
		log.Fatal("Erro ao criar socket:", err)
	}
	defer syscall.Close(fd)

	for {
		buf := make([]byte, 65536)
		n, _, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.Fatal("Erro ao receber do socket:", err)
		}
		if n > 0 {
			processPacket(buf[:n])
		}
	}
}

func processPacket(packet []byte) {
	if len(packet) > 14 {
		ethType := ntohs(uint16(packet[12])<<8 | uint16(packet[13]))
		if ethType == 0x0800 && len(packet) >= 34 {
			ipHeaderLen := int(packet[14]&0x0F) * 4
			if packet[23] == 0x06 {
				tcpHeaderStart := 14 + ipHeaderLen
				if len(packet) > tcpHeaderStart {
					tcpHeaderLen := int((packet[tcpHeaderStart+12] >> 4) * 4)
					payloadStart := tcpHeaderStart + tcpHeaderLen
					if len(packet) > payloadStart && bytes.Contains(packet[payloadStart:], []byte("SSH-")) {
						fmt.Println("Conexão SSH detectada")
					}
				}
			}
		}
	}
}

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func ntohs(i uint16) uint16 {
	return htons(i)
}
