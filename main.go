package main

import (
	"crawler/pkg/clock"
	"crawler/pkg/crawler"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	crawlerUrl := os.Getenv("CRAWLER_URL")
	size := os.Getenv("SIZE")
	tgChannel := os.Getenv("TELEGRAM_CHANNEL")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err == nil {
		msgCooldown := time.Minute * 15
		lastSend := time.UnixMilli(0)
		log.Info().Msg("Got telegram credentials. Start crawler ...")

		go func() {
			for range time.Tick(time.Minute * 1) {
				log.Info().Msg("Is in stock ...")

				cli := crawler.New(crawlerUrl)
				inStock := cli.InStock(size)

				if inStock && lastSend.Add(msgCooldown).Before(clock.Now()) {
					log.Info().Msg("Yes it was in stock.")
					msg := tgbotapi.NewMessageToChannel(tgChannel, fmt.Sprintf("Yay! %s is in stock. GO GO GO - %s", size, crawlerUrl))
					_, err := bot.Send(msg)
					if err != nil {
						log.Error().Err(err)
					}
					lastSend = clock.Now()
				} else {
					log.Info().Msg("No it was not in stock.")
				}
			}
		}()
	}
}
