import { Web3Storage, getFilesFromPath } from 'web3.storage'
import path from 'path';

import { fileURLToPath } from 'url';
import { dirname } from 'path';

function getAccessToken() {
    return process.env.WEB3STORAGE_TOKEN
}

export function makeStorageClient() {
    return new Web3Storage({ token: getAccessToken() })
}

async function uploadDirectory() {
    const client = makeStorageClient()

    const __filename = fileURLToPath(import.meta.url);
    const current_dir = dirname(__filename);

    const localDirectory = path.join(current_dir, '..', 'counter_contract');

    console.log("entered inside uploadDirectory fn")

    try {
        // Get the list of files from the local directory
        const files = await getFilesFromPath(localDirectory);
        console.log("have counter contract files now")

        if (files.length === 0) {
            throw new Error('No files found in the specified directory.');
        }

        // Upload the files to Web3.Storage
        const cid = await client.put(files);

        console.log(`Directory uploaded to Web3.Storage. CID: ${cid}`);
        return cid;
    } catch (error) {
        console.error('Error uploading the directory:', error);
    }
}

async function main() {

    const cid = await uploadDirectory();
    console.log("CID is ", cid);

}

main()
