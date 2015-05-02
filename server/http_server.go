package server

import "github.com/cwninja/media-controller/tv"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import jsonEncoding "encoding/json"

func GetRouter(tv * tv.TV) http.Handler {
  r := mux.NewRouter()
  r.Methods("GET").Path("/status").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    status := tv.Status()
    data, _ := jsonEncoding.Marshal(status)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(data)
  })

  r.Methods("POST").Path("/stop").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tv.Stop()
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte("{}"))
  })

  r.Methods("POST").Path("/seek").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if seekPercent, e := strconv.ParseFloat(r.PostFormValue("to"), 64); e == nil {
      length := tv.Length()
      seekPosition := seekPercent * float64(length) / 100.0
      tv.SeekTo(int(seekPosition))
    } else if seekBy, e := strconv.Atoi(r.PostFormValue("by")); e == nil {
      tv.SeekBy(seekBy)
    } else {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(422)
      w.Write([]byte("{\"error\":\"WHAT?\"}"))
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte("{}"))
  })

  r.Methods("POST").Path("/play").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    url := r.PostFormValue("url")
    startPos, _ := strconv.Atoi(r.PostFormValue("pos"))
    tv.PlayFrom(url, startPos)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte("{}"))
  })

  r.Methods("POST").Path("/pause").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tv.Pause()
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte("{}"))
  })

  return r
}
