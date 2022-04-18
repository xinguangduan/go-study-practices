package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func main() {
	fmt.Println("Download Started")

	fileUrl := "https://dr.sf-converter.com/download?id=233a7879601fc8b9b61d4ee727ec165c18d7031adc7c86644a59a1fe0b3c20aa&payload=1*eJzVlM%2BP20QUx%2F%2BXlXZOJPXvsStZlZN1Um%2BySTbZbDa5WI7tJJN4bMe%2FZr0IaSWglG1XFIE4AAIJCS4UoZ6oCog%2Fppste%2BJfYOxdSnvsCRpZn5k8%2B82b997M992tOEgj2x1G3tbtrUWShPHtW7cIIdU8SJN06lbtAN8iVmIv7mSqNSJtdjnTGJHZeufG03De2NFKHRS8HjGKuEqlEvuVRSjClS9F1XkQzD03Q44blEuVs9Cz8qllr%2B64xyGKXJWVRIbleYbhgYvUNexY4%2BDA0k0p9VK5x0vNKNdsgEKVY5Qqy8AqxwhVyAHkqEFFay3Wvia1Viw32Ed19jhNInl%2F7DG1VSs8mfZm2slqt%2BWPejXsAZRYc5UVGHCdtXqTJYjcdUp3EseemrsxwAvVXgKMVZ7d5uqcArCvvpIVtd38y30YARyrVkptkZMCnKmYAqkcCD2VEwDyUWIT35mGscpznMjQ2KGtCoF9WGcDlqlYMNubBD6KIp8dTvqay4IsjDKVBRhhVy2rvM01cCgAP1ZnODzxCRFtjmAz8mvRjtbQcdwEcxSVW7c911cFQZR5jgVOGqkcFKsCLwIPJ7TQrAg5GYq8IjMKK4LSRotPfzIDZhmtsQBWrhtaHsrccsEZbRJNhKEdgnxRDTqFnCABWx3pNZDQt6LAcwJP96fm0wNl2MatzmB2UCE0UyuyaHmu%2B0ydXUSBwgJOAdqOophlM4oKvmxDYQ1tyqIWdChqQQc%2FpqCpUhaJ0oGmSEmTA94%2F0fCicMAFii9w4YPLRbIifFgs%2FkpbqCOaq1qTr5vHGunvGwvNWEu8kDtZ1sPz1iRoOPlYSzvysZ4P9F2vVa%2BYeliBnVqrsaehWqp3M58c%2BjAV64mxc0wWONeFSo%2F4ph9B2%2FWGliGOEmxCss3v0AeUAbtrZn9XJv05DdgNzXxEak2js24JPu5rS6vf00fBLBFF5PFSvuzbnXR6uEeONaRr0CDraWLHDbHSMiJt4OxqFW00ljsBTOeDXktP09y%2FawzJcEzD0btaXru38q7yErCKWUynxfFjeaGkWFIqCQtKTHk4uZLXB1UoCUvKBWFBXhFKiiWlkrDk9Vvl7RGHskNvKg6szEn0xsov1YGXpFfVQWEEhYOSyP%2FX6nDd9%2F%2BbPmjGfNweGQfQDJrrHafb0LCyMrG8kpzWiSMQPYRd2Djiolk7jPW60eQP14OaPQj3zE4Smd0us6px8z4xUMou4niSEg%2FVB1J%2FlBy%2BrguFEO3WQ3M8kBK%2FTnhuJ8dyz5xMsJV29ibdxXiuKEdH3kGbZK48okI0s0k%2FXSrN4V57uIvcu%2FKSNZawrvQWMZZGbE1v9tvjsUam3I0QUWVIUOK5VBYunp5e%2FfDFn%2B%2F%2Ffnl2dvHs2ebR481nP7348oPnp2eb%2Bz%2B%2B%2BPb7zS9Pnp8%2B%2BOu3h5tPHtAPLp6eb%2B6dX310fnX6XWH8%2FOfLrz%2B%2BvP%2Fpxa%2BPNx8%2B2Tx6SH03X32z%2BeMeDXLd4VhLtm7%2Fqx%2Fv%2FQ0n7%2Blo*1650111403*9a97cfd21337e6bb"
	err := DownloadFile("9.mp4", fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
