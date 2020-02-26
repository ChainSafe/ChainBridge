/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const EmitterContract = artifacts.require("Emitter");

contract('Emitter - [Deployment]', async (accounts) => {
    let EmitterInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        EmitterInstance = await EmitterContract.new();
    });

    it('should deploy safe', async () => {
        const emitterAddress = EmitterInstance.address;
        const owner = await EmitterInstance.owner.call();
        assert.strictEqual(owner, emitterAddress);
    });
});