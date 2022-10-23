# Shelly exporter for prometheus and grafana

This exporter exports some metrics from Shelly to prometheus.

![ShellyPIC](https://shelly.cloud/wp-content/uploads/2020/06/Shelly_HT.png)
![ShellyPIC](https://shelly.cloud/wp-content/uploads/2020/06/Shelly_2.5.png)

## Shelly API

Shelly API: https://shelly-api-docs.shelly.cloud/#shelly-h-amp-t

Shelly cloud API access: https://shelly.cloud/documents/developers/shelly_cloud_api_access.pdf

## shelly-metrics.json

Determine "auth_key", "id" and "url" of your device via Shelly cloud API access and update in shelly-metrics.json.
"shelly_name" and "name" can be determined on your own.

    {
    "account":{
        "auth_key": "...",
        "url": "..."
    },
    "products": [
        {
            "type": "ht",
            "export": {
                "isok": "isok",
                "temperature": "data.device_status.tmp.value",
                "humidity": "data.device_status.hum.value",
                "battery": "data.device_status.bat.value",
                "has_update": "data.device_status.update.has_update",
                "firmware": "data.device_status.getinfo.fw_info.fw",
                "mac": "data.device_status.mac",
                "updated": "data.device_status._updated"
            },
            "devices": [
                {
                    "id": "956b54",                  
                    "shelly_name": "S1",
                    "name": "Indoor"
                },
                {
                    "id": "9574a8",                   
                    "shelly_name": "S2",
                    "name": "Outdoor"
                }
            ]
        },
        ...

## Build

    go get github.com/aexel90/shelly_exporter/
    cd $GOPATH/src/github.com/aexel90/shelly_exporter
    go install

## Execution

Usage:

    $GOPATH/bin/shelly_exporter -h

    Usage of ./shelly_exporter:
        -listen-address string
                The address to listen on for HTTP requests. (default "127.0.0.1:9784")
        -metrics-file string
                The JSON file with the metric definitions. (default "shelly-metrics.json")
        -test
                Test configured metrics

### Execute:

    $GOPATH/go/bin/shelly_exporter -metrics-file $GOPATH/go/bin/shelly-metrics.json

### Testing:

    $GOPATH/go/bin/shelly_exporter -metrics-file $GOPATH/go/bin/shelly-metrics.json -test

    Metric: shelly_ht_info
    - Exporter Result:
    - Exporter Result 0:
        - humidity="53"
        - mac="E09806956B54"
        - name="Indoor"
        - isok="true"
        - shelly_name="S1"
        - temperature="22.88"
        - battery="85"
        - has_update="false"
        - firmware=""
        - updated="2021-11-23 19:43:54"
    - Exporter Result 1:
        - has_update="false"
        - firmware="20210323-104951/v1.10.1-gf276b51"
        - isok="true"
        - name="Outdoor"
        - temperature="6.38"
        ...

## Docker

        cp shelly-metrics.json.template shelly-metrics.json
        vi shelly-metrics.json
        docker-compose up -d --build

## Grafana Dashboard

### weather widget

https://weatherwidget.io/

    <!doctype html> <html lang="de">
    <head>
    </head>
    <body>
    <a class="weatherwidget-io" href="https://forecast7.com/de/51d0513d74/dresden/" data-label_1="DRESDEN" data-theme="original" data-highcolor="#88d976" >DRESDEN</a>
    <script>
    !function(d,s,id){
        var js,fjs=d.getElementsByTagName(s)[0];
        if(!d.getElementById(id)){
        js=d.createElement(s);
        js.id=id;
        js.src='https://weatherwidget.io/js/widget.min.js';
        fjs.parentNode.insertBefore(js,fjs);
        setInterval('__weatherwidget_init()', 1800000) <!-- refresh widget every 30 minutes (1800000 milliseconds): -->
        }
    }(document,'script','weatherwidget-io-js');
    </script>
    </body>
    </html>

### dashboard

Grafana-ID: 13739

https://grafana.com/grafana/dashboards/13739

![Grafana](https://raw.githubusercontent.com/aexel90/shelly_exporter/main/grafana/screenshot.jpg)
![Grafana Rollo](https://raw.githubusercontent.com/aexel90/shelly_exporter/main/grafana/screenshot_rollo.jpg)