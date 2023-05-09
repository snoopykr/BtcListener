package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/exp/slices"
)

type ReturnStruct struct {
	Result ResultStruct `json:"result"`
	Error  any          `json:"error"`
	ID     any          `json:"id"`
}

type ResultStruct struct {
	Hash              string     `json:"hash"`
	Confirmations     int        `json:"confirmations"`
	Height            int        `json:"height"`
	Version           int        `json:"version"`
	VersionHex        string     `json:"versionHex"`
	Merkleroot        string     `json:"merkleroot"`
	Time              int        `json:"time"`
	Mediantime        int        `json:"mediantime"`
	Nonce             int        `json:"nonce"`
	Bits              string     `json:"bits"`
	Difficulty        float64    `json:"difficulty"`
	Chainwork         string     `json:"chainwork"`
	NTx               int        `json:"nTx"`
	Previousblockhash string     `json:"previousblockhash"`
	Strippedsize      int        `json:"strippedsize"`
	Size              int        `json:"size"`
	Weight            int        `json:"weight"`
	Tx                []TxStruct `json:"tx"`
}

type TxStruct struct {
	Txid     string       `json:"txid"`
	Hash     string       `json:"hash"`
	Version  int          `json:"version"`
	Size     int          `json:"size"`
	Vsize    int          `json:"vsize"`
	Weight   int          `json:"weight"`
	Locktime int          `json:"locktime"`
	Vin      []VinStruct  `json:"vin"`
	Vout     []VoutStruct `json:"vout"`
	Hex      string       `json:"hex"`
}
type VinStruct struct {
	Coinbase    string   `json:"coinbase"`
	Txinwitness []string `json:"txinwitness"`
	Sequence    int64    `json:"sequence"`
}

type VoutStruct struct {
	Value        float64            `json:"value"`
	N            int                `json:"n"`
	ScriptPubKey ScriptPubKeyStruct `json:"ScriptPubKey"`
}

type ScriptPubKeyStruct struct {
	Asm     string `json:"asm"`
	Hex     string `json:"hex"`
	Address string `json:"address,omitempty"`
	Type    string `json:"type"`
}

var RpcUrl = "http://hessegg:hessegg@192.168.219.107:8332"
var FileName = "./Bitlistener.data"

var DB *sql.DB
var address []string

func main() {
	DB, _ = sql.Open("mysql", "root:1234qwer@tcp(192.168.219.107:13306)/Inae")

	gocron.Every(5).Minutes().Do(cronJob)
	<-gocron.Start()

	//cronJob()
}

func cronJob() {
	address, _ = getAddress()

	latest := getLatest()
	height := getBlockCount()

	latestInt, _ := strconv.Atoi(latest)
	heightInt, _ := strconv.Atoi(height)

	for latestInt < heightInt {
		latestInt += 1

		hash := getBlockHash(latestInt)
		getBlock(hash)
	}

	saveLatest(height)
}

func getAddress() ([]string, error) {
	var result []string

	query := `SELECT address FROM btc_address`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := ""

		err := rows.Scan(&item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

func saveLatest(height string) {
	err := ioutil.WriteFile(FileName, []byte(height), 0)
	if err != nil {
		log.Fatalf("Write File: %v", err)
	}
}

func getLatest() string {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Fatalf("Read File: %v", err)
	}

	return string(data)
}

func getBlockCount() string {
	data, err := json.Marshal(map[string]interface{}{
		"method": "getblockcount",
		"id":     1,
		"params": []interface{}{},
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}

	resp, err := http.Post(RpcUrl, "application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fmt.Sprintf("%v", result["result"])
}

func getBlockHash(height int) string {
	data, err := json.Marshal(map[string]interface{}{
		"method": "getblockhash",
		"id":     1,
		"params": []interface{}{height},
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}

	resp, err := http.Post(RpcUrl, "application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fmt.Sprintf("%v", result["result"])
}

func getBlock(blockhash string) {
	data, err := json.Marshal(map[string]interface{}{
		"method": "getblock",
		"id":     1,
		"params": []interface{}{blockhash, 2},
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}

	resp, err := http.Post(RpcUrl, "application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	result := ReturnStruct{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	dbtx, err := DB.Begin()
	if err != nil {
		log.Fatalf("DB Transaction: %v", err)
	}

	defer dbtx.Rollback()

	query := `INSERT INTO btc_history(height, hash, txid, address, value, create_dt) VALUES (?, ?, ?, ?, ?, NOW())`

	for _, tx := range result.Result.Tx {
		for _, vout := range tx.Vout {
			if vout.Value != 0 {
				idx := slices.IndexFunc(address, func(c string) bool { return c == vout.ScriptPubKey.Address })
				if idx != -1 {
					_, err := dbtx.Exec(query, result.Result.Height, result.Result.Hash, tx.Txid, vout.ScriptPubKey.Address, vout.Value)
					if err != nil {
						log.Fatalf("DB Insert: %v", err)
					}
				}
			}
		}
	}

	err = dbtx.Commit()
	if err != nil {
		log.Fatalf("DB Commit: %v", err)
	}
}
