const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

async function main() {
    try {
        const ccpPath = path.resolve(__dirname, '../network/connection-RBI.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        const identity = await wallet.get('appUser');
        if (!identity) {
            console.log('An identity for "appUser" does not exist in the wallet');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork('mychannel');
        const contract = network.getContract('lendingChaincode');

        // TODO: Implement logic to request a loan
        console.log('Submitting transaction: RequestLoan...');
        // TODO: Call contract.submitTransaction() to request a loan
        console.log('Loan request submitted successfully.');

        // TODO: Implement logic to approve the loan request
        console.log('Approving Loan...');
        // TODO: Call contract.submitTransaction() to approve a loan
        console.log('Loan approved successfully.');

        // TODO: Implement logic to repay the loan partially
        console.log('Repaying Loan...');
        // TODO: Call contract.submitTransaction() to repay the loan
        console.log('Loan repayment made successfully.');

        // TODO: Implement logic to query loan details
        console.log('Querying loan details...');
        // TODO: Call contract.evaluateTransaction() to fetch loan details
        console.log('Loan details fetched successfully.');

        await gateway.disconnect();
    } catch (error) {
        console.error(`Error: ${error}`);
    }
}

main();
