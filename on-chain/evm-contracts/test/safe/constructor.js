/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const SafeContract = artifacts.require("Safe");

contract('Safe - [Deployment]', async (accounts) => {
    const contractOwner = accounts[0];

    let SafeInstance;

    before(async () => {
        SafeInstance = await SafeContract.new(contractOwner);
    });

    it('should set constructor values', async () => {
        let owner = await SafeInstance.owner.call();
        assert.strictEqual(contractOwner, owner);
    });
});
