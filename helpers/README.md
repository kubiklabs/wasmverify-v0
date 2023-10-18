### Guide to run the WasmVerifier blockchain and test it

WasmVerifier is a blockchain built using Ignite CLI which integrates the wasm module for supporting cosmwasm smart contract and have verifier module which is the core module created for on-chain verification of smart contract. This tutorial will help you in running and testing the flow for wasmverifier.
Follow [the guide](https://docs.ignite.com/welcome/install) to install ignite cli and go lang in your system. 

Follow the guide written in below steps -

#### 1. Clone the github repository and move into the directory - 

Open a terminal and  paste the below command to **clone** the repository in your system.

``` 
git clone https://github.com/kubiklabs/wasmverify-v0.git 
```

Move into the wasmverify-v0 directory with the below command.

``` cd wasmverify-v0 ```

#### 2. Setup and run the blockchain

i.  Compile the repository and build the verifierd demon.Run
```
ignite chain build
```

It will compile the wasmverifier blockchain and create verifierd demon which will be used to run and interact with the blockchain.

ii. Create genesis file and setup accounts. Run scripts in the docker directory. You first have to give executable permission for these scripts in your system.

Run 

```
chmod +x ./docker/setup_verifierd.sh
```

```
chmod +x ./docker/run_verifierd.sh
```

It will give executable access to these files.

Now run the below command to setup the blockchain.

```
./docker/setup_verifierd.sh alice
```

The name alice that we passed in the above command will create a dummy account name alice which will be used for using the verifier module as a user.

It will initialise the blockchain's data directory, creates genesis file and create a account named validator and stake tokens to the blockchain to make it a validator. You can also see the mnemonics and account address in the terminal. Store them somewhere.

Till now, you have initialised the chain and have a validator account with the name **validator** and a dummy account with the name alice. Now you can start the blockchain.Run

```
./docker/run_verifierd.sh
```

You chain is running now.

For On-chain contract verification, we first have to upload the smart contract on the blockchain against which we will verify the off-chain contract code.

#### 3. Upload the counter contract on-chain

WasmVerifier repository contains counter contract in the **wasmverify-v0/helpers/counter_contract** directory. We will first compile the contract and then upload the wasm file on the blockchain.

Open a new terminal and move to **wasmverify-v0/helpers/counter_contract** directory.

Run this docker command to compile the contract. It will create a artifacts directory in the here(**wasmverify-v0/helpers/counter_contract**) in which compiled wasm file along with the checksum.txt(It contains the hash of the wasm file).

Start the docker before running the below command. Now run,

```
docker run --rm -v "$(pwd)":/code \
      --mount type=volume,source="$(basename "$(pwd)")_cache",target=/target \
      --mount type=volume,source=registry_cache,target=/usr/local/cargo/registry \
      cosmwasm/workspace-optimizer:0.14.0
```

Now, Upload the compiled wasm to the chain.
Run the command from the same directory (**wasmverify-v0/helpers/counter_contract**)-

```
verifierd tx wasm store /artifacts/contracts/counter_contract.wasm --chain-id testing --from alice --gas=60000000
```

You will see something like this in your terminal. It's a prompt to sign the tx.
```
auth_info:
  fee:
	amount: []
	gas_limit: "60000000"
	granter: ""
	payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /cosmwasm.wasm.v1.MsgStoreCode
	instantiate_permission: null
	sender: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd
	wasm_byte_code: H4sIAAAAAAAA/+y9DbhdV3UYuPf……………………………………………………………………………………………………………………………………../98GNxyBXsCAA==
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
```

Enter **y** to proceed. The final tx will be something like this.

```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5CA6633AF8A2069C2BF3E9AA877BE4820F626F382C8C27B5741F5A79947559D6
```
You can also query the transaction by the txhash to verify the contract was indeed deployed(if tx was successful). Run

```
verifierd query tx ${txhash}
```

Here txhash is the hash obtained in the above transaction(The last line from the terminal)

Output will look something like this -

```
verifierd query tx 5CA6633AF8A2069C2BF3E9AA877BE4820F626F382C8C27B5741F5A79947559D6
code: 0
codespace: ""
data: 124E0A262F636F736D7761736D2E7761736D2E76312E4D736753746F7265436F6465526573706F6E7365122408011220E9EB006A6709157E1E436CAD5745B2AD973A52C82A30691AD8B092DC92811355
events:
- attributes:
  - index: true
	key: fee
	value: ""
  - index: true
	key: fee_payer
	value: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd
  type: tx
- attributes:
  - index: true
	key: acc_seq
	value: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd/2
  type: tx
- attributes:
  - index: true
	key: signature
	value: /+7iaLIqxaEvwTlunkoeXmjJ53Amgl6q/Az8XcIz5r89e2Bbet5Y6w0ry3THEB+EIM30NGW8I9HFCchK1Fup3Q==
  type: tx
- attributes:
  - index: true
	key: action
	value: /cosmwasm.wasm.v1.MsgStoreCode
  - index: true
	key: sender
	value: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd
  - index: true
	key: module
	value: wasm
  type: message
- attributes:
  - index: true
	key: code_checksum
	value: e9eb006a6709157e1e436cad5745b2ad973a52c82a30691ad8b092dc92811355
  - index: true
	key: code_id
	value: "1"
  type: store_code
gas_used: "1068721"
gas_wanted: "60000000"
height: "2147"
info: ""
logs:
- events:
  - attributes:
	- key: action
  	value: /cosmwasm.wasm.v1.MsgStoreCode
	- key: sender
  	value: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd
	- key: module
  	value: wasm
	type: message
  - attributes:
	- key: code_checksum
  	value: e9eb006a6709157e1e436cad5745b2ad973a52c82a30691ad8b092dc92811355
	- key: code_id
  	value: "1"
	type: store_code
  log: ""
  msg_index: 0
raw_log: '[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"/cosmwasm.wasm.v1.MsgStoreCode"},{"key":"sender","value":"veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd"},{"key":"module","value":"wasm"}]},{"type":"store_code","attributes":[{"key":"code_checksum","value":"e9eb006a6709157e1e436cad5745b2ad973a52c82a30691ad8b092dc92811355"},{"key":"code_id","value":"1"}]}]}]'
timestamp: "2023-08-21T08:33:37Z"
tx:
  '@type': /cosmos.tx.v1beta1.Tx
  auth_info:
	fee:
  	amount: []
  	gas_limit: "60000000"
  	granter: ""
  	payer: ""
	signer_infos:
	- mode_info:
    	single:
      	mode: SIGN_MODE_DIRECT
  	public_key:
    	'@type': /cosmos.crypto.secp256k1.PubKey
    	key: AgHxkh3di4daAa8jyNOqItRF0q4LHZZES0fxVA2sgO36
  	sequence: "2"
	tip: null
  body:
	extension_options: []
	memo: ""
	messages:
	- '@type': /cosmwasm.wasm.v1.MsgStoreCode
  	instantiate_permission: null
  	sender: veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd
  	wasm_byte_code: H4sIAAAAAAAA/+y9DbhdV3UYuPfa59y/c8595/3o6enHeJ8bd3o1YxXRcd5ThNPRfuN………………………………………………………………………….ysPxwUbIAQl6INzTxkkE20kPT7BKEG4nU0HANXBRUSRv0NMMxCEJwI3CNRBwIMi1Awh5QfR6OWmx8Gc8SAiDXtgE20TYDCTYQjGKOXi/BOB6AQhAKklgonIB0FiSxIKrav7uCMnghpAMrgrJYGVIBjMDMvi/AQAA//98GNxyBXsCAA==
	non_critical_extension_options: []
	timeout_height: "0"
  signatures:
  - /+7iaLIqxaEvwTlunkoeXmjJ53Amgl6q/Az8XcIz5r89e2Bbet5Y6w0ry3THEB+EIM30NGW8I9HFCchK1Fup3Q==
txhash: 5CA6633AF8A2069C2BF3E9AA877BE4820F626F382C8C27B5741F5A79947559D6
```

#### 3. upload off-chain contract on IPFS node

At this point of time, we have contract uploaded on the blockchain. Now we can start the off-chain code verification process.

To start the process, we first have to upload our contract code on IPFS node so that validators can access the contract code.

To upload the contract on IPFS, move to  **wasmverify-v0/helpers/ipfs_scripts** directory and run the below commands in the terminal - 

```
// to install all the dependecies
npm install
```

Now in the upload_contract.js file you can see get getAccessToken function which return the web3 storage token. Here you can paste your your own access token key. For testing you can keep it as it is. Now run,

```
node ./upload_contract.js
```

After contract gets uploaded to the IPFS, you will be able to see a CID in the terminal something like this **bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq**. Store this cid somewhere.

#### 4. Set Blocktime in the blockchain
The process for downloading, compiling and uploading the hash to the blockchain takes time, so we will have to set a sufficient time for validator to upload the hash data to the chain(right now anyone can change blocktimme, poc only)

Make a transaction to set block time.

Run

```
verifierd tx verifier update-block-time 50 25 10 --from validator --chain-id testing
```

10 is for calculating the hash, 15 is for submitting the prevote msg and 20 is for sending the vote msg to the chain.

You will prompted to sign the txn. Type Y and enter. Now you can see txhash in your terminal.
```
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /verifier.verifier.MsgUpdateBlockTime
    compilationBlocks: "10"
    creator: veri1h66ek0f65mxv2k08n9cdxgjdx488lmkn9msyp3
    prevoteBlocks: "10"
    voteBlocks: "10"
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5C4318E66B0D5A012BE6CDAA2094CBF3ED5AF019114DDD94C7C143F29B50BE6E
```

you can check if the txn is successful by running
```
//txhash = 5C4318E66B0D5A012BE6CDAA2094CBF3ED5AF019114DDD94C7C143F29B50BE6E
verifierd q tx {txhash}
```

#### 5. Apply for verification process(step 1)

The verification process consist of two steps.
In the first step, you will make the transaction and send the CID obtained in the above step that you got after uploading contract to ipfs. This txn will be your application to verify the contract. You will receive an **application id** in reponse to this txn.(**CHECK HOW TO SEE RESPOSE RECEIVED FROM THE TRANSACTION**)

To apply for verification make the transaction, run the command -

```
verifierd tx verifier apply-verify-application "{Your_CID}" --chain-id testing --from alice
```

You will be prompted to sign the transaction, type y and enter. You will get response like this
```
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /verifier.verifier.MsgVerifyContract
    contractURL: bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq
    creator: veri1h9706x06yfdnjr8dtzxa59ut6df8f7r4gxuh7k
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 92AD3FF7F7B609B5A8C7957926F53F5C6CE314FB292F8A6930BF994E7853E826
``` 
Again query the txhash and check the details about the transaction.
`` verifierd query tx ${txhash}`

You have now applied for the application and you have received application id for your application. You can query the status for you application.

To check your application-info -

```
verifierd q verifier contract-info {application id}
```

This is the first application so id will 1.

Something like this will pop up in your terminal.

```
contractinfo:
  Id: "1"
  assignedVerificationBlockHeight: "48"
  codeId: "0"
  offchainCodeHash: ""
  offchainCodeUrl: bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq
```

**assignedVerificationBlockHeight** is the height after which you can make your second txn for final contract verification. Till this height validators have already sent the hash data to the chain and  final hash gets uploaded on chain.
**codeId** represents the code-id for which this application's offchain code matches. Here you can see it is zero, it represents that it has not been verified yet.(codeId is a non zero value)
**offchainCodeHash** is the finalize hash which gets updated at the end of assignedVerificationBlockHeight block. It is the final hash on which consensus has happended and this is the hash for the offchain contract code.
**offchainCodeUrl** is the cid that you provided while applying for the verification application.

#### 6. Validators action

<!-- ##### 6.1 Query the current pending contract

Validators keeps on querying for the current pending contract.In case of pending contract they start the action.
Now you also have the validator key. So you will have to make the prevote and vote msg to the blockchain.

Check for pending contract.Run 

``verifierd q verifier current-pending-contract``

If there is any pending contract, you can see the application id for it.
Now you can see
```
"id" : 1
```

Query this id to get the information for the application.Run

``verifierd q verifier contract-info 1``

You can see the info like this -
```
contractinfo:
  Id: "1"
  assignedVerificationBlockHeight: "248"
  codeId: "0"
  offchainCodeHash: ""
  offchainCodeUrl: bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq
```

Copy the offchainCodeUrl and  -->

##### 6.1 Download and compile the contract

Validators run a script which keeps quering for pending application and if it gets any it downlaods and compile the contract.

To run the script, open a new terminal and move to **wasmverify-v0/helper/ipfs_scripts** directory and run -

```
node ./validator_helper.js
``` 

It will download the contract in the ipfs_scripts/download directory and compile the contract. 

##### 6.2 Create prevote msg for prevote

Validators have to take hash of the contract codehash along with some salt and its address.
To create the prevote msg, open a new terminal move to **/wasmverify-v0/helpers/prevotemsggen** directory and run the command 
```
go run prevoteHashGenerator.go
```
It will prompt to ask salt codehash and address of the validator to create the prevote msg. Enter any random string, for example "thisIsSalt" then the codeHash which you can get in **helpers/ipfs_scripts/downloads/artifacts/checksum.txt** file. Copy (it will be something like **8230425a1828d8b593c7e1cee39600219d6cd795d47ea6ac9a86b5e603da35cd**) and paste in the terminal and then space and the validator address which was shown in  the terminal while running the setup_verifier.sh script, paste here in the termnal and enter. You can now see some thing like **2117efa95e4d89684b1e33e9bf7ac345f37aca4d** in your terminal. Copy this (hash).

##### 6.3 Make prevote message in the blockchain

To make prevote txn we will have to run 
```
verifierd tx verifier aggregate-code-hash-prevote {application-id} {prevote msg} --chain-id testing --from validator
``` 
In our case, it will be
```
verifierd tx verifier aggregate-code-hash-prevote 1 2117efa95e4d89684b1e33e9bf7ac345f37aca4d --chain-id testing --from validator
``` 

Again type Y and enter. Now you have made the prevote msg.

##### 6.4 Make the vote message in the blockchain

To make vote txn we will have to run 
```
verifierd tx verifier aggregate-code-hash-vote {application-id} {salt} {codehash} --chain-id testing --from validator
``` 
In our case, it will 
```
verifierd tx verifier aggregate-code-hash-vote 1 thisIsSalt 8230425a1828d8b593c7e1cee39600219d6cd795d47ea6ac9a86b5e603da35cd --chain-id testing --from validator
``` 

Again type Y and enter. Now you have made the vote msg.

Now wait for the blockheight to reach **assignedVerificationBlockHeight**.

After blockheight > assignedVerificationBlockHeight

#### 7. Final verification step(step 2 by user)

Now user alice will make the second transaction to verify the onchain code. Before it you can query the application info and check the current status.

```
verifierd q tx contract-info 1
```
You can see the info like this -

```
contractinfo:
  Id: "1"
  assignedVerificationBlockHeight: "48"
  codeId: "0"
  offchainCodeHash: "8230425a1828d8b593c7e1cee39600219d6cd795d47ea6ac9a86b5e603da35cd"
  offchainCodeUrl: bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq
```

**offchainCodeHash** has been updated now.

Make the final verification txn.

```
verifierd tx verifier final-verification {application id} {code id} --from validator --chain-id testing
```

In our case, application id is 1 and code id is also 1. So ,
```
verifierd tx verifier final-verification 1 1 --from validator --chain-id testing
```

If the code-id's corrosponding onchain codehash matches the offchaincodehash finalised by validators then teh txn will become successful and the code id will then get updated in the applciation. If ypu query the application you can see

```
contractinfo:
  Id: "1"
  assignedVerificationBlockHeight: "48"
  codeId: "1"
  offchainCodeHash: "8230425a1828d8b593c7e1cee39600219d6cd795d47ea6ac9a86b5e603da35cd"
  offchainCodeUrl: bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq
```

the **codeId** is now  updated. 


