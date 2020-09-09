#!/bin/sh
sudo cp E-Leave_Server.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable E-Leave_Server.service
sudo systemctl start E-Leave_Server.service
sudo journalctl -f -u E-Leave_Server.service