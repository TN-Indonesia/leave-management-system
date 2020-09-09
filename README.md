# E-Leave

## Prerequisites to Run Client With Systemd
* Install server
    ```
    sudo npm install -g serve
    ```
* Open project directory (client)
* Build project (npm run build)
    ```
    npm run build
    ```

## How to Run With Systemd
* Open project directory (server / client)
* Exec run_systemd.sh
    ```
    ./run_systemd.sh
    ```

## How to Stop Running Systemd Services
* Open project directory (server / client)
* Exec stop_systemd.sh
    ```
    ./stop_systemd.sh
    ```

# How to Read Log
* Client :
    ```
    sudo journalctl -f -i E-Leave_Client
    ```
* Server :
    ```
    sudo journalctl -f -i E-Leave_Server
    ```