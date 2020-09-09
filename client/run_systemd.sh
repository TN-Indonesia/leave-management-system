#!/bin/sh
sudo cp E-Leave_Client.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable E-Leave_Client.service
sudo systemctl start E-Leave_Client.service
sudo journalctl -f -u E-Leave_Client.service