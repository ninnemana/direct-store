package brands_test

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/ninnemana/direct-store/curt/brands"
	"github.com/ninnemana/direct-store/curt/client"
)

var (
	lg *logging.Client
	cl *client.Client
)

func TestMain(m *testing.M) {

	var err error

	lg, err = logging.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer lg.Close()
	lgr := lg.Logger("brands-service-test")

	cl, err = client.New(lgr, "https", "goapi.curtmfg.com", "9300f7bc-2ca6-11e4-8758-42010af0fd79", "")
	if err != nil {
		log.Fatal(err)
		return
	}

	m.Run()
}

func TestNew(t *testing.T) {
	cfg := brands.Config{}
	_, err := brands.New(cfg)
	if err == nil {
		t.Error("should have errored on invalid client.Client")
	}

	cfg.Client = &client.Client{}

	_, err = brands.New(cfg)
	if err == nil {
		t.Error("should have errored on invalid logging.Logger")
	}

	cfg.Log = logging.Logger{}
	_, err = brands.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cfg.Log = new(logging.Logger)
	_, err = brands.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cfg.Log, err = logging.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = brands.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestList(t *testing.T) {
	cfg := brands.Config{
		Client: cl,
		Log:    lg,
	}

	s, err := brands.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	bds, err := s.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(bds) == 0 {
		t.Error("brands were empty")
		return
	}
}

func TestGet(t *testing.T) {
	cfg := brands.Config{
		Client: cl,
		Log:    lg,
	}

	s, err := brands.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	bds, err := s.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(bds) == 0 {
		t.Error("brands were empty")
		return
	}

	b, err := s.Get(bds[0].ID)
	if err != nil {
		t.Error(err)
		return
	}

	if b.ID != bds[0].ID {
		t.Errorf("Brand request returned ID '%d', expected '%d'", b.ID, bds[0].ID)
		return
	}
}
