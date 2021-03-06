// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package bdb implements an instance of walletdb that uses boltdb for the backing
datastore.

Usage

This package is only a driver to the walletdb package and provides the database
type of "bdb".  The parameters accepted by the Open and Create functions are the
database path as a string and the database options:

	db, err := walletdb.Open("bdb", "path/to/database.db", opts bbolt.Options)
	if err != nil {
		// Handle error
	}
    opts := &bbolt.Options{
        // bbolt options
    }
	db, err := walletdb.Create("bdb", "path/to/database.db", opts bbolt.Options)
	if err != nil {
		// Handle error
	}
*/
package bdb
