package rpc

type NacBaseMessage struct {
	Flag    bool   `json:"flag"`
	Message string `json:"message"`
	Data    interface{}
}

type NacNewAccountData struct {
	PrivateKey string `json:"privateKey"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Salt       string `json:"salt"`
	Name       string `json:"name"`
	Mnemonic   string `json:"mnemonic"`
}

type NacBlockDetailData struct {
	BlocDataRow struct {
		HeightText string `json:"heightText"`
		HashText   string `json:"hashText"`
		TimeText   string `json:"timeText"`
		MinerText  string `json:"minerText"`
		RawBlock   struct {
			Instance         int    `json:"instance"`
			Height           int    `json:"height"`
			Timestamp        int64  `json:"timestamp"`
			Miner            string `json:"miner"`
			BlockReward      int    `json:"blockReward"`
			CollectMined     int    `json:"collectMined"`
			Size             int    `json:"size"`
			ParentHash       string `json:"parentHash"`
			Version          int    `json:"version"`
			TransactionsRoot string `json:"transactionsRoot"`
			TxVolume         int    `json:"txVolume"`
			GasUsed          int    `json:"gasUsed"`
			GasLimit         int    `json:"gasLimit"`
			GasMinimum       int    `json:"gasMinimum"`
			GasMaximum       int    `json:"gasMaximum"`
			ExtraData        struct {
				Client         string `json:"client"`
				DataCenter     string `json:"dataCenter"`
				DappInvoking   string `json:"dappInvoking"`
				GasDestroy     int    `json:"gasDestroy"`
				UninstallAward int    `json:"uninstallAward"`
				GasAward       int    `json:"gasAward"`
				BleedValue     int    `json:"bleedValue"`
			} `json:"extraData"`
			Hash      string `json:"hash"`
			MinedSign string `json:"minedSign"`
		} `json:"rawBlock"`
	} `json:"blocDataRow"`
	TxList []string `json:"txList"`
}

type NacTxDetail struct {
	HashText     string `json:"hashText"`
	TxHeightText string `json:"txHeightText"`
	DateTimeText string `json:"dateTimeText"`
	FromText     string `json:"fromText"`
	ToText       string `json:"toText"`
	AmountText   string `json:"amountText"`
	StatusText   string `json:"statusText"`
	FeeText      string `json:"feeText"`
	RawTx        struct {
		Instance       int           `json:"instance"`
		Version        int           `json:"version"`
		Timestamp      int64         `json:"timestamp"`
		Token          int           `json:"token"`
		From           string        `json:"from"`
		To             string        `json:"to"`
		Value          int           `json:"value"`
		Gas            int           `json:"gas"`
		GasType        int           `json:"gasType"`
		GasLimit       int           `json:"gasLimit"`
		GasInputData   []interface{} `json:"gasInputData"`
		TxHeight       int           `json:"txHeight"`
		TxType         int           `json:"txType"`
		Context        string        `json:"context"`
		Remarks        string        `json:"remarks"`
		BlockCondition int           `json:"blockCondition"`
		Hash           string        `json:"hash"`
		SenderSign     string        `json:"senderSign"`
		Status         int           `json:"status"`
		EventStatus    int           `json:"eventStatus"`
		EventInfo      string        `json:"eventInfo"`
		EventStates    []struct {
			StateType       int    `json:"stateType"`
			Tx              string `json:"tx"`
			Address         string `json:"address"`
			Before          int64  `json:"before"`
			After           int64  `json:"after"`
			StateDifference int    `json:"stateDifference"`
			Hash            string `json:"hash"`
		} `json:"eventStates"`
		BleedValue int           `json:"bleedValue"`
		MinedSign  string        `json:"minedSign"`
		InputData  []interface{} `json:"inputData"`
		Output     struct {
			Instance int    `json:"instance"`
			TargetTx string `json:"targetTx"`
			Amount   int    `json:"amount"`
		} `json:"output"`
		ChangeTx    string `json:"changeTx"`
		BlockHeight int    `json:"blockHeight"`
	} `json:"rawTx"`
}

type NacBalance struct {
	Instance        int         `json:"instance"`
	TokenBalanceMap interface{} `json:"tokenBalanceMap"`
}

type NacWithdrawData struct {
	Mail interface{} `json:"mail"`
	Hash string      `json:"hash"`
}
