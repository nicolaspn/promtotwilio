# Prometheus alert with text message

This is a simple and stupid program that will receive webhooks from [Prometheus](https://prometheus.io/) to send them as text message (using [Twilio](https://www.twilio.com/)) with the summary of the alert.


## Configuration

It needs 5 arguments:

- `PORT` - Listener port
- `SID` - Twilio Account SID
- `TOKEN` - Twilio Auth Token
- `SENDER` - Phone number managed by Twilio (friendly name)


## API

`/`: ping promtotwilio application. Returns 200 OK if application works fine.

`/send?receiver=<rcv>`: send Prometheus firing alerts from payload to a rcv if specified, or to default receiver, represented by RECEIVER environment variable. If none is specified, status code 400 BadRequest is returned.

## Test it

To send test sms to a phone +zxxxyyyyyyy use the following command (please notice `%2B` symbols, representing a url encoded `+` sign)

```bash
$ curl -H "Content-Type: application/json" -X POST -d \
'{"version":"2","status":"firing","alerts":[{"annotations":{"summary":"Server down"},"startsAt":"2016-03-19T05:54:01Z"}]}' \
http://localhost:9090/send?receiver=%2Bzxxxyyyyyyy
```

## Configuration example

fmt.Println("Usage:", os.Args[0], "PORT", "SID", "TOKEN", "SENDER")

./promtotwilio_app 9090 XXXX AAA 0623456789


Here's the AlertManager config where `sms` is the host 

```yml
route:
  receiver: 'admin'

receivers:
- name: 'admin'
  webhook_configs:
  - url: 'http://sms:9090/send'
```
