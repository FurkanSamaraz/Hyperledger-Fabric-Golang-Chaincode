# Hyperledger-Fabric-Golang-Chaincode

![image](https://user-images.githubusercontent.com/92402372/171406041-ee5e1eca-3479-4cce-a726-59bf2ab24af1.png)

# Hyperledger Nasıl Çalışır?
Hyperledger bir kriptopara projesi değildir. Dolayısıyla Fabric ağı kriptopara transferi için değildir. Ağda varlıklar(asset) vardır. Bu varlıkların maddi değeri olabilir. Yalnızca veri barındıran varlıklar da olabilir.

Fabric’in çalışma prensiplerini aşağıdaki maddeler ile açıklayabiliriz:

İşlemler, chaincode’lar(smart contract) ile güvence altına alınır ve tüm katılımcılar Docker konteynırları çalıştırarak ağa bağlanır.

İşlemler kriptopara olmadan gerçekleştirilir.

Tüm işlemler gizli ve güvenlidir. Bilgiler, yalnızca ağdaki katılımcıların fikir birliği(consensus) ile güncellenir.

Tüm işlemlerin içeriği kriptoludur. Böylece içerik tüm kullanıcılar tarafından görüntülenemez.

Katılımcılar ağa erişim sağlamak için üyelik servislerine kimlik kanıtı yapmak zorundadır.

# Indiriniz:
Curl
Docker 
Docker-compose
Go
Node.js and npm
Expressjs
Python

# - Hyperledger install
Bir klasör açınız ve hyperlender fabric indiriniz.

curl -sSL https://bit.ly/2ysbOFE | bash -s
 
 
# - Test ağına git
cd fabric-samples/test-network

# - Hali hazır da çalışan bir test ağınız var ise temizleyin.
./network.sh down

# - docker' i temizleyin
docker rm -f $(docker ps -aq)

# - Ağı Bağlatın
network.sh Kabuk komut dosyasını kullanarak Fabric test ağını başlatırız -ca ise sertifikalar içindir.

./network.sh up createChannel -c mychannel -ca

# - Chaincode Dağıtın
./network.sh Ardından zincir kod adı ve dil seçenekleriyle çağırarak akıllı sözleşmeyi içeren zincir kod paketini dağıtalım. asset-transfer-basic/chaincode-go/ yoluna giderek kontratı yazınız. yazmaz iseniz hazırda var olan örnek chaincode kullanılacaktır.

./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go


# - Örnek uygulama çalıştır ve chaincode(kontrata) bağlan
cd asset-transfer-basic/application-gateway-go

go run 


# - Postman 
# - Init
http://localhost:8080/api/init

![Ekran Resmi 2022-06-01 15 33 21](https://user-images.githubusercontent.com/92402372/171405598-8495a9ac-f82b-4675-8310-f0df3ddbe77b.png)

# - Get
http://localhost:8080/api/get

![Ekran Resmi 2022-06-01 15 33 34](https://user-images.githubusercontent.com/92402372/171405678-f12be428-2457-4f69-8fc6-07a4b79e8d5c.png)

# - Create
http://localhost:8080/api/create

![Ekran Resmi 2022-06-01 15 33 47](https://user-images.githubusercontent.com/92402372/171405713-e893556e-f43c-4b83-9d39-bbc419730188.png)

# - Main Modules in Hyperledger Fabric Node SDK
Bir sonraki bölümde, bazı kodların üzerinden geçeceğiz. Bu nedenle, ön koşul olarak, Hyperledger Fabric SDK'da kullanılan modüller, çeşitli sınıflar ve yöntemler hakkında temel fikre sahip olmamız gerekir. Bu nedenle bu bölümün atlanması önerilmez. SDK üç ana modülden oluşur

In the next section, we are going to walk through some code. so as a pre-requisite, we need to have basic idea on modules, various classes, and methods used in Hyperledger Fabric SDK. So skipping this section is not recommended.
SDK consists of three major modules

# - 1.fabric-network:
Bu modül, konuşlandırılan zincir koduyla etkileşim için üst düzey bir API sağlar. Etkileşim, işlemlerin ve sorguların sunulmasını içerir.

This module provides a high-level API for interacting with the deployed chaincode. Interaction includes submission of transactions and queries.

# - 2.fabric-ca-client:
Bu modül, kullanıcı kayıtları ve ağ üzerinde yeniden kayıt gibi çeşitli sertifika işlemleri sağlar.

This module provides various certification operations like user registrations and enrollments, re-enrollments on the network.

# - 3.fabric-client:
Bu modül, kanal oluşturma, eşleri kanala birleştirme, zincir kodunu yükleme ve başlatma, işlemleri gönderme ve zincir kodunu sorgulama, işlem ayrıntılarını sorgulama, blok yüksekliği, kurulu zincir kodları ve daha fazlası gibi ağ düzeyinde işlemler sağlar. yapı-istemci birçok işlemi gerçekleştirebilse de, istemci uygulamalarının Yapı ağına dağıtılan akıllı sözleşmelerle etkileşim kurması için yapı-ağı önerilen API'dir. çünkü istemci uygulamaları için anahtar sınıfların çoğu yalnızca yapı ağından devralınabilir. Fabric-network Fabric-network'deki Anahtar Sınıflar ve yöntemler üç ana sınıfa sahiptir.

This module provides network-level operations like creating channels, joining peers to the channel, installing and Instantiation of chaincode, submitting transactions and querying chaincode, querying transaction details, block height, installed chaincodes and many more.
even though fabric-client is capable of doing many operations, fabric-network is recommended API for client applications to interact with smart contracts deployed to Fabric network. because the majority of key classes for client applications are only inheritable from fabric-network.
Key Classes and methods in fabric-network
fabric-network has three main classes

# - 1. Gateway
Yapı ağı modülündeki Ağ Geçidi Sınıfı, çalışan yapı ağlarına bağlanmak ve bunlarla etkileşim kurmak için kullanılır. bu sınıf çeşitli yöntemler içerir. onlar,

Gateway Class in fabric-network module is used to connect and interact with running fabric networks. this class includes various methods. those are,

# a. connect:

Bu yöntem, mevcut kullanıcı veya Yönetici kimliği kullanılarak bağlantı profilinde tanımlanan eşlere ve bunların IP adreslerine dayalı olarak çalışan yapı ağını birbirine bağlar.

This method connects running fabric-network based on the peers and their IP addresses defined in connection profile using existed user or Admin identity.

# b. disconnect:

Bu yöntem, çalışan yapı ağı bağlantısını keser ve önbelleği temizler.

This method disconnects from running fabric-network and cleans up the cache.

# c. getClient:

Bu yöntem, geçerli kayıtlı istemci ayrıntılarını bir nesne olarak döndürür.

This method returns current registered client details as an object.

# d. getNetwork:

Bu yöntem, ağdaki belirli bir kanalla iletişim kurar.

This method communicates with a specified channel in the network.

# e. getContract:

getContract yönteminin, bağlantı profilinde tanımlanan ağın üstündeki kanala dağıtılan belirli bir zincir koduna erişimi olacaktır.

getContract method will have access to a particular chaincode deployed to channel on top of the network defined in the connection profile.

# f. sendTransaction:

sendTransaction yöntemi, belirli bir zincir kodu yöntemini gönderir ve eşlere (onaylayıcılara) itiraz eder.

submitTransaction method will submit a specified chaincode method and args to the peers(endorsers).

# g. evaluateTransaction: 

evaluateTransaction, HTTP isteklerinde GET yöntemine benzer, yalnızca defter durumunu okuyabilir ve chaincode'da sorgulama yöntemleri için kullanılır.

evaluateTransaction is similar to GET method in HTTP requests, it can only read the ledger state and used for query methods in chaincode.

# - 2. FileSystemWallet

Bu sınıf, kumaş dosya sisteminde kalıcı olan bir Identity cüzdan uygulamasını tanımlar. Bu sınıf, var olan, içe aktarılan, silinen bazı yaygın yöntemleri içerir.

This class defines the implementation of an Identity wallet that persists to the fabric file system. this class includes some common methods exists, import, delete.

# a. exists:

Bu yöntem, sağlanan kimliğin dosya sisteminde olup olmadığını kontrol eder.

This method checks whether provided identity exists in the file system or not.

# b. import:

Bu yöntem, oluşturulan PKI ve x509 sertifikalarını ve anahtarlarını, katılımcının kimliği altında dosya sistemi cüzdanına aktarır.

This method imports generated PKI and x509 certificates and keys into the filesystem wallet under the identity of participant.

# c. delete:

Bu yöntem, belirli bir kullanıcının kimliğini dosya sistemi cüzdanından siler.

This method deletes the identity of a particular user from the filesystem wallet.

# - 3. X509WalletMixin

Temel olarak CA, PKI formatında (ortak anahtar altyapısı) kayıt sertifikaları sağlar. Bu sınıf, PKI kullanıcı Sertifikaları kullanılarak X.509 kimlik bilgisi modelinde kimlik oluşturmak için kullanılır. Bu sınıf, Kimlik oluşturmak için createIdentity() yöntemini içerir. Bu X.509 sertifikası, ağdaki işlemleri imzalamak için kullanılır. Fabric-ca-client içindeki anahtar yöntemler Fabric-ca-client CA işlemleri için kullanılan birkaç yaygın yönteme sahiptir. bunlar kayıt ol, kayıt ol, yeniden kayıt ol.

Basically, CA provides enrollment certificates in PKI format(public key infrastructure). This class is used for creating identity in X.509 credential model using PKI user Certificates. This class includes createIdentity() method for creating Identity. This X.509 certificate is used for signing transactions in the network.
Key methods in fabric-ca-client
fabric-ca-client has few common methods used for CA operations. these are register, enroll, re-enroll.

# a. register:

Bu yöntem, yeni Katılımcıları kaydetmek için kullanılır. kayıt başarılı olduğunda, kullanıcı sırrını döndürür. Bu sır, kayıt olurken sağlanmalıdır.

This method is used for registering new Participants. when registration is successful, it returns user secret. This secret needs to be provided while enrolling.

# b. enroll:

Bu yöntem, kayıtlı Katılımcıları ağa kaydetmek için kullanılır. Kayıt olmak için, kullanıcının önce kayıt olması gerekir. kayıt başarılı olursa, bu yöntem kullanıcının PKI tabanlı Sertifikasını ve Özel anahtarını döndürür.

This method is used for enrolling registered Participants in the network. in order to enroll, the user must be registered first. if enrollment succeeded, this method will return PKI based Certificate and Private key of the user.

# c. reenroll:

Bir sertifikanın süresinin dolduğu veya güvenliğinin ihlal edildiği durumlar olabilir (bu nedenle iptal edilmesi gerekir). Yani bu, yeniden kayıt devreye girdiğinde ve bu yöntemi kullanarak yeni sertifikalar almak için CA ile aynı kimliği tekrar kaydettirirsiniz.

There could be cases when a certificate expires or gets compromised (so it has to be revoked). So this is when re-enrollment comes into the picture and you enroll the same identity again with the CA to get new certificates using this method.
