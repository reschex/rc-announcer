# rc-announcer
[![Go Report Card](https://goreportcard.com/badge/github.com/reschex/rc-announcer)](https://goreportcard.com/report/github.com/reschex/rc-announcer)

## Dev Environment

### config

To configure the rc-announcer, create a `config` file in the project root with the following content:

``` bash
export RC_URL=https://<URL>
export RC_AUTH_TOKEN=
export RC_USER_ID=
export RC_USER_NAME=
export RC_USER_PW=
```

RC_URL is mandatory.
To authenticate to RocketChat, you need either the RC_USER_ID & RC_AUTH_TOKEN or the RC_USER_NAME & RC_USER_PW. 
If both are provided, rc-announcer will use the TOKEN.

### start dev environment

`vagrant up && vagrant ssh`

### build

`make build`

### run

`make run`

## Helm deploy

`helm upgrade  rc-announcer ./k8s/rc-announcer --install --wait --namespace=monitoring --set-string rocketchat.RC_AUTH_TOKEN=????`

## Usage

3 Endpoints are available

### /

POSTing a http request at the service root will echo the request into the log

### /grafana/{channel}

Setup a [Grafana Notification Channel](https://grafana.com/docs/alerting/notifications/#webhook) as:

|||
| -------------|-----------------|
| URL | http://rc-announcer.monitoring:8080/grafana/\<target channel\> |
| Type | webhook |
| Http Method | POST |
| Username/Password| can be left blank |
| Include image | supported but optional |

**Note:** The URL is based on the helm deploy command and assuming that Grafana also runs in K8s, the format is http://helm-deployment-name.namespace:8080/...

### /announce/{channel}

Sends a pre-formated message into the {channel} like so:

`curl -X POST rc-announcer.monitoring:8080/announce/\<target channel\> -d '{ "text": "this is a test" }'`