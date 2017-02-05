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

One line install and Usage (requires i2c to be enabled)

`curl https://github.com/mdp/delo/releases/download/v0.0.1/delo_Linux_arm.tar.gz | tar xvz && echo -e "Installed" | ./delo`

### Resin.io i2c setup

`modprobe i2c-dev`

**Inside of Docker**

`CMD modprobe i2c-dev && echo -e "Test output" | ./delo`

### i2c on Raspbian Jessie

1. Run sudo raspi-config.
1. Use the down arrow to select 9 Advanced Options
1. Arrow down to A7 I2C.
1. Select yes when it asks you to enable I2C
1. Also select yes when it tasks about automatically loading the kernel module.
1. Use the right arrow to select the 'Finish' button.
1. Select yes when it asks to reboot.

## License - MIT
Copyright (c) 2017 Mark Percival

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
