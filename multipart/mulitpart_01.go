package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"
)

func main() {
	url := "https://pvtcuiapi.lenovo.com.cn/v20180430/events"
	data := `{"event": {"header":{"namespace": "SpeechRecognizer","name": "TranslateText","messageId": "34471018-2c1e-4610-8124-2d142b2wwwww","dialogRequestId": "b13dba86-232b-4ecf-9e34-e432d142b2wwwww"        },        "payload":        {"text": "新型冠状病毒引起了世界范围大流行","language":{"source": "auto","target":"en-US"}}}}"`
	token := "eyJlbmMiOiJBMTI4R0NNIiwiYWxnIjoiZGlyIn0..qpiJwCj3A8lFtNAD.J9HfsIY8mgxM3mpH4LwoeOkeTXuaYs8APrMeDkaYMjGghgjJcELpCP7Wg2tku9xCmJ3mkhgTE1DIq8INEzPPI7HU94vM6gV66LDopVQG9FKc_cJHnPhNX85myOn3jRIXqwNvmIFh5NRTapztldXwLTJDBkD8Bce8u3w8SwFE4ja4o8XtPyizf3XFFahKcW0G2vfIKQaZEpZmWtDXZ83VVHS9lyibnm7PyU10DyQEKLj7WzF7nbvag7eiUTGd6I-MM_UEGyHKCT3-7O6eOq9XCB4-h744WGXkqrevmGuiuZiCSnKAHJMkGggqVB9yq5fm1F7-QKjOEuteypF_MXjtdI3dkWkOox7YXif59-MUT3AYz5-Jd6KDwSdnQrLgxJb7eE7eWVt_LnVSwqXyi6WIugcq2dCOUxy_xvMyaWHpPu7haT7CH4-CNAnrFDdC3Yaqv-CVRLdjDbnT6E4UUsYgXMFZntN-OuPo0v_DUth1uDmGsiKDvty6DHlOdikWOSSIxpXCnAU1i3laxZvCGTuaUraDImEAfRB5fH3hRfnpYOFRn0CN_vuV33e86P9CKAw_KOJ2OhDXBNTnXYRdpOopy9SIb9l4jrzja1vniI8PBKU3lKHnE1gTIR91ZXRCFxJc8dONfT6n1Y45f2_7LQpg3Pd1POtUe-UzL90PS9o2uFw2wzTzvMnzTLbjbcEdZIEKPNhZzcGvONlDbjKJCgNMeQ8IDB0VIg_E6XExnKcBSECu7pbK2yz-epPbPVybzxSCi2GlOlEUe_63wHdlenBZ2F07i4G_W5JYh9eOuVWwf6c5sFr2JrN3c3TNzAJDn-gKix_4zgyBVYeyXgCA2-oONOnVR2MBaXLpSto3O66zEM0U3avUGQbcz58N7gxbHRlpyG07vmhxy2UoZpjXOupNmlGfD29znYf5oXDxgVi4JjqXFGp-Eo9n7NwxB7DRyfn2i2eVJ0JUZwkH0qetqWQ7A65puqxUGr_Mfr6Dog.OkWlFGt818m5DXGiEkJaog"

	// Metadata content.
	metadata := `{"title": "hello world", "description": "Multipart related upload test"}`

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part.
	metadataHeader := textproto.MIMEHeader{}
	metadataHeader.Set("Content-Type", "application/json")
	metadataHeader.Set("Content-ID", "metadata")
	part, err := writer.CreatePart(metadataHeader)
	if err != nil {
		log.Fatalf("Error writing metadata headers: %v", err)
	}
	part.Write([]byte(metadata))
	part.Write([]byte(data))

	// Media Files.
	//for _, mediaFilename := range positionalArgs {
	//	mediaData, errRead := ioutil.ReadFile(mediaFilename)
	//	if errRead != nil {
	//		log.Fatalf("Error reading media file: %v", errRead)
	//	}
	//	mediaHeader := textproto.MIMEHeader{}
	//	mediaHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\".", mediaFilename))
	//	mediaHeader.Set("Content-ID", "media")
	//	mediaHeader.Set("Content-Filename", mediaFilename)
	//
	//	mediaPart, err := writer.CreatePart(mediaHeader)
	//	if err != nil {
	//		log.Fatalf("Error writing media headers: %v", errRead)
	//	}
	//
	//	if _, err := io.Copy(mediaPart, bytes.NewReader(mediaData)); err != nil {
	//		log.Fatalf("Error writing media: %v", errRead)
	//	}
	//}

	// Close multipart writer.
	if err := writer.Close(); err != nil {
		log.Fatalf("Error closing multipart writer: %v", err)
	}

	// Request Content-Type with boundary parameter.
	contentType := fmt.Sprintf("multipart/related; boundary=%s", writer.Boundary())

	// Initialize HTTP Request and headers.
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalf("Error initializing a request: %v", err)
	}
	r.Header.Set("Content-Type", contentType)
	r.Header.Set("Accept", "*/*")
	r.Header.Set("authorization", token)

	// HTTP Client.
	client := &http.Client{Timeout: 180 * time.Second}
	rsp, err := client.Do(r)
	if err != nil {
		log.Fatalf("Error making a request: %v", err)
	}

	// Check response status code.
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	} else {
		log.Print("Request was a success")
	}
	var b []byte
	rsp.Body.Read(b)

	fmt.Printf("Got response %d: %s %s", rsp.StatusCode, rsp.Proto, string(b))

}
