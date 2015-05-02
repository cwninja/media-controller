package tv

import "net/http"
import "io"
import "os"
import "io/ioutil"
import "fmt"
import "log"
import "errors"

func (tv * TV) sendSoapMessage(body io.Reader, action string) (responseBody []byte, err error) {
  request, err := http.NewRequest("POST", tv.Url, body)
  if err != nil {
    log.Fatalf("Could not build HTTP request: %s", err)
  }

  request.Header.Set("Content-Type", "text/xml")
  request.Header.Set("SOAPACTION", fmt.Sprintf(`"%s"`, action))

  response, err := tv.Client.Do(request)
  if err != nil {
    return
  }

  if response.StatusCode != 200 {
    io.Copy(os.Stderr, response.Body)
    return responseBody, errors.New(response.Status)
  }

  defer response.Body.Close()

  responseBody, err = ioutil.ReadAll(response.Body)
  if err != nil {
    return responseBody, err
  }

  return responseBody, nil
}
