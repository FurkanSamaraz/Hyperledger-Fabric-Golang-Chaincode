# Hyperledger-Fabric-Golang-Chaincode

Indiriniz:
Curl
Docker 
Docker-compose
Go
Node.js and npm
Expressjs
Python


Bir klasör açınız ve hyperlender fabric indiriniz.
 curl -sSL https://bit.ly/2ysbOFE | bash -s
 
 
- Test ağına git
cd fabric-samples/test-network

Hali hazır da çalışan bir test ağınız var ise temizleyin.
./network.sh down

— docker' i temizleyin
docker rm -f $(docker ps -aq)

- Ağı Bağlatın
network.sh Kabuk komut dosyasını kullanarak Fabric test ağını başlatırız -ca ise sertifikalar içindir.
./network.sh up createChannel -c mychannel -ca

- Chaincode Dağıtın
./network.sh Ardından zincir kod adı ve dil seçenekleriyle çağırarak akıllı sözleşmeyi içeren zincir kod paketini dağıtalım. asset-transfer-basic/chaincode-go/ yoluna giderek kontratı yazınız. yazmaz iseniz hazırda var olan örnek chaincode kullanılacaktır.

./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go


Örnek uygulama çalıştır ve chaincode(kontrata) bağlan
cd asset-transfer-basic/application-gateway-go

go run 
