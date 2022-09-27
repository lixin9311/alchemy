package alchemy

import (
	"encoding/json"
	"time"
)

type NotifyEvent struct {
	WebhookID string            `json:"webhookId"`
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"createdAt"`
	Type      string            `json:"type"`
	Event     AddressWatchEvent `json:"event"`
}

type Network string

const (
	ARB_MAINNET Network = "ARB_MAINNET"
	ARB_GOERLI  Network = "ARB_GOERLI"
)

type TxnCategory string

const (
	TxnExternal TxnCategory = "external"
	TxnInternal TxnCategory = "internal"
	TxnErc721   TxnCategory = "erc721"
	TxnErc1155  TxnCategory = "erc1155"
	TxnErc20    TxnCategory = "erc20"
	TxnToken    TxnCategory = "token"
)

type AddressWatchEvent struct {
	Network  Network            `json:"network"`
	Activity []*AddressActivity `json:"activity"`
}

type ContractInfo struct {
	RawValue string `json:"rawValue"` // raw transfer value (hex string). Omitted if erc721 transfer
	Address  string `json:"address"`  // contract address (hex string). Omitted if external or internal transfer
	Decimals int    `json:"decimals"` // contract decimal (hex string). Omitted if not defined in the contract and not available from other sources.
}

type Log struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`      // the block where the transfer occurred (hex string)
	TransactionHash  string   `json:"transactionHash"`  // the transaction hash (hex string)
	TransactionIndex string   `json:"transactionIndex"` // the transaction index (hex string)
	BlockHash        string   `json:"blockHash"`        // the block hash (hex string)
	LogIndex         string   `json:"logIndex"`         // the log index (hex string)
	Removed          bool     `json:"removed"`          // true if the log was removed, due to a chain reorganization
}

type AddressActivity struct {
	FromAddress      string           `json:"fromAddress"`
	ToAddress        string           `json:"toAddress"`
	BlockNum         string           `json:"blockNum"`         // the block where the transfer occurred (hex string)
	Hash             string           `json:"hash"`             // the transaction hash (hex string)
	Category         TxnCategory      `json:"category"`         // the category of the transfer (e.g. external, internal, erc721, erc1155, erc20, or token)
	Value            float64          `json:"value"`            // converted asset transfer value as a number (raw value divided by contract decimal). Omitted if erc721 transfer or contract decimal is not available.
	Asset            string           `json:"asset"`            // ETH or the token's symbol. Omitted if not defined in the contract and not available from other sources.
	Erc721TokenID    string           `json:"erc721TokenId"`    // the erc721 token id (hex string). Omitted if not an erc721 transfer.
	Erc1155Metadata  *json.RawMessage `json:"erc1155Metadata"`  // A list of objects containing the ERC1155 tokenId (hex string) and value (hex string). Omitted if not an ERC1155 transfer
	RawContract      ContractInfo     `json:"rawContract"`      // the contract info for the transfer
	TypeTraceAddress string           `json:"typeTraceAddress"` // the type of internal transfer (call, staticcall, create, suicide) followed by the trace address (ex. call_0_1).Omitted if not internal transfer. (note you can use this as a unique id for internal transfers since they will have the same parent hash)
	Log              Log              `json:"log"`
}
