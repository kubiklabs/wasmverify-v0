import { Web3Storage, getFilesFromPath } from 'web3.storage'
import fs from 'fs';
import path from 'path';
import { exec } from 'child_process'
import axios from "axios";

function getAccessToken() {
    return 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDBCYmFGNTg1ZGVEZUI4YTcwZjE2QjdEZDBFNjE0RTliZjhkMzhEODUiLCJpc3MiOiJ3ZWIzLXN0b3JhZ2UiLCJpYXQiOjE2OTUzODYxNjc0OTcsIm5hbWUiOiJ2ZXJpZmllclRva2VuIn0.GDAdg-s6ZCLONU9l2p8BBzgi6z-9HGHiPaXN2VVFSVU'
    //   return process.env.WEB3STORAGE_TOKEN
}

export function makeStorageClient() {
    return new Web3Storage({ token: getAccessToken() })
}

export async function retrieve(cid) {

    const client = makeStorageClient()
    console.log("web3storage client cretaed")

    const res = await client.get(cid); // Web3Response
    if (res == null) {
        console.log("Error fetching file from IPFS");
    }

    const __filename = fileURLToPath(import.meta.url);
    const current_dir = dirname(__filename);
    const downloadDirectory = path.join(current_dir, 'downloads');

    // Ensure the download directory exists
    if (!fs.existsSync(downloadDirectory)) {
        fs.mkdirSync(downloadDirectory);
    }
    console.log("directory created")

    const files = await res?.files();
    console.log("files received from response");

    if (files != null) {
        for (const file of files) {

            const destinationPath = path.join(downloadDirectory, file.name);
            // Write the file data to the destination file
            console.log((await file.text()).length, "the length  of the file")

            fs.writeFileSync(destinationPath, await file.text());
            console.log("File downloaded and stored in the directory")

        }

    }
}

function runDockerCommand() {
    const dockerCommand = `
      docker run --rm -v /home/adarsh/my_files/arufa_research/newWeb3demo/downloads/upload:/code \
      --mount type=volume,source="$(basename "$(pwd)")_cache",target=/target \
      --mount type=volume,source=registry_cache,target=/usr/local/cargo/registry \
      cosmwasm/workspace-optimizer:0.14.0
    `;

    exec(dockerCommand, { shell: true }, function (error, stdout, stderr) {
        if (error) {
            console.error(`Error executing Docker command: ${error}`);
            return;
        }

        console.log(`Docker command output:\n${stdout}`);
        console.error(`Docker command errors:\n${stderr}`);
    });

}

const url = "http://localhost:1317";

async function main() {

    // stargate client created to query blockchain
    const verifierClient = async () => {
        await StargateClient.connect(url)
    }
    let complete_count = 0;
    // Keep quering the contract and if any pending contract found it will compile and save in the download directory
    while (true) {
        // querying current pending contract id
        await axios.get(url + "/verifier/verifier/current_pending_contract").
        then(async (requestIdres) => {
            console.log("currrent pending contract id is: ", requestIdres.data.id);
            
            if (requestIdres.data.id != 0 && complete_count < requestIdres.data.id) {
    
                // Pending contract query by it's id
                const pendingContractres = await axios.get(url + `/verifier/verifier/contract_info/${requestIdres.data.id}`)
                console.log("current pending contract is : ", pendingContractres.data);
    
                // const cid = "bafybeifg7wrwmd5sjobcvo4ctzkrse6adq3hfwvdq64bf2ika3ruys4inq";
                const cid = pendingContractres.data.PendingContracts.offchainCodeUrl;
                console.log(`CID is ${cid}`)
    
                // Download contract in contract directory from ipfs
                await retrieve(cid);
    
                // compile the contracts downloaded in the download directory
                console.log("starting docker compilation")
                runDockerCommand();
                complete_count++;
            }
        })
    }

}

main()

/*
1. How to check if the files uploaded are in correct pattern for compilation
2. Require a cargo.toml and cargo.lock file in workspace and contracts in contract folder 
*/