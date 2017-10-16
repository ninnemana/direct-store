package categories_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/ninnemana/direct-store/curt/categories"
	"github.com/ninnemana/direct-store/curt/client"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func TestNew(t *testing.T) {
	cfg := categories.Config{}
	_, err := categories.New(cfg)
	if err == nil {
		t.Errorf("should have errored on invalid client.Client")
		return
	}

	cfg.Client = &client.Client{}

	_, err = categories.New(cfg)
	if err == nil {
		t.Errorf("should have errored on invalid logging.Logger")
		return
	}

	cfg.Log = logging.Logger{}
	_, err = categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cfg.Log = new(logging.Logger)
	_, err = categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cfg.Log, err = logging.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestList(t *testing.T) {
	lg, err := logging.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	defer lg.Close()

	client := client.Client{
		Schema:    "https",
		Hostname:  "goapi.curtmfg.com",
		PublicKey: "9300f7bc-2ca6-11e4-8758-42010af0fd79",
	}

	switch appengine.IsDevAppServer() {
	case false:
		client.Client = http.DefaultClient
	default:
		client.Client = urlfetch.Client(context.Background())
	}

	cfg := categories.Config{
		Client: &client,
		Log:    lg,
	}

	s, err := categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = s.List()
	if err != nil {
		t.Error(err)
		return
	}
}
