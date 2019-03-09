// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// EBC Foundation Go
	"enode://a5c3516384be24fce445067cc0a0c221f1dc6e5b74ce691dc18a69c902cae4b987a266ed29fcc15a52e670c161cd2566f3d75921f5a7ca9f3c2a5ca261edf8e5@108.61.192.135:62688",
	"enode://b60d0dabb4118ac2e676931396a97fe4a2cdfff6722de2866066f46545d6a78d12d1edff2054bd7825691463993da9f0cb050bca9827be9c151ddb3ff8d461e9@139.180.204.158:62688",
	"enode://c1e63308b71328d4f3618ae592d751368c2f9d478029305f4ab93a3df43093481db7cad8138d3a40d58c7df57b3d36709ab0ac997dd242fa1a35ed1c170c1c85@212.237.36.86:62688",
	"enode://691f4597b914db62cea8ccc6f8a42e4c5438a89b7834f34a0fd2b04508bfe209d67420cec13ef84ac0b63397341bb33f2a499958474e3cc3f51bc0c1a3a192c8@80.211.214.252:30303",
	"enode://f2a16b76922636c05ec5429a5f0f612c696a09cbeb8386766a5ee035bb96010ff84996cc35fa2cc057d7b797ee52346e904d556dc9f47213e9b08c1265baa9bb@81.2.252.58:62688",
	"enode://fe9988f1e8748fb6ff7ffb8d19100cd3c73f024f97751f69273c75e5eaa45c719d553010abaf881262ad8b23a4d3459bc541bd4018d97881c5115366f4c7b943@80.211.213.154:30303",
	"enode://a3c66a4e0b5c9fead8ffd45bd5dbdb04ed96678aeac6c49fede710d572b7978c71651493e1e0a4cc28dcfd2cdac7a6ee8c79506779ffd1a8b75c97fca1ab9a11@85.255.3.94:62688",
	"enode://f9802b18f0b59c951b24edf60481b7325708e9b158cdcd60c9a7f0b405a55c5ba17e933dbc8a0c2c8fa7e7787337d9bf52068f9e75bdfebdb897e42e8733ba31@85.255.9.89:30303",
	"enode://b3f41c5290af58e9c0cb7ec04994db98fe9d80ce7a400d439e374b47eed29dd64b9e8404bd7daaef361c72f422ad3d3249caa699e0badaec1b947150ccd6c653@94.177.180.38:30303",
	"enode://ab6b5206166f051c1796d3668da8ae897c0d83e21ba662703d782279bdfbc13cde0d29a01e1cf697e4083d307aa51079135021f65505bc1761596484fa212b64@45.76.206.141:62688",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
