/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ReceiverContract = artifacts.require("Receiver");

contract('Receiver - [Deployment]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            2         // relayer threshold
        )
    });

    it('should set constructor values', async () => {
        let relayers = await ReceiverInstance.TotalRelayers.call();
        assert.strictEqual(parseInt(relayers, 10), 2);

        let depositThreshold = await ReceiverInstance.DepositThreshold.call();
        assert.strictEqual(parseInt(depositThreshold, 10), 2);

        let relayerThreshold = await ReceiverInstance.RelayerThreshold.call();
        assert.strictEqual(parseInt(relayerThreshold, 10), 2);

        let relayer1 = await ReceiverInstance.Relayers(v1);
        assert.strictEqual(relayer1, true);

        let relayer2 = await ReceiverInstance.Relayers(v2);
        assert.strictEqual(relayer2, true);
    });
});