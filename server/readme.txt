1. before bee run , call -> source .env

2. Migrate use  
bee migrate -driver=postgres -conn="postgres://postgres:root@127.0.0.1:5432/e_leave?sslmode=disable"

3. change server client /home/tnisindo/Documents/work/e_leave/src/client/src/store/Actions/types.js
// export const ROOT_API = 'http://35.197.155.131:8080'; // GCP
export const ROOT_API = 'http://localhost:8080'; // local