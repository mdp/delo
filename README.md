![Oled](https://cloud.githubusercontent.com/assets/2868/22624086/01516f08-eb26-11e6-9fa9-130884d50ace.jpg)

# Delo
A utility for writing text to an OLED

## Example

With the 128x32 OLED:  
`echo -e "192.168.0.124\nSome status\nAnother Line\nGoodbye World" | delo`

With the 128x64 OLED:  
`echo -e "192.168.0.124\nSome status\nAnother Line\nGoodbye World" | delo -height 64`

## Installation

The goal of 'delo' is to create a binary that can quickly write text
to a OLED, such as the [OLED from 52Pi](http://wiki.52pi.com/index.php/0.96_OLED(English))

`curl https://github.com/mdp/delo/releases/download/v0.0.1/delo_Linux_arm.tar.gz | tar xvz`
`echo -e "Installed" | ./delo`

### Resin.io instruction

**Inside of Docker**  
`CMD modprobe i2c-dev && echo -e "Test output" | ./delo`

### Installation on Linux

**Enable i2c under Raspbian Jessie**  

1. Run sudo raspi-config.
1. Use the down arrow to select 9 Advanced Options
1. Arrow down to A7 I2C.
1. Select yes when it asks you to enable I2C
1. Also select yes when it tasks about automatically loading the kernel module.
1. Use the right arrow to select the 'Finish' button.
1. Select yes when it asks to reboot.

## License
MIT
