package categories_test

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/ninnemana/direct-store/curt/categories"
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

	lgr := lg.Logger("categories-service-test")
	cl, err = client.New(lgr, "https", "goapi.curtmfg.com", "9300f7bc-2ca6-11e4-8758-42010af0fd79", "")
	if err != nil {
		log.Fatal(err)
		return
	}

	m.Run()
}

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

	cfg := categories.Config{
		Client: cl,
		Log:    lg,
	}

	s, err := categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cats, err := s.List(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if len(cats) == 0 {
		t.Error("categories were empty")
		return
	}

	brands := map[int][]string{}
	for _, cat := range cats {
		if _, ok := brands[cat.Brand.ID]; !ok {
			brands[cat.Brand.ID] = []string{}
		}

		brands[cat.Brand.ID] = append(brands[cat.Brand.ID], cat.Title)
	}

	if len(brands) < 2 {
		t.Error("failed to retrieve more than one brand")
	}

	brandID := cats[0].Brand.ID
	cats, err = s.List(&categories.ListParams{
		BrandID: &brandID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	brands = map[int][]string{}
	for _, cat := range cats {
		if _, ok := brands[cat.Brand.ID]; !ok {
			brands[cat.Brand.ID] = []string{}
		}

		brands[cat.Brand.ID] = append(brands[cat.Brand.ID], cat.Title)
	}

	if len(brands) != 1 {
		t.Error("failed to retrieve one brand")
	}
}

func TestGet(t *testing.T) {

	cfg := categories.Config{
		Client: cl,
		Log:    lg,
	}

	s, err := categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cats, err := s.List(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if len(cats) == 0 {
		t.Error("categories were empty")
		return
	}

	cat, err := s.Get(nil)
	if err == nil {
		t.Error("should fail when retrieving a category with nil params")
	}

	cat, err = s.Get(&categories.GetParams{
		ID: &cats[0].ID,
	})
	if err != nil {
		t.Error(err)
	}
	if cat.ID != cats[0].ID {
		t.Error("returned the wrong category")
	}
}

func TestGetParts(t *testing.T) {

	cfg := categories.Config{
		Client: cl,
		Log:    lg,
	}

	s, err := categories.New(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	cats, err := s.List(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if len(cats) == 0 {
		t.Error("categories were empty")
		return
	}

	parts, err := s.GetParts(nil)
	if err == nil {
		t.Error("should fail when retrieving a category with nil params")
	}

	parts, err = s.GetParts(&categories.GetParams{
		ID: &cats[0].ID,
	})
	if err != nil {
		t.Error(err)
	}
	if len(parts.Parts) == 0 {
		t.Error("should have returned parts")
	}
}
