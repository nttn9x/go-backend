package elastic

import (
	"context"
	"fmt"
	"net/http"

	"peakvise/common"
	"peakvise/models"

	elastic "gopkg.in/olivere/elastic.v6"
)

// InitElasticSearch hello
func initElasticSearch() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(common.AppConfig.Elastic.URL))
	return client, err

}

// SearchEs hello
func SearchEs(w http.ResponseWriter) {
	nameIndex := fmt.Sprintf("%s_%s", common.AppConfig.Elastic.Prefix, "articles")

	client, err := initElasticSearch()
	defer client.CloseIndex(nameIndex)

	if err != nil {
		common.RespondWithError(w, err, "", http.StatusInternalServerError)
	} else {
		searchResult, _ := client.Search().Index(nameIndex).Size(100).Query(elastic.NewMatchAllQuery()).Do(context.Background())
		response := &models.EsResult{
			PageIndex: 0,
			PageSize:  100,
			Total:     searchResult.Hits.TotalHits,
			Result:    searchResult.Hits}

		common.RespondWithJSON(w, models.Response{Data: response})
	}
}
