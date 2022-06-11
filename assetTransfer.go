package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID         = "Org1MSP"
	cryptoPath    = "../../test-network/organizations/peerOrganizations/org1.example.com"
	certPath      = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	keyPath       = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath   = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint  = "localhost:7051"
	gatewayPeer   = "peer0.org1.example.com"
	channelName   = "mychannel"
	chaincodeName = "basic"
)

var getall GetAll
var cre Create
var initt InitLed

var contract *client.Contract
var network *client.Network
var gateway *client.Gateway
var now = time.Now()
var assetId = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

type Create struct {
	Ad    string `json:"ad"`
	Soyad string `json:"soyad"`
	Yas   string `json:"yas"`
}
type InitLed struct {
	InitLed string `json:"initledger"`
}
type GetAll struct {
	Get string `json:"get"`
}

func ApiInit(c *fiber.Ctx) error {
	network = gateway.GetNetwork(channelName)
	contract = network.GetContract(chaincodeName)
	fmt.Println("initLedger:")
	c.Response().Header.Set("Content-Type", "application/json")
	apiInit := c.FormValue("initleger")
	initt.InitLed = apiInit

	initLedger(contract)
	fmt.Printf("*** ISLEM BASARILI TAMAMLANDI.\n")
	byte, _ := json.MarshalIndent(initt, "", "\t")
	c.JSON(byte)
	return c.JSON(byte)
}

func ApiGet(c *fiber.Ctx) error {
	network = gateway.GetNetwork(channelName)
	contract = network.GetContract(chaincodeName)
	fmt.Println("getallassets:")
	c.Response().Header.Set("Content-Type", "application/json")

	apiget := c.FormValue("get")
	getall.Get = apiget
	getAllAssets(contract)
	fmt.Printf("*** ISLEM BASARILI TAMAMLANDI.\n")
	byte, _ := json.MarshalIndent(getall, "", "\t")

	return c.JSON(byte)
}

func ApiCreate(c *fiber.Ctx) error {
	network = gateway.GetNetwork(channelName)
	contract = network.GetContract(chaincodeName)
	fmt.Println("create:")

	c.Response().Header.Set("Content-Type", "application/json")
	ad := c.FormValue("ad")
	yas := c.FormValue("yas")
	soyad := c.FormValue("soyad")
	cre.Ad = ad
	cre.Soyad = soyad
	cre.Yas = yas

	createAsset(contract)
	fmt.Printf("*** ISLEM BASARILI TAMAMLANDI.\n")
	byte, _ := json.MarshalIndent(cre, "", "\t")

	return c.JSON(byte)
}

func main() {
	log.Println("============ application-golang starts ============")

	// gRPC istemci bağlantısı, bu uç noktaya yönelik tüm Ağ Geçidi bağlantıları tarafından paylaşılmalıdır.
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Belirli bir istemci kimliği için bir Ağ Geçidi bağlantısı oluşturun
	gateway, _ = client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Farklı gRPC çağrıları için varsayılan zaman aşımları
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)

	defer gateway.Close()

	network = gateway.GetNetwork(channelName)
	contract = network.GetContract(chaincodeName)
	app := fiber.New()

	app.Post("/api/init", ApiInit)
	app.Post("/api/get", ApiGet)

	app.Post("/api/create", ApiCreate)

	//r := mux.NewRouter()

	//r.HandleFunc("/api/init", apiIinit).Methods("POST")
	//r.HandleFunc("/api/get", apiGet).Methods("POST")
	//r.HandleFunc("/api/create", apiCreate)
	app.Listen(":8080")
	//http.ListenAndServe(":8080", r)
}

//********************************************************************************************************************************
//newGrpcConnection, Ağ Geçidi sunucusuna bir gRPC bağlantısı oluşturur.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("GRPC BAGLANTISI OLUSTURULAMADI.: %w", err))
	}

	return connection
}

// newIdentity, bir X.509 sertifikası kullanarak bu Ağ Geçidi bağlantısı için bir istemci kimliği oluşturur.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("SERTIFIKA DOSYASI OKUNAMADI.: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign, özel bir anahtar kullanarak bir mesaj özetinden dijital imza oluşturan bir işlev oluşturur.
func newSign() identity.Sign {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("OZEL ANAHTARLI DOSYA OKUNAMADI.: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("OZEL ANAHTARLI DOSYA OKUNAMADI.: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

//**********************************************************************************************************************************
// Defter de varlık oluşturma.
func initLedger(contract *client.Contract) {
	fmt.Printf("ISLEM GONDER: InitLedger ISLEVI, DEFTER DE ILK VARLIK KUMESINI DONDURUR.\n")

	_, err := contract.SubmitTransaction(initt.InitLed)
	if err != nil {
		panic(fmt.Errorf("ISLEM GONDERILEMEDI!!: %w", err))
	}

	fmt.Printf("*** ISLEM BASARIYLA TAMAMLANDI.\n")
}

// Defter durumunu sorgulamak için bir işlemi değerlendirin.
func getAllAssets(contract *client.Contract) {
	fmt.Println("ISLEM DEGERLENDIR GetAllAssets ISLEVI DEFTERDE Kİ TUM MEVCUT VARLIKLARI LISTELER.")

	evaluateResult, err := contract.EvaluateTransaction(getall.Get)
	if err != nil {
		panic(fmt.Errorf("ISLEM GONDERILEMEDI!!: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** SONUC:%s\n", result)
}

//Deftere taahhüt edilene kadar bloke ederek eşzamanlı olarak bir işlem gönderin.
func createAsset(contract *client.Contract) {
	fmt.Printf("GONDERIM ISLEMI KIMLIK,AD,YAS,SOYAD DEGISKENLERI ILE YENI VARLIK OLUSTURULUR. \n")

	_, err := contract.SubmitTransaction("CreateAsset", assetId, cre.Ad, cre.Yas, cre.Soyad)
	if err != nil {
		panic(fmt.Errorf("ISLEM GONDERILEMEDI!!: %w", err))
	}

	fmt.Printf("*** ISLEM BASARIYLA TAMAMLANDI.\n")
}

//JSON verilerini biçimlendir
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, " ", ""); err != nil {
		panic(fmt.Errorf("JSON AYRISTIRILAMADI.: %w", err))
	}
	return prettyJSON.String()
}
