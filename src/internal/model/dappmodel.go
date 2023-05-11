package model

type DappModel struct {
	DappRecords []*DappEntity
}

type DappEntity struct {
	DappId   string
	DappName string
	ApiKey   string
}

func NewDappModel() (*DappModel, error) {
	return &DappModel{
		DappRecords: initDappData(),
	}, nil
}

func (m *DappModel) FindDappByApiKey(apiKey string) (*DappEntity, error) {
	var dapp *DappEntity
	for _, record := range m.DappRecords {
		if record.ApiKey == apiKey {
			dapp = record
			break
		}
	}

	if dapp == nil {
		return nil, ErrNotFound
	}

	return dapp, nil
}

func (m *DappModel) FindDapps() ([]*DappEntity, error) {

	return m.DappRecords, nil
}

func initDappData() []*DappEntity {
	dappRecords := make([]*DappEntity, 1)
	dappRecords[0] = &DappEntity{
		DappId:   "DappId",
		DappName: "DappName",
		ApiKey:   "ApiKey",
	}
	return dappRecords
}
