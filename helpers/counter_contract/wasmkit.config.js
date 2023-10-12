const cosmos_localnet_accounts = [
  {
    name: 'alice',
    address: 'veri1pplud7095lpfs0jnqcal3wh059gsd4ufsm5ycd',
    mnemonic: 'define cereal surround useful brief flower viable left laugh thunder alley turtle improve news invite situate south dose empower crisp fetch will rookie basic'
  },
  {
    name: 'bob',
    address: 'veri1k5j8rxpaa7kgnleju7x85gwcvk0w7nkwg8ywwv',
    mnemonic: 'print kit van romance strike march law until adult neither scatter video chaos dream bubble perfect issue chapter shrimp remove music promote gorilla wage'
  }
];

// Default list covers most of the supported network
// Networks which are not required can be removed from here
const networks = {
  cosmos_localnet: {
    endpoint: 'http://localhost:26657/',
    chainId: 'verifier',
    accounts: cosmos_localnet_accounts,
    fees: {
      upload: {
        amount: [{ amount: "750000", denom: "stake" }],
        gas: "3000000",
      },
      init: {
        amount: [{ amount: "250000", denom: "stake" }],
        gas: "1000000",
      },
      exec: {
        amount: [{ amount: "250000", denom: "stake" }],
        gas: "1000000",
      }
    },
  },
};

module.exports = {
  networks: {
    default: networks.cosmos_localnet,
  },
   
  localnetworks: {
    cosmos: {
      docker_image: "uditgulati0/juno-node",
      rpc_port: 26657,
      rest_port: 1317,
      flags: ["GAS_LIMIT=10000000", "STAKE_TOKEN=token", "TIMEOUT_COMMIT=6s"],
      docker_command: "./setup_and_run.sh veri1njjftpelthhegcl8x6rppu0req02duguaqle5z",
    },
  },
  mocha: {
    timeout: 60000
  },
  rust: {
    version: "1.71.1",
  },
  commands: {
    compile: "RUSTFLAGS='-C link-arg=-s' cargo build --release --target wasm32-unknown-unknown",
    schema: "cargo run --example schema",
  }
};
