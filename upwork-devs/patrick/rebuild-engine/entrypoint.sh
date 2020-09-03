#!/bin/sh

envsubst < /home/glasswall/Config.ini.tmpl > /home/glasswall/Config.ini
glasswallCLI -config=/home/glasswall/Config.ini -xmlconfig=/home/glasswall/Config.xml

