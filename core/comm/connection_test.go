/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package comm

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"github.com/hyperledger/fabric/core/config"
	"google.golang.org/grpc"
)

func TestConnection_Correct(t *testing.T) {
	config.SetupTestConfig("./../../peer")
	viper.Set("ledger.blockchain.deploy-system-chaincode", "false")
	var tmpConn *grpc.ClientConn
	var err error
	if TLSEnabled() {
		tmpConn, err = NewClientConnectionWithAddress(viper.GetString("peer.address"), true, true, InitTLSForPeer())
	}
	tmpConn, err = NewClientConnectionWithAddress(viper.GetString("peer.address"), true, false, nil)
	if err != nil {
		t.Fatalf("error connection to server at host:port = %s\n", viper.GetString("peer.address"))
	}

	tmpConn.Close()
}

func TestConnection_WrongAddress(t *testing.T) {
	config.SetupTestConfig("./../../peer")
	viper.Set("ledger.blockchain.deploy-system-chaincode", "false")
	viper.Set("peer.address", "0.0.0.0:7052")
	var tmpConn *grpc.ClientConn
	var err error
	if TLSEnabled() {
		tmpConn, err = NewClientConnectionWithAddress(viper.GetString("peer.address"), true, true, InitTLSForPeer())
	}
	tmpConn, err = NewClientConnectionWithAddress(viper.GetString("peer.address"), true, false, nil)
	if err == nil {
		fmt.Printf("error connection to server -  at host:port = %s\n", viper.GetString("peer.address"))
		t.Error("error connection to server - connection should fail")
		tmpConn.Close()
	}
}
