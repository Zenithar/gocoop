# GoCoop

![Coop](https://github.com/fallais/gocoop/blob/master/assets/coop.png)

**GoCoop** is a tool written in **Go** and **Angular** that helps you to manage your chicken coop !

## Objectives

The main objective is to protect the chickens against the **hungry fox** or the **greedy weasel**. To do so, we need to automaticaly open and close the door of the chicken coop, with two options :

- At a fixed time (for example *08h30*)
- Based on the sunset and sunrise (for example *30min after sunrise*)

> If you worry about using the sun based condition, be sure that the chickens always go to sleep at sunset. As the sentence says : **go to bed with the chickens**.

## Components

### The motor

I use the `Nextrox 37mm 12V 15RPM`. I chose this motor because of its torque : **250 N*cm**

![Nextrox](https://github.com/fallais/gocoop/blob/master/assets/nextrox.jpg)

### The motor driver

I use the `L293D`. It is capable of handling two motors.

![L293D](https://github.com/fallais/gocoop/blob/master/assets/L293D.jpg)

I will be using the following pins :

- Enable 1 : To enable the motor, to make it turn.
- Input 1 : To set the rotation way
- Input 2 : To set the rotation way
- Output 1 : To motor
- Output 2 : To motor
- All the GND pins

### The GPIO pins

Here are the GPIO pins that are used :

- 2 : Power (5V)
- 6 : Ground
- 23 : connected to the **Input 1** of the **L293D**
- 24 : connected to the **Input 2** of the **L293D**
- 25 : connected to the **Enable 1** of the **L293D**

The GPIO of my **Raspberry 3 B+** are located as follow :

![GPIO](https://github.com/fallais/gocoop/blob/master/assets/gpios.png)

### The global schema

All compoents are now connected together with a breadboard.

![Schema](https://github.com/fallais/gocoop/blob/master/assets/schema.png)

## Is it tested ?

Sure, I tried to do my best to add package tests because chickens deserv the best ! Moreover, I have been using it for more than one year at home, it has never failed since the begining.

## Interface

It also comes with an interface built with **Angular** to manage the coop.

Protected by a login.  
![Login](https://github.com/fallais/gocoop/blob/master/assets/login.png)

With a dashboard.  
![dashboard](https://github.com/fallais/gocoop/blob/master/assets/dashboard.png)

## Installation

First, I installed **raspbian-lite** on the Raspberry.  
Then, I updated all the packages.  
Finally, I installed **docker** with the convenience script.  
And, a basic **logrotate** configuration.

The Raspberry is ready to run the Docker container.

## Usage

### Docker

Deploy with a `docker-compose`.

```yaml
version: "3"

services:
  redis:
    image: redis:alpine
    container_name: redis
    restart: always
    networks:
      main:
        aliases:
          - redis

  gocoop:
    image: fallais/gocoop
    container_name: gocoop
    restart: always
    volumes:
      - /data/docker/config.yml:/usr/bin/config.yml
      - /sys:/sys
    networks:
      main:
        aliases:
          - gocoop
    
    caddy:


networks:
  main:
    driver: bridge
```

#### Parameters

- config.yml : mandtory
- port : mandatory

### Configuration

The configuration file must look like this below :

```yaml
general:
  gui_username: admin
  gui_password: admin
  private_key: keykey
  redis_host: localhost:6379
  redis_password:  
coop:
  latitude: 43.388352
  longitude: 1.277914
  opening:
    mode: "time_based"
    value: "08h00"
  closing:
    mode: "sun_based"
    value: "30m"
door:
  openening_duration: "65s"
  closing_duration: "60s"
cameras:
  outdoor: http://185.39.101.26/mjpg/video.mjpg
  indoor: http://187.157.229.132/mjpg/video.mjpg
```

#### Modes and values

Two modes are available :

- Time based (fixed time) : `time_based`
  - Value must be something like **HHhMM** : `08h00`
- Sun based (based on the sunrise and sunset) : `sun_based`
  - Value must be a valid Golang duration : `45m`

## Production ready

It is actually also used by a friend who have **160 chickens**. Below an overview of how it looks like.

![Door](https://github.com/fallais/gocoop/blob/master/assets/door.jpg)

## Licence

I do not set any licence as of now. Please do not use code for commercial purpose, instead, contact me and we could work together. Chickens are against money problems..

### Ideas

- [ ] Use **kill switches** to get rid of opening and closing durations
- [ ] **Water warming** in case of low temperature
- [ ] Egg counting with **OpenCV**
- [ ] Omelet maker (no, just joking)
