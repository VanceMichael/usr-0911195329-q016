package main
import("encoding/json";"log";"net/http";"os";"example.com/meme-context/internal/health")
func main(){m:=http.NewServeMux();m.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){json.NewEncoder(w).Encode(health.Current())});p:=os.Getenv("PORT");if p==""{p="8080"};log.Fatal(http.ListenAndServe(":"+p,m))}
