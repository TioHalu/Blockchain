// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.


// Package yagn diperlukan untuk membuat consensus
package consensus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// memdefenisikan methode yang akan digunakan
type ChainHeaderReader interface {
	// mengambil rantai blockchain
	Config() *params.ChainConfig

	// Mengambil header dari chain lokal
	CurrentHeader() *types.Header

	// Mengambil header dari blok dari database dengan hash dan nomor
	GetHeader(hash common.Hash, number uint64) *types.Header

	// Mengambil header block dari database hanya dengan nomor
	GetHeaderByNumber(number uint64) *types.Header

	//Mengambil header block dari database hanya dengan hash
	GetHeaderByHash(hash common.Hash) *types.Header

	// Mengambil semua data yang sulit dari database dengan hash dan nomor
	GetTd(hash common.Hash, number uint64) *big.Int
}

// ChainReader mendefenisikan kumpulan kecil methode untuk mengakses chain lokal
// verifikasi header
type ChainReader interface {
	ChainHeaderReader

	// mengambil blok dari database dengan hash dan nomor
	GetBlock(hash common.Hash, number uint64) *types.Block
}

// algoritma consensus
type Engine interface {
	//penulis akan mengambil alamat ethereum dari akun yang mungkin akan terjadi error jika terdapat perbedaan header hal ini akan di dasari dari signature
	Author(header *types.Header) (common.Address, error)

	//melakukan verifikasi pada header yang memiliki signature dengan metode verifikasi seal
	VerifyHeader(chain ChainHeaderReader, header *types.Header, seal bool) error

	//memverifikasi sekumpulan header secara bersamaan sehingga bisa untuk membatalkan operasi dan mengambil verivikasi yang ter urut
	VerifyHeaders(chain ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error)

	//memberikan informasi bahwa verifikasi yang dilakukan sesuai dengan consesus
	VerifyUncles(chain ChainReader, block *types.Block) error

	//menginisialisasi consensus dari blok header
	Prepare(chain ChainHeaderReader, header *types.Header) error

	//melakukan modifikasi setelah transaksi dilakukan 
	Finalize(chain ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
		uncles []*types.Header)

	//melakukan modifikasi setelah transaksi dilakukan , header blok dan data mungkin akan diperbaharui
	FinalizeAndAssemble(chain ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
		uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error)

	//seal akan menghasilkan permintaan penyegelan baru untuk blok dan pada metode ini akan mengirimkan hasil asyn kembali
	Seal(chain ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error

	//seal akan mengembalikan hash dari sebuah blok sebelum disegel
	SealHash(header *types.Header) common.Hash

	//ini adalah algoritma untuk membuat teka teki Proof of Work
	CalcDifficulty(chain ChainHeaderReader, time uint64, parent *types.Header) *big.Int

	// mengembalikan API RPC
	APIs(chain ChainHeaderReader) []rpc.API

	// penutup
	Close() error
}

// Pow menjadi pembuktian consensus pada saat melakukan transaksi
type PoW interface {
	Engine

	//mengembalikan nilai hash
	Hashrate() float64
}
