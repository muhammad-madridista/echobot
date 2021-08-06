package main

import(
  "os"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "github.com/PaulSonOfLars/gotgbot"
  "github.com/PaulSonOfLars/gotgbot/ext"
  "github.com/PaulSonOfLars/gotgbot/handlers"
  "github.com/PaulSonOfLars/gotgbot/handlers/Filters"


)

func main() {
  log:=zap.NewProductionEncoderConfig()
  // log.EncoderLevel=zapcore.CapitalLevelEncoder
  // log.EncoderTime=zapcore.RFC3339TimeEncoder

  logger:=zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(log),os.Stdout,zap.InfoLevel))

  updater,err:=gotgbot.NewUpdater(logger,"1838567969:AAGusqqs0x82uzowrWh5XJT9uuzbKavf2i4")

  if err!=nil{
    logger.Panic("UPDATER FAILED TO START")
    return
  }
  logger.Sugar().Info("UPDATER STARTED SUCCESFULLY")
  updater.StartCleanPolling()
  updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text,echo))
  updater.Idle()
}

func echo(b ext.Bot,u *gotgbot.Update) error {
  b.SendMessage(u.EffectiveChat.Id, u.EffectiveMessage.Text)
  return nil
}
