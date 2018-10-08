package elastic

import (
	"context"

	"github.com/nttn9x/go-backend/models"
	"gopkg.in/olivere/elastic.v6"
)

// InitElasticSearch hello
func InitElasticSearch() *elastic.Client {
	client, _ := elastic.NewClient(elastic.SetURL("http://es-local.xena.local:9204"))
	return client

}

// SearchEs hello
func SearchEs() *models.EsResult {
	client := InitElasticSearch()
	defer client.CloseIndex("peakvise_articles")
	searchResult, _ := client.Search().Index("peakvise_articles").Query(elastic.NewMatchAllQuery()).Do(context.Background())
	response := &models.EsResult{
		PageIndex: 0,
		PageSize:  10,
		Total:     searchResult.Hits.TotalHits,
		Result:    searchResult.Hits}
	return response
}
