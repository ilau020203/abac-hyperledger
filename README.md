# ABAC Hyperledger Fabric Chaincode






```bash
peer chaincode invoke -C mychannel -n abac -c '{"Args":["InitLedger", "A", "100", "B", "200"]}'

peer chaincode invoke -C mychannel -n abac -c '{"Args":["InvokeTransfer", "A", "B", "50"]}'

peer chaincode query -C mychannel -n abac -c '{"Args":["Query", "A"]}'
``` 