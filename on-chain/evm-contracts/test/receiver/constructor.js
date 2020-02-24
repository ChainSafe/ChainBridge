/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ReceiverContract = artifacts.require("Receiver");

contract('Receiver - [Deployment]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('should set constructor values', async () => {
        let validators = await ReceiverInstance.TotalValidators.call();
        assert.strictEqual(parseInt(validators, 10), 2);

        let depositThreshold = await ReceiverInstance.DepositThreshold.call();
        assert.strictEqual(parseInt(depositThreshold, 10), 2);

        let validatorThreshold = await ReceiverInstance.ValidatorThreshold.call();
        assert.strictEqual(parseInt(validatorThreshold, 10), 2);

        let validator1 = await ReceiverInstance.Validators(v1);
        assert.strictEqual(validator1, true);

        let validator2 = await ReceiverInstance.Validators(v2);
        assert.strictEqual(validator2, true);
    });
});