#!/bin/bash

curl -s https://raw.githubusercontent.com/FGRibreau/mailchecker/master/list.json | grep -v "  //"
